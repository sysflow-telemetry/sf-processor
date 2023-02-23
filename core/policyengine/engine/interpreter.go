//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
// Andreas Schade <san@zurich.ibm.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package engine implements a rules engine for telemetry records.
package engine

import (
	"sync"
	"time"

	"github.com/paulbellamy/ratecounter"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source"
)

// PolicyInterpreter defines a rules engine for SysFlow data streams.
type PolicyInterpreter[R any] struct {

	// Input policy language compiler
	pc policy.PolicyCompiler[R]

	// Configuration
	config Config

	// Prefilter
	prefilter source.Prefilter[R]

	// Record contextualizer
	ctx source.Contextualizer[R]

	// Parsed rule and filter object maps
	rules   []policy.Rule[R]
	filters []policy.Filter[R]

	// Worker channel and waitgroup
	workerCh chan R
	wg       *sync.WaitGroup

	// Callback for sending records downstream
	out func(R)

	// Worker pool size
	concurrency int

	// Action Handler
	ah *ActionHandler[R]

	// Rate counter
	rc       *ratecounter.RateCounter
	lastRcTs time.Time
}

// NewPolicyInterpreter constructs a new interpreter instance.
func NewPolicyInterpreter[R any](conf Config, pc policy.PolicyCompiler[R], pf source.Prefilter[R], ctx source.Contextualizer[R], out func(R)) *PolicyInterpreter[R] {
	pi := new(PolicyInterpreter[R])
	pi.pc = pc
	if pi.prefilter = pf; pf == nil {
		pi.prefilter = source.NewDefaultPrefilter[R]()
	}
	if pi.ctx = ctx; ctx == nil {
		pi.ctx = source.NewDefaultContextualizer[R]()
	}
	pi.config = conf
	pi.concurrency = conf.Concurrency
	pi.rules = make([]policy.Rule[R], 0)
	pi.filters = make([]policy.Filter[R], 0)
	pi.out = out
	pi.ah = NewActionHandler[R](conf)

	// This should only be used for benchmarking the engine
	if logger.IsEnabled(logger.Perf) {
		pi.rc = ratecounter.NewRateCounter(1 * time.Second)
		pi.lastRcTs = time.Now()
	}
	return pi
}

// StartWorkers creates the worker pool.
func (pi *PolicyInterpreter[R]) StartWorkers() {
	logger.Trace.Printf("Starting policy engine's thread pool with %d workers", pi.concurrency)
	pi.workerCh = make(chan R, pi.concurrency)
	pi.wg = new(sync.WaitGroup)
	pi.wg.Add(pi.concurrency)
	for i := 0; i < pi.concurrency; i++ {
		go pi.worker()
	}
}

// StopWorkers stops the worker pool and waits for all tasks to finish.
func (pi *PolicyInterpreter[R]) StopWorkers() {
	logger.Trace.Println("Stopping policy engine's thread pool")
	close(pi.workerCh)
	pi.wg.Wait()
}

// Compile parses and interprets a set of input policies defined in paths.
func (pi *PolicyInterpreter[R]) Compile(paths ...string) (err error) {
	if pi.rules, pi.filters, err = pi.pc.Compile(paths...); err != nil {
		return err
	}
	if logger.IsEnabled(logger.Perf) {
		if pi.config.BenchRuleIndex >= 0 && pi.config.BenchRuleIndex < len(pi.rules) {
			pi.rules = append(make([]policy.Rule[R], 0), pi.rules[pi.config.BenchRuleIndex])
		} else if pi.config.BenchRulesetSize >= 0 && pi.config.BenchRulesetSize <= len(pi.rules) {
			pi.rules = append(make([]policy.Rule[R], 0), pi.rules[:pi.config.BenchRulesetSize]...)
		}
		logger.Perf.Printf("Benchmarking %d rule(s)", len(pi.rules))
		for _, r := range pi.rules {
			logger.Perf.Printf("Rule Name: %s, Description: %-50s", r.Name, r.Desc)
		}
	}
	logger.Info.Printf("Policy engine loaded %d rules and %d prefilters", len(pi.rules), len(pi.filters))
	pi.ah.CheckActions(pi.rules)
	return nil
}

// ProcessAsync queues the record for processing in the worker pool.
func (pi *PolicyInterpreter[R]) ProcessAsync(r R) {
	pi.workerCh <- r
	if logger.IsEnabled(logger.Perf) && time.Since(pi.lastRcTs) > (15*time.Second) {
		logger.Perf.Println("Policy engine rate (events/sec): ", pi.rc.Rate())
		pi.lastRcTs = time.Now()
	}
}

// Asynchronous worker thread: apply all compiled policies, enrich matching records, and send records downstream.
func (pi *PolicyInterpreter[R]) worker() {
	for {
		// Fetch record
		r, ok := <-pi.workerCh
		if !ok {
			logger.Trace.Println("Worker channel closed. Shutting down.")
			break
		}

		// Increment rate counter
		if logger.IsEnabled(logger.Perf) {
			pi.rc.Incr(1)
		}

		// Drop record if any drop rule applied
		if pi.evalFilters(r) {
			continue
		}

		// Enrich mode is non-blocking: Push record even if no rule matches
		match := (pi.config.Mode == EnrichMode)

		// Apply rules
		for _, rule := range pi.rules {
			if rule.Enabled && pi.prefilter.IsApplicable(r, rule) && rule.Condition.Eval(r) {
				if pi.ctx != nil {
					pi.ctx.AddRules(r, rule)
				}
				pi.ah.HandleActions(rule, r)
				match = true
			}
		}

		// Push record if a rule matches (or if mode is enrich)
		if match && pi.out != nil {
			pi.out(r)
		}
	}
	pi.wg.Done()
}

// EvalFilters executes compiled policy filters against record r.
func (pi *PolicyInterpreter[R]) evalFilters(r R) bool {
	for _, f := range pi.filters {
		if f.Enabled && f.Condition.Eval(r) {
			return true
		}
	}
	return false
}

// sampleRules is used in performance benchmarks to randomly sample a subset of rules.
// func (pi *PolicyInterpreter[R]) sampleRules(n int) []policy.Rule[R] {
// 	rand.Seed(time.Now().Unix())
// 	permutation := rand.Perm(len(pi.rules))
// 	rules := make([]policy.Rule[R], 0)
// 	for i := 0; i < n && i < len(pi.rules); i++ {
// 		rules = append(rules, pi.rules[permutation[i]])
// 	}
// 	return rules
// }

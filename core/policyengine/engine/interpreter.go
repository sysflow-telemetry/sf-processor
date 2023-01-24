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

	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy/falco/lang/parser"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source"
)

// PolicyInterpreter defines a rules engine for SysFlow data streams.
type PolicyInterpreter[R any] struct {
	*parser.BaseSfplListener

	// Input policy language compiler
	pc policy.PolicyCompiler[R]

	// Prefilter
	prefilter Prefilter[R]

	// Record contextualizer
	ctx source.Contextualizer[R]

	// Mode of the policy interpreter
	mode Mode

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
}

// NewPolicyInterpreter constructs a new interpreter instance.
func NewPolicyInterpreter[R any](conf Config, out func(R)) *PolicyInterpreter[R] {
	pi := new(PolicyInterpreter[R])
	// pi.pc = should we concretize compiler and operations object in configuration object? or in policy engine?
	// pi.prefilter = ...
	// pi.ctx = ...
	//pc := falco.NewPolicyCompiler(flatrecord.NewOperations())
	//pi.pc = pc
	pi.mode = conf.Mode
	pi.concurrency = conf.Concurrency
	pi.rules = make([]policy.Rule[R], 0)
	pi.filters = make([]policy.Filter[R], 0)
	pi.out = out
	pi.ah = NewActionHandler[R](conf)
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
	pi.ah.CheckActions(pi.rules)
	return nil
}

// ProcessAsync queues the record for processing in the worker pool.
func (pi *PolicyInterpreter[R]) ProcessAsync(r R) {
	pi.workerCh <- r
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

		// Drop record if any drop rule applied.
		if pi.evalFilters(r) {
			continue
		}

		// Enrich mode is non-blocking: Push record even if no rule matches
		match := (pi.mode == EnrichMode)

		// Apply rules
		for _, rule := range pi.rules {
			if rule.Enabled && pi.prefilter.IsApplicable(r, rule) && rule.Condition.Eval(r) {
				pi.ctx.AddRules(r, rule)
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

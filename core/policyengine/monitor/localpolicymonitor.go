//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
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

// Package monitor implements a policy monitor for the policy engine.
package monitor

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"io"
	"os"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/sysflow-telemetry/sf-apis/go/ioutils"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/engine"
)

// LocalPolicyMonitor is an object that monitors the local policy file
// directory for changes and compiles a new policy engine if changes occur.
type LocalPolicyMonitor[R any] struct {
	config    engine.Config
	interChan chan *engine.PolicyInterpreter[R]
	watcher   *fsnotify.Watcher
	started   bool
	done      chan bool
	policies  map[string][]byte
	out       func(R)
}

// NewLocalPolicyMonitor returns a new policy monitor object given an engine configuration.
func NewLocalPolicyMonitor[R any](config engine.Config, out func(R)) (PolicyMonitor[R], error) {
	lpm := &LocalPolicyMonitor[R]{config: config, interChan: make(chan *engine.PolicyInterpreter[R], 10), started: false,
		done: make(chan bool), policies: make(map[string][]byte), out: out}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Error.Printf("Unable to create policy watcher object %v", err)
		return nil, err
	}
	lpm.watcher = watcher
	err = lpm.CheckForPolicyUpdate()
	if err != nil {
		return nil, err
	}
	return lpm, err
}

// GetInterpreterChan returns a channel of the policy engine after they have been built.
// This channel can be checked for policy engines that are ready to be used.
func (p *LocalPolicyMonitor[R]) GetInterpreterChan() chan *engine.PolicyInterpreter[R] {
	return p.interChan
}

func (p *LocalPolicyMonitor[R]) dequeueFileEvents() int {
	count := 0
	i := 0
	for i < 1000 {
		select {
		case ev := <-p.watcher.Events:
			logger.Trace.Printf("Queued Event %#v, Operation: %s\n", ev, ev.Op.String())
			if hasModifiedYaml(ev) {
				count++
			}
		default:
			time.Sleep(10 * time.Millisecond)
			i++
		}
	}
	return count
}

func hasModifiedYaml(event fsnotify.Event) bool {
	result := false
	if (event.Op == fsnotify.Create || event.Op == fsnotify.Remove ||
		event.Op == fsnotify.Write || event.Op == fsnotify.Rename) && (strings.HasSuffix(event.Name, ".yaml") ||
		strings.HasSuffix(event.Name, ".yml")) {
		result = true
	}
	return result
}

func checksum(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		logger.Error.Printf("Unable to open file %s for checksum, %v", path, err)
		return nil, err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		logger.Error.Printf("Unable to calculate sha256 for file %s, %v", path, err)
		return nil, err
	}
	return h.Sum(nil), nil
}

func (p *LocalPolicyMonitor[R]) calculateChecksum() (bool, []string, error) {
	paths, err := ioutils.ListFilePaths(p.config.PoliciesPath, ".yaml")
	if err != nil {
		return false, nil, err
	}
	if len(paths) == 0 {
		p.policies = make(map[string][]byte)
		return false, make([]string, 0), errors.New("no policy files with extension .yaml found in policy directory: " + p.config.PoliciesPath)
	}
	newPolicies := make(map[string][]byte)
	changes := false
	for _, policy := range paths {
		cs, err := checksum(policy)
		if err != nil {
			p.policies = make(map[string][]byte)
			return false, nil, err
		}
		if val, ok := p.policies[policy]; ok {
			if !bytes.Equal(val, cs) {
				changes = true
			}
		} else {
			changes = true
		}
		newPolicies[policy] = cs
	}

	if len(p.policies) != len(newPolicies) {
		changes = true
	}
	p.policies = newPolicies
	return changes, paths, nil
}

// StartMonitor starts a thread to monitor the local policy directory.
func (p *LocalPolicyMonitor[R]) StartMonitor() error {
	if p.started {
		return nil
	}
	go func() {
		for {
			yamlCount := 0
			select {
			case <-p.done:
				logger.Trace.Printf("Policy monitor received done event... exiting...")
				return
			// watch for events
			case event := <-p.watcher.Events:
				logger.Trace.Printf("Event: %#v, Operation: %s\n", event, event.Op.String())
				if hasModifiedYaml(event) {
					yamlCount++
				}
				yamlCount += p.dequeueFileEvents()
				logger.Trace.Printf("Received %d more file events.\n", yamlCount)
				if yamlCount > 0 {
					changes, policyFiles, err := p.calculateChecksum()
					if err != nil {
						if policyFiles != nil && len(policyFiles) == 0 {
							logger.Error.Printf("There are no policy files in the policy path %s. Waiting for policies to be added.", p.config.PoliciesPath)
							continue
						} else {
							logger.Error.Printf("Unable to calculate checksums on policies.. attempting to compile policies")
						}
					}
					if changes || err != nil {
						logger.Info.Println("Checking for policy update")
						p.CheckForPolicyUpdate() //nolint:errcheck
					}
				}
			// watch for errors
			case err := <-p.watcher.Errors:
				logger.Error.Printf("Error while watching policy directory %s, %v", p.config.PoliciesPath, err)
			}
		}
	}()
	p.started = true
	if err := p.watcher.Add(p.config.PoliciesPath); err != nil {
		logger.Error.Printf("Unable to add watch to directory %s, %v", p.config.PoliciesPath, err)
		return err
	}
	return nil
}

// StopMonitor sends a signal to exit the monitor thread.
func (p *LocalPolicyMonitor[R]) StopMonitor() error {
	p.started = false
	p.done <- true
	return nil
}

// CheckForPolicyUpdate creates a new policy engine based on updated policies.
func (p *LocalPolicyMonitor[R]) CheckForPolicyUpdate() error {
	paths, err := ioutils.ListFilePaths(p.config.PoliciesPath, ".yaml")
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		return errors.New("no policy files with extension .yaml found in policy directory: " + p.config.PoliciesPath)
	}
	logger.Info.Println("Creating new policy interpreter")
	pi := engine.NewPolicyInterpreter[R](p.config, p.out)
	logger.Info.Println("Attempting to compile new policy")
	err = pi.Compile(paths...)
	if err != nil {
		logger.Error.Printf("Unable to compile policy files in directory %s. Not using new policy files. %v", p.config.PoliciesPath, err)
		return err
	}
	select {
	case p.interChan <- pi:
		logger.Info.Printf("Pushed new policy interpreter on channel")
	default:
		logger.Error.Printf("Unable to push new policy interpreter to policy thread.")
	}

	return nil
}

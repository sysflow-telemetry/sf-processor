//go:build flatrecord
// +build flatrecord

//
// Copyright (C) 2023 IBM Corporation.
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

// Package source implements a backend for policy compilers.
package common

import (
	"github.com/sysflow-telemetry/sf-processor/core/flattener"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source/flatrecord"
)

// Record is the record associated with the data source (backend) that will
// be used with the rules compiler.  The policy engine is defined
// with a "common.Record" type.  We use type aliasing to swap in
// different data sources currently. We will support multipe sources
// in the future. All types defined here are specific to the SysFlow flatrecord backend.
type Record = flatrecord.Record

// Channel is the backend specific channel wrapper object used to send data
// to the policy engine
type Channel = flattener.FlatChannel

// NewOperations specifies a constructor for the backend specific operations
// object used with the policy engine
var NewOperations = flatrecord.NewOperations

// NewPrefilter specifies a constructor for the backend specific prefilter
// object used with the policy engine
var NewPrefilter = flatrecord.NewPrefilter

// NewContextualizer specifies a constructor for the backend specific contextualizer
// object used with the policy engine
var NewContextualizer = flatrecord.NewContextualizer

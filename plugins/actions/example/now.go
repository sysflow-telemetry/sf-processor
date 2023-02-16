//
// Copyright (C) 2021 IBM Corporation.
//
// Authors:
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
//
package main

import (
	"strconv"
	"time"

	"github.com/sysflow-telemetry/sf-processor/core/policyengine/engine"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source/flatrecord"
)

type MyAction struct{}

func (a *MyAction) GetName() string {
	return "now"
}

func (a *MyAction) GetFunc() engine.ActionFunc[*flatrecord.Record] {
	return addMyTag
}

func addMyTag(r *flatrecord.Record) error {
	r.Ctx.AddTags("now_in_nanos:" + strconv.FormatInt(time.Now().UnixNano(), 10))
	return nil
}

var Action MyAction

// This function is not run when module is used as a plugin.
func main() {}

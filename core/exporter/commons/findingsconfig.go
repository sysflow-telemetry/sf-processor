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
//
package commons

import "strconv"

// Configuration keys.
const (
	FindingsApiKeyConfigKey       string = "findings.apikey"
	FindingsUrlConfigKey          string = "findings.url"
	FindingsAccountIDConfigKey    string = "findings.accountid"
	FindingsProviderIDConfigKey   string = "findings.provider"
	FindingsSqlQueryUrlConfigKey  string = "findings.sqlqueryurl"
	FindingsSqlQueryCrnConfigKey  string = "findings.sqlquerycrn"
	FindingsRegionConfigKey       string = "findings.region"
	FindingsPathConfigKey         string = "findings.path"
	FindingsPoolCapacityConfigKey string = "findings.pool.capacity"
	FindingsPoolMaxAgeConfigKey   string = "findings.pool.maxage"
)

// FindingsConfig holds IBM Findings API specific configuration.
type FindingsConfig struct {
	FindingsApiKey       string
	FindingsUrl          string
	FindingsAccountID    string
	FindingsProviderID   string
	FindingsSqlQueryUrl  string
	FindingsSqlQueryCrn  string
	FindingsRegion       string
	FindingsPath         string
	FindingsPoolCapacity int
	FindingsPoolMaxAge   int
}

// CreateFindingsConfig creates a new config object from config dictionary.
func CreateFindingsConfig(bc Config, conf map[string]interface{}) (c FindingsConfig, err error) {
	// default values
	c = FindingsConfig{
		FindingsUrl:          "https://us-south.secadvisor.cloud.ibm.com/findings",
		FindingsSqlQueryUrl:  "https://us.sql-query.cloud.ibm.com/sqlquery",
		FindingsPath:         "/mnt/occurrences",
		FindingsPoolCapacity: 250,
		FindingsPoolMaxAge:   1440} // 24 hours (specified in minutes)

	// parse config map
	if v, ok := conf[FindingsApiKeyConfigKey].(string); ok {
		c.FindingsApiKey = v
	} else if bc.VaultEnabled {
		s, err := bc.secrets.GetDecoded(FindingsApiKeyConfigKey)
		if err != nil {
			return c, err
		}
		c.FindingsApiKey = string(s)
	}
	if v, ok := conf[FindingsAccountIDConfigKey].(string); ok {
		c.FindingsAccountID = v
	}
	if v, ok := conf[FindingsUrlConfigKey].(string); ok {
		c.FindingsUrl = v
	}
	if v, ok := conf[FindingsProviderIDConfigKey].(string); ok {
		c.FindingsProviderID = v
	}
	if v, ok := conf[FindingsSqlQueryUrlConfigKey].(string); ok {
		c.FindingsSqlQueryUrl = v
	}
	if v, ok := conf[FindingsSqlQueryCrnConfigKey].(string); ok {
		c.FindingsSqlQueryCrn = v
	}
	if v, ok := conf[FindingsRegionConfigKey].(string); ok {
		c.FindingsRegion = v
	}
	if v, ok := conf[FindingsPathConfigKey].(string); ok {
		c.FindingsPath = v
	}
	if v, ok := conf[FindingsPoolCapacityConfigKey].(string); ok {
		c.FindingsPoolCapacity, err = strconv.Atoi(v)
		if err != nil {
			return c, err
		}
	}
	if v, ok := conf[FindingsPoolMaxAgeConfigKey].(string); ok {
		c.FindingsPoolMaxAge, err = strconv.Atoi(v)
		if err != nil {
			return c, err
		}
	}
	return
}

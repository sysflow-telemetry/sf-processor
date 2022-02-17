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

// Package commons defines common facilities for exporters.
package commons

import (
	"strconv"

	"github.com/IBM/scc-go-sdk/v3/findingsv1"
)

// Configuration keys.
const (
	FindingsAPIKeyConfigKey       string = "findings.apikey"
	FindingsURLConfigKey          string = "findings.url"
	FindingsAccountIDConfigKey    string = "findings.accountid"
	FindingsProviderIDConfigKey   string = "findings.provider"
	FindingsRegionConfigKey       string = "findings.region"
	FindingsSQLQueryURLConfigKey  string = "findings.sqlqueryurl"
	FindingsSQLQueryCrnConfigKey  string = "findings.sqlquerycrn"
	FindingsS3RegionConfigKey     string = "findings.s3region"
	FindingsS3BucketConfigKey     string = "findings.s3bucket"
	FindingsS3PrefixConfigKey     string = "findings.s3prefix"
	FindingsPathConfigKey         string = "findings.path"
	FindingsPoolCapacityConfigKey string = "findings.pool.capacity"
	FindingsPoolMaxAgeConfigKey   string = "findings.pool.maxage"
)

// FindingsConfig holds IBM Findings API specific configuration.
type FindingsConfig struct {
	FindingsAPIKey       string
	FindingsURL          string
	FindingsAccountID    string
	FindingsProviderID   string
	FindingsSQLQueryURL  string
	FindingsSQLQueryCrn  string
	FindingsRegion       string
	FindingsS3Region     string
	FindingsS3Bucket     string
	FindingsS3Prefix     string
	FindingsPath         string
	FindingsPoolCapacity int
	FindingsPoolMaxAge   int
}

// CreateFindingsConfig creates a new config object from config dictionary.
func CreateFindingsConfig(bc Config, conf map[string]interface{}) (c FindingsConfig, err error) {
	// default values
	c = FindingsConfig{
		FindingsURL:          findingsv1.DefaultServiceURL,
		FindingsSQLQueryURL:  "https://us.sql-query.cloud.ibm.com/sqlquery",
		FindingsPath:         "/mnt/occurrences",
		FindingsPoolCapacity: 250,
		FindingsPoolMaxAge:   1440} // 24 hours (specified in minutes)

	// parse config map
	if v, ok := conf[FindingsAPIKeyConfigKey].(string); ok {
		c.FindingsAPIKey = v
	} else if bc.VaultEnabled && bc.Transport == FindingsTransport {
		s, err := bc.GetSecret(FindingsAPIKeyConfigKey)
		if err != nil {
			return c, err
		}
		c.FindingsAPIKey = string(s)
	}
	if v, ok := conf[FindingsAccountIDConfigKey].(string); ok {
		c.FindingsAccountID = v
	}
	if v, ok := conf[FindingsURLConfigKey].(string); ok {
		c.FindingsURL = v
	}
	if v, ok := conf[FindingsProviderIDConfigKey].(string); ok {
		c.FindingsProviderID = v
	}
	if v, ok := conf[FindingsSQLQueryURLConfigKey].(string); ok {
		c.FindingsSQLQueryURL = v
	}
	if v, ok := conf[FindingsSQLQueryCrnConfigKey].(string); ok {
		c.FindingsSQLQueryCrn = v
	}
	if v, ok := conf[FindingsRegionConfigKey].(string); ok {
		c.FindingsRegion = v
	}
	if v, ok := conf[FindingsS3RegionConfigKey].(string); ok {
		c.FindingsS3Region = v
	}
	if v, ok := conf[FindingsS3PrefixConfigKey].(string); ok {
		c.FindingsS3Prefix = v
	}
	if v, ok := conf[FindingsS3BucketConfigKey].(string); ok {
		c.FindingsS3Bucket = v
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

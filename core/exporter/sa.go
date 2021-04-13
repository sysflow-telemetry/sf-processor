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
package exporter

import (
	"fmt"

	"github.com/IBM/go-sdk-core/v3/core"
	cqueue "github.com/enriquebris/goconcurrentqueue"
	"github.com/ibm-cloud-security/security-advisor-sdk-go/findingsapiv1"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/findings"
)

const (
	details             = "Occurrence Details"
	queryUrlFmt         = "%s/?instance_crn=%s&statement=%s"
	exportQueueCapacity = 2
)

// Severity type for enumeration.
type Severity int

// Severity enumeration.
const (
	SeverityLow Severity = iota
	SMedium
	SHigh
)

// String returns the string representation of a severity instance.
func (s Severity) String() string {
	return [...]string{"LOW", "MEDIUM", "HIGH"}[s]
}

// Certainty type for enumeration.
type Certainty int

// Certainty enumeration.
const (
	CertaintyLow Certainty = iota
	CertaintyMedium
	CertaintyHigh
)

// String returns the string representation of a severity instance.
func (s Certainty) String() string {
	return [...]string{"LOW", "MEDIUM", "HIGH"}[s]
}

// Occurrence object.
type Occurrence struct {
	ID         string
	ShortDescr string
	LongDescr  string
	Severity   Severity
	Certainty  Certainty
	ResType    string
	ResName    string
	AlertQuery string
	NoteID     string
}

// SAClient implements a custom client for IBM Cloud Security and Compliance Insights.
type SAClient struct {
	exportQueue *cqueue.FIFO
	AccountID   string
	ProviderID  string
	ApiKey      string
	SAUrl       string
	SqlQueryUrl string
	SqlQueryCrn string
	Region      string
}

// NewSAClient is a constructor for SAClient.
func NewSAClient(config Config) *SAClient {
	queue := cqueue.NewFIFO()
	queue.Enqueue(cmap.New())
	return &SAClient{AccountID: config.SAAccountID,
		exportQueue: queue,
		ProviderID:  config.SAProviderID,
		ApiKey:      config.SAApiKey,
		SAUrl:       config.SAUrl,
		SqlQueryUrl: config.SASqlQueryUrl,
		SqlQueryCrn: config.SASqlQueryCrn,
		Region:      config.Region}
}

// AddAlert adds alert to export queue.
func (s *SAClient) AddAlert(event Event) {
	// if r, ok := event.(TelemetryRecord); ok {
	// 	r.Container[]
	// 	s.exportQueue.Get(0)[]
	// }

}

//CreateFindingOccurrence creates a new occurrence of type finding.
func (s *SAClient) CreateOccurrence(occ *Occurrence) error {
	service, err := findings.NewFindingsApi(s.ApiKey, s.SAUrl)
	if err != nil {
		logger.Error.Printf("Error while creating Findings API wrapper %v", err)
		return err
	}

	noteName := fmt.Sprintf("%s/providers/%s/notes/%s", s.AccountID, s.ProviderID, occ.NoteID)
	var nextStep []findingsapiv1.RemediationStep
	if occ.AlertQuery != "" {
		nextStep = []findingsapiv1.RemediationStep{{
			Title: core.StringPtr(details),
			URL:   core.StringPtr(fmt.Sprintf(queryUrlFmt, s.SqlQueryUrl, s.SqlQueryCrn, occ.AlertQuery))},
		}
	}
	finding := findingsapiv1.Finding{Severity: core.StringPtr(occ.Severity.String()), Certainty: core.StringPtr(occ.Certainty.String()), NextSteps: nextStep}
	context := findingsapiv1.Context{Region: core.StringPtr(s.Region), ResourceType: core.StringPtr(occ.ResType), ResourceName: core.StringPtr(occ.ResName)}

	var options = service.NewCreateCustomOccurrenceOptions(s.AccountID, s.ProviderID, noteName, occ.ID)
	options.SetFinding(&finding)
	options.SetContext(&context)
	options.SetLongDescription(occ.LongDescr)
	options.SetShortDescription(occ.ShortDescr)

	result, response, err := service.CreateCustomOccurrence(options)
	if err != nil {
		logger.Error.Println("Failed to create occurrence: ", err)
		logger.Error.Println(response.Result)
		return err
	}
	logger.Info.Println(response.StatusCode)
	logger.Info.Println(*result.ID)
	return nil
}

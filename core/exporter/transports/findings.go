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

// Package transports implements transports for telemetry data.
package transports

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v3/common"
	"github.com/IBM/scc-go-sdk/v3/findingsv1"
	"github.com/pkg/errors"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/commons"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/encoders"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/utils"
)

const (
	details     = "Finding Context"
	queryURLFmt = "%s/?instance_crn=%s&statement=%s"
)

// FindingsAPIProto implements a custom client for IBM Cloud Security and Compliance Insights.
type FindingsAPIProto struct {
	AccountID   string
	ProviderID  string
	APIKey      string
	FindingsURL string
	SQLQueryURL string
	SQLQueryCrn string
	Region      string
}

// NewFindingsAPIProto is a constructor for FindingsAPIProto.
func NewFindingsAPIProto(conf commons.Config) TransportProtocol {
	return &FindingsAPIProto{AccountID: conf.FindingsAccountID,
		ProviderID:  conf.FindingsProviderID,
		APIKey:      conf.FindingsAPIKey,
		FindingsURL: conf.FindingsURL,
		SQLQueryURL: conf.FindingsSQLQueryURL,
		SQLQueryCrn: conf.FindingsSQLQueryCrn,
		Region:      conf.FindingsRegion}
}

// Init intializes a new protocol object.
func (s *FindingsAPIProto) Init() error {
	return nil
}

// Test tests the transport protocol.
func (s *FindingsAPIProto) Test() (bool, error) {
	service, err := NewFindingsAPI(s.AccountID, s.APIKey, s.FindingsURL)
	if err != nil {
		return false, errors.Wrap(err, "failed to instantiate Findings API")
	}
	return service.CheckAPIConfiguration(s.ProviderID)
}

// Export does nothing.
func (s *FindingsAPIProto) Export(data []commons.EncodedData) (err error) {
	for _, d := range data {
		if occ, ok := d.(*encoders.Occurrence); ok {
			if err = s.CreateOccurrence(occ); err != nil {
				return
			}
		} else {
			return errors.New("Expected Occurrence object as exported data")
		}
	}
	return
}

// Register registers the protocol object with the exporter.
func (s *FindingsAPIProto) Register(eps map[commons.Transport]TransportProtocolFactory) {
	eps[commons.FindingsTransport] = NewFindingsAPIProto
}

// Cleanup cleans up the protocol object.
func (s *FindingsAPIProto) Cleanup() {}

// CreateOccurrence creates a new occurrence of type finding.
func (s *FindingsAPIProto) CreateOccurrence(occ *encoders.Occurrence) error {
	service, err := NewFindingsAPI(s.AccountID, s.APIKey, s.FindingsURL)
	if err != nil {
		return err
	}

	noteName := fmt.Sprintf("%s/providers/%s/notes/%s", s.AccountID, s.ProviderID, occ.NoteID())
	var nextStep []findingsv1.RemediationStep
	if occ.AlertQuery != "" {
		nextStep = []findingsv1.RemediationStep{{
			Title: core.StringPtr(details),
			URL:   core.StringPtr(fmt.Sprintf(queryURLFmt, s.SQLQueryURL, s.SQLQueryCrn, occ.AlertQuery))},
		}
	}
	finding := findingsv1.Finding{Severity: core.StringPtr(occ.Severity.String()), Certainty: core.StringPtr(occ.Certainty.String()), NextSteps: nextStep}
	context := findingsv1.Context{Region: core.StringPtr(s.Region), ResourceType: core.StringPtr(occ.ResType), ResourceName: core.StringPtr(occ.ResName)}

	var options = service.Service.NewCreateOccurrenceOptions(s.ProviderID, noteName, findingsv1.CreateOccurrenceOptionsKindFindingConst, occ.ID)
	options.SetFinding(&finding)
	options.SetContext(&context)
	options.SetLongDescription(occ.LongDescr)
	options.SetShortDescription(occ.ShortDescr)

	result, response, err := service.CreateOccurrence(options)
	if err != nil {
		if response != nil {
			logger.Error.Println(response.Result)
		}
		return errors.Wrap(err, "error while creating occurrence")
	}

	logger.Trace.Println(response.StatusCode)
	logger.Trace.Println(*result.ID)

	return nil
}

// FindingsAPI implements an API for IBM Findings.
type FindingsAPI struct {
	Service *findingsv1.FindingsV1
}

// NewFindingsAPI constructs an instance of FindingsAPI with passed in options.
func NewFindingsAPI(accountID string, apiKey string, url string) (service *FindingsAPI, err error) {
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
	}
	serviceOptions := &findingsv1.FindingsV1Options{
		URL:           findingsv1.DefaultServiceURL,
		Authenticator: authenticator,
		AccountID:     core.StringPtr(accountID),
	}
	var baseService *findingsv1.FindingsV1
	baseService, err = findingsv1.NewFindingsV1(serviceOptions)
	if err != nil {
		return service, errors.Wrap(err, "couldn't instantiate base service for Findings API")
	}
	service = &FindingsAPI{
		Service: baseService,
	}
	return
}

// CheckAPIConfiguration checks Findings API connectivity and access.
func (s *FindingsAPI) CheckAPIConfiguration(providerID string) (pass bool, err error) {
	listNotesOptions := s.Service.NewListNotesOptions(providerID)
	listNotesResult, listNotesResponse, err := s.Service.ListNotes(listNotesOptions)
	if err != nil {
		return false, errors.Wrap(err, "couldn't list notes using Findings API")
	}

	if listNotesResponse.StatusCode != 200 {
		return false, errors.Wrapf(err, "bad response code while checking Findings API: %d", listNotesResponse.StatusCode)
	}

	ids := utils.NewSet()
	for _, n := range listNotesResult.Notes {
		id, err := json.Marshal(n.ID)
		if err != nil {
			return false, errors.Wrap(err, "can't decode note ID")
		}
		ids.Add(string(id[1 : len(id)-1]))
	}
	req := encoders.NoteIDs()
	if !req.IsSubset(ids) {
		return false, errors.Errorf("Provider doesn't contain required note IDs: %v", req)
	}
	return true, nil
}

// CreateOccurrence : Create an occurrence
// Create an occurrence to denote the existence of a particular type of finding.
//
// An occurrence describes provider-specific details of a note and contains vulnerability details, remediation steps,
// and other general information.
// Extends https://github.com/IBM/scc-go-sdk/blob/41f22b39e9ceea47d5c0c0a5515d9eaf5fee23d0/v3/findingsv1/findings_v1.go#L806
// to include short and long descriptions.
func (findings *FindingsAPI) CreateOccurrence(createOccurrenceOptions *findingsv1.CreateOccurrenceOptions) (result *findingsv1.APIOccurrence, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createOccurrenceOptions, "createOccurrenceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createOccurrenceOptions, "createOccurrenceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id":  *findings.Service.AccountID,
		"provider_id": *createOccurrenceOptions.ProviderID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(context.Background())
	builder.EnableGzipCompression = findings.Service.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(findings.Service.Service.Options.URL, `/v1/{account_id}/providers/{provider_id}/occurrences`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createOccurrenceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("findings", "V1", "CreateOccurrence")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createOccurrenceOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createOccurrenceOptions.TransactionID))
	}
	if createOccurrenceOptions.ReplaceIfExists != nil {
		builder.AddHeader("Replace-If-Exists", fmt.Sprint(*createOccurrenceOptions.ReplaceIfExists))
	}

	body := make(map[string]interface{})
	if createOccurrenceOptions.NoteName != nil {
		body["note_name"] = createOccurrenceOptions.NoteName
	}
	if createOccurrenceOptions.Kind != nil {
		body["kind"] = createOccurrenceOptions.Kind
	}
	if createOccurrenceOptions.ID != nil {
		body["id"] = createOccurrenceOptions.ID
	}
	if createOccurrenceOptions.ResourceURL != nil {
		body["resource_url"] = createOccurrenceOptions.ResourceURL
	}
	if createOccurrenceOptions.Remediation != nil {
		body["remediation"] = createOccurrenceOptions.Remediation
	}
	if createOccurrenceOptions.CreateTime != nil {
		body["create_time"] = createOccurrenceOptions.CreateTime
	}
	if createOccurrenceOptions.UpdateTime != nil {
		body["update_time"] = createOccurrenceOptions.UpdateTime
	}
	if createOccurrenceOptions.Context != nil {
		body["context"] = createOccurrenceOptions.Context
	}
	if createOccurrenceOptions.Finding != nil {
		body["finding"] = createOccurrenceOptions.Finding
	}
	if createOccurrenceOptions.Kpi != nil {
		body["kpi"] = createOccurrenceOptions.Kpi
	}
	if createOccurrenceOptions.ReferenceData != nil {
		body["reference_data"] = createOccurrenceOptions.ReferenceData
	}
	if createOccurrenceOptions.LongDescription != nil {
		body["long_description"] = createOccurrenceOptions.LongDescription
	}
	if createOccurrenceOptions.ShortDescription != nil {
		body["short_description"] = createOccurrenceOptions.ShortDescription
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = findings.Service.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, findingsv1.UnmarshalAPIOccurrence)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

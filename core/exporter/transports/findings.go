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
package transports

import (
	"errors"
	"fmt"

	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	"github.com/ibm-cloud-security/security-advisor-sdk-go/common"
	"github.com/ibm-cloud-security/security-advisor-sdk-go/findingsapiv1"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/commons"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/encoders"
)

const (
	kind        = "FINDING"
	details     = "Occurrence Details"
	queryUrlFmt = "%s/?instance_crn=%s&statement=%s"
)

// FindingsApiProto implements a custom client for IBM Cloud Security and Compliance Insights.
type FindingsApiProto struct {
	AccountID   string
	ProviderID  string
	ApiKey      string
	FindingsUrl string
	SqlQueryUrl string
	SqlQueryCrn string
	Region      string
}

// NewFindingsApiProto is a constructor for FindingsApiProto.
func NewFindingsApiProto(conf commons.Config) TransportProtocol {
	return &FindingsApiProto{AccountID: conf.FindingsAccountID,
		ProviderID:  conf.FindingsProviderID,
		ApiKey:      conf.FindingsApiKey,
		FindingsUrl: conf.FindingsUrl,
		SqlQueryUrl: conf.FindingsSqlQueryUrl,
		SqlQueryCrn: conf.FindingsSqlQueryCrn,
		Region:      conf.FindingsRegion}
}

// Init intializes a new null protocol object.
func (s *FindingsApiProto) Init() error {
	return nil
}

// Export does nothing.
func (s *FindingsApiProto) Export(data commons.EncodedData) error {
	if occ, ok := data.(*encoders.Occurrence); ok {
		return s.CreateOccurrence(occ)
	}
	return errors.New("Expected Occurrence object as exported data")
}

// Register registers the null protocol object with the exporter.
func (s *FindingsApiProto) Register(eps map[commons.Transport]TransportProtocolFactory) {
	eps[commons.FindingsTransport] = NewFindingsApiProto
}

// Cleanup cleans up the null protocol object.
func (s *FindingsApiProto) Cleanup() {}

//CreateFindingOccurrence creates a new occurrence of type finding.
func (s *FindingsApiProto) CreateOccurrence(occ *encoders.Occurrence) error {
	service, err := NewFindingsApi(s.ApiKey, s.FindingsUrl)
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

// FindingsApi implements an API for IBM Findings.
type FindingsApi struct {
	Service *core.BaseService
}

// NewFindingsApi constructs an instance of FindingsApi with passed in options.
func NewFindingsApi(apiKey string, url string) (service *FindingsApi, err error) {
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
	}

	serviceOptions := &core.ServiceOptions{
		URL:           findingsapiv1.DefaultServiceURL,
		Authenticator: authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if url != "" {
		err = baseService.SetServiceURL(url)
		if err != nil {
			return
		}
	}

	service = &FindingsApi{
		Service: baseService,
	}

	return
}

// CreateCustomOccurrence creates a new `Occurrence`. Use this method to create `Occurrences` for a resource.
func (s *FindingsApi) CreateCustomOccurrence(createOccurrenceOptions *CreateCustomOccurrenceOptions) (result *ApiCustomOccurrence, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createOccurrenceOptions, "createOccurrenceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createOccurrenceOptions, "createOccurrenceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "providers", "occurrences"}
	pathParameters := []string{*createOccurrenceOptions.AccountID, *createOccurrenceOptions.ProviderID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(s.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createOccurrenceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("findings_api", "V1", "CreateOccurrence")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
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

	response, err = s.Service.Request(request, new(ApiCustomOccurrence))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*ApiCustomOccurrence)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response")
		}
	}

	return
}

// NewCreateCustomOccurrenceOptions instantiates CreateCustomOccurrenceOptions
func (s *FindingsApi) NewCreateCustomOccurrenceOptions(accountID string, providerID string, noteName string, ID string) *CreateCustomOccurrenceOptions {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	return &CreateCustomOccurrenceOptions{
		AccountID:  core.StringPtr(accountID),
		ProviderID: core.StringPtr(providerID),
		NoteName:   core.StringPtr(noteName),
		Kind:       core.StringPtr(kind),
		ID:         core.StringPtr(ID),
		Headers:    headers,
	}
}

// CreateCustomOccurrenceOptions is the CreateCustomOccurrence options.
type CreateCustomOccurrenceOptions struct {

	// Account ID.
	AccountID *string `json:"account_id" validate:"required"`

	// Part of `parent`. This contains the provider_id for example: providers/{provider_id}.
	ProviderID *string `json:"provider_id" validate:"required"`

	// An analysis note associated with this image, in the form "{account_id}/providers/{provider_id}/notes/{note_id}" This
	// field can be used as a filter in list requests.
	NoteName *string `json:"note_name" validate:"required"`

	// A one sentence description of this `Note`.
	ShortDescription *string `json:"short_description" validate:"required"`

	// A detailed description of this `Note`.
	LongDescription *string `json:"long_description" validate:"required"`

	// Output only. This explicitly denotes which of the `Occurrence` details are specified.
	// This field can be used as a filter in list requests.
	Kind *string `json:"kind" validate:"required"`

	ID *string `json:"id" validate:"required"`

	// The unique URL of the resource, image or the container, for which the `Occurrence` applies. For example,
	// https://gcr.io/provider/image@sha256:foo. This field can be used as a filter in list requests.
	ResourceURL *string `json:"resource_url,omitempty"`

	Remediation *string `json:"remediation,omitempty"`

	// Output only. The time this `Occurrence` was created.
	CreateTime *strfmt.DateTime `json:"create_time,omitempty"`

	// Output only. The time this `Occurrence` was last updated.
	UpdateTime *strfmt.DateTime `json:"update_time,omitempty"`

	// Details about the context of this `Occurrence`.
	Context *findingsapiv1.Context `json:"context,omitempty"`

	// Details of the occurrence of a finding.
	Finding *findingsapiv1.Finding `json:"finding,omitempty"`

	// Details of the occurrence of a KPI.
	Kpi *findingsapiv1.Kpi `json:"kpi,omitempty"`

	// It allows replacing an existing occurrence when set to true.
	ReplaceIfExists *bool `json:"Replace-If-Exists,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// SetContext allows user to set Context
func (options *CreateCustomOccurrenceOptions) SetContext(context *findingsapiv1.Context) *CreateCustomOccurrenceOptions {
	options.Context = context
	return options
}

// SetFinding allows user to set Finding
func (options *CreateCustomOccurrenceOptions) SetFinding(finding *findingsapiv1.Finding) *CreateCustomOccurrenceOptions {
	options.Finding = finding
	return options
}

// SetShortDescription allows user to set ShortDescription
func (options *CreateCustomOccurrenceOptions) SetShortDescription(shortDescription string) *CreateCustomOccurrenceOptions {
	options.ShortDescription = core.StringPtr(shortDescription)
	return options
}

// SetLongDescription allows user to set LongDescription
func (options *CreateCustomOccurrenceOptions) SetLongDescription(longDescription string) *CreateCustomOccurrenceOptions {
	options.LongDescription = core.StringPtr(longDescription)
	return options
}

// ApiCustomOccurrence includes information about analysis occurrences.
type ApiCustomOccurrence struct {

	// The unique URL of the resource, image or the container, for which the `Occurrence` applies. For example,
	// https://gcr.io/provider/image@sha256:foo. This field can be used as a filter in list requests.
	ResourceURL *string `json:"resource_url,omitempty"`

	// An analysis note associated with this image, in the form "{account_id}/providers/{provider_id}/notes/{note_id}" This
	// field can be used as a filter in list requests.
	NoteName *string `json:"note_name" validate:"required"`

	// A one sentence description of this `Note`.
	ShortDescription *string `json:"short_description" validate:"required"`

	// A detailed description of this `Note`.
	LongDescription *string `json:"long_description" validate:"required"`

	// Output only. This explicitly denotes which of the `Occurrence` details are specified.
	// This field can be used as a filter in list requests.
	Kind *string `json:"kind" validate:"required"`

	Remediation *string `json:"remediation,omitempty"`

	// Output only. The time this `Occurrence` was created.
	CreateTime *strfmt.DateTime `json:"create_time,omitempty"`

	// Output only. The time this `Occurrence` was last updated.
	UpdateTime *strfmt.DateTime `json:"update_time,omitempty"`

	ID *string `json:"id" validate:"required"`

	//OccurrenceID of the occurrence
	OccurrenceID *string `json:"occurrence_id,omitempty"`

	//ProviderID of the occurrence
	ProviderID *string `json:"provider_id,omitempty"`

	//Name of the occurrence
	Name *string `json:"name,omitempty"`

	// Details about the context of this `Occurrence`.
	Context *findingsapiv1.Context `json:"context,omitempty"`

	// Details of the occurrence of a finding.
	Finding *findingsapiv1.Finding `json:"finding,omitempty"`

	// Details of the occurrence of a KPI.
	Kpi *findingsapiv1.Kpi `json:"kpi,omitempty"`
}

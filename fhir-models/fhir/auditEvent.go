// Copyright 2019 The Samply Development Community
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

package fhir

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// AuditEvent is documented here http://hl7.org/fhir/StructureDefinition/AuditEvent
type AuditEvent struct {
	Id                *string            `json:"id,omitempty"`
	Meta              *Meta              `json:"meta,omitempty"`
	ImplicitRules     *string            `json:"implicitRules,omitempty"`
	Language          *string            `json:"language,omitempty"`
	Text              *Narrative         `json:"text,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Type              Coding             `json:"type"`
	Subtype           []Coding           `json:"subtype,omitempty"`
	Action            *AuditEventAction  `json:"action,omitempty"`
	Period            *Period            `json:"period,omitempty"`
	Recorded          string             `json:"recorded"`
	Outcome           *AuditEventOutcome `json:"outcome,omitempty"`
	OutcomeDesc       *string            `json:"outcomeDesc,omitempty"`
	PurposeOfEvent    []CodeableConcept  `json:"purposeOfEvent,omitempty"`
	Agent             []AuditEventAgent  `json:"agent"`
	Source            AuditEventSource   `json:"source"`
	Entity            []AuditEventEntity `json:"entity,omitempty"`
}
type AuditEventAgent struct {
	Id                *string                 `json:"id,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept        `json:"type,omitempty"`
	Role              []CodeableConcept       `json:"role,omitempty"`
	Who               *Reference              `json:"who,omitempty"`
	AltId             *string                 `json:"altId,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	Requestor         bool                    `json:"requestor"`
	Location          *Reference              `json:"location,omitempty"`
	Policy            []string                `json:"policy,omitempty"`
	Media             *Coding                 `json:"media,omitempty"`
	Network           *AuditEventAgentNetwork `json:"network,omitempty"`
	PurposeOfUse      []CodeableConcept       `json:"purposeOfUse,omitempty"`
}
type AuditEventAgentNetwork struct {
	Id                *string                     `json:"id,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Address           *string                     `json:"address,omitempty"`
	Type              *AuditEventAgentNetworkType `json:"type,omitempty"`
}
type AuditEventSource struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Site              *string     `json:"site,omitempty"`
	Observer          Reference   `json:"observer"`
	Type              []Coding    `json:"type,omitempty"`
}
type AuditEventEntity struct {
	Id                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	What              *Reference               `json:"what,omitempty"`
	Type              *Coding                  `json:"type,omitempty"`
	Role              *Coding                  `json:"role,omitempty"`
	Lifecycle         *Coding                  `json:"lifecycle,omitempty"`
	SecurityLabel     []Coding                 `json:"securityLabel,omitempty"`
	Name              *string                  `json:"name,omitempty"`
	Description       *string                  `json:"description,omitempty"`
	Query             *string                  `json:"query,omitempty"`
	Detail            []AuditEventEntityDetail `json:"detail,omitempty"`
}
type AuditEventEntityDetail struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
}
type OtherAuditEvent AuditEvent

// MarshalJSON marshals the given AuditEvent as JSON into a byte slice
func (r AuditEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAuditEvent
		ResourceType string `json:"resourceType"`
	}{
		OtherAuditEvent: OtherAuditEvent(r),
		ResourceType:    "AuditEvent",
	})
}

// UnmarshalAuditEvent unmarshals a AuditEvent.
func UnmarshalAuditEvent(b []byte) (AuditEvent, error) {
	var auditEvent AuditEvent
	if err := json.Unmarshal(b, &auditEvent); err != nil {
		return auditEvent, err
	}
	return auditEvent, nil
}

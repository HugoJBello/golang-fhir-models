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

// ActivityDefinition is documented here http://hl7.org/fhir/StructureDefinition/ActivityDefinition
type ActivityDefinition struct {
	Id                           *string                          `json:"id,omitempty"`
	Meta                         *Meta                            `json:"meta,omitempty"`
	ImplicitRules                *string                          `json:"implicitRules,omitempty"`
	Language                     *string                          `json:"language,omitempty"`
	Text                         *Narrative                       `json:"text,omitempty"`
	Extension                    []Extension                      `json:"extension,omitempty"`
	ModifierExtension            []Extension                      `json:"modifierExtension,omitempty"`
	Url                          *string                          `json:"url,omitempty"`
	Identifier                   []Identifier                     `json:"identifier,omitempty"`
	Version                      *string                          `json:"version,omitempty"`
	Name                         *string                          `json:"name,omitempty"`
	Title                        *string                          `json:"title,omitempty"`
	Subtitle                     *string                          `json:"subtitle,omitempty"`
	Status                       PublicationStatus                `json:"status"`
	Experimental                 *bool                            `json:"experimental,omitempty"`
	Date                         *string                          `json:"date,omitempty"`
	Publisher                    *string                          `json:"publisher,omitempty"`
	Contact                      []ContactDetail                  `json:"contact,omitempty"`
	Description                  *string                          `json:"description,omitempty"`
	UseContext                   []UsageContext                   `json:"useContext,omitempty"`
	Jurisdiction                 []CodeableConcept                `json:"jurisdiction,omitempty"`
	Purpose                      *string                          `json:"purpose,omitempty"`
	Usage                        *string                          `json:"usage,omitempty"`
	Copyright                    *string                          `json:"copyright,omitempty"`
	ApprovalDate                 *string                          `json:"approvalDate,omitempty"`
	LastReviewDate               *string                          `json:"lastReviewDate,omitempty"`
	EffectivePeriod              *Period                          `json:"effectivePeriod,omitempty"`
	Topic                        []CodeableConcept                `json:"topic,omitempty"`
	Author                       []ContactDetail                  `json:"author,omitempty"`
	Editor                       []ContactDetail                  `json:"editor,omitempty"`
	Reviewer                     []ContactDetail                  `json:"reviewer,omitempty"`
	Endorser                     []ContactDetail                  `json:"endorser,omitempty"`
	RelatedArtifact              []RelatedArtifact                `json:"relatedArtifact,omitempty"`
	Library                      []string                         `json:"library,omitempty"`
	Kind                         *RequestResourceType             `json:"kind,omitempty"`
	Profile                      *string                          `json:"profile,omitempty"`
	Code                         *CodeableConcept                 `json:"code,omitempty"`
	Intent                       *RequestIntent                   `json:"intent,omitempty"`
	Priority                     *RequestPriority                 `json:"priority,omitempty"`
	DoNotPerform                 *bool                            `json:"doNotPerform,omitempty"`
	Location                     *Reference                       `json:"location,omitempty"`
	Participant                  []ActivityDefinitionParticipant  `json:"participant,omitempty"`
	Quantity                     *Quantity                        `json:"quantity,omitempty"`
	Dosage                       []Dosage                         `json:"dosage,omitempty"`
	BodySite                     []CodeableConcept                `json:"bodySite,omitempty"`
	SpecimenRequirement          []Reference                      `json:"specimenRequirement,omitempty"`
	ObservationRequirement       []Reference                      `json:"observationRequirement,omitempty"`
	ObservationResultRequirement []Reference                      `json:"observationResultRequirement,omitempty"`
	Transform                    *string                          `json:"transform,omitempty"`
	DynamicValue                 []ActivityDefinitionDynamicValue `json:"dynamicValue,omitempty"`
}
type ActivityDefinitionParticipant struct {
	Id                *string               `json:"id,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Type              ActionParticipantType `json:"type"`
	Role              *CodeableConcept      `json:"role,omitempty"`
}
type ActivityDefinitionDynamicValue struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Path              string      `json:"path"`
	Expression        Expression  `json:"expression"`
}
type OtherActivityDefinition ActivityDefinition

// MarshalJSON marshals the given ActivityDefinition as JSON into a byte slice
func (r ActivityDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherActivityDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherActivityDefinition: OtherActivityDefinition(r),
		ResourceType:            "ActivityDefinition",
	})
}

// UnmarshalActivityDefinition unmarshals a ActivityDefinition.
func UnmarshalActivityDefinition(b []byte) (ActivityDefinition, error) {
	var activityDefinition ActivityDefinition
	if err := json.Unmarshal(b, &activityDefinition); err != nil {
		return activityDefinition, err
	}
	return activityDefinition, nil
}

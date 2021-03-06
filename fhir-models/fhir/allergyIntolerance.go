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

// AllergyIntolerance is documented here http://hl7.org/fhir/StructureDefinition/AllergyIntolerance
type AllergyIntolerance struct {
	Id                 *string                        `json:"id,omitempty"`
	Meta               *Meta                          `json:"meta,omitempty"`
	ImplicitRules      *string                        `json:"implicitRules,omitempty"`
	Language           *string                        `json:"language,omitempty"`
	Text               *Narrative                     `json:"text,omitempty"`
	Extension          []Extension                    `json:"extension,omitempty"`
	ModifierExtension  []Extension                    `json:"modifierExtension,omitempty"`
	Identifier         []Identifier                   `json:"identifier,omitempty"`
	ClinicalStatus     *CodeableConcept               `json:"clinicalStatus,omitempty"`
	VerificationStatus *CodeableConcept               `json:"verificationStatus,omitempty"`
	Type               *AllergyIntoleranceType        `json:"type,omitempty"`
	Category           []AllergyIntoleranceCategory   `json:"category,omitempty"`
	Criticality        *AllergyIntoleranceCriticality `json:"criticality,omitempty"`
	Code               *CodeableConcept               `json:"code,omitempty"`
	Patient            Reference                      `json:"patient"`
	Encounter          *Reference                     `json:"encounter,omitempty"`
	RecordedDate       *string                        `json:"recordedDate,omitempty"`
	Recorder           *Reference                     `json:"recorder,omitempty"`
	Asserter           *Reference                     `json:"asserter,omitempty"`
	LastOccurrence     *string                        `json:"lastOccurrence,omitempty"`
	Note               []Annotation                   `json:"note,omitempty"`
	Reaction           []AllergyIntoleranceReaction   `json:"reaction,omitempty"`
}
type AllergyIntoleranceReaction struct {
	Id                *string                     `json:"id,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Substance         *CodeableConcept            `json:"substance,omitempty"`
	Manifestation     []CodeableConcept           `json:"manifestation"`
	Description       *string                     `json:"description,omitempty"`
	Onset             *string                     `json:"onset,omitempty"`
	Severity          *AllergyIntoleranceSeverity `json:"severity,omitempty"`
	ExposureRoute     *CodeableConcept            `json:"exposureRoute,omitempty"`
	Note              []Annotation                `json:"note,omitempty"`
}
type OtherAllergyIntolerance AllergyIntolerance

// MarshalJSON marshals the given AllergyIntolerance as JSON into a byte slice
func (r AllergyIntolerance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAllergyIntolerance
		ResourceType string `json:"resourceType"`
	}{
		OtherAllergyIntolerance: OtherAllergyIntolerance(r),
		ResourceType:            "AllergyIntolerance",
	})
}

// UnmarshalAllergyIntolerance unmarshals a AllergyIntolerance.
func UnmarshalAllergyIntolerance(b []byte) (AllergyIntolerance, error) {
	var allergyIntolerance AllergyIntolerance
	if err := json.Unmarshal(b, &allergyIntolerance); err != nil {
		return allergyIntolerance, err
	}
	return allergyIntolerance, nil
}

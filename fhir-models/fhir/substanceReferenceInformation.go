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

// SubstanceReferenceInformation is documented here http://hl7.org/fhir/StructureDefinition/SubstanceReferenceInformation
type SubstanceReferenceInformation struct {
	Id                *string                                       `json:"id,omitempty"`
	Meta              *Meta                                         `json:"meta,omitempty"`
	ImplicitRules     *string                                       `json:"implicitRules,omitempty"`
	Language          *string                                       `json:"language,omitempty"`
	Text              *Narrative                                    `json:"text,omitempty"`
	Extension         []Extension                                   `json:"extension,omitempty"`
	ModifierExtension []Extension                                   `json:"modifierExtension,omitempty"`
	Comment           *string                                       `json:"comment,omitempty"`
	Gene              []SubstanceReferenceInformationGene           `json:"gene,omitempty"`
	GeneElement       []SubstanceReferenceInformationGeneElement    `json:"geneElement,omitempty"`
	Classification    []SubstanceReferenceInformationClassification `json:"classification,omitempty"`
	Target            []SubstanceReferenceInformationTarget         `json:"target,omitempty"`
}
type SubstanceReferenceInformationGene struct {
	Id                 *string          `json:"id,omitempty"`
	Extension          []Extension      `json:"extension,omitempty"`
	ModifierExtension  []Extension      `json:"modifierExtension,omitempty"`
	GeneSequenceOrigin *CodeableConcept `json:"geneSequenceOrigin,omitempty"`
	Gene               *CodeableConcept `json:"gene,omitempty"`
	Source             []Reference      `json:"source,omitempty"`
}
type SubstanceReferenceInformationGeneElement struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Element           *Identifier      `json:"element,omitempty"`
	Source            []Reference      `json:"source,omitempty"`
}
type SubstanceReferenceInformationClassification struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Domain            *CodeableConcept  `json:"domain,omitempty"`
	Classification    *CodeableConcept  `json:"classification,omitempty"`
	Subtype           []CodeableConcept `json:"subtype,omitempty"`
	Source            []Reference       `json:"source,omitempty"`
}
type SubstanceReferenceInformationTarget struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Target            *Identifier      `json:"target,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Interaction       *CodeableConcept `json:"interaction,omitempty"`
	Organism          *CodeableConcept `json:"organism,omitempty"`
	OrganismType      *CodeableConcept `json:"organismType,omitempty"`
	AmountType        *CodeableConcept `json:"amountType,omitempty"`
	Source            []Reference      `json:"source,omitempty"`
}
type OtherSubstanceReferenceInformation SubstanceReferenceInformation

// MarshalJSON marshals the given SubstanceReferenceInformation as JSON into a byte slice
func (r SubstanceReferenceInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceReferenceInformation
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceReferenceInformation: OtherSubstanceReferenceInformation(r),
		ResourceType:                       "SubstanceReferenceInformation",
	})
}

// UnmarshalSubstanceReferenceInformation unmarshals a SubstanceReferenceInformation.
func UnmarshalSubstanceReferenceInformation(b []byte) (SubstanceReferenceInformation, error) {
	var substanceReferenceInformation SubstanceReferenceInformation
	if err := json.Unmarshal(b, &substanceReferenceInformation); err != nil {
		return substanceReferenceInformation, err
	}
	return substanceReferenceInformation, nil
}

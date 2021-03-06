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

// Coverage is documented here http://hl7.org/fhir/StructureDefinition/Coverage
type Coverage struct {
	Id                *string                      `json:"id,omitempty"`
	Meta              *Meta                        `json:"meta,omitempty"`
	ImplicitRules     *string                      `json:"implicitRules,omitempty"`
	Language          *string                      `json:"language,omitempty"`
	Text              *Narrative                   `json:"text,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                 `json:"identifier,omitempty"`
	Status            FinancialResourceStatusCodes `json:"status"`
	Type              *CodeableConcept             `json:"type,omitempty"`
	PolicyHolder      *Reference                   `json:"policyHolder,omitempty"`
	Subscriber        *Reference                   `json:"subscriber,omitempty"`
	SubscriberId      *string                      `json:"subscriberId,omitempty"`
	Beneficiary       Reference                    `json:"beneficiary"`
	Dependent         *string                      `json:"dependent,omitempty"`
	Relationship      *CodeableConcept             `json:"relationship,omitempty"`
	Period            *Period                      `json:"period,omitempty"`
	Payor             []Reference                  `json:"payor"`
	Class             []CoverageClass              `json:"class,omitempty"`
	Order             *int                         `json:"order,omitempty"`
	Network           *string                      `json:"network,omitempty"`
	CostToBeneficiary []CoverageCostToBeneficiary  `json:"costToBeneficiary,omitempty"`
	Subrogation       *bool                        `json:"subrogation,omitempty"`
	Contract          []Reference                  `json:"contract,omitempty"`
}
type CoverageClass struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Value             string          `json:"value"`
	Name              *string         `json:"name,omitempty"`
}
type CoverageCostToBeneficiary struct {
	Id                *string                              `json:"id,omitempty"`
	Extension         []Extension                          `json:"extension,omitempty"`
	ModifierExtension []Extension                          `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept                     `json:"type,omitempty"`
	Exception         []CoverageCostToBeneficiaryException `json:"exception,omitempty"`
}
type CoverageCostToBeneficiaryException struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Period            *Period         `json:"period,omitempty"`
}
type OtherCoverage Coverage

// MarshalJSON marshals the given Coverage as JSON into a byte slice
func (r Coverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCoverage
		ResourceType string `json:"resourceType"`
	}{
		OtherCoverage: OtherCoverage(r),
		ResourceType:  "Coverage",
	})
}

// UnmarshalCoverage unmarshals a Coverage.
func UnmarshalCoverage(b []byte) (Coverage, error) {
	var coverage Coverage
	if err := json.Unmarshal(b, &coverage); err != nil {
		return coverage, err
	}
	return coverage, nil
}

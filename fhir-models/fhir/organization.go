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

// Organization is documented here http://hl7.org/fhir/StructureDefinition/Organization
type Organization struct {
	Id                *string               `json:"id,omitempty"`
	Meta              *Meta                 `json:"meta,omitempty"`
	ImplicitRules     *string               `json:"implicitRules,omitempty"`
	Language          *string               `json:"language,omitempty"`
	Text              *Narrative            `json:"text,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `json:"identifier,omitempty"`
	Active            *bool                 `json:"active,omitempty"`
	Type              []CodeableConcept     `json:"type,omitempty"`
	Name              *string               `json:"name,omitempty"`
	Alias             []string              `json:"alias,omitempty"`
	Telecom           []ContactPoint        `json:"telecom,omitempty"`
	Address           []Address             `json:"address,omitempty"`
	PartOf            *Reference            `json:"partOf,omitempty"`
	Contact           []OrganizationContact `json:"contact,omitempty"`
	Endpoint          []Reference           `json:"endpoint,omitempty"`
}
type OrganizationContact struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Purpose           *CodeableConcept `json:"purpose,omitempty"`
	Name              *HumanName       `json:"name,omitempty"`
	Telecom           []ContactPoint   `json:"telecom,omitempty"`
	Address           *Address         `json:"address,omitempty"`
}
type OtherOrganization Organization

// MarshalJSON marshals the given Organization as JSON into a byte slice
func (r Organization) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherOrganization
		ResourceType string `json:"resourceType"`
	}{
		OtherOrganization: OtherOrganization(r),
		ResourceType:      "Organization",
	})
}

// UnmarshalOrganization unmarshals a Organization.
func UnmarshalOrganization(b []byte) (Organization, error) {
	var organization Organization
	if err := json.Unmarshal(b, &organization); err != nil {
		return organization, err
	}
	return organization, nil
}

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

// Patient is documented here http://hl7.org/fhir/StructureDefinition/Patient
type Patient struct {
	Id                   *string                `json:"id,omitempty"`
	Meta                 *Meta                  `json:"meta,omitempty"`
	ImplicitRules        *string                `json:"implicitRules,omitempty"`
	Language             *string                `json:"language,omitempty"`
	Text                 *Narrative             `json:"text,omitempty"`
	Extension            []Extension            `json:"extension,omitempty"`
	ModifierExtension    []Extension            `json:"modifierExtension,omitempty"`
	Identifier           []Identifier           `json:"identifier,omitempty"`
	Active               *bool                  `json:"active,omitempty"`
	Name                 []HumanName            `json:"name,omitempty"`
	Telecom              []ContactPoint         `json:"telecom,omitempty"`
	Gender               *AdministrativeGender  `json:"gender,omitempty"`
	BirthDate            *string                `json:"birthDate,omitempty"`
	Address              []Address              `json:"address,omitempty"`
	MaritalStatus        *CodeableConcept       `json:"maritalStatus,omitempty"`
	Photo                []Attachment           `json:"photo,omitempty"`
	Contact              []PatientContact       `json:"contact,omitempty"`
	Communication        []PatientCommunication `json:"communication,omitempty"`
	GeneralPractitioner  []Reference            `json:"generalPractitioner,omitempty"`
	ManagingOrganization *Reference             `json:"managingOrganization,omitempty"`
	Link                 []PatientLink          `json:"link,omitempty"`
}
type PatientContact struct {
	Id                *string               `json:"id,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Relationship      []CodeableConcept     `json:"relationship,omitempty"`
	Name              *HumanName            `json:"name,omitempty"`
	Telecom           []ContactPoint        `json:"telecom,omitempty"`
	Address           *Address              `json:"address,omitempty"`
	Gender            *AdministrativeGender `json:"gender,omitempty"`
	Organization      *Reference            `json:"organization,omitempty"`
	Period            *Period               `json:"period,omitempty"`
}
type PatientCommunication struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Language          CodeableConcept `json:"language"`
	Preferred         *bool           `json:"preferred,omitempty"`
}
type PatientLink struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Other             Reference   `json:"other"`
	Type              LinkType    `json:"type"`
}
type OtherPatient Patient

// MarshalJSON marshals the given Patient as JSON into a byte slice
func (r Patient) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPatient
		ResourceType string `json:"resourceType"`
	}{
		OtherPatient: OtherPatient(r),
		ResourceType: "Patient",
	})
}

// UnmarshalPatient unmarshals a Patient.
func UnmarshalPatient(b []byte) (Patient, error) {
	var patient Patient
	if err := json.Unmarshal(b, &patient); err != nil {
		return patient, err
	}
	return patient, nil
}

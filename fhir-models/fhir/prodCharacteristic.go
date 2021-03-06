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

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// ProdCharacteristic is documented here http://hl7.org/fhir/StructureDefinition/ProdCharacteristic
type ProdCharacteristic struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Height            *Quantity        `json:"height,omitempty"`
	Width             *Quantity        `json:"width,omitempty"`
	Depth             *Quantity        `json:"depth,omitempty"`
	Weight            *Quantity        `json:"weight,omitempty"`
	NominalVolume     *Quantity        `json:"nominalVolume,omitempty"`
	ExternalDiameter  *Quantity        `json:"externalDiameter,omitempty"`
	Shape             *string          `json:"shape,omitempty"`
	Color             []string         `json:"color,omitempty"`
	Imprint           []string         `json:"imprint,omitempty"`
	Image             []Attachment     `json:"image,omitempty"`
	Scoring           *CodeableConcept `json:"scoring,omitempty"`
}

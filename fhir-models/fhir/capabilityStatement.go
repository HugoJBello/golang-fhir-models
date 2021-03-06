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

// CapabilityStatement is documented here http://hl7.org/fhir/StructureDefinition/CapabilityStatement
type CapabilityStatement struct {
	Id                  *string                            `json:"id,omitempty"`
	Meta                *Meta                              `json:"meta,omitempty"`
	ImplicitRules       *string                            `json:"implicitRules,omitempty"`
	Language            *string                            `json:"language,omitempty"`
	Text                *Narrative                         `json:"text,omitempty"`
	Extension           []Extension                        `json:"extension,omitempty"`
	ModifierExtension   []Extension                        `json:"modifierExtension,omitempty"`
	Url                 *string                            `json:"url,omitempty"`
	Version             *string                            `json:"version,omitempty"`
	Name                *string                            `json:"name,omitempty"`
	Title               *string                            `json:"title,omitempty"`
	Status              PublicationStatus                  `json:"status"`
	Experimental        *bool                              `json:"experimental,omitempty"`
	Date                string                             `json:"date"`
	Publisher           *string                            `json:"publisher,omitempty"`
	Contact             []ContactDetail                    `json:"contact,omitempty"`
	Description         *string                            `json:"description,omitempty"`
	UseContext          []UsageContext                     `json:"useContext,omitempty"`
	Jurisdiction        []CodeableConcept                  `json:"jurisdiction,omitempty"`
	Purpose             *string                            `json:"purpose,omitempty"`
	Copyright           *string                            `json:"copyright,omitempty"`
	Kind                CapabilityStatementKind            `json:"kind"`
	Instantiates        []string                           `json:"instantiates,omitempty"`
	Imports             []string                           `json:"imports,omitempty"`
	Software            *CapabilityStatementSoftware       `json:"software,omitempty"`
	Implementation      *CapabilityStatementImplementation `json:"implementation,omitempty"`
	FhirVersion         FHIRVersion                        `json:"fhirVersion"`
	Format              []string                           `json:"format"`
	PatchFormat         []string                           `json:"patchFormat,omitempty"`
	ImplementationGuide []string                           `json:"implementationGuide,omitempty"`
	Rest                []CapabilityStatementRest          `json:"rest,omitempty"`
	Messaging           []CapabilityStatementMessaging     `json:"messaging,omitempty"`
	Document            []CapabilityStatementDocument      `json:"document,omitempty"`
}
type CapabilityStatementSoftware struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Version           *string     `json:"version,omitempty"`
	ReleaseDate       *string     `json:"releaseDate,omitempty"`
}
type CapabilityStatementImplementation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       string      `json:"description"`
	Url               *string     `json:"url,omitempty"`
	Custodian         *Reference  `json:"custodian,omitempty"`
}
type CapabilityStatementRest struct {
	Id                *string                                      `json:"id,omitempty"`
	Extension         []Extension                                  `json:"extension,omitempty"`
	ModifierExtension []Extension                                  `json:"modifierExtension,omitempty"`
	Mode              RestfulCapabilityMode                        `json:"mode"`
	Documentation     *string                                      `json:"documentation,omitempty"`
	Security          *CapabilityStatementRestSecurity             `json:"security,omitempty"`
	Resource          []CapabilityStatementRestResource            `json:"resource,omitempty"`
	Interaction       []CapabilityStatementRestInteraction         `json:"interaction,omitempty"`
	SearchParam       []CapabilityStatementRestResourceSearchParam `json:"searchParam,omitempty"`
	Operation         []CapabilityStatementRestResourceOperation   `json:"operation,omitempty"`
	Compartment       []string                                     `json:"compartment,omitempty"`
}
type CapabilityStatementRestSecurity struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Cors              *bool             `json:"cors,omitempty"`
	Service           []CodeableConcept `json:"service,omitempty"`
	Description       *string           `json:"description,omitempty"`
}
type CapabilityStatementRestResource struct {
	Id                *string                                      `json:"id,omitempty"`
	Extension         []Extension                                  `json:"extension,omitempty"`
	ModifierExtension []Extension                                  `json:"modifierExtension,omitempty"`
	Type              ResourceType                                 `json:"type"`
	Profile           *string                                      `json:"profile,omitempty"`
	SupportedProfile  []string                                     `json:"supportedProfile,omitempty"`
	Documentation     *string                                      `json:"documentation,omitempty"`
	Interaction       []CapabilityStatementRestResourceInteraction `json:"interaction,omitempty"`
	Versioning        *ResourceVersionPolicy                       `json:"versioning,omitempty"`
	ReadHistory       *bool                                        `json:"readHistory,omitempty"`
	UpdateCreate      *bool                                        `json:"updateCreate,omitempty"`
	ConditionalCreate *bool                                        `json:"conditionalCreate,omitempty"`
	ConditionalRead   *ConditionalReadStatus                       `json:"conditionalRead,omitempty"`
	ConditionalUpdate *bool                                        `json:"conditionalUpdate,omitempty"`
	ConditionalDelete *ConditionalDeleteStatus                     `json:"conditionalDelete,omitempty"`
	ReferencePolicy   []ReferenceHandlingPolicy                    `json:"referencePolicy,omitempty"`
	SearchInclude     []string                                     `json:"searchInclude,omitempty"`
	SearchRevInclude  []string                                     `json:"searchRevInclude,omitempty"`
	SearchParam       []CapabilityStatementRestResourceSearchParam `json:"searchParam,omitempty"`
	Operation         []CapabilityStatementRestResourceOperation   `json:"operation,omitempty"`
}
type CapabilityStatementRestResourceInteraction struct {
	Id                *string                `json:"id,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Code              TypeRestfulInteraction `json:"code"`
	Documentation     *string                `json:"documentation,omitempty"`
}
type CapabilityStatementRestResourceSearchParam struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Name              string          `json:"name"`
	Definition        *string         `json:"definition,omitempty"`
	Type              SearchParamType `json:"type"`
	Documentation     *string         `json:"documentation,omitempty"`
}
type CapabilityStatementRestResourceOperation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Definition        string      `json:"definition"`
	Documentation     *string     `json:"documentation,omitempty"`
}
type CapabilityStatementRestInteraction struct {
	Id                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	Code              SystemRestfulInteraction `json:"code"`
	Documentation     *string                  `json:"documentation,omitempty"`
}
type CapabilityStatementMessaging struct {
	Id                *string                                        `json:"id,omitempty"`
	Extension         []Extension                                    `json:"extension,omitempty"`
	ModifierExtension []Extension                                    `json:"modifierExtension,omitempty"`
	Endpoint          []CapabilityStatementMessagingEndpoint         `json:"endpoint,omitempty"`
	ReliableCache     *int                                           `json:"reliableCache,omitempty"`
	Documentation     *string                                        `json:"documentation,omitempty"`
	SupportedMessage  []CapabilityStatementMessagingSupportedMessage `json:"supportedMessage,omitempty"`
}
type CapabilityStatementMessagingEndpoint struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Protocol          Coding      `json:"protocol"`
	Address           string      `json:"address"`
}
type CapabilityStatementMessagingSupportedMessage struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Mode              EventCapabilityMode `json:"mode"`
	Definition        string              `json:"definition"`
}
type CapabilityStatementDocument struct {
	Id                *string      `json:"id,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Mode              DocumentMode `json:"mode"`
	Documentation     *string      `json:"documentation,omitempty"`
	Profile           string       `json:"profile"`
}
type OtherCapabilityStatement CapabilityStatement

// MarshalJSON marshals the given CapabilityStatement as JSON into a byte slice
func (r CapabilityStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCapabilityStatement
		ResourceType string `json:"resourceType"`
	}{
		OtherCapabilityStatement: OtherCapabilityStatement(r),
		ResourceType:             "CapabilityStatement",
	})
}

// UnmarshalCapabilityStatement unmarshals a CapabilityStatement.
func UnmarshalCapabilityStatement(b []byte) (CapabilityStatement, error) {
	var capabilityStatement CapabilityStatement
	if err := json.Unmarshal(b, &capabilityStatement); err != nil {
		return capabilityStatement, err
	}
	return capabilityStatement, nil
}

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

// MedicationKnowledge is documented here http://hl7.org/fhir/StructureDefinition/MedicationKnowledge
type MedicationKnowledge struct {
	Id                         *string                                         `json:"id,omitempty"`
	Meta                       *Meta                                           `json:"meta,omitempty"`
	ImplicitRules              *string                                         `json:"implicitRules,omitempty"`
	Language                   *string                                         `json:"language,omitempty"`
	Text                       *Narrative                                      `json:"text,omitempty"`
	Extension                  []Extension                                     `json:"extension,omitempty"`
	ModifierExtension          []Extension                                     `json:"modifierExtension,omitempty"`
	Code                       *CodeableConcept                                `json:"code,omitempty"`
	Status                     *string                                         `json:"status,omitempty"`
	Manufacturer               *Reference                                      `json:"manufacturer,omitempty"`
	DoseForm                   *CodeableConcept                                `json:"doseForm,omitempty"`
	Amount                     *Quantity                                       `json:"amount,omitempty"`
	Synonym                    []string                                        `json:"synonym,omitempty"`
	RelatedMedicationKnowledge []MedicationKnowledgeRelatedMedicationKnowledge `json:"relatedMedicationKnowledge,omitempty"`
	AssociatedMedication       []Reference                                     `json:"associatedMedication,omitempty"`
	ProductType                []CodeableConcept                               `json:"productType,omitempty"`
	Monograph                  []MedicationKnowledgeMonograph                  `json:"monograph,omitempty"`
	Ingredient                 []MedicationKnowledgeIngredient                 `json:"ingredient,omitempty"`
	PreparationInstruction     *string                                         `json:"preparationInstruction,omitempty"`
	IntendedRoute              []CodeableConcept                               `json:"intendedRoute,omitempty"`
	Cost                       []MedicationKnowledgeCost                       `json:"cost,omitempty"`
	MonitoringProgram          []MedicationKnowledgeMonitoringProgram          `json:"monitoringProgram,omitempty"`
	AdministrationGuidelines   []MedicationKnowledgeAdministrationGuidelines   `json:"administrationGuidelines,omitempty"`
	MedicineClassification     []MedicationKnowledgeMedicineClassification     `json:"medicineClassification,omitempty"`
	Packaging                  *MedicationKnowledgePackaging                   `json:"packaging,omitempty"`
	DrugCharacteristic         []MedicationKnowledgeDrugCharacteristic         `json:"drugCharacteristic,omitempty"`
	Contraindication           []Reference                                     `json:"contraindication,omitempty"`
	Regulatory                 []MedicationKnowledgeRegulatory                 `json:"regulatory,omitempty"`
	Kinetics                   []MedicationKnowledgeKinetics                   `json:"kinetics,omitempty"`
}
type MedicationKnowledgeRelatedMedicationKnowledge struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Reference         []Reference     `json:"reference"`
}
type MedicationKnowledgeMonograph struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Source            *Reference       `json:"source,omitempty"`
}
type MedicationKnowledgeIngredient struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	IsActive          *bool       `json:"isActive,omitempty"`
	Strength          *Ratio      `json:"strength,omitempty"`
}
type MedicationKnowledgeCost struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Source            *string         `json:"source,omitempty"`
	Cost              Money           `json:"cost"`
}
type MedicationKnowledgeMonitoringProgram struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Name              *string          `json:"name,omitempty"`
}
type MedicationKnowledgeAdministrationGuidelines struct {
	Id                     *string                                                             `json:"id,omitempty"`
	Extension              []Extension                                                         `json:"extension,omitempty"`
	ModifierExtension      []Extension                                                         `json:"modifierExtension,omitempty"`
	Dosage                 []MedicationKnowledgeAdministrationGuidelinesDosage                 `json:"dosage,omitempty"`
	PatientCharacteristics []MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics `json:"patientCharacteristics,omitempty"`
}
type MedicationKnowledgeAdministrationGuidelinesDosage struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Dosage            []Dosage        `json:"dosage"`
}
type MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Value             []string    `json:"value,omitempty"`
}
type MedicationKnowledgeMedicineClassification struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type"`
	Classification    []CodeableConcept `json:"classification,omitempty"`
}
type MedicationKnowledgePackaging struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Quantity          *Quantity        `json:"quantity,omitempty"`
}
type MedicationKnowledgeDrugCharacteristic struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
}
type MedicationKnowledgeRegulatory struct {
	Id                  *string                                     `json:"id,omitempty"`
	Extension           []Extension                                 `json:"extension,omitempty"`
	ModifierExtension   []Extension                                 `json:"modifierExtension,omitempty"`
	RegulatoryAuthority Reference                                   `json:"regulatoryAuthority"`
	Substitution        []MedicationKnowledgeRegulatorySubstitution `json:"substitution,omitempty"`
	Schedule            []MedicationKnowledgeRegulatorySchedule     `json:"schedule,omitempty"`
	MaxDispense         *MedicationKnowledgeRegulatoryMaxDispense   `json:"maxDispense,omitempty"`
}
type MedicationKnowledgeRegulatorySubstitution struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Allowed           bool            `json:"allowed"`
}
type MedicationKnowledgeRegulatorySchedule struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Schedule          CodeableConcept `json:"schedule"`
}
type MedicationKnowledgeRegulatoryMaxDispense struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Quantity          Quantity    `json:"quantity"`
	Period            *Duration   `json:"period,omitempty"`
}
type MedicationKnowledgeKinetics struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	AreaUnderCurve    []Quantity  `json:"areaUnderCurve,omitempty"`
	LethalDose50      []Quantity  `json:"lethalDose50,omitempty"`
	HalfLifePeriod    *Duration   `json:"halfLifePeriod,omitempty"`
}
type OtherMedicationKnowledge MedicationKnowledge

// MarshalJSON marshals the given MedicationKnowledge as JSON into a byte slice
func (r MedicationKnowledge) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationKnowledge
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationKnowledge: OtherMedicationKnowledge(r),
		ResourceType:             "MedicationKnowledge",
	})
}

// UnmarshalMedicationKnowledge unmarshals a MedicationKnowledge.
func UnmarshalMedicationKnowledge(b []byte) (MedicationKnowledge, error) {
	var medicationKnowledge MedicationKnowledge
	if err := json.Unmarshal(b, &medicationKnowledge); err != nil {
		return medicationKnowledge, err
	}
	return medicationKnowledge, nil
}

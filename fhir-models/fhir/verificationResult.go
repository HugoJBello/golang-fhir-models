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

// VerificationResult is documented here http://hl7.org/fhir/StructureDefinition/VerificationResult
type VerificationResult struct {
	Id                *string                           `json:"id,omitempty"`
	Meta              *Meta                             `json:"meta,omitempty"`
	ImplicitRules     *string                           `json:"implicitRules,omitempty"`
	Language          *string                           `json:"language,omitempty"`
	Text              *Narrative                        `json:"text,omitempty"`
	Extension         []Extension                       `json:"extension,omitempty"`
	ModifierExtension []Extension                       `json:"modifierExtension,omitempty"`
	Target            []Reference                       `json:"target,omitempty"`
	TargetLocation    []string                          `json:"targetLocation,omitempty"`
	Need              *CodeableConcept                  `json:"need,omitempty"`
	Status            string                            `json:"status"`
	StatusDate        *string                           `json:"statusDate,omitempty"`
	ValidationType    *CodeableConcept                  `json:"validationType,omitempty"`
	ValidationProcess []CodeableConcept                 `json:"validationProcess,omitempty"`
	Frequency         *Timing                           `json:"frequency,omitempty"`
	LastPerformed     *string                           `json:"lastPerformed,omitempty"`
	NextScheduled     *string                           `json:"nextScheduled,omitempty"`
	FailureAction     *CodeableConcept                  `json:"failureAction,omitempty"`
	PrimarySource     []VerificationResultPrimarySource `json:"primarySource,omitempty"`
	Attestation       *VerificationResultAttestation    `json:"attestation,omitempty"`
	Validator         []VerificationResultValidator     `json:"validator,omitempty"`
}
type VerificationResultPrimarySource struct {
	Id                  *string           `json:"id,omitempty"`
	Extension           []Extension       `json:"extension,omitempty"`
	ModifierExtension   []Extension       `json:"modifierExtension,omitempty"`
	Who                 *Reference        `json:"who,omitempty"`
	Type                []CodeableConcept `json:"type,omitempty"`
	CommunicationMethod []CodeableConcept `json:"communicationMethod,omitempty"`
	ValidationStatus    *CodeableConcept  `json:"validationStatus,omitempty"`
	ValidationDate      *string           `json:"validationDate,omitempty"`
	CanPushUpdates      *CodeableConcept  `json:"canPushUpdates,omitempty"`
	PushTypeAvailable   []CodeableConcept `json:"pushTypeAvailable,omitempty"`
}
type VerificationResultAttestation struct {
	Id                        *string          `json:"id,omitempty"`
	Extension                 []Extension      `json:"extension,omitempty"`
	ModifierExtension         []Extension      `json:"modifierExtension,omitempty"`
	Who                       *Reference       `json:"who,omitempty"`
	OnBehalfOf                *Reference       `json:"onBehalfOf,omitempty"`
	CommunicationMethod       *CodeableConcept `json:"communicationMethod,omitempty"`
	Date                      *string          `json:"date,omitempty"`
	SourceIdentityCertificate *string          `json:"sourceIdentityCertificate,omitempty"`
	ProxyIdentityCertificate  *string          `json:"proxyIdentityCertificate,omitempty"`
	ProxySignature            *Signature       `json:"proxySignature,omitempty"`
	SourceSignature           *Signature       `json:"sourceSignature,omitempty"`
}
type VerificationResultValidator struct {
	Id                   *string     `json:"id,omitempty"`
	Extension            []Extension `json:"extension,omitempty"`
	ModifierExtension    []Extension `json:"modifierExtension,omitempty"`
	Organization         Reference   `json:"organization"`
	IdentityCertificate  *string     `json:"identityCertificate,omitempty"`
	AttestationSignature *Signature  `json:"attestationSignature,omitempty"`
}
type OtherVerificationResult VerificationResult

// MarshalJSON marshals the given VerificationResult as JSON into a byte slice
func (r VerificationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherVerificationResult
		ResourceType string `json:"resourceType"`
	}{
		OtherVerificationResult: OtherVerificationResult(r),
		ResourceType:            "VerificationResult",
	})
}

// UnmarshalVerificationResult unmarshals a VerificationResult.
func UnmarshalVerificationResult(b []byte) (VerificationResult, error) {
	var verificationResult VerificationResult
	if err := json.Unmarshal(b, &verificationResult); err != nil {
		return verificationResult, err
	}
	return verificationResult, nil
}

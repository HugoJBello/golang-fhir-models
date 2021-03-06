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

// PaymentNotice is documented here http://hl7.org/fhir/StructureDefinition/PaymentNotice
type PaymentNotice struct {
	Id                *string                      `json:"id,omitempty"`
	Meta              *Meta                        `json:"meta,omitempty"`
	ImplicitRules     *string                      `json:"implicitRules,omitempty"`
	Language          *string                      `json:"language,omitempty"`
	Text              *Narrative                   `json:"text,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                 `json:"identifier,omitempty"`
	Status            FinancialResourceStatusCodes `json:"status"`
	Request           *Reference                   `json:"request,omitempty"`
	Response          *Reference                   `json:"response,omitempty"`
	Created           string                       `json:"created"`
	Provider          *Reference                   `json:"provider,omitempty"`
	Payment           Reference                    `json:"payment"`
	PaymentDate       *string                      `json:"paymentDate,omitempty"`
	Payee             *Reference                   `json:"payee,omitempty"`
	Recipient         Reference                    `json:"recipient"`
	Amount            Money                        `json:"amount"`
	PaymentStatus     *CodeableConcept             `json:"paymentStatus,omitempty"`
}
type OtherPaymentNotice PaymentNotice

// MarshalJSON marshals the given PaymentNotice as JSON into a byte slice
func (r PaymentNotice) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPaymentNotice
		ResourceType string `json:"resourceType"`
	}{
		OtherPaymentNotice: OtherPaymentNotice(r),
		ResourceType:       "PaymentNotice",
	})
}

// UnmarshalPaymentNotice unmarshals a PaymentNotice.
func UnmarshalPaymentNotice(b []byte) (PaymentNotice, error) {
	var paymentNotice PaymentNotice
	if err := json.Unmarshal(b, &paymentNotice); err != nil {
		return paymentNotice, err
	}
	return paymentNotice, nil
}

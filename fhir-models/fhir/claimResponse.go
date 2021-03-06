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

// ClaimResponse is documented here http://hl7.org/fhir/StructureDefinition/ClaimResponse
type ClaimResponse struct {
	Id                   *string                         `json:"id,omitempty"`
	Meta                 *Meta                           `json:"meta,omitempty"`
	ImplicitRules        *string                         `json:"implicitRules,omitempty"`
	Language             *string                         `json:"language,omitempty"`
	Text                 *Narrative                      `json:"text,omitempty"`
	Extension            []Extension                     `json:"extension,omitempty"`
	ModifierExtension    []Extension                     `json:"modifierExtension,omitempty"`
	Identifier           []Identifier                    `json:"identifier,omitempty"`
	Status               FinancialResourceStatusCodes    `json:"status"`
	Type                 CodeableConcept                 `json:"type"`
	SubType              *CodeableConcept                `json:"subType,omitempty"`
	Use                  Use                             `json:"use"`
	Patient              Reference                       `json:"patient"`
	Created              string                          `json:"created"`
	Insurer              Reference                       `json:"insurer"`
	Requestor            *Reference                      `json:"requestor,omitempty"`
	Request              *Reference                      `json:"request,omitempty"`
	Outcome              ClaimProcessingCodes            `json:"outcome"`
	Disposition          *string                         `json:"disposition,omitempty"`
	PreAuthRef           *string                         `json:"preAuthRef,omitempty"`
	PreAuthPeriod        *Period                         `json:"preAuthPeriod,omitempty"`
	PayeeType            *CodeableConcept                `json:"payeeType,omitempty"`
	Item                 []ClaimResponseItem             `json:"item,omitempty"`
	AddItem              []ClaimResponseAddItem          `json:"addItem,omitempty"`
	Adjudication         []ClaimResponseItemAdjudication `json:"adjudication,omitempty"`
	Total                []ClaimResponseTotal            `json:"total,omitempty"`
	Payment              *ClaimResponsePayment           `json:"payment,omitempty"`
	FundsReserve         *CodeableConcept                `json:"fundsReserve,omitempty"`
	FormCode             *CodeableConcept                `json:"formCode,omitempty"`
	Form                 *Attachment                     `json:"form,omitempty"`
	ProcessNote          []ClaimResponseProcessNote      `json:"processNote,omitempty"`
	CommunicationRequest []Reference                     `json:"communicationRequest,omitempty"`
	Insurance            []ClaimResponseInsurance        `json:"insurance,omitempty"`
	Error                []ClaimResponseError            `json:"error,omitempty"`
}
type ClaimResponseItem struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	ItemSequence      int                             `json:"itemSequence"`
	NoteNumber        []int                           `json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `json:"adjudication"`
	Detail            []ClaimResponseItemDetail       `json:"detail,omitempty"`
}
type ClaimResponseItemAdjudication struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Category          CodeableConcept  `json:"category"`
	Reason            *CodeableConcept `json:"reason,omitempty"`
	Amount            *Money           `json:"amount,omitempty"`
	Value             *string          `json:"value,omitempty"`
}
type ClaimResponseItemDetail struct {
	Id                *string                            `json:"id,omitempty"`
	Extension         []Extension                        `json:"extension,omitempty"`
	ModifierExtension []Extension                        `json:"modifierExtension,omitempty"`
	DetailSequence    int                                `json:"detailSequence"`
	NoteNumber        []int                              `json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication    `json:"adjudication,omitempty"`
	SubDetail         []ClaimResponseItemDetailSubDetail `json:"subDetail,omitempty"`
}
type ClaimResponseItemDetailSubDetail struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	SubDetailSequence int                             `json:"subDetailSequence"`
	NoteNumber        []int                           `json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `json:"adjudication,omitempty"`
}
type ClaimResponseAddItem struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	ItemSequence      []int                           `json:"itemSequence,omitempty"`
	DetailSequence    []int                           `json:"detailSequence,omitempty"`
	SubdetailSequence []int                           `json:"subdetailSequence,omitempty"`
	Provider          []Reference                     `json:"provider,omitempty"`
	ProductOrService  CodeableConcept                 `json:"productOrService"`
	Modifier          []CodeableConcept               `json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept               `json:"programCode,omitempty"`
	Quantity          *Quantity                       `json:"quantity,omitempty"`
	UnitPrice         *Money                          `json:"unitPrice,omitempty"`
	Factor            *string                         `json:"factor,omitempty"`
	Net               *Money                          `json:"net,omitempty"`
	BodySite          *CodeableConcept                `json:"bodySite,omitempty"`
	SubSite           []CodeableConcept               `json:"subSite,omitempty"`
	NoteNumber        []int                           `json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `json:"adjudication,omitempty"`
	Detail            []ClaimResponseAddItemDetail    `json:"detail,omitempty"`
}
type ClaimResponseAddItemDetail struct {
	Id                *string                               `json:"id,omitempty"`
	Extension         []Extension                           `json:"extension,omitempty"`
	ModifierExtension []Extension                           `json:"modifierExtension,omitempty"`
	ProductOrService  CodeableConcept                       `json:"productOrService"`
	Modifier          []CodeableConcept                     `json:"modifier,omitempty"`
	Quantity          *Quantity                             `json:"quantity,omitempty"`
	UnitPrice         *Money                                `json:"unitPrice,omitempty"`
	Factor            *string                               `json:"factor,omitempty"`
	Net               *Money                                `json:"net,omitempty"`
	NoteNumber        []int                                 `json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication       `json:"adjudication,omitempty"`
	SubDetail         []ClaimResponseAddItemDetailSubDetail `json:"subDetail,omitempty"`
}
type ClaimResponseAddItemDetailSubDetail struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	ProductOrService  CodeableConcept                 `json:"productOrService"`
	Modifier          []CodeableConcept               `json:"modifier,omitempty"`
	Quantity          *Quantity                       `json:"quantity,omitempty"`
	UnitPrice         *Money                          `json:"unitPrice,omitempty"`
	Factor            *string                         `json:"factor,omitempty"`
	Net               *Money                          `json:"net,omitempty"`
	NoteNumber        []int                           `json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `json:"adjudication,omitempty"`
}
type ClaimResponseTotal struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Category          CodeableConcept `json:"category"`
	Amount            Money           `json:"amount"`
}
type ClaimResponsePayment struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              CodeableConcept  `json:"type"`
	Adjustment        *Money           `json:"adjustment,omitempty"`
	AdjustmentReason  *CodeableConcept `json:"adjustmentReason,omitempty"`
	Date              *string          `json:"date,omitempty"`
	Amount            Money            `json:"amount"`
	Identifier        *Identifier      `json:"identifier,omitempty"`
}
type ClaimResponseProcessNote struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Number            *int             `json:"number,omitempty"`
	Type              *NoteType        `json:"type,omitempty"`
	Text              string           `json:"text"`
	Language          *CodeableConcept `json:"language,omitempty"`
}
type ClaimResponseInsurance struct {
	Id                  *string     `json:"id,omitempty"`
	Extension           []Extension `json:"extension,omitempty"`
	ModifierExtension   []Extension `json:"modifierExtension,omitempty"`
	Sequence            int         `json:"sequence"`
	Focal               bool        `json:"focal"`
	Coverage            Reference   `json:"coverage"`
	BusinessArrangement *string     `json:"businessArrangement,omitempty"`
	ClaimResponse       *Reference  `json:"claimResponse,omitempty"`
}
type ClaimResponseError struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	ItemSequence      *int            `json:"itemSequence,omitempty"`
	DetailSequence    *int            `json:"detailSequence,omitempty"`
	SubDetailSequence *int            `json:"subDetailSequence,omitempty"`
	Code              CodeableConcept `json:"code"`
}
type OtherClaimResponse ClaimResponse

// MarshalJSON marshals the given ClaimResponse as JSON into a byte slice
func (r ClaimResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClaimResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherClaimResponse: OtherClaimResponse(r),
		ResourceType:       "ClaimResponse",
	})
}

// UnmarshalClaimResponse unmarshals a ClaimResponse.
func UnmarshalClaimResponse(b []byte) (ClaimResponse, error) {
	var claimResponse ClaimResponse
	if err := json.Unmarshal(b, &claimResponse); err != nil {
		return claimResponse, err
	}
	return claimResponse, nil
}

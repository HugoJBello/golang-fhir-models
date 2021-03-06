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

// RequestGroup is documented here http://hl7.org/fhir/StructureDefinition/RequestGroup
type RequestGroup struct {
	Id                    *string              `json:"id,omitempty"`
	Meta                  *Meta                `json:"meta,omitempty"`
	ImplicitRules         *string              `json:"implicitRules,omitempty"`
	Language              *string              `json:"language,omitempty"`
	Text                  *Narrative           `json:"text,omitempty"`
	Extension             []Extension          `json:"extension,omitempty"`
	ModifierExtension     []Extension          `json:"modifierExtension,omitempty"`
	Identifier            []Identifier         `json:"identifier,omitempty"`
	InstantiatesCanonical []string             `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string             `json:"instantiatesUri,omitempty"`
	BasedOn               []Reference          `json:"basedOn,omitempty"`
	Replaces              []Reference          `json:"replaces,omitempty"`
	GroupIdentifier       *Identifier          `json:"groupIdentifier,omitempty"`
	Status                RequestStatus        `json:"status"`
	Intent                RequestIntent        `json:"intent"`
	Priority              *RequestPriority     `json:"priority,omitempty"`
	Code                  *CodeableConcept     `json:"code,omitempty"`
	Subject               *Reference           `json:"subject,omitempty"`
	Encounter             *Reference           `json:"encounter,omitempty"`
	AuthoredOn            *string              `json:"authoredOn,omitempty"`
	Author                *Reference           `json:"author,omitempty"`
	ReasonCode            []CodeableConcept    `json:"reasonCode,omitempty"`
	ReasonReference       []Reference          `json:"reasonReference,omitempty"`
	Note                  []Annotation         `json:"note,omitempty"`
	Action                []RequestGroupAction `json:"action,omitempty"`
}
type RequestGroupAction struct {
	Id                  *string                           `json:"id,omitempty"`
	Extension           []Extension                       `json:"extension,omitempty"`
	ModifierExtension   []Extension                       `json:"modifierExtension,omitempty"`
	Prefix              *string                           `json:"prefix,omitempty"`
	Title               *string                           `json:"title,omitempty"`
	Description         *string                           `json:"description,omitempty"`
	TextEquivalent      *string                           `json:"textEquivalent,omitempty"`
	Priority            *RequestPriority                  `json:"priority,omitempty"`
	Code                []CodeableConcept                 `json:"code,omitempty"`
	Documentation       []RelatedArtifact                 `json:"documentation,omitempty"`
	Condition           []RequestGroupActionCondition     `json:"condition,omitempty"`
	RelatedAction       []RequestGroupActionRelatedAction `json:"relatedAction,omitempty"`
	Participant         []Reference                       `json:"participant,omitempty"`
	Type                *CodeableConcept                  `json:"type,omitempty"`
	GroupingBehavior    *ActionGroupingBehavior           `json:"groupingBehavior,omitempty"`
	SelectionBehavior   *ActionSelectionBehavior          `json:"selectionBehavior,omitempty"`
	RequiredBehavior    *ActionRequiredBehavior           `json:"requiredBehavior,omitempty"`
	PrecheckBehavior    *ActionPrecheckBehavior           `json:"precheckBehavior,omitempty"`
	CardinalityBehavior *ActionCardinalityBehavior        `json:"cardinalityBehavior,omitempty"`
	Resource            *Reference                        `json:"resource,omitempty"`
	Action              []RequestGroupAction              `json:"action,omitempty"`
}
type RequestGroupActionCondition struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Kind              ActionConditionKind `json:"kind"`
	Expression        *Expression         `json:"expression,omitempty"`
}
type RequestGroupActionRelatedAction struct {
	Id                *string                `json:"id,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	ActionId          string                 `json:"actionId"`
	Relationship      ActionRelationshipType `json:"relationship"`
}
type OtherRequestGroup RequestGroup

// MarshalJSON marshals the given RequestGroup as JSON into a byte slice
func (r RequestGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRequestGroup
		ResourceType string `json:"resourceType"`
	}{
		OtherRequestGroup: OtherRequestGroup(r),
		ResourceType:      "RequestGroup",
	})
}

// UnmarshalRequestGroup unmarshals a RequestGroup.
func UnmarshalRequestGroup(b []byte) (RequestGroup, error) {
	var requestGroup RequestGroup
	if err := json.Unmarshal(b, &requestGroup); err != nil {
		return requestGroup, err
	}
	return requestGroup, nil
}

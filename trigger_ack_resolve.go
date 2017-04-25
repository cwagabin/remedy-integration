// Copyright:: Copyright (c) 2016-2017 PagerDuty, Inc.
// License:: Apache License, Version 2.0

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type IncidentBody struct {
	Details string `json:"details,omitempty"`
	Type    string `json:"type,omitempty"`
}

type IncidentIDReference struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

type IncidentAssignee struct {
	Assignee *IncidentIDReference `json:"assignee,omitempty"`
}

type Incident struct {
	IncidentKey      string               `json:"incident_key,omitempty"`
	Type             string               `json:"type,omitempty"`
	Title            string               `json:"title,omitempty"`
	Body             *IncidentBody        `json:"body,omitempty"`
	Status           string               `json:"status,omitempty"`
	Service          *IncidentIDReference `json:"service,omitempty"`
	EscalationPolicy *IncidentIDReference `json:"escalation_policy,omitempty"`
	Assignments      []IncidentAssignee   `json:"assignments,omitempty"`
}

type IncidentEntry struct {
	Entry Incident `json:"incident"`
}

func setIDReference(id string, idType string) *IncidentIDReference {
	if id == "" {
		return nil
	}
	idRef := IncidentIDReference{}
	idRef.ID = id
	idRef.Type = idType
	return &idRef
}

func setDetails(details string, detailsType string) *IncidentBody {
	if details == "" {
		return nil
	}
	idRef := IncidentBody{}
	idRef.Details = details
	idRef.Type = detailsType
	return &idRef
}

func setAssignments(id string, idType string, assignments []IncidentAssignee) []IncidentAssignee {
	if id == "" {
		return nil
	}
	assignee := &IncidentAssignee{}
	assignee.Assignee = setIDReference(id, idType)
	return append(assignments, *assignee)
}

func triggerAckResolveChangePolicy(c *Config) {
	if c.IncidentID == "" {
		create(c)
	} else {
		update(c)
	}
}

func create(c *Config) {
	url := fmt.Sprintf("%s/incidents", c.apiEndpoint())

	incident := &Incident{}
	incident.Type = c.Type
	incident.IncidentKey = c.IncidentKey
	incident.Title = c.Description
	incident.Body = setDetails(c.Details, "incident_body")
	incident.Status = c.Status
	incident.Service = setIDReference(c.ServiceID, "service_reference")
	incident.EscalationPolicy = setIDReference(c.EscalationPolicyID, "escalation_policy_reference")
	incident.Assignments = setAssignments(c.Assignee, "user_reference", incident.Assignments)

	incidentEntry := &IncidentEntry{}
	incidentEntry.Entry = *incident

	payload, err := json.Marshal(incidentEntry)
	failIf(err)

	req := &HTTPRequest{
		Method: "POST",
		URL:    &url,
		Data:   bytes.NewBuffer(payload),
	}

	_, _, err = httpRequest(c, req)
	failIf(err)
}

func update(c *Config) {
	url := fmt.Sprintf("%s/incidents/%s", c.apiEndpoint(), c.IncidentID)

	incident := &Incident{}
	incident.Type = c.Type
	incident.Status = c.Status
	incident.Title = c.Description
	incident.EscalationPolicy = setIDReference(c.EscalationPolicyID, "escalation_policy_reference")
	incident.Assignments = setAssignments(c.Assignee, "user_reference", incident.Assignments)
	incidentEntry := &IncidentEntry{}
	incidentEntry.Entry = *incident
	payload, err := json.Marshal(incidentEntry)
	failIf(err)

	req := &HTTPRequest{
		Method: "PUT",
		URL:    &url,
		Data:   bytes.NewBuffer(payload),
	}
	_, _, err = httpRequest(c, req)
	failIf(err)
}

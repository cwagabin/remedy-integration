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
	"encoding/json"
	"fmt"
)

type ILE struct {
	Type             string
	Summary          string
	AgentName        string
	AgentType        string
	ChannelType      string
	NotificationType string
	UserName         string
	AssignedUser     string
	ChannelSummary   string
	Message          string
}

func (i *ILE) FriendlyMessage() *string {
	var agentName string
	var channelType string
	var friendlyMessage string

	channelTypes := map[string]string{
		"api":                       "API",
		"auto":                      "auto",
		"sms":                       "SMS",
		"phone":                     "phone",
		"email":                     "email",
		"short_email":               "email",
		"website":                   "website",
		"ios_push_notification":     "push notification",
		"android_push_notification": "push notification",
	}

	if i.AgentType == "service" {
		agentName = "API"
	} else {
		agentName = i.AgentName
	}

	switch i.Type {
	case "acknowledge_log_entry":
		{
			friendlyMessage = fmt.Sprintf("Acknowledged by %s", agentName)
		}
	case "resolve_log_entry":
		{
			friendlyMessage = fmt.Sprintf("Resolved by %s", agentName)
		}
	case "trigger_log_entry":
		{
			friendlyMessage = fmt.Sprintf("Triggered by %s", agentName)
		}
	case "notify_log_entry":
		{
			friendlyMessage = fmt.Sprintf("Notified %s", i.UserName)
		}
	case "assign_log_entry":
		{
			friendlyMessage = fmt.Sprintf("Assigned to %s", i.AssignedUser)
		}
	case "escalate_log_entry":
		{
			friendlyMessage = fmt.Sprintf("Escalated to %s", i.AssignedUser)
		}
	case "delegate_log_entry":
		{
			friendlyMessage = fmt.Sprintf("Delegated Default by %s", agentName)
		}
	case "annotate_log_entry":
		{
			friendlyMessage = fmt.Sprintf("Note added by %s\n=>%s", agentName, i.ChannelSummary)
		}
	case "status_update_log_entry":
		{
			friendlyMessage = fmt.Sprintf("%s\n=>%s", i.Summary, i.Message)
		}
	default:
		{
			friendlyMessage = fmt.Sprintf(i.Summary)
		}
	}

	if i.ChannelType != "" {
		channelType = i.ChannelType
	} else {
		channelType = i.NotificationType
	}

	if val, ok := channelTypes[channelType]; ok {
		friendlyMessage = fmt.Sprintf("%s via %s", friendlyMessage, val)
	}

	return &friendlyMessage
}

// Returns an empty map if the JSON key is missing
// Allows us to chase down empty paths and only check for nil at the end
func mapOrNewMap(m map[string]interface{}, k string) map[string]interface{} {
	if val, ok := m[k]; ok {
		return val.(map[string]interface{})
	} else {
		return make(map[string]interface{})
	}
}

// Returns an empty string if the JSON key is missing
// Allows us to chase down empty paths and only check for nil at the end
func stringOrNewString(m map[string]interface{}, k string) string {
	if val, ok := m[k]; ok {
		return val.(string)
	} else {
		return ""
	}
}

func mapArrayOrNewMap(m map[string]interface{}, k string) map[string]interface{} {
	if val, ok := m[k]; ok {
		arr := val.([]interface{})
		if len(arr) > 0 {
			return arr[0].(map[string]interface{})
		}
	}
	return make(map[string]interface{})
}

func parseIles(res []byte) ([]*ILE, error) {
	var jsonMap map[string]interface{}
	var iles []*ILE

	if err := json.Unmarshal(res, &jsonMap); err != nil {
		return iles, err
	}

	rawILEs := jsonMap["log_entries"].([]interface{})
	for _, rawILE := range rawILEs {
		ile := rawILE.(map[string]interface{})

		ileAgent := mapOrNewMap(ile, "agent")
		ileUser := mapOrNewMap(ile, "user")
		ileAssignedUser := mapArrayOrNewMap(ile, "assignees")
		ileChannel := mapOrNewMap(ile, "channel")
		ileNotification := mapOrNewMap(ile, "notification")

		currentILE := &ILE{
			Type:             stringOrNewString(ile, "type"),
			Summary:          stringOrNewString(ile, "summary"),
			AgentName:        stringOrNewString(ileAgent, "summary"),
			AgentType:        stringOrNewString(ileAgent, "type"),
			ChannelType:      stringOrNewString(ileChannel, "type"),
			NotificationType: stringOrNewString(ileNotification, "type"),
			UserName:         stringOrNewString(ileUser, "summary"),
			AssignedUser:     stringOrNewString(ileAssignedUser, "summary"),
			ChannelSummary:   stringOrNewString(ileChannel, "summary"),
			Message:          stringOrNewString(ile, "message"),
		}

		iles = append(iles, currentILE)
	}

	return iles, nil
}

func getIles(c *Config) (messages []string) {
	url := fmt.Sprintf("%s/incidents/%s/log_entries", c.apiEndpoint(), c.IncidentID)
	req := &HTTPRequest{
		Method: "GET",
		URL:    &url,
		Data:   nil,
	}
	res, _, err := httpRequest(c, req)
	failIf(err)

	ileSlice, err := parseIles(res)
	failIf(err)

	for _, ile := range ileSlice {
		messages = append(messages, *ile.FriendlyMessage())
	}

	return
}

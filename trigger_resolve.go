// Copyright:: Copyright (c) 2016 PagerDuty, Inc.
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
)

type Event struct {
	ServiceKey  string `json:"service_key"`
	IncidentKey string `json:"incident_key"`
	EventType   string `json:"event_type"`
	Description string `json:"description"`
	Details     string `json:"details"`
}

func triggerResolve(c *Config) {
	event := &Event{}
	eventEndpoint := "https://events.pagerduty.com/generic/2010-04-15/create_event.json"

	event.EventType = c.Mode
	event.ServiceKey = c.ServiceKey
	event.IncidentKey = c.IncidentKey
	event.Description = c.Description
	event.Details = c.Details

	payload, err := json.Marshal(event)
	failIf(err)

	req := &HttpRequest{
		Method: "POST",
		Url:    &eventEndpoint,
		Data:   bytes.NewBuffer(payload),
	}

	_, err = httpRequest(c, req)
	failIf(err)
}

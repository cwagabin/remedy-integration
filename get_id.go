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
	"net/url"
	"time"
)

func getID(c *Config) string {
	req := &HTTPRequest{}
	url := fmt.Sprintf("%s/incidents?incident_key=%s&sort_by=%s", c.apiEndpoint(), c.IncidentKey, url.QueryEscape("created_at:DESC"))
	req.URL = &url
	req.Method = "GET"
	res, _, err := httpRequest(c, req)
	failIf(err)

	// Store JSON response into a map
	var jsonMap map[string]interface{}
	failIf(json.Unmarshal(res, &jsonMap))

	// Pull out the first Incident ID returned, return empty string if no incidents found
	incidents := jsonMap["incidents"].([]interface{})
	if len(incidents) < 1 {
		return ""
	}
	firstIncident := incidents[0].(map[string]interface{})
	return firstIncident["id"].(string)
}

func getValidateID(c *Config) string {
	req := &HTTPRequest{}
	url := fmt.Sprintf("%s/incidents/%s", c.apiEndpoint(), c.IncidentID)
	req.URL = &url
	req.Method = "GET"
	res, statusCode, err := httpRequest(c, req)
	if statusCode == 404 {
		// Incident is not found
		return ""
	}
	failIf(err)

	// Store JSON response into a map
	var jsonMap map[string]interface{}
	failIf(json.Unmarshal(res, &jsonMap))

	incident := jsonMap["incident"].(map[string]interface{})
	return incident["id"].(string)
}

func getIDStatus(c *Config) (id string, status string) {
	req := &HTTPRequest{}
	url := fmt.Sprintf("%s/incidents?incident_key=%s&sort_by=%s", c.apiEndpoint(), c.IncidentKey, url.QueryEscape("created_at:DESC"))
	req.URL = &url
	req.Method = "GET"
	res, _, err := httpRequest(c, req)
	failIf(err)

	// Store JSON response into a map
	var jsonMap map[string]interface{}
	failIf(json.Unmarshal(res, &jsonMap))

	// Pull out the first Incident ID returned, return empty string if no incidents found
	incidents := jsonMap["incidents"].([]interface{})
	if len(incidents) < 1 {
		return "", ""
	}
	firstIncident := incidents[0].(map[string]interface{})
	return firstIncident["id"].(string), firstIncident["status"].(string)
}

func getIncidentStatus(c *Config) string {
	req := &HTTPRequest{}
	url := fmt.Sprintf("%s/incidents/%s", c.apiEndpoint(), c.IncidentID)
	req.URL = &url
	req.Method = "GET"
	res, statusCode, err := httpRequest(c, req)
	if statusCode == 404 {
		// Service is not found
		return ""
	}
	failIf(err)

	// Store JSON response into a map
	var jsonMap map[string]interface{}
	failIf(json.Unmarshal(res, &jsonMap))

	incident := jsonMap["incident"].(map[string]interface{})
	return incident["status"].(string)
}

func getIDRetry(c *Config, maxRetries int, delayTime int) (id string, status string) {
	attempt := 1
	for {
		id, status = getIDStatus(c)
		if id != "" && (c.Status != status && status == "resolved") {
			// This is to prevent the 2-call process from finding the old PD incident that is already resolved
			id = ""
			status = ""
		}
		if (id != "") || (attempt >= maxRetries && id == "") {
			return id, status
		}
		time.Sleep(time.Duration(delayTime) * time.Second)
		attempt++
	}
}

func testAPIConnection(c *Config) string {
	req := &HTTPRequest{}
	urlString := fmt.Sprintf("%s/incidents/count", c.apiEndpoint())
	req.URL = &urlString
	req.Method = "GET"
	_, _, err := httpRequest(c, req)
	if err != nil {
		return "failed"
	}
	return "ok"
}

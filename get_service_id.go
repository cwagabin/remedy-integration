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
)

func getServiceID(c *Config) string {
	req := &HTTPRequest{}
	urlString := fmt.Sprintf("%s/services?name=%s", c.apiEndpoint(), url.QueryEscape(c.ServiceName))
	req.URL = &urlString
	req.Method = "GET"
	res, _, err := httpRequest(c, req)
	failIf(err)

	// Store JSON response into a map
	var jsonMap map[string]interface{}
	failIf(json.Unmarshal(res, &jsonMap))

	// Loop through to get the matching service name and return the id (not performing exact match)
	services := jsonMap["services"].([]interface{})
	for i := 0; i < len(services); i++ {
		service := services[i].(map[string]interface{})
		name := service["name"].(string)
		if name == c.ServiceName {
			return service["id"].(string)
		}
	}
	return ""
}

func getServiceEscalationID(c *Config) string {
	req := &HTTPRequest{}
	urlString := fmt.Sprintf("%s/services/%s", c.apiEndpoint(), c.ServiceID)
	req.URL = &urlString
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

	//Check if a service is found
	service := jsonMap["service"].(map[string]interface{})
	escalationPolicy := service["escalation_policy"].(map[string]interface{})
	return escalationPolicy["id"].(string)
}

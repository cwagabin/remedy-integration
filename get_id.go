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

import "encoding/json"
import "fmt"

func getID(c *Config) string {
	req := &HTTPRequest{}
	url := fmt.Sprintf("%s/incidents?incident_key=%s", c.apiEndpoint(), c.IncidentKey)
	req.URL = &url
	req.Method = "GET"
	res, err := httpRequest(c, req)
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

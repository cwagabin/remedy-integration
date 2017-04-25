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

func validateEscalationPolicyID(c *Config) string {
	req := &HTTPRequest{}
	url := fmt.Sprintf("%s/escalation_policies/%s", c.apiEndpoint(), c.EscalationPolicyID)
	req.URL = &url
	req.Method = "GET"
	res, statusCode, err := httpRequest(c, req)
	if statusCode == 404 {
		// Escalation policy is not found
		return ""
	}
	failIf(err)

	// Store JSON response into a map
	var jsonMap map[string]interface{}
	failIf(json.Unmarshal(res, &jsonMap))

	escalationPolicy := jsonMap["escalation_policy"].(map[string]interface{})
	return escalationPolicy["id"].(string)
}

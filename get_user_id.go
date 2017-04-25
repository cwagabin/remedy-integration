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

func getUserID(c *Config) string {
	req := &HTTPRequest{}
	url := fmt.Sprintf("%s/users?query=%s", c.apiEndpoint(), c.UserEmail)
	req.URL = &url
	req.Method = "GET"
	res, _, err := httpRequest(c, req)
	failIf(err)

	// Store JSON response into a map
	var jsonMap map[string]interface{}
	failIf(json.Unmarshal(res, &jsonMap))

	// Pull out the first User ID returned, return empty string if no user found
	users := jsonMap["users"].([]interface{})
	if len(users) < 1 {
		return ""
	}
	firstUser := users[0].(map[string]interface{})
	return firstUser["id"].(string)
}

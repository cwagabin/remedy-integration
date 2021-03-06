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
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type HTTPRequest struct {
	Method string
	URL    *string
	Data   io.Reader
}

func httpRequest(c *Config, r *HTTPRequest) (responseBody []byte, statusCode int, err error) {
	req, err := http.NewRequest(r.Method, *r.URL, r.Data)
	if err != nil {
		return
	}

	// Add authentication headers
	token := fmt.Sprintf("Token token=%s", c.APIKey)
	req.Header.Set("Authorization", token)
	req.Header.Set("Accept", "application/vnd.pagerduty+json;version=2")
	req.Header.Set("Content-type", "application/json")
	clientHeader := fmt.Sprintf("\"%s\" <%s>", c.Client, c.ClientURL)
	req.Header.Set("X-PagerDuty-Client", clientHeader)
	if c.Requester != "" {
		req.Header.Set("From", c.Requester)
	}

	// Fire request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	// Only 200 (ok) and 201 (created) are allowed
	responseBody, err = ioutil.ReadAll(response.Body)
	statusCode = response.StatusCode
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		message := fmt.Sprintf("Got HTTP status code %s%s", response.Status, parseResponseBody(responseBody))
		err = errors.New(message)
		return
	}

	return
}

func retryHTTPRequest(c *Config, r *HTTPRequest, maxRetries int, delayTime time.Duration) (body []byte, statusCode int, err error) {
	attempt := 1
	for {
		body, statusCode, err = httpRequest(c, r)
		if (err == nil) || (attempt >= maxRetries && err != nil) {
			return
		}
		time.Sleep(delayTime * time.Second)
		attempt++
	}
}

func parseResponseBody(responseBody []byte) string {
	var jsonMap map[string]interface{}
	if json.Unmarshal(responseBody, &jsonMap) != nil {
		return ""
	}
	errEntry := jsonMap["error"].(map[string]interface{})
	return fmt.Sprintf(": %s", errEntry["message"].(string))
}

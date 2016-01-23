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

import "errors"
import "fmt"
import "io"
import "io/ioutil"
import "net/http"

type HttpRequest struct {
	Method string
	Url    *string
	Data   io.Reader
}

func httpRequest(c *Config, r *HttpRequest) (response_body []byte, err error) {
	req, err := http.NewRequest(r.Method, *r.Url, r.Data)
	if err != nil {
		return
	}

	// Add authentication headers
	token := fmt.Sprintf("Token token=%s", c.ApiKey)
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-type", "application/json")

	// Fire request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	// Only 200's allowed
	if response.Status != "200 OK" {
		message := fmt.Sprintf("Got HTTP status code %s", response.Status)
		err = errors.New(message)
		return
	}

	response_body, err = ioutil.ReadAll(response.Body)
	return
}

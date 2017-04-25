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
	"errors"
	"flag"
	"fmt"
	"time"
)

type Config struct {
	APIKey             string
	IncidentKey        string
	IncidentID         string
	Description        string
	Details            string
	Mode               string
	Type               string
	Status             string
	ServiceID          string
	ServiceName        string
	EscalationPolicyID string
	Requester          string
	Assignee           string
	UserEmail          string
	Retry              int
	DelayTime          int
	Client             string
	ClientURL          string
}

func (c *Config) apiEndpoint() (url string) {
	url = "https://api.pagerduty.com"
	return url
}

func createGetIDConfig(args []string) (c *Config, err error) {
	c = &Config{}
	flags := flag.NewFlagSet("get-id", flag.ExitOnError)
	flags.StringVar(&c.APIKey, "api-key", "", "Your PagerDuty API key (required)")
	flags.StringVar(&c.IncidentKey, "incident-key", "", "The PagerDuty incident key (required)")
	flags.IntVar(&c.DelayTime, "delay-time", 0, "The Delay Time (in seconds) for Retry")
	err = flags.Parse(args)
	if err != nil {
		return
	}

	if c.DelayTime > 0 {
		time.Sleep(time.Duration(c.DelayTime) * time.Second)
	}

	c.Mode = "get-id"
	required := map[string]*string{
		"ap-key":       &c.APIKey,
		"incident-key": &c.IncidentKey,
	}
	err = checkConfig(c, required)
	return
}

func createGetValidateIdConfig(args []string) (c *Config, err error) {
	c = &Config{}
	flags := flag.NewFlagSet("validate-escalation-policy-id", flag.ExitOnError)
	flags.StringVar(&c.APIKey, "api-key", "", "Your PagerDuty API key (required)")
	flags.StringVar(&c.IncidentID, "incident-id", "", "The PagerDuty Incident ID")
	err = flags.Parse(args)
	if err != nil {
		return
	}

	c.Mode = "get-validate-id"
	required := map[string]*string{
		"ap-key":      &c.APIKey,
		"incident-id": &c.IncidentID,
	}
	err = checkConfig(c, required)
	return
}

func createGetServiceIDConfig(args []string) (c *Config, err error) {
	c = &Config{}
	flags := flag.NewFlagSet("get-service-id", flag.ExitOnError)
	flags.StringVar(&c.APIKey, "api-key", "", "Your PagerDuty API key (required)")
	flags.StringVar(&c.ServiceName, "service-name", "", "The PagerDuty Service Nanme (required)")
	err = flags.Parse(args)
	if err != nil {
		return
	}

	c.Mode = "get-service-id"
	required := map[string]*string{
		"ap-key":       &c.APIKey,
		"service-name": &c.ServiceName,
	}
	err = checkConfig(c, required)
	return
}

func createGetServiceEscalationIDConfig(args []string) (c *Config, err error) {
	c = &Config{}
	flags := flag.NewFlagSet("get-service-escalation-id", flag.ExitOnError)
	flags.StringVar(&c.APIKey, "api-key", "", "Your PagerDuty API key (required)")
	flags.StringVar(&c.ServiceID, "service-id", "", "The PagerDuty Service ID (required)")
	err = flags.Parse(args)
	if err != nil {
		return
	}

	c.Mode = "get-service-escalation-id"
	required := map[string]*string{
		"ap-key":     &c.APIKey,
		"service-id": &c.ServiceID,
	}
	err = checkConfig(c, required)
	return
}

func createGetUserIDConfig(args []string) (c *Config, err error) {
	c = &Config{}
	flags := flag.NewFlagSet("get-user-id", flag.ExitOnError)
	flags.StringVar(&c.APIKey, "api-key", "", "Your PagerDuty API key (required)")
	flags.StringVar(&c.UserEmail, "user-email", "", "The PagerDuty user email (required)")
	err = flags.Parse(args)
	if err != nil {
		return
	}

	c.Mode = "get-user-id"
	required := map[string]*string{
		"ap-key":     &c.APIKey,
		"user-email": &c.UserEmail,
	}
	err = checkConfig(c, required)
	return
}

func createValidateEscalationPolicyIDConfig(args []string) (c *Config, err error) {
	c = &Config{}
	flags := flag.NewFlagSet("validate-escalation-policy-id", flag.ExitOnError)
	flags.StringVar(&c.APIKey, "api-key", "", "Your PagerDuty API key (required)")
	flags.StringVar(&c.EscalationPolicyID, "escalation-policy-id", "", "The PagerDuty Escalation Policy ID (required)")
	err = flags.Parse(args)
	if err != nil {
		return
	}

	c.Mode = "validate-escalation-policy-id"
	required := map[string]*string{
		"ap-key":               &c.APIKey,
		"escalation-policy-id": &c.EscalationPolicyID,
	}
	err = checkConfig(c, required)
	return
}

func createGetIlesConfig(args []string) (c *Config, err error) {
	c = &Config{}
	flags := flag.NewFlagSet("get-iles", flag.ExitOnError)
	flags.StringVar(&c.APIKey, "api-key", "", "Your PagerDuty API key (required)")
	flags.StringVar(&c.IncidentID, "incident-id", "", "The PagerDuty Incident ID")
	err = flags.Parse(args)
	if err != nil {
		return
	}

	c.Mode = "get-iles"
	required := map[string]*string{
		"api-key":     &c.APIKey,
		"incident-id": &c.IncidentID,
	}
	err = checkConfig(c, required)
	return
}

func createTestConnectionConfig(args []string) (c *Config, err error) {
	c = &Config{}
	flags := flag.NewFlagSet("test-connection", flag.ExitOnError)
	flags.StringVar(&c.APIKey, "api-key", "", "Your PagerDuty API key (required)")
	err = flags.Parse(args)
	if err != nil {
		return
	}

	c.Mode = "test-connection"
	required := map[string]*string{
		"api-key": &c.APIKey,
	}
	err = checkConfig(c, required)
	return
}

func createTriggerAckResolveChangePolicyConfig(mode string, args []string) (c *Config, err error) {
	c = &Config{}
	flags := flag.NewFlagSet("trigger", flag.ExitOnError)
	flags.StringVar(&c.APIKey, "api-key", "", "Your PagerDuty API key (required)")
	flags.StringVar(&c.IncidentID, "incident-id", "", "The PagerDuty Incident ID (required)")
	flags.StringVar(&c.IncidentKey, "incident-key", "", "The PagerDuty Incident key")
	flags.StringVar(&c.Description, "description", "", "A PagerDuty Incident description (required)")
	flags.StringVar(&c.Details, "details", "", "PagerDuty Incident details")
	flags.StringVar(&c.ServiceID, "service-id", "", "PagerDuty Incident Service ID")
	flags.StringVar(&c.EscalationPolicyID, "escalation-policy-id", "", "PagerDuty Escalation Policy iD")
	flags.StringVar(&c.Requester, "requester", "", "PagerDuty Requester email address")
	flags.StringVar(&c.Assignee, "assignee", "", "The PagerDuty Assignee User ID")
	flags.IntVar(&c.Retry, "retry", 0, "The Number of Times to Retry")
	flags.IntVar(&c.DelayTime, "delay-time", 0, "The Delay Time (in seconds) for Retry")
	flags.StringVar(&c.Client, "client", "BMC Remedy", "The client information defaults to BMC Remedy")
	flags.StringVar(&c.ClientURL, "client-url", "https://www.bmc.com", "The client url")
	err = flags.Parse(args)
	if err != nil {
		return
	}

	switch mode {
	case "resolve":
		{
			c.Mode = "resolve"
			c.Status = "resolved"
		}
	case "acknowledge":
		{
			c.Mode = "acknowledge"
			c.Status = "acknowledged"
		}
	default:
		{
			c.Mode = "trigger"
			c.Status = "triggered"
		}
	}

	if c.Retry == 0 && c.DelayTime > 0 {
		time.Sleep(time.Duration(c.DelayTime) * time.Second)
	}

	if c.IncidentID == "" {
		// Find the incident id using incident key
		cGetID := &Config{}
		cGetID.Mode = "get-id"
		cGetID.APIKey = c.APIKey
		cGetID.IncidentKey = c.IncidentKey
		required := map[string]*string{
			"ap-key":       &c.APIKey,
			"incident-key": &c.IncidentKey,
		}
		err = checkConfig(cGetID, required)
		if err != nil {
			return
		}
		var status string
		if c.Retry > 0 {
			// Introduce retries in case there is a timing issue on the 2-call process
			c.IncidentID, status = getIDRetry(cGetID, c.Retry, c.DelayTime)
		} else {
			c.IncidentID, status = getIDStatus(cGetID)
		}
		if c.Status != status && status == "resolved" {
			c.IncidentID = ""
		}
	} else {
		status := getIncidentStatus(c)
		if status == c.Status {
			// Incident status has not changed
			c.Status = ""
		} else {
			// For reopen, if the current incident is already resolved and
			// the subcommand specifies trigger or acknowledge mode, then
			// it will create a PD incident
			if status == "resolved" {
				c.IncidentID = ""
			}
		}
	}

	if c.IncidentID != "" {
		// Update the PD incident
		c.Type = "incident_reference"
		required := map[string]*string{
			"api-key":     &c.APIKey,
			"incident-id": &c.IncidentID,
			"requester":   &c.Requester,
		}
		err = checkConfig(c, required)
	} else {
		// Create the PD incident
		c.Type = "incident"
		required := map[string]*string{
			"api-key":      &c.APIKey,
			"incident-key": &c.IncidentKey,
			"description":  &c.Description,
			"requester":    &c.Requester,
		}
		err = checkConfig(c, required)
	}
	return
}

// Check for empty string amongst required args
func checkConfig(c *Config, required map[string]*string) (err error) {
	for k, v := range required {
		if *v == "" {
			msg := fmt.Sprintf("Argument error: %s is required for the %s subcommand", k, c.Mode)
			err = errors.New(msg)
		}
	}

	return
}

package main

import "testing"

func TestCreateGetIlesConfig(t *testing.T) {
	args := []string{
		"-api-key",
		"NOT_A_REAL_KEY",
		"-incident-id",
		"MY_INCIDENT_ID",
	}

	_, err := createGetIlesConfig(args)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateGetIdConfig(t *testing.T) {
	args := []string{
		"-api-key",
		"NOT_A_REAL_KEY",
		"-incident-key",
		"MY_INCIDENT_KEY",
	}

	_, err := createGetIDConfig(args)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateGetServiceIdConfig(t *testing.T) {
	args := []string{
		"-api-key",
		"NOT_A_REAL_KEY",
		"-service-name",
		"MY_SERVICE_NAME",
	}

	_, err := createGetServiceIDConfig(args)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateGetServiceEscalationPolicyIdConfig(t *testing.T) {
	args := []string{
		"-api-key",
		"NOT_A_REAL_KEY",
		"-service-id",
		"MY_SERVICE_ID",
	}

	_, err := createGetServiceEscalationIDConfig(args)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateValidateEscalationPolicyIdConfig(t *testing.T) {
	args := []string{
		"-api-key",
		"NOT_A_REAL_KEY",
		"-escalation-policy-id",
		"MY_ESCALATION_POLICY_ID",
	}

	_, err := createValidateEscalationPolicyIDConfig(args)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateGetUserIdConfig(t *testing.T) {
	args := []string{
		"-api-key",
		"NOT_A_REAL_KEY",
		"-user-email",
		"MY_USER_EMAIL",
	}

	_, err := createGetUserIDConfig(args)
	if err != nil {
		t.Error(err)
	}
}

func TestConfigCheckConfig(t *testing.T) {
	args := []string{
		"-api-key",
		"NOT_A_REAL_KEY",
		"-incident-key",
		"MY_INCIDENT_KEY",
	}

	c, err := createGetIDConfig(args)
	if err != nil {
		t.Error(err)
	}

	required := map[string]*string{
		"api-key":      &c.APIKey,
		"incident-key": &c.IncidentKey,
	}

	// Should pass w/o error
	err = checkConfig(c, required)
	if err != nil {
		t.Error(err)
	}

	// Should return an error
	c.IncidentKey = ""
	err = checkConfig(c, required)
	if err == nil {
		t.Error(err)
	}
}

func TestConfigApiEndpoint(t *testing.T) {
	args := []string{
		"-api-key",
		"NOT_A_REAL_KEY",
		"-incident-id",
		"MY_INCIDENT_ID",
	}

	c, err := createGetIlesConfig(args)
	if err != nil {
		t.Error(err)
	}

	endpoint := c.apiEndpoint()
	if endpoint != "https://api.pagerduty.com" {
		t.Error("*Config.ApiEndpoint failure: %s", endpoint)
	}
}

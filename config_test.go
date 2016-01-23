package main

import "testing"

func TestCreateResolveConfig(t *testing.T) {
	args := []string{
		"-service-key",
		"NOT_A_REAL_KEY",
		"-incident-key",
		"MY_INCIDENT_KEY",
		"-description",
		"My Description",
	}

	_, err := createResolveConfig(&args)
	if err != nil {
		t.Error(err)
	}

	// Try w/o optional params
	args = []string{
		"-service-key",
		"NOT_A_REAL_KEY",
		"-incident-key",
		"NOT_A_REAL_KEY",
	}
	_, err = createResolveConfig(&args)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateTriggerConfig(t *testing.T) {
	args := []string{
		"-service-key",
		"NOT_A_REAL_KEY",
		"-incident-key",
		"MY_INCIDENT_KEY",
		"-description",
		"My Description",
		"-details",
		"My Details",
	}

	_, err := createTriggerConfig(&args)
	if err != nil {
		t.Error(err)
	}

	// Try w/o optional params
	args = []string{
		"-service-key",
		"NOT_A_REAL_KEY",
		"-description",
		"My Description",
	}
	_, err = createTriggerConfig(&args)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateGetIlesConfig(t *testing.T) {
	args := []string{
		"-api-key",
		"NOT_A_REAL_KEY",
		"-subdomain",
		"foo",
		"-incident-id",
		"MY_INCIDENT_ID",
	}

	_, err := createGetIlesConfig(&args)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateGetIdConfig(t *testing.T) {
	args := []string{
		"-api-key",
		"NOT_A_REAL_KEY",
		"-subdomain",
		"foo",
		"-incident-key",
		"MY_INCIDENT_KEY",
	}

	_, err := createGetIDConfig(&args)
	if err != nil {
		t.Error(err)
	}
}

func TestConfigCheckConfig(t *testing.T) {
	args := []string{
		"-api-key",
		"NOT_A_REAL_KEY",
		"-subdomain",
		"foo",
		"-incident-key",
		"MY_INCIDENT_KEY",
	}

	c, err := createGetIDConfig(&args)
	if err != nil {
		t.Error(err)
	}

	required := map[string]*string{
		"api-key":      &c.APIKey,
		"subdomain":    &c.Subdomain,
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
		"-subdomain",
		"foo",
		"-incident-id",
		"MY_INCIDENT_ID",
	}

	c, err := createGetIlesConfig(&args)
	if err != nil {
		t.Error(err)
	}

	endpoint := c.apiEndpoint()
	if endpoint != "https://foo.pagerduty.com/api/v1" {
		t.Error("*Config.ApiEndpoint failure: %s", endpoint)
	}
}

package main

import "errors"
import "flag"
import "fmt"

type Config struct {
  ApiKey      string
  Subdomain   string
  ServiceKey  string
  IncidentKey string
  IncidentId  string
  Description string
  Details     string
  Mode        string
}

func (c *Config) apiEndpoint() (url string) {
  url = fmt.Sprintf("https://%s.pagerduty.com/api/v1", c.Subdomain)
  return url
}

func createGetIdConfig(args *[]string) (c *Config, err error) {
  c = &Config{}
  flags := flag.NewFlagSet("get-id", flag.ExitOnError)
  flags.StringVar(&c.ApiKey, "api-key", "", "Your PagerDuty API key (required)")
  flags.StringVar(&c.Subdomain, "subdomain", "", "Your PagerDuty subdomain (required)")
  flags.StringVar(&c.IncidentKey, "incident-key", "", "The PagerDuty incident key (required)")
  err = flags.Parse(*args)
  if err != nil {
    return
  }

  c.Mode = "get-id"
  required := map[string]*string{
    "ap-key": &c.ApiKey,
    "subdomain": &c.Subdomain,
    "incident-key": &c.IncidentKey,
  }
  err = checkConfig(c, required)
  return
}

func createGetIlesConfig(args *[]string) (c *Config, err error) {
  c = &Config{}
  flags := flag.NewFlagSet("get-iles", flag.ExitOnError)
  flags.StringVar(&c.ApiKey, "api-key", "", "Your PagerDuty API key (required)")
  flags.StringVar(&c.Subdomain, "subdomain", "", "Your PagerDuty subdomain (required)")
  flags.StringVar(&c.IncidentId, "incident-id", "", "The PagerDuty Incident ID")
  err = flags.Parse(*args)
  if err != nil {
    return
  }

  c.Mode = "get-iles"
  required := map[string]*string{
    "api-key": &c.ApiKey,
    "subdomain": &c.Subdomain,
    "incident-id": &c.IncidentId,
  }
  err = checkConfig(c, required)
  return
}

func createTriggerConfig(args *[]string) (c *Config, err error) {
  c = &Config{}
  flags := flag.NewFlagSet("trigger", flag.ExitOnError)
  flags.StringVar(&c.ServiceKey, "service-key", "", "Your PagerDuty Service key (required)")
  flags.StringVar(&c.IncidentKey, "incident-key", "", "A PagerDuty Incident key")
  flags.StringVar(&c.Description, "description", "", "A PagerDuty Incident description (required)")
  flags.StringVar(&c.Details, "details", "", "PagerDuty Incident details")
  err = flags.Parse(*args)
  if err != nil {
    return
  }

  c.Mode = "trigger"
  required := map[string]*string{
    "service-key": &c.ServiceKey,
    "description": &c.Description,
  }
  err = checkConfig(c, required)
  return
}

func createResolveConfig(args *[]string) (c *Config, err error) {
  c = &Config{}
  flags := flag.NewFlagSet("resolve", flag.ExitOnError)
  flags.StringVar(&c.ServiceKey, "service-key", "", "Your PagerDuty Service key (required)")
  flags.StringVar(&c.IncidentKey, "incident-key", "", "The PagerDuty Incident key (required)")
  flags.StringVar(&c.Description, "description", "", "A description")
  err = flags.Parse(*args)
  if err != nil {
    return
  }

  c.Mode = "resolve"
  required := map[string]*string{
    "service-key": &c.ServiceKey,
    "incident-key": &c.IncidentKey,
  }
  err = checkConfig(c, required)
  return
}

// Check for empty string amongst required args
func checkConfig(c *Config, required map[string]*string) (err error) {
  for k, v := range required {
    if *v == "" {
      msg := fmt.Sprintf("Argument error: %s is required for the %s subommand", k, c.Mode)
      err = errors.New(msg)
    }
  }

  return
}

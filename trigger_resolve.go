package main

import "bytes"
import "encoding/json"

type Event struct {
  ServiceKey  string  `json:"service_key"`
  IncidentKey string  `json:"incident_key"`
  EventType   string  `json:"event_type"`
  Description string  `json:"description"`
  Details     string  `json:"details"`
}

func triggerResolve(c *Config) {
  event := &Event{}
  eventEndpoint := "https://events.pagerduty.com/generic/2010-04-15/create_event.json"

  event.EventType = c.Mode
  event.ServiceKey = c.ServiceKey
  event.IncidentKey = c.IncidentKey
  event.Description = c.Description
  event.Details = c.Details

  payload, err := json.Marshal(event)
  failIf(err)

  req := &HttpRequest{
    Method: "POST",
    Url: &eventEndpoint,
    Data: bytes.NewBuffer(payload),
  }

  _, err = httpRequest(c, req)
  failIf(err)
}

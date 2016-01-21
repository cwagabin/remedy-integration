package main

import "encoding/json"
import "fmt"

func getId(c *Config) string {
  req := &HttpRequest{}
  url := fmt.Sprintf("%s/incidents?incident_key=%s", c.apiEndpoint(), c.IncidentKey)
  req.Url = &url
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
  first_incident := incidents[0].(map[string]interface{})
  return first_incident["id"].(string)
}

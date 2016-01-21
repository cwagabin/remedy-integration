package main

import "encoding/json"
import "fmt"

type ILE struct {
  Type             string
  AgentName        string
  AgentType        string
  ChannelType      string
  NotificationType string
  UserName         string
  AssignedUser     string
}

func (i *ILE) FriendlyMessage() *string {
  var agentName string
  var channelType string
  var friendlyMessage string

  channelTypes := map[string]string{
    "sms":                       "SMS",
    "phone":                     "phone",
    "email":                     "email",
    "short_email":               "email",
    "website":                   "website",
    "ios_push_notification":     "push notification",
    "android_push_notification": "push notification",
  }

  if i.AgentType == "service" {
    agentName = "API"
  } else {
    agentName = i.AgentName
  }

  switch i.Type {
  case "acknowledge":
    {
      friendlyMessage = fmt.Sprintf("Acknowledged by %s", agentName)
    }
  case "resolve":
    {
      friendlyMessage = fmt.Sprintf("Resolved by %s", agentName)
    }
  case "trigger":
    {
      friendlyMessage = fmt.Sprintf("Triggered by %s", agentName)
    }
  case "notify":
    {
      friendlyMessage = fmt.Sprintf("Notified %s", i.UserName)
    }
  case "assign":
    {
      friendlyMessage = fmt.Sprintf("Assigned to %s", i.AssignedUser)
    }
  case "escalate":
    {
      friendlyMessage = fmt.Sprintf("Escalated to %s", i.AssignedUser)
    }
  }

  if i.ChannelType != "" {
    channelType = i.ChannelType
  } else {
    channelType = i.NotificationType
  }

  if val, ok := channelTypes[channelType]; ok {
    friendlyMessage = fmt.Sprintf("%s via %s", friendlyMessage, val)
  }

  return &friendlyMessage
}

// Returns an empty map if the JSON key is missing
// Allows us to chase down empty paths and only check for nil at the end
func mapOrNewMap(m map[string]interface{}, k string) (map[string]interface{}) {
  if val, ok := m[k]; ok {
    return val.(map[string]interface{})
  } else {
    return make(map[string]interface{})
  }
}

// Returns an empty string if the JSON key is missing
// Allows us to chase down empty paths and only check for nil at the end
func stringOrNewString(m map[string]interface{}, k string) (string) {
  if val, ok := m[k]; ok {
    return val.(string)
  } else {
    return ""
  }
}

func parseIles(res []byte) ([]*ILE, error) {
  var jsonMap map[string]interface{}
  var iles []*ILE

  if err := json.Unmarshal(res, &jsonMap); err != nil {
    return iles, err
  }

  rawILEs := jsonMap["log_entries"].([]interface{})
  for _, rawILE := range rawILEs {
    ile := rawILE.(map[string]interface{})

    ileAgent := mapOrNewMap(ile, "agent")
    ileUser := mapOrNewMap(ile, "user")
    ileAssignedUser := mapOrNewMap(ile, "assigned_user")
    ileChannel := mapOrNewMap(ile, "channel")
    ileNotification := mapOrNewMap(ile, "notification")

    currentILE := &ILE {
      Type: stringOrNewString(ile, "type"),
      AgentName: stringOrNewString(ileAgent, "name"),
      AgentType: stringOrNewString(ileAgent, "type"),
      ChannelType: stringOrNewString(ileChannel, "type"),
      NotificationType: stringOrNewString(ileNotification, "type"),
      UserName: stringOrNewString(ileUser, "name"),
      AssignedUser: stringOrNewString(ileAssignedUser, "name"),
    }

    iles = append(iles, currentILE)
  }

  return iles, nil
}

func getIles(c *Config) (messages []string) {
  url := fmt.Sprintf("%s/incidents/%s/log_entries", c.apiEndpoint(), c.IncidentId)
  req := &HttpRequest{
    Method: "GET",
    Url: &url,
    Data: nil,
  }
  res, err := httpRequest(c, req)
  failIf(err)

  ileSlice, err := parseIles(res)
  failIf(err)

  for _, ile := range ileSlice {
    messages = append(messages, *ile.FriendlyMessage())
  }

  return
}

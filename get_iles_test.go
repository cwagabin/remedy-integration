package main

import (
	"fmt"
	"testing"
)

func TestGetIlesmapOrNewMap(t *testing.T) {
	populatedMapData := make(map[string]interface{})
	populatedMapData["bar"] = "baz"
	populatedMap := make(map[string]interface{})
	populatedMap["foo"] = populatedMapData

	result := mapOrNewMap(populatedMap, "foo")
	if len(result) < 1 {
		msg := fmt.Sprintf("mapOrNewMap failed to extract valid key from map. Got: %v", result)
		t.Error(msg)
	}

	emptyMap := make(map[string]interface{})
	result = mapOrNewMap(emptyMap, "foo")
	if len(result) != 0 {
		msg := fmt.Sprintf("mapOrNewMap failed to return a new empty map. Got: %v", result)
		t.Error(msg)
	}
}

func TestGetIlesstringOrNewString(t *testing.T) {
	populatedMap := make(map[string]interface{})
	populatedMap["foo"] = "bar"
	emptyMap := make(map[string]interface{})

	result := stringOrNewString(populatedMap, "foo")
	if result != "bar" {
		msg := fmt.Sprintf("stringOrNewString failed to extract valid string from map. Got: %v", result)
		t.Error(msg)
	}

	result = stringOrNewString(emptyMap, "foo")
	if result != "" {
		msg := fmt.Sprintf("stringOrNewString failed to return a new empty string. Got: %v", result)
		t.Error(msg)
	}
}

func TestGetIlesParseILEs(t *testing.T) {
	response := `{
    "log_entries": [
      {
        "id": "PPV5KG7",
        "type": "resolve",
        "created_at": "2013-03-06T21:33:15Z",
        "note": "Fixed",
        "agent": {
          "id": "PT23IWX",
          "name": "Tim Wright",
          "email": "tim@acme.com",
          "time_zone": "Eastern Time (US & Canada)",
          "color": "purple",
          "role": "owner",
          "avatar_url": "https://secure.gravatar.com/avatar/923a2b907dc04244e9bb5576a42e70a7.png?d=mm&r=PG",
          "user_url": "/users/PT23IWX",
          "invitation_sent": false,
          "marketing_opt_out": false,
          "type": "user"
        },
        "channel": {
          "type": "website"
        }
      },
      {
        "id": "PVALU1K",
        "type": "notify",
        "created_at": "2013-03-06T21:08:43Z",
        "user": {
          "id": "PPSFHH7",
          "name": "Bob Smith",
          "email": "bob@acme.com",
          "time_zone": "Eastern Time (US & Canada)",
          "color": "red",
          "role": "user",
          "avatar_url": "https://secure.gravatar.com/avatar/78e9fa7d6d2ddba416ad7534eb1403d0.png?d=mm&r=PG",
          "user_url": "/users/PPSFHH7",
          "invitation_sent": true,
          "marketing_opt_out": false
        },
        "notification": {
          "type": "email",
          "address": "bob@acme.com",
          "status": "success"
        }
      },
      {
        "id": "P8MDTZ8",
        "type": "notify",
        "created_at": "2013-03-06T21:00:30Z",
        "user": {
          "id": "PPSFHH7",
          "name": "Bob Smith",
          "email": "bob@acme.com",
          "time_zone": "Eastern Time (US & Canada)",
          "color": "red",
          "role": "user",
          "avatar_url": "https://secure.gravatar.com/avatar/78e9fa7d6d2ddba416ad7534eb1403d0.png?d=mm&r=PG",
          "user_url": "/users/PPSFHH7",
          "invitation_sent": true,
          "marketing_opt_out": false
        },
        "notification": {
          "type": "email",
          "address": "bob@acme.com",
          "status": "success"
        }
      },
      {
        "id": "PHCXXHR",
        "type": "escalate",
        "created_at": "2013-03-06T21:00:30Z",
        "agent": {
          "type": "service"
        },
        "channel": {
          "type": "timeout"
        },
        "assigned_user": {
          "id": "PPSFHH7",
          "name": "Bob Smith",
          "email": "bob@acme.com",
          "time_zone": "Eastern Time (US & Canada)",
          "color": "red",
          "role": "user",
          "avatar_url": "https://secure.gravatar.com/avatar/78e9fa7d6d2ddba416ad7534eb1403d0.png?d=mm&r=PG",
          "user_url": "/users/PPSFHH7",
          "invitation_sent": true,
          "marketing_opt_out": false
        },
        "note": null
      },
      {
        "id": "PE5XGDA",
        "type": "notify",
        "created_at": "2013-03-06T20:59:31Z",
        "user": {
          "id": "PT23IWX",
          "name": "Tim Wright",
          "email": "tim@acme.com",
          "time_zone": "Eastern Time (US & Canada)",
          "color": "purple",
          "role": "owner",
          "avatar_url": "https://secure.gravatar.com/avatar/923a2b907dc04244e9bb5576a42e70a7.png?d=mm&r=PG",
          "user_url": "/users/PT23IWX",
          "invitation_sent": false,
          "marketing_opt_out": false
        },
        "notification": {
          "type": "sms",
          "address": "+1 956-821-0372",
          "status": "success"
        }
      },
      {
        "id": "P9DWJ9J",
        "type": "notify",
        "created_at": "2013-03-06T20:59:30Z",
        "user": {
          "id": "PT23IWX",
          "name": "Tim Wright",
          "email": "tim@acme.com",
          "time_zone": "Eastern Time (US & Canada)",
          "color": "purple",
          "role": "owner",
          "avatar_url": "https://secure.gravatar.com/avatar/923a2b907dc04244e9bb5576a42e70a7.png?d=mm&r=PG",
          "user_url": "/users/PT23IWX",
          "invitation_sent": false,
          "marketing_opt_out": false
        },
        "notification": {
          "type": "email",
          "address": "tim@pagerduty.com",
          "status": "success"
        }
      },
      {
        "id": "P36T0K8",
        "type": "unacknowledge",
        "created_at": "2013-03-06T20:59:30Z",
        "agent": {
          "type": "service"
        },
        "channel": {
          "type": "timeout"
        }
      },
      {
        "id": "PKU5VMR",
        "type": "acknowledge",
        "created_at": "2013-03-06T20:29:30Z",
        "agent": {
          "id": "PT23IWX",
          "name": "Tim Wright",
          "email": "tim@acme.com",
          "time_zone": "Eastern Time (US & Canada)",
          "color": "purple",
          "role": "owner",
          "avatar_url": "https://secure.gravatar.com/avatar/923a2b907dc04244e9bb5576a42e70a7.png?d=mm&r=PG",
          "user_url": "/users/PT23IWX",
          "invitation_sent": false,
          "marketing_opt_out": false,
          "type": "user"
        },
        "channel": {
          "type": "phone"
        }
      },
      {
        "id": "P36HB17",
        "type": "notify",
        "created_at": "2013-03-06T20:28:51Z",
        "user": {
          "id": "PT23IWX",
          "name": "Tim Wright",
          "email": "tim@acme.com",
          "time_zone": "Eastern Time (US & Canada)",
          "color": "purple",
          "role": "owner",
          "avatar_url": "https://secure.gravatar.com/avatar/923a2b907dc04244e9bb5576a42e70a7.png?d=mm&r=PG",
          "user_url": "/users/PT23IWX",
          "invitation_sent": false,
          "marketing_opt_out": false
        },
        "notification": {
          "type": "email",
          "address": "tim@pagerduty.com",
          "status": "success"
        }
      },
      {
        "id": "PBJBF9Q",
        "type": "notify",
        "created_at": "2013-03-06T20:28:50Z",
        "user": {
          "id": "PT23IWX",
          "name": "Tim Wright",
          "email": "tim@acme.com",
          "time_zone": "Eastern Time (US & Canada)",
          "color": "purple",
          "role": "owner",
          "avatar_url": "https://secure.gravatar.com/avatar/923a2b907dc04244e9bb5576a42e70a7.png?d=mm&r=PG",
          "user_url": "/users/PT23IWX",
          "invitation_sent": false,
          "marketing_opt_out": false
        },
        "notification": {
          "type": "phone",
          "address": "+1 956-821-0372",
          "status": "success"
        }
      },
      {
        "id": "PRD9J7P",
        "type": "assign",
        "created_at": "2013-03-06T20:28:46Z",
        "agent": {
          "type": "service"
        },
        "channel": {
          "type": "auto"
        },
        "assigned_user": {
          "id": "PT23IWX",
          "name": "Tim Wright",
          "email": "tim@acme.com",
          "time_zone": "Eastern Time (US & Canada)",
          "color": "purple",
          "role": "owner",
          "avatar_url": "https://secure.gravatar.com/avatar/923a2b907dc04244e9bb5576a42e70a7.png?d=mm&r=PG",
          "user_url": "/users/PT23IWX",
          "invitation_sent": false,
          "marketing_opt_out": false
        },
        "note": null
      },
      {
        "id": "PVPXJJC",
        "type": "trigger",
        "created_at": "2013-03-06T20:28:46Z",
        "agent": {
          "id": "PT23IWX",
          "name": "Tim Wright",
          "email": "tim@acme.com",
          "time_zone": "Eastern Time (US & Canada)",
          "color": "purple",
          "role": "owner",
          "avatar_url": "https://secure.gravatar.com/avatar/923a2b907dc04244e9bb5576a42e70a7.png?d=mm&r=PG",
          "user_url": "/users/PT23IWX",
          "invitation_sent": false,
          "marketing_opt_out": false,
          "type": "user"
        },
        "channel": {
          "summary": "Martian's are attacking",
          "type": "web_trigger"
        }
      }
    ],
    "limit": 100,
    "offset": 0,
    "total": 12
  }`

	var err error
	var expectedIle ILE

	ileSlice, err := parseIles([]byte(response))
	if err != nil {
		msg := fmt.Sprintf("Encountered an error while parsing ILE response: %v", err)
		t.Error(msg)
	}

	if len(ileSlice) != 12 {
		msg := fmt.Sprintf("Unexpected number of ILEs in parsed response. Got: %v", len(ileSlice))
		t.Error(msg)
	}

	// Test parsing of various ILE types
	resolveIle := ileSlice[0]
	friendlyMessage := resolveIle.FriendlyMessage()
	expectedIle.Type = "resolve"
	expectedIle.AgentName = "Tim Wright"
	expectedIle.AgentType = "user"
	expectedIle.ChannelType = "website"
	expectedIle.NotificationType = ""
	expectedIle.UserName = ""
	expectedIle.AssignedUser = ""
	if *resolveIle != expectedIle {
		msg := fmt.Sprintf("Resolve ILE parsed incorrectly. Got: %v", resolveIle)
		t.Error(msg)
	}
	if *friendlyMessage != "Resolved by Tim Wright via website" {
		msg := fmt.Sprintf("Error generating friendly resolve message. Got: %v", *friendlyMessage)
		t.Error(msg)
	}

	notifyIle := ileSlice[1]
	friendlyMessage = notifyIle.FriendlyMessage()
	expectedIle.Type = "notify"
	expectedIle.AgentName = ""
	expectedIle.AgentType = ""
	expectedIle.ChannelType = ""
	expectedIle.NotificationType = "email"
	expectedIle.UserName = "Bob Smith"
	expectedIle.AssignedUser = ""
	if *notifyIle != expectedIle {
		msg := fmt.Sprintf("Notify ILE parsed incorrectly. Got: %v", notifyIle)
		t.Error(msg)
	}
	if *friendlyMessage != "Notified Bob Smith via email" {
		msg := fmt.Sprintf("Error generating friendly notify message. Got: %v", *friendlyMessage)
		t.Error(msg)
	}

	escalateIle := ileSlice[3]
	friendlyMessage = escalateIle.FriendlyMessage()
	expectedIle.Type = "escalate"
	expectedIle.AgentName = ""
	expectedIle.AgentType = "service"
	expectedIle.ChannelType = "timeout"
	expectedIle.NotificationType = ""
	expectedIle.UserName = ""
	expectedIle.AssignedUser = "Bob Smith"
	if *escalateIle != expectedIle {
		msg := fmt.Sprintf("Escalate ILE parsed incorrectly. Got: %v", escalateIle)
		t.Error(msg)
	}
	if *friendlyMessage != "Escalated to Bob Smith" {
		msg := fmt.Sprintf("Error generating friendly escalate message. Got: %v", *friendlyMessage)
		t.Error(msg)
	}

	// No friendly message for unacks
	unackIle := ileSlice[6]
	expectedIle.Type = "unacknowledge"
	expectedIle.AgentName = ""
	expectedIle.AgentType = "service"
	expectedIle.ChannelType = "timeout"
	expectedIle.NotificationType = ""
	expectedIle.UserName = ""
	expectedIle.AssignedUser = ""
	if *unackIle != expectedIle {
		msg := fmt.Sprintf("Unack ILE parsed incorrectly. Got: %v", unackIle)
		t.Error(msg)
	}

	ackIle := ileSlice[7]
	friendlyMessage = ackIle.FriendlyMessage()
	expectedIle.Type = "acknowledge"
	expectedIle.AgentName = "Tim Wright"
	expectedIle.AgentType = "user"
	expectedIle.ChannelType = "phone"
	expectedIle.NotificationType = ""
	expectedIle.UserName = ""
	expectedIle.AssignedUser = ""
	if *ackIle != expectedIle {
		msg := fmt.Sprintf("Ack ILE parsed incorrectly. Got: %v", ackIle)
		t.Error(msg)
	}
	if *friendlyMessage != "Acknowledged by Tim Wright via phone" {
		msg := fmt.Sprintf("Error generating friendly ack message. Got: %v", *friendlyMessage)
		t.Error(msg)
	}

	assignIle := ileSlice[10]
	friendlyMessage = assignIle.FriendlyMessage()
	expectedIle.Type = "assign"
	expectedIle.AgentName = ""
	expectedIle.AgentType = "service"
	expectedIle.ChannelType = "auto"
	expectedIle.NotificationType = ""
	expectedIle.UserName = ""
	expectedIle.AssignedUser = "Tim Wright"
	if *assignIle != expectedIle {
		msg := fmt.Sprintf("Assign ILE parsed incorrectly. Got: %v", assignIle)
		t.Error(msg)
	}
	if *friendlyMessage != "Assigned to Tim Wright" {
		msg := fmt.Sprintf("Error generating friendly assign message. Got %v", *friendlyMessage)
		t.Error(msg)
	}

	triggerIle := ileSlice[11]
	friendlyMessage = triggerIle.FriendlyMessage()
	expectedIle.Type = "trigger"
	expectedIle.AgentName = "Tim Wright"
	expectedIle.AgentType = "user"
	expectedIle.ChannelType = "web_trigger"
	expectedIle.NotificationType = ""
	expectedIle.UserName = ""
	expectedIle.AssignedUser = ""
	if *triggerIle != expectedIle {
		msg := fmt.Sprintf("Trigger ILE parsed incorrectly. Got: %v", triggerIle)
		t.Error(msg)
	}
	if *friendlyMessage != "Triggered by Tim Wright" {
		msg := fmt.Sprintf("Error generating friendly trigger message. Got: %v", *friendlyMessage)
		t.Error(msg)
	}
}

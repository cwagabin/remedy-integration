package main

import "fmt"
import "testing"

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
        "type": "resolve_log_entry",
        "summary": "Resolved by Tim Wright",
        "created_at": "2013-03-06T21:33:15Z",
        "self": "https://api.pagerduty.com/log_entries/PPV5KG7",
	"html_url": null,
        "note": "Fixed",
        "agent": {
          "id": "PT23IWX",
          "type": "user_reference",
          "summary": "Tim Wright",
	  "self": "https://api.pagerduty.com/users/PT23IWX",
	  "html_url": "https://effect-tech.pagerduty.com/users/PT23IWX"
        },
        "channel": {
          "type": "website"
        },
        "service": {
	  "id": "P4672E4",
	  "type": "service_reference",
	  "summary": "Test Service",
	  "self": "https://api.pagerduty.com/services/P4672E4",
	  "html_url": "https://effect-tech.pagerduty.com/services/P4672E4"
	},
	"incident": {
	  "id": "P31FVLZ",
	  "type": "incident_reference",
	  "summary": "[#12] GO009: Create from go program updated",
	  "self": "https://api.pagerduty.com/incidents/P31FVLZ",
	  "html_url": "https://effect-tech.pagerduty.com/incidents/P31FVLZ"
	},
	"teams": [],
	"contexts": [],
	"event_details": {}
      },
      {
        "id": "PVALU1K",
        "type": "notify_log_entry",
        "summary": "Notified Bob Smith by email",
	"self": "https://api.pagerduty.com/log_entries/PVALU1K",
	"html_url": null,
        "created_at": "2013-03-06T21:08:43Z",
        "channel": {
	  "type": "auto",
	  "notification": {
	    "type": "email",
	    "address": "bob@acme.com",
	    "status": "success"
	  }
	},
	"service": {
	  "id": "P4672E4",
	  "type": "service_reference",
	  "summary": "Test Service",
	  "self": "https://api.pagerduty.com/services/P4672E4",
	  "html_url": "https://effect-tech.pagerduty.com/services/P4672E4"
	},
	"incident": {
	  "id": "P31FVLZ",
	  "type": "incident_reference",
	  "summary": "[#12] GO009: Create from go program updated",
	  "self": "https://api.pagerduty.com/incidents/P31FVLZ",
	  "html_url": "https://effect-tech.pagerduty.com/incidents/P31FVLZ"
	},
	"teams": [],
	"contexts": [],
        "user": {
          "id": "PPSFHH7",
          "type": "user_reference",
          "summary": "Bob Smith",
          "self": "https://api.pagerduty.com/users/PVALU1K",
	  "html_url": "https://effect-tech.pagerduty.com/users/PVALU1K"
        }
      },
      {
        "id": "P8MDTZ8",
        "type": "notify_log_entry",
        "summary": "Notified Bob Smith by email",
	"self": "https://api.pagerduty.com/log_entries/PVALU1K",
	"html_url": null,
        "created_at": "2013-03-06T21:00:30Z",
        "channel": {
	  "type": "auto",
	  "notification": {
	    "type": "email",
	    "address": "bob@acme.com",
	    "status": "success"
	  }
	},
	"service": {
	  "id": "P4672E4",
	  "type": "service_reference",
	  "summary": "Test Service",
	  "self": "https://api.pagerduty.com/services/P4672E4",
	  "html_url": "https://effect-tech.pagerduty.com/services/P4672E4"
	},
	"incident": {
	  "id": "P31FVLZ",
	  "type": "incident_reference",
	  "summary": "[#12] GO009: Create from go program updated",
	  "self": "https://api.pagerduty.com/incidents/P31FVLZ",
	  "html_url": "https://effect-tech.pagerduty.com/incidents/P31FVLZ"
	},
	"teams": [],
	"contexts": [],
        "user": {
          "id": "PPSFHH7",
          "type": "user_reference",
          "summary": "Bob Smith",
          "self": "https://api.pagerduty.com/users/PVALU1K",
	  "html_url": "https://effect-tech.pagerduty.com/users/PVALU1K"
        }
      },
      {
        "id": "PHCXXHR",
        "type": "escalate_log_entry",
        "summary": "Escalated to Bob Smith by timeout",
        "self": "https://api.pagerduty.com/log_entries/PHCXXHR",
	"html_url": null,
        "created_at": "2013-03-06T21:00:30Z",
        "agent": {
	  "id": "P4672E4",
	  "type": "service_reference",
	  "summary": "Test Service",
	  "self": "https://api.pagerduty.com/services/P4672E4",
	  "html_url": "https://effect-tech.pagerduty.com/services/P4672E4"
        },
        "channel": {
          "type": "auto"
        },
        "service": {
	  "id": "P4672E4",
	  "type": "service_reference",
	  "summary": "Test Service",
	  "self": "https://api.pagerduty.com/services/P4672E4",
	  "html_url": "https://effect-tech.pagerduty.com/services/P4672E4"
	},
	"incident": {
	  "id": "P31FVLZ",
	  "type": "incident_reference",
	  "summary": "[#12] GO009: Create from go program updated",
	  "self": "https://api.pagerduty.com/incidents/P31FVLZ",
	  "html_url": "https://effect-tech.pagerduty.com/incidents/P31FVLZ"
	},
	"teams": [],
	"contexts": [],
        "assignees": [{
          "id": "PPSFHH7",
          "type": "user_reference",
          "summary": "Bob Smith",
          "self": "https://api.pagerduty.com/users/PE5NWQP",
	  "html_url": "https://effect-tech.pagerduty.com/users/PPSFHH7"
        }]
      },
      {
        "id": "PE5XGDA",
        "type": "notify_log_entry",
        "summary": "Notified Tim Wright by sms",
	"self": "https://api.pagerduty.com/log_entries/R5ILAUXL2SVD4UMK7QX70QB9BG",
	"html_url": null,
        "created_at": "2013-03-06T20:59:31Z",
	"channel": {
	  "type": "auto",
	  "notification": {
	    "type": "sms",
	    "address": "+1 956-821-0372",
	    "status": "success"
	  }
	},
	"service": {
	  "id": "P4672E4",
	  "type": "service_reference",
	  "summary": "Test Service",
	  "self": "https://api.pagerduty.com/services/P4672E4",
	  "html_url": "https://effect-tech.pagerduty.com/services/P4672E4"
	},
	"incident": {
	  "id": "P31FVLZ",
	  "type": "incident_reference",
	  "summary": "[#12] GO009: Create from go program updated",
	  "self": "https://api.pagerduty.com/incidents/P31FVLZ",
	  "html_url": "https://effect-tech.pagerduty.com/incidents/P31FVLZ"
	},
	"teams": [],
	"contexts": [],
        "user": {
          "id": "PT23IWX",
          "type": "user_reference",
          "summary": "Tim Wright",
          "self": "https://api.pagerduty.com/users/PT23IWX",
	  "html_url": "https://effect-tech.pagerduty.com/users/PT23IWX"
        }
      },
      {
        "id": "P9DWJ9J",
        "type": "notify_log_entry",
        "summary": "Notified Tim Wright by email",
	"self": "https://api.pagerduty.com/log_entries/P9DWJ9J",
	"html_url": null,
        "created_at": "2013-03-06T20:59:30Z",
        "channel": {
	  "type": "auto",
	  "notification": {
	    "type": "email",
	    "address": "tim@pagerduty.com",
	    "status": "success"
	  }
	},
	"service": {
	  "id": "P4672E4",
	  "type": "service_reference",
	  "summary": "Test Service",
	  "self": "https://api.pagerduty.com/services/P4672E4",
	  "html_url": "https://effect-tech.pagerduty.com/services/P4672E4"
	},
	"incident": {
	  "id": "P31FVLZ",
	  "type": "incident_reference",
	  "summary": "[#12] GO009: Create from go program updated",
	  "self": "https://api.pagerduty.com/incidents/P31FVLZ",
	  "html_url": "https://effect-tech.pagerduty.com/incidents/P31FVLZ"
	},
	"teams": [],
	"contexts": [],
        "user": {
          "id": "PT23IWX",
          "type": "user_reference",
          "summary": "Tim Wright",
          "self": "https://api.pagerduty.com/users/PT23IWX",
	  "html_url": "https://effect-tech.pagerduty.com/users/PT23IWX"
        }
      },
      {
        "id": "P36T0K8",
        "type": "unacknowledge_log_entry",
        "summary": "Unacknowledged by timeout",
	"self": "https://api.pagerduty.com/log_entries/P36T0K8",
	"html_url": null,
        "created_at": "2013-03-06T20:59:30Z",
        "agent": {
          "id": "P4672E4",
	  "type": "service_reference",
	  "summary": "Test Service",
	  "self": "https://api.pagerduty.com/services/P4672E4",
	  "html_url": "https://effect-tech.pagerduty.com/services/P4672E4"
        },
        "channel": {
          "type": "timeout"
        },
        "service": {
	  "id": "P4672E4",
	  "type": "service_reference",
	  "summary": "Test Service",
	  "self": "https://api.pagerduty.com/services/P4672E4",
	  "html_url": "https://effect-tech.pagerduty.com/services/P4672E4"
	},
	"incident": {
	  "id": "P31FVLZ",
	  "type": "incident_reference",
	  "summary": "[#12] GO009: Create from go program updated",
	  "self": "https://api.pagerduty.com/incidents/P31FVLZ",
	  "html_url": "https://effect-tech.pagerduty.com/incidents/P31FVLZ"
	},
	"teams": [],
	"contexts": []
      },
      {
        "id": "PKU5VMR",
        "type": "acknowledge_log_entry",
        "summary": "Acknowledged by Tim Wright",
	"self": "https://api.pagerduty.com/log_entries/PKU5VMR",
	"html_url": null,
        "created_at": "2013-03-06T20:29:30Z",
        "agent": {
          "id": "PT23IWX",
          "summary": "Tim Wright",
	  "self": "https://api.pagerduty.com/users/PT23IWX",
	  "html_url": "https://effect-tech.pagerduty.com/users/PT23IWX"
        },
        "channel": {
          "type": "phone"
        },
        "service": {
	  "id": "P4672E4",
	  "type": "service_reference",
	  "summary": "Test Service",
	  "self": "https://api.pagerduty.com/services/P4672E4",
	  "html_url": "https://effect-tech.pagerduty.com/services/P4672E4"
	},
	"incident": {
	  "id": "P31FVLZ",
	  "type": "incident_reference",
	  "summary": "[#12] GO009: Create from go program updated",
	  "self": "https://api.pagerduty.com/incidents/P31FVLZ",
	  "html_url": "https://effect-tech.pagerduty.com/incidents/P31FVLZ"
	},
	"teams": [],
	"contexts": [],
	"event_details": {}
      },
      {
        "id": "P36HB17",
        "type": "notify_log_entry",
        "summary": "Notified Tim Wright by email",
	"self": "https://api.pagerduty.com/log_entries/P36HB17",
	"html_url": null,
        "created_at": "2013-03-06T20:28:51Z",
        "channel": {
	  "type": "auto",
	  "notification": {
	    "type": "email",
	    "address": "tim@pagerduty.com",
	    "status": "success"
	  }
	},
	"service": {
	  "id": "P4672E4",
	  "type": "service_reference",
	  "summary": "Test Service",
	  "self": "https://api.pagerduty.com/services/P4672E4",
	  "html_url": "https://effect-tech.pagerduty.com/services/P4672E4"
	},
	"incident": {
	  "id": "P31FVLZ",
	  "type": "incident_reference",
	  "summary": "[#12] GO009: Create from go program updated",
	  "self": "https://api.pagerduty.com/incidents/P31FVLZ",
	  "html_url": "https://effect-tech.pagerduty.com/incidents/P31FVLZ"
	},
	"teams": [],
	"contexts": [],
        "user": {
          "id": "PT23IWX",
	  "type": "user_reference",
	  "summary": "Tim Wright",
	  "self": "https://api.pagerduty.com/users/PTNQU99",
	  "html_url": "https://effect-tech.pagerduty.com/users/PTNQU99"
        }
      },
      {
        "id": "PBJBF9Q",
        "type": "notify_log_entry",
        "summary": "Notified Tim Wright by phone",
	"self": "https://api.pagerduty.com/log_entries/PBJBF9Q",
	"html_url": null,
        "created_at": "2013-03-06T20:28:50Z",
        "channel": {
	  "type": "auto",
	  "notification": {
	    "type": "phone",
	    "address": "+1 956-821-0372",
	    "status": "success"
	  }
	},
        "service": {
	  "id": "P4672E4",
	  "type": "service_reference",
	  "summary": "Test Service",
	  "self": "https://api.pagerduty.com/services/P4672E4",
	  "html_url": "https://effect-tech.pagerduty.com/services/P4672E4"
	},
	"incident": {
	  "id": "P31FVLZ",
	  "type": "incident_reference",
	  "summary": "[#12] GO009: Create from go program updated",
	  "self": "https://api.pagerduty.com/incidents/P31FVLZ",
	  "html_url": "https://effect-tech.pagerduty.com/incidents/P31FVLZ"
	},
	"teams": [],
	"contexts": [],
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
        "type": "assign_log_entry",
        "summary": "Assigned to Tim Wright by timeout",
	"self": "https://api.pagerduty.com/log_entries/PRD9J7P",
	"html_url": null,
        "created_at": "2013-03-06T20:28:46Z",
        "agent": {
	  "id": "P4672E4",
	  "type": "service_reference",
	  "summary": "Test Service",
	  "self": "https://api.pagerduty.com/services/P4672E4",
  	  "html_url": "https://effect-tech.pagerduty.com/services/P4672E4"
        },
        "channel": {
          "type": "auto"
        },
        "service": {
	  "id": "P4672E4",
	  "type": "service_reference",
	  "summary": "Test Service",
	  "self": "https://api.pagerduty.com/services/P4672E4",
	  "html_url": "https://effect-tech.pagerduty.com/services/P4672E4"
	},
	"incident": {
	  "id": "P31FVLZ",
	  "type": "incident_reference",
	  "summary": "[#12] GO009: Create from go program updated",
	  "self": "https://api.pagerduty.com/incidents/P31FVLZ",
	  "html_url": "https://effect-tech.pagerduty.com/incidents/P31FVLZ"
	},
	"teams": [],
	"contexts": [],
        "assignees": [{
          "id": "PT23IWX",
          "type": "user_reference",
          "summary": "Tim Wright",
          "self": "https://api.pagerduty.com/users/PT23IWX",
	  "html_url": "https://effect-tech.pagerduty.com/users/PT23IWX"
        }]
      },
      {
        "id": "PVPXJJC",
        "type": "trigger_log_entry",
        "summary": "Triggered through the website",
	"self": "https://api.pagerduty.com/log_entries/PVPXJJC",
	"html_url": "https://effect-tech.pagerduty.com/incidents/P31FVLZ/log_entries/PVPXJJC",
        "created_at": "2013-03-06T20:28:46Z",
        "agent": {
          "id": "PT23IWX",
	  "type": "service_reference",
          "summary": "Tim Wright",
	  "self": "https://api.pagerduty.com/services/PT23IWX",
	  "html_url": "https://effect-tech.pagerduty.com/services/PT23IWX"
        },
        "channel": {
          "summary": "Martian's are attacking",
          "type": "web"
        },
        "service": {
	  "id": "P4672E4",
	  "type": "service_reference",
	  "summary": "Test Service",
	  "self": "https://api.pagerduty.com/services/P4672E4",
	  "html_url": "https://effect-tech.pagerduty.com/services/P4672E4"
	},
	"incident": {
	  "id": "P31FVLZ",
	  "type": "incident_reference",
	  "summary": "[#12] GO009: Create from go program updated",
	  "self": "https://api.pagerduty.com/incidents/P31FVLZ",
	  "html_url": "https://effect-tech.pagerduty.com/incidents/P31FVLZ"
	},
	"teams": [],
	"contexts": [],
	"event_details": {
	  "description": "Test incident"
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
	expectedIle.Type = "resolve_log_entry"
	expectedIle.AgentName = "Tim Wright"
	expectedIle.AgentType = "user_reference"
	expectedIle.ChannelType = "website"
	expectedIle.NotificationType = ""
	expectedIle.UserName = ""
	expectedIle.AssignedUser = ""
	expectedIle.Summary = "Resolved by Tim Wright"
	expectedIle.ChannelSummary = ""
	expectedIle.Message = ""
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
	expectedIle.Type = "notify_log_entry"
	expectedIle.Summary = "Notified Bob Smith by email"
	expectedIle.AgentName = ""
	expectedIle.AgentType = ""
	expectedIle.ChannelType = "auto"
	expectedIle.NotificationType = ""
	expectedIle.UserName = "Bob Smith"
	expectedIle.AssignedUser = ""
	expectedIle.ChannelSummary = ""
	expectedIle.Message = ""
	if *notifyIle != expectedIle {
		msg := fmt.Sprintf("Notify ILE parsed incorrectly. Got: %v", notifyIle)
		t.Error(msg)
	}
	if *friendlyMessage != "Notified Bob Smith via auto" {
		msg := fmt.Sprintf("Error generating friendly notify message. Got: %v", *friendlyMessage)
		t.Error(msg)
	}

	escalateIle := ileSlice[3]
	friendlyMessage = escalateIle.FriendlyMessage()
	expectedIle.Type = "escalate_log_entry"
	expectedIle.Summary = "Escalated to Bob Smith by timeout"
	expectedIle.AgentName = "Test Service"
	expectedIle.AgentType = "service_reference"
	expectedIle.ChannelType = "auto"
	expectedIle.NotificationType = ""
	expectedIle.UserName = ""
	expectedIle.AssignedUser = "Bob Smith"
	expectedIle.ChannelSummary = ""
	expectedIle.Message = ""
	if *escalateIle != expectedIle {
		msg := fmt.Sprintf("Escalate ILE parsed incorrectly. Got: %v", escalateIle)
		t.Error(msg)
	}
	if *friendlyMessage != "Escalated to Bob Smith via auto" {
		msg := fmt.Sprintf("Error generating friendly escalate message. Got: %v", *friendlyMessage)
		t.Error(msg)
	}

	// No friendly message for unacks
	unackIle := ileSlice[6]
	expectedIle.Type = "unacknowledge_log_entry"
	expectedIle.Summary = "Unacknowledged by timeout"
	expectedIle.AgentName = "Test Service"
	expectedIle.AgentType = "service_reference"
	expectedIle.ChannelType = "timeout"
	expectedIle.NotificationType = ""
	expectedIle.UserName = ""
	expectedIle.AssignedUser = ""
	expectedIle.ChannelSummary = ""
	expectedIle.Message = ""
	if *unackIle != expectedIle {
		msg := fmt.Sprintf("Unack ILE parsed incorrectly. Got: %v", unackIle)
		t.Error(msg)
	}

	ackIle := ileSlice[7]
	friendlyMessage = ackIle.FriendlyMessage()
	expectedIle.Type = "acknowledge_log_entry"
	expectedIle.Summary = "Acknowledged by Tim Wright"
	expectedIle.AgentName = "Tim Wright"
	expectedIle.AgentType = ""
	expectedIle.ChannelType = "phone"
	expectedIle.NotificationType = ""
	expectedIle.UserName = ""
	expectedIle.AssignedUser = ""
	expectedIle.ChannelSummary = ""
	expectedIle.Message = ""
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
	expectedIle.Type = "assign_log_entry"
	expectedIle.Summary = "Assigned to Tim Wright by timeout"
	expectedIle.AgentName = "Test Service"
	expectedIle.AgentType = "service_reference"
	expectedIle.ChannelType = "auto"
	expectedIle.NotificationType = ""
	expectedIle.UserName = ""
	expectedIle.AssignedUser = "Tim Wright"
	expectedIle.ChannelSummary = ""
	expectedIle.Message = ""
	if *assignIle != expectedIle {
		msg := fmt.Sprintf("Assign ILE parsed incorrectly. Got: %v", assignIle)
		t.Error(msg)
	}
	if *friendlyMessage != "Assigned to Tim Wright via auto" {
		msg := fmt.Sprintf("Error generating friendly assign message. Got %v", *friendlyMessage)
		t.Error(msg)
	}

	triggerIle := ileSlice[11]
	friendlyMessage = triggerIle.FriendlyMessage()
	expectedIle.Type = "trigger_log_entry"
	expectedIle.Summary = "Triggered through the website"
	expectedIle.AgentName = "Tim Wright"
	expectedIle.AgentType = "service_reference"
	expectedIle.ChannelType = "web"
	expectedIle.NotificationType = ""
	expectedIle.UserName = ""
	expectedIle.AssignedUser = ""
	expectedIle.ChannelSummary = "Martian's are attacking"
	expectedIle.Message = ""
	if *triggerIle != expectedIle {
		msg := fmt.Sprintf("Trigger ILE parsed incorrectly. Got: %v", triggerIle)
		t.Error(msg)
	}
	if *friendlyMessage != "Triggered by Tim Wright" {
		msg := fmt.Sprintf("Error generating friendly trigger message. Got: %v", *friendlyMessage)
		t.Error(msg)
	}
}

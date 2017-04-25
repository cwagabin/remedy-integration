# PagerDuty-Remedy Integration
The PagerDuty Remedy integration is a binary that provides various helper functions to the BMC Remedy tool. The behavior and requirements of this tool are dictated solely by BMC Remedy behavior, so use of it as a general tool is discouraged.

## Main Subcommands
Currently, there are four main functions required to make the Remedy integration go. They are documented below.

### Trigger
The `trigger` subcommand is used to trigger a new PagerDuty Incident. Example:
```
./pd-remedy trigger --api-key KEY --incident-key REMEDY_INCIDENT_NUMBER --description "This is a new incident" --details "This is the incident details"  --service-id PAGER_DUTY_SERVICE_ID --assignee ASSIGNEE_USER_ID --requester USER_EMAIL --client REMEDY_CLIENT_VALUE --client-url REMEDY_INCIDENT_URL 
```
It requires `api-key`, `incident-key`, `description`, `service-id` and `requester`. `assignee`, `details`, `client`, and `client-url` are optional. If not specified, `client` will default to "BMC Remedy" and `client-url` will default to https://www.bmc.com.

### Acknowledge
The `acknowledge` subcommand is used to acknowledge open PagerDuty Incidents. Example:
```
./pd-remedy acknowledge --api-key KEY --incident-id PAGER_DUTY_INCIDENT_ID --requester USER_EMAIL --client REMEDY_CLIENT_VALUE --client-url REMEDY_INCIDENT_URL --retry NUMBER_OF_TIMES_TO_RETRY --delay-time NUMBER_OF_SECONDS
```
It requires `api-key`, `incident-id` and `requester`. If the `incident-id` is not known `incident-key` can be used. `client-url`, `retry` and `delay-time` are optional.
If not specified, `client` will default to "BMC Remedy" and `client-url` will default to https://www.bmc.com. 
`retry` and `delay-time` allow user to specify the number of times to try to find the existing PagerDuty incident and the number of seconds to delay before the next retry attempt.

### Resolve
The `resolve` subcommand is used to resolve open PagerDuty Incidents. Example:
```
./pd-remedy resolve --api-key KEY --incident-id PAGER_DUTY_INCIDENT_ID --requester USER_EMAIL --client REMEDY_CLIENT_VALUE --client-url REMEDY_INCIDENT_URL --retry NUMBER_OF_TIMES_TO_RETRY --delay-time NUMBER_OF_SECONDS
```
It requires `api-key`, `incident-id` and `requester`. If the `incident-id` is not known `incident-key` can be used. `client`, `client-url`, `retry` and `delay-time` are optional. 
If not specified, `client` will default to "BMC Remedy" and `client-url` will default to https://www.bmc.com. 
`retry` and `delay-time` allow user to specify the number of times to try to find the existing PagerDuty incident and the number of seconds to delay before the next retry attempt.

### Get-Iles
The `get-iles` subcommand is used to retrieve the history and status of a particular PagerDuty incident using Incident ID. An error is thrown if the Incident ID is invalid.
```
./pd-remedy get-iles --api-key KEY --incident-id PAGER_DUTY_INCIDENT_ID
```

## Secondary Subcommands
These are functions to retrieve or validate id's. 

### Get-Id
The `get-id` subcommand is used to translate a PagerDuty Incident Key into a PagerDuty Incident ID. The Incident ID is required during incident-specific API actions (such as retrieving incident updates), and is useful for generating links to the PagerDuty Incident after it has been submitted. `delay-time` is optional. It allows user to specify the number of seconds to delay before performing the subcommand. *PLEASE NOTE:* This subcommand will return an empty string and exit with code 0 if no incidents match the given key. This is a Remedy-ism.
```
./pd-remedy get-id --api-key KEY --incident-key REMEDY_INCIDENT_NUMBER --delay-time NUMBER_OF_SECONDS
```

### Get-Validate-Id
The `get-validate-id` subcommand is used to verify a PagerDuty Incident ID. The Incident ID is required during incident-specific API actions (such as retrieving incident updates), and is useful for generating links to the PagerDuty Incident after it has been submitted. *PLEASE NOTE:* This subcommand will return an empty string and exit with code 0 if no incidents match the given key. If a match is found, it will return the validated PagerDuty Incident ID. This is a Remedy-ism.
```
./pd-remedy get-validate-id --api-key KEY --incident-id PAGER_DUTY_INCIDENT_ID
```

### Get-Service-Id
The `get-service-id` subcommand is used to translate a PagerDuty Service Name into a PagerDuty Service ID. The Service ID is required during incident-specific API actions such as creating incident. *PLEASE NOTE:* This subcommand will return an empty string and exit with code 0 if no services match the given key. This is a Remedy-ism.
```
./pd-remedy get-service-id --api-key KEY --service-name PAGER_DUTY_SERVICE_NAME
```

### Get-Service-Escalation-Id
The `get-service-escalation-id` subcommand is used to retrieve the PagerDuty Escalation Policy ID corresponding to a PagerDuty Service ID. The Escalation Policy ID is required during incident-specific API actions such as reassigning incident. *PLEASE NOTE:* This subcommand will return an empty string and exit with code 0 if no services match the given key. This is a Remedy-ism.
```
./pd-remedy get-service-escalation-id --api-key KEY --service-id PAGER_DUTY_SERVICE_ID
```

### Validate-Escalation-Policy-Id
The `validate-escalation-policy-id` subcommand is used to validate a PagerDuty Escalation Policy ID. If a match is found, it will return the validated Escalation Policy ID. *PLEASE NOTE:* This subcommand will return an empty string and exit with code 0 if no escalation policy match the given ID. This is a Remedy-ism.
```
./pd-remedy validate-escalation-policy-id --api-key KEY --escalation-policy-id PAGER_DUTY_ESCALATION_POLICY_ID
```

### Get-User-Id
The `get-user-id` subcommand is used to translate a PagerDuty User Email Address into a PagerDuty User ID. The User ID is required during incident-specific API actions such as setting value for assignee during incident updates. *PLEASE NOTE:* This subcommand will return an empty string and exit with code 0 if no users match the given key. This is a Remedy-ism.
```
./pd-remedy get-user-id --api-key KEY --user-email USER_EMAIL
```

### Test-Connection
The `test-connection` subcommand is used to test the connection to the PagerDuty instance. Returns "ok" if connection is working or "failed" otherwise.
```
./pd-remedy test-connection --api-key KEY 
```

## Reassignment
Incidents can be updated due to reassignment. Reassignment can mean a change of escalation policy or reassignment to a different assignee. NOTE: Cannot perform reassignment using both assignee and escalation policy at the same time. 

For the examples below, `incident-key` can be used if `incident-id` is not available. The mode `trigger` can be set to `acknowledge` if needed.

### Example for Changing Escalation Policy
This will set the PagerDuty incident status to `triggered` and PagerDuty will reassign the incident based on the escalation policy.  The examples use `acknowledge` but `trigger` can also be used.
```
pd-remedy.exe acknowledge --api-key KEY --incident-id PAGER_DUTY_INCIDENT_ID --escalation-policy-id PAGER_DUTY_ESCALATION_POLICY_ID --requester USER_EMAIL --client REMEDY_CLIENT_VALUE  --client-url REMEDY_INCIDENT_URL
```

### Example for Changing Assignee
The status of the PagerDuty incident remains as `acknowledged` when assignee is provided.
```
pd-remedy.exe acknowledge --api-key KEY --incident-id PAGER_DUTY_INCIDENT_ID --assignee ASSIGNEE_USER_ID --requester USER_EMAIL --client REMEDY_CLIENT_VALUE  --client-url REMEDY_INCIDENT_URL
```

## License
[Apache 2](http://www.apache.org/licenses/LICENSE-2.0)

## Contributing
1. Fork this repo and clone it to your workstation.
1. Create a feature branch for your change.
1. Write code and tests.
1. Update the docs.
1. Push your feature branch to github and open a pull request against master.

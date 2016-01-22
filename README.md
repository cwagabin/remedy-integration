# PagerDuty-Remedy Integration
The PagerDuty Remedy integration is a binary that provides various helper functions to the BMC Remedy tool. The behavior and requirements of this tool are dictated solely by BMC Remedy behavior, so use of it as a general tool is discouraged.

## Subcommands
Currently, there are four distinct functions required to make the Remedy integration go. They are documented below.

### Trigger
The `trigger` subcommand is used to trigger a new PagerDuty Incident. Example:
```
./pd-remedy trigger --service-key KEY --description "This is a new incident"
```
It optionally takes `--incident-id` for deduplication and `--details` for a detailed description of the event. Setting `--incident-id` to the BMC Remedy Incident ID is strongly encouraged.

### Resolve
The `resolve` subcommand is used to resolve open PagerDuty Incidents. Example:
```
./pd-remedy resolve --service-key KEY --incident-key KEY
```
`incident-key` is the optional key defined when the PagerDuty Incident is triggered, and is required for the Resolve action. Using the BMC Remedy Incident ID as the `incident-key` when triggering the PagerDuty Incident is recommended for ease of use.

### Get-Id
The `get-id` subcommand is used to translate a PagerDuty Incident Key into a PagerDuty Incident ID. The Incident ID is required during incident-specific API actions (such as retrieving incident updates), and is useful for generating links to the PagerDuty Incident after it has been submitted. *PLEASE NOTE:* This subcommand will return an empty string and exit with code 0 if no incidents match the given key. This is a Remedy-ism.
```
./pd-remedy get-id --api-key KEY --subdomain SUBDOMAIN --incident-key KEY
```

### Get-Iles
The `get-iles` subcommand is used to retrieve the history and status of a particular PagerDuty incident using its ID. An error is thrown if the Incident ID is invalid.
```
./pd-remedy get-iles --api-key KEY --subdomain SUBDOMAIN --incident-id ID
```

## Contributing
1. Fork this repo and clone it to your workstation.
1. Create a feature branch for your change.
1. Write code and tests.
1. Update the docs.
1. Push your feature branch to github and open a pull request against master.

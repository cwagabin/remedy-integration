// Copyright:: Copyright (c) 2016-2017 PagerDuty, Inc.
// License:: Apache License, Version 2.0

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
)

func main() {
	mode := os.Args[1]
	args := os.Args[2:]
	switch mode {
	case "trigger":
		{
			c, err := createTriggerAckResolveChangePolicyConfig(mode, args)
			failIf(err)
			triggerAckResolveChangePolicy(c)
		}
	case "acknowledge":
		{
			c, err := createTriggerAckResolveChangePolicyConfig(mode, args)
			failIf(err)
			triggerAckResolveChangePolicy(c)
		}
	case "resolve":
		{
			c, err := createTriggerAckResolveChangePolicyConfig(mode, args)
			failIf(err)
			triggerAckResolveChangePolicy(c)
		}
	case "get-id":
		{
			c, err := createGetIDConfig(args)
			failIf(err)
			fmt.Println(getID(c))
		}
	case "get-validate-id":
		{
			c, err := createGetValidateIdConfig(args)
			failIf(err)
			fmt.Println(getValidateID(c))
		}
	case "get-iles":
		{
			c, err := createGetIlesConfig(args)
			failIf(err)

			// Print messages in reverse, trigger first
			messages := getIles(c)
			for i := len(messages) - 1; i >= 0; i-- {
				fmt.Println(messages[i])
			}
		}
	case "get-service-id":
		{
			c, err := createGetServiceIDConfig(args)
			failIf(err)
			fmt.Println(getServiceID(c))
		}
	case "get-service-escalation-id":
		{
			c, err := createGetServiceEscalationIDConfig(args)
			failIf(err)
			fmt.Println(getServiceEscalationID(c))
		}
	case "validate-escalation-policy-id":
		{
			c, err := createValidateEscalationPolicyIDConfig(args)
			failIf(err)
			fmt.Println(validateEscalationPolicyID(c))
		}
	case "get-user-id":
		{
			c, err := createGetUserIDConfig(args)
			failIf(err)
			fmt.Println(getUserID(c))
		}
	case "test-connection":
		{
			c, err := createTestConnectionConfig(args)
			failIf(err)
			fmt.Println(testAPIConnection(c))
		}
	default:
		{
			usage()
		}
	}
}

func usage() {
	usageStatement := fmt.Sprintf("Usage: %s [subcommand] [options]", os.Args[0])
	fmt.Fprintln(os.Stderr, usageStatement)
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Subcommands:")
	fmt.Fprintln(os.Stderr, "\ttrigger")
	fmt.Fprintln(os.Stderr, "\tacknowledge")
	fmt.Fprintln(os.Stderr, "\tresolve")
	fmt.Fprintln(os.Stderr, "\tget-id")
	fmt.Fprintln(os.Stderr, "\tget-validate-id")
	fmt.Fprintln(os.Stderr, "\tget-iles")
	fmt.Fprintln(os.Stderr, "\tget-service-id")
	fmt.Fprintln(os.Stderr, "\tget-service-escalation-id")
	fmt.Fprintln(os.Stderr, "\tget-user-id")
	fmt.Fprintln(os.Stderr, "\tvalidate-escalation-policy-id")
	fmt.Fprintln(os.Stderr, "\ttest-connection")
	os.Exit(1)
}

func failIf(e error) {
	if e != nil {
		msg := fmt.Sprintf("Error: %s", e.Error())
		fmt.Fprintln(os.Stderr, msg)
		os.Exit(1)
	}
}

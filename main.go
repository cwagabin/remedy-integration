// Copyright:: Copyright (c) 2016 PagerDuty, Inc.
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

import "fmt"
import "os"

func main() {
	mode := os.Args[1]
	args := os.Args[2:]
	switch mode {
	case "trigger":
		{
			c, err := createTriggerConfig(&args)
			failIf(err)
			triggerResolve(c)
		}
	case "resolve":
		{
			c, err := createResolveConfig(&args)
			failIf(err)
			triggerResolve(c)
		}
	case "get-id":
		{
			c, err := createGetIDConfig(&args)
			failIf(err)
			fmt.Println(getID(c))
		}
	case "get-iles":
		{
			c, err := createGetIlesConfig(&args)
			failIf(err)

			// Print messages in reverse, trigger first
			messages := getIles(c)
			for i := len(messages) - 1; i >= 0; i-- {
				fmt.Println(messages[i])
			}
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
	fmt.Fprintln(os.Stderr, "\tresolve")
	fmt.Fprintln(os.Stderr, "\tget-id")
	fmt.Fprintln(os.Stderr, "\tget-iles")
	os.Exit(1)
}

func failIf(e error) {
	if e != nil {
		msg := fmt.Sprintf("Error: %s", e.Error())
		fmt.Fprintln(os.Stderr, msg)
		os.Exit(1)
	}
}

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
        c, err := createGetIdConfig(&args)
        failIf(err)
        fmt.Println(getId(c))
      }
    case "get-iles":
      {
        c, err := createGetIlesConfig(&args)
        failIf(err)

        // Print messages in reverse, trigger first
        messages := getIles(c)
        for i := len(messages)-1; i >= 0; i-- {
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

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/dherbst/sentry-web-api"
)

var funcMap map[string]func(context.Context)

func init() {
	funcMap = map[string]func(context.Context){
		"help":     Usage,
		"version":  Version,
		"orgs":     Organizations,
		"projects": Projects,
		"events":   Events,
	}
}

// Version prints the version from the sentry.GitHash out and exits.
func Version(ctx context.Context) {
	fmt.Printf("Version: %v %v\n", sentry.Version, sentry.GitHash)
}

// Usage prints how to invoke `sentry` from the command line.
func Usage(ctx context.Context) {
	fmt.Printf(`
Usage:

sentry version                              ; prints the commit version
sentry orgs                                 ; prints the Organizations you can access
sentry projects                             ; prints the Projects you can access
sentry events org_slug project_slug         ; prints the events in the project
`)

}

// Organizations prints the orgs you have access to.
func Organizations(ctx context.Context) {
	client := sentry.NewClient("", 0, "")
	orgs, err := client.OrganizationList(false, "")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	result, err := json.MarshalIndent(orgs, "", "  ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(string(result))
}

// Projects prints the projects you have access to.
func Projects(ctx context.Context) {
	client := sentry.NewClient("", 0, "")
	projects, err := client.ProjectList("")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	result, err := json.MarshalIndent(projects, "", "  ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(string(result))
}

// Events prints the events in the project.  Requires org_slug and project_slug
func Events(ctx context.Context) {
	orgSlug := flag.Arg(1)
	projectSlug := flag.Arg(2)

	client := sentry.NewClient("", 0, "")
	events, err := client.EventList(orgSlug, projectSlug, false, "")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	result, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(string(result))
}

func main() {
	flag.Parse()

	ctx := context.Background()

	command := flag.Arg(0)
	if command == "" || command == "help" {
		Usage(ctx)
		return
	}

	f := funcMap[command]
	if f == nil {
		fmt.Println("Unknown command")
		Usage(ctx)
		return
	}

	f(ctx)
}

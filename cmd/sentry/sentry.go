package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/dherbst/sentry"
)

var funcMap map[string]func(context.Context)

func init() {
	funcMap = map[string]func(context.Context){
		"help":    Usage,
		"version": Version,
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
`)

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

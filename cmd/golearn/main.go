package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zhravan/golearn/internal/cli"
)

func main() {
	// Allow -h without subcommands to show help
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, cli.HelpText())
	}

	if err := cli.Execute(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package cli

import (
	"errors"
	"fmt"
)

// HelpText returns usage information for the CLI.
func HelpText() string {
	return `golearn - Rustlings-like Go learning CLI

Usage:
  golearn list                List available exercises
  golearn verify [name]       Run tests for a specific exercise (or all)
  golearn hint [name]         Show hints for an exercise
  golearn watch               Watch files and re-run tests on change
  golearn progress            Show progress
  golearn reset [name]        Reset exercise to starter state
  golearn init                Initialize exercises (copy templates)
  golearn help                Show this help
`
}

// Execute runs the CLI with arguments.
func Execute(args []string) error {
	if len(args) == 0 {
		fmt.Println(HelpText())
		return nil
	}

	switch args[0] {
	case "help", "-h", "--help":
		fmt.Println(HelpText())
		return nil
	case "list":
		return runList()
	case "verify":
		var name string
		if len(args) > 1 {
			name = args[1]
		}
		return runVerify(name)
	case "hint":
		if len(args) < 2 {
			return errors.New("hint requires an exercise name")
		}
		return runHint(args[1])
	case "watch":
		return runWatch()
	case "progress":
		return runProgress()
	case "reset":
		var name string
		if len(args) > 1 {
			name = args[1]
		}
		return runReset(name)
	case "init":
		return runInit()
	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

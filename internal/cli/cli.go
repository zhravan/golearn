package cli

import (
	"errors"
	"fmt"

	"github.com/shravan20/golearn/internal/cli/theme"
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
  golearn init [repo] [dir]   Initialize workspace: clone exercises repo or copy built-in templates
  golearn help                Show this help

Global options:
  --no-color                  Disable ANSI colors (honors NO_COLOR)
  --theme=<name>              Theme: default | high-contrast | monochrome
  --screen-reader, --sr       Optimize for screen readers; avoid screen clears
`
}

// Execute runs the CLI with arguments.
func Execute(args []string) error {
	// Parse global accessibility/theming flags and envs
	args = theme.Setup(args)
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
		var repo, dir string
		if len(args) > 1 {
			repo = args[1]
		}
		if len(args) > 2 {
			dir = args[2]
		}
		return runInit(repo, dir)
	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

package theme

import (
	"os"
	"strings"
)

// Lightweight, dependency-free theming and accessibility helpers for CLI output.

type themeName string

const (
	defaultTheme      themeName = "default"
	highContrastTheme themeName = "high-contrast"
	monochromeTheme   themeName = "monochrome"
)

// config is process-wide configuration derived from env vars and global flags.
var config = struct {
	colorsEnabled    bool
	screenReaderMode bool
	selectedTheme    themeName
}{
	colorsEnabled:    shouldEnableColorByDefault(),
	screenReaderMode: isTruthy(os.Getenv("GOLEARN_SCREEN_READER")),
	selectedTheme:    resolveTheme(os.Getenv("GOLEARN_THEME")),
}

// Setup parses and applies global flags. It returns args with the recognized
// global options removed so command parsing can proceed unchanged.
func Setup(args []string) []string {
	filtered := make([]string, 0, len(args))
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == "--no-color":
			config.colorsEnabled = false
			continue
		case strings.HasPrefix(arg, "--theme="):
			name := strings.TrimPrefix(arg, "--theme=")
			config.selectedTheme = resolveTheme(name)
			continue
		case arg == "--screen-reader" || arg == "--sr":
			config.screenReaderMode = true
			// For many screen readers, color does not help and dynamic effects can hinder
			config.colorsEnabled = false
			continue
		}
		filtered = append(filtered, arg)
	}
	return filtered
}

func resolveTheme(name string) themeName {
	switch strings.ToLower(strings.TrimSpace(name)) {
	case string(highContrastTheme), "high", "hc":
		return highContrastTheme
	case string(monochromeTheme), "mono", "bw":
		return monochromeTheme
	case string(defaultTheme), "", "auto":
		return defaultTheme
	default:
		return defaultTheme
	}
}

func shouldEnableColorByDefault() bool {
	// NO_COLOR standard disables colors
	if _, ok := os.LookupEnv("NO_COLOR"); ok {
		return false
	}
	// FORCE_COLOR enables colors regardless of TTY (useful in CI)
	if _, ok := os.LookupEnv("FORCE_COLOR"); ok {
		return true
	}
	// Basic TTY heuristic without external deps
	if os.Getenv("TERM") == "" || os.Getenv("TERM") == "dumb" {
		return false
	}
	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

// Public toggles
func ColorsEnabled() bool    { return config.colorsEnabled && config.selectedTheme != monochromeTheme }
func ScreenReaderMode() bool { return config.screenReaderMode }
func SelectedTheme() string  { return string(config.selectedTheme) }

// Style helpers return formatted strings. They never panic and degrade gracefully.
func Heading(s string) string { return applyStyle(s, styleHeading) }
func Success(s string) string { return applyStyle(s, styleSuccess) }
func Error(s string) string   { return applyStyle(s, styleError) }
func Hint(s string) string    { return applyStyle(s, styleHint) }
func Muted(s string) string   { return applyStyle(s, styleMuted) }
func Emph(s string) string    { return applyStyle(s, styleEmphasis) }

type style int

const (
	styleHeading style = iota
	styleSuccess
	styleError
	styleHint
	styleMuted
	styleEmphasis
)

func applyStyle(s string, st style) string {
	if !config.colorsEnabled || config.selectedTheme == monochromeTheme || s == "" {
		return s
	}
	start, end := styleCodes(st, config.selectedTheme)
	if start == "" {
		return s
	}
	return start + s + end
}

func styleCodes(st style, tn themeName) (start, end string) {
	// ANSI SGR codes. Keep to foreground and intensity for broad compatibility.
	// end reset
	end = "\x1b[0m"
	if tn == highContrastTheme {
		switch st {
		case styleHeading:
			return "\x1b[1;97m", end // bold bright white
		case styleSuccess:
			return "\x1b[1;92m", end // bold bright green
		case styleError:
			return "\x1b[1;91m", end // bold bright red
		case styleHint:
			return "\x1b[1;93m", end // bold bright yellow
		case styleMuted:
			return "\x1b[2;37m", end // dim white
		case styleEmphasis:
			return "\x1b[1m", end // bold
		}
	} else {
		// default theme
		switch st {
		case styleHeading:
			return "\x1b[1;96m", end // bold bright cyan
		case styleSuccess:
			return "\x1b[32m", end // green
		case styleError:
			return "\x1b[31m", end // red
		case styleHint:
			return "\x1b[33m", end // yellow
		case styleMuted:
			return "\x1b[2;37m", end // dim white
		case styleEmphasis:
			return "\x1b[1m", end // bold
		}
	}
	return "", ""
}

// MaybeClearScreen returns true if the screen was cleared. It avoids clearing
// for screen readers and non-TTYs to keep output linear and accessible.
func MaybeClearScreen() bool {
	if ScreenReaderMode() || !ColorsEnabled() {
		// Treat non-color scenario as likely non-TTY; avoid dynamic effects
		return false
	}
	os.Stdout.WriteString("\x1b[2J\x1b[H")
	return true
}

func isTruthy(v string) bool {
	if v == "" {
		return false
	}
	v = strings.ToLower(strings.TrimSpace(v))
	return v == "1" || v == "true" || v == "yes" || v == "on"
}

// Package ansi allows for advanced terminal text manipulation.
package ansi

import (
	"fmt"
	"io"
)

type AnsiColor struct {
	bold  bool
	value int
}

var (
	// Black sets the ansi color black.
	Black = AnsiColor{false, 30}
	// BoldBlack sets the ansi color black and bold.
	BoldBlack = AnsiColor{true, 30}
	// Red sets the ansi color red.
	Red = AnsiColor{false, 31}
	// BoldRed sets the ansi color red and bold.
	BoldRed = AnsiColor{true, 31}
	// Green sets the ansi color green.
	Green = AnsiColor{false, 32}
	// BoldGreen sets the ansi color green and bold.
	BoldGreen = AnsiColor{true, 32}
	// Yellow sets the ansi color yellow.
	Yellow = AnsiColor{false, 33}
	// BoldYellow sets the ansi color yellow and bold.
	BoldYellow = AnsiColor{true, 33}
	// Blue sets the ansi color blue.
	Blue = AnsiColor{false, 34}
	// BoldBlue sets the ansi color blue and bold.
	BoldBlue = AnsiColor{true, 34}
	// Purple sets the ansi color purple.
	Purple = AnsiColor{false, 35}
	// BoldPurple sets the ansi color purple and bold.
	BoldPurple = AnsiColor{true, 35}
	// Cyan sets the ansi color cyan.
	Cyan = AnsiColor{false, 36}
	// BoldCyan sets the ansi color cyan and bold.
	BoldCyan = AnsiColor{true, 36}
	// White sets the ansi color white.
	White = AnsiColor{false, 37}
	// BoldWhite sets the ansi color white and bold.
	BoldWhite = AnsiColor{true, 37}
	// Default sets the ansi color to the default color.
	Default = AnsiColor{false, 39}

	// Bold sets the text to be bold.
	Bold = AnsiColor{false, 1}
	// NotBold sets the text to be not bold.
	NotBold = AnsiColor{false, 22}
	// Underline sets the text to be underlined.
	Underline = AnsiColor{false, 4}
	// NotUnderline sets the text to be not underlined.
	NotUnderline = AnsiColor{false, 24}
	// Reversed sets the text to be reversed.
	Reversed = AnsiColor{false, 7}
	// NotReversed sets the text to be not reversed.
	NotReversed = AnsiColor{false, 27}
)

var (
	fgColor     = Default
	bgColor     = Default
	isBold      = false
	isUnderline = false
	isReversed  = false
)

// SetFg sets the foreground text color.
func SetFg(fg AnsiColor) {
	fgColor = fg
	if fg.bold {
		isBold = true
	}
	applySettings()
}

// SetBg sets the background text color.
func SetBg(bg AnsiColor) {
	bgColor = bg
	applySettings()
}

// SetBold sets the text to be bold or not.
func SetBold(bold bool) {
	isBold = bold
	applySettings()
}

// SetUnderline sets the text to be underlined or not.
func SetUnderline(underline bool) {
	isUnderline = underline
	applySettings()
}

// SetReversed sets the text to be reversed or not.
func SetReversed(reversed bool) {
	isReversed = reversed
	applySettings()
}

// SetAll sets the foreground, background, bold, underline and reversed properties at once.
func SetAll(fg, bg AnsiColor, bold, underline, reversed bool) {
	fgColor = fg
	bgColor = bg
	isBold = bold
	isUnderline = underline
	isReversed = reversed
	applySettings()
}

// ResetAll resets all properties to default.
func ResetAll() {
	SetAll(Default, Default, false, false, false)
	fmt.Print("\033[K")
}

func applySettings() {
	b := 22
	if isBold {
		b = 1
	}
	u := 24
	if isUnderline {
		u = 4
	}
	r := 27
	if isReversed {
		r = 7
	}
	fmt.Printf("\033[0;%d;%d;%d;%d;%dm", fgColor.value, bgColor.value+10, b, u, r)
}

// Print is a replacement for fmt.Print, which accepts an AnsiColor as the first argument.
func Print(col AnsiColor, s ...any) {
	fmt.Printf("\033[0;%dm", col.value)
	if col.bold {
		fmt.Printf("\033[1m")
	}
	fmt.Print(s...)
	applySettings()
}

// Println is a replacement for fmt.Println, which accepts an AnsiColor as the first argument.
func Println(col AnsiColor, s ...any) {
	fmt.Printf("\033[0;%dm", col.value)
	if col.bold {
		fmt.Printf("\033[1m")
	}
	fmt.Println(s...)
	applySettings()
}

// Printf is a replacement for fmt.Printf, which accepts an AnsiColor as the first argument.
func Printf(col AnsiColor, s string, args ...any) {
	fmt.Printf("\033[0;%dm", col.value)
	if col.bold {
		fmt.Printf("\033[1m")
	}
	fmt.Printf(s, args...)
	applySettings()
}

// Fprint is a replacement for fmt.Fprint, which accepts an AnsiColor as the first argument.
func Fprint(output io.Writer, col AnsiColor, s ...any) {
	fmt.Fprintf(output, "\033[0;%dm", col.value)
	if col.bold {
		fmt.Fprint(output, "\033[1m")
	}
	fmt.Fprint(output, s...)
	applySettings()
}

// Fprintln is a replacement for fmt.Fprintln, which accepts an AnsiColor as the first argument.
func Fprintln(output io.Writer, col AnsiColor, s ...any) {
	fmt.Fprintf(output, "\033[0;%dm", col.value)
	if col.bold {
		fmt.Fprintf(output, "\033[1m")
	}
	fmt.Fprintln(output, s...)
	applySettings()
}

// Fprintf is a replacement for fmt.Fprintf, which accepts an AnsiColor as the first argument.
func Fprintf(output io.Writer, col AnsiColor, s string, args ...any) {
	fmt.Fprintf(output, "\033[0;%dm", col.value)
	if col.bold {
		fmt.Fprintf(output, "\033[1m")
	}
	fmt.Fprintf(output, s, args...)
	applySettings()
}

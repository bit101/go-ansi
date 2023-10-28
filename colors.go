// Package ansi allows for advanced terminal text manipulation.
package ansi

import (
	"fmt"
	"io"
)

type ansiColor struct {
	bold  bool
	value int
}

var (
	// Black sets the ansi color black.
	Black = ansiColor{false, 30}
	// BoldBlack sets the ansi color black and bold.
	BoldBlack = ansiColor{true, 30}
	// Red sets the ansi color red.
	Red = ansiColor{false, 31}
	// BoldRed sets the ansi color red and bold.
	BoldRed = ansiColor{true, 31}
	// Green sets the ansi color green.
	Green = ansiColor{false, 32}
	// BoldGreen sets the ansi color green and bold.
	BoldGreen = ansiColor{true, 32}
	// Yellow sets the ansi color yellow.
	Yellow = ansiColor{false, 33}
	// BoldYellow sets the ansi color yellow and bold.
	BoldYellow = ansiColor{true, 33}
	// Blue sets the ansi color blue.
	Blue = ansiColor{false, 34}
	// BoldBlue sets the ansi color blue and bold.
	BoldBlue = ansiColor{true, 34}
	// Purple sets the ansi color purple.
	Purple = ansiColor{false, 35}
	// BoldPurple sets the ansi color purple and bold.
	BoldPurple = ansiColor{true, 35}
	// Cyan sets the ansi color cyan.
	Cyan = ansiColor{false, 36}
	// BoldCyan sets the ansi color cyan and bold.
	BoldCyan = ansiColor{true, 36}
	// White sets the ansi color white.
	White = ansiColor{false, 37}
	// BoldWhite sets the ansi color white and bold.
	BoldWhite = ansiColor{true, 37}
	// Default sets the ansi color to the default color.
	Default = ansiColor{false, 39}

	// Bold sets the text to be bold.
	Bold = ansiColor{false, 1}
	// NotBold sets the text to be not bold.
	NotBold = ansiColor{false, 22}
	// Underline sets the text to be underlined.
	Underline = ansiColor{false, 4}
	// NotUnderline sets the text to be not underlined.
	NotUnderline = ansiColor{false, 24}
	// Reversed sets the text to be reversed.
	Reversed = ansiColor{false, 7}
	// NotReversed sets the text to be not reversed.
	NotReversed = ansiColor{false, 27}
)

var (
	fgColor     = Default
	bgColor     = Default
	isBold      = false
	isUnderline = false
	isReversed  = false
)

// SetFg sets the foreground text color.
func SetFg(fg ansiColor) {
	fgColor = fg
	if fg.bold {
		isBold = true
	}
	applySettings()
}

// SetBg sets the background text color.
func SetBg(bg ansiColor) {
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
func SetAll(fg, bg ansiColor, bold, underline, reversed bool) {
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

// Print is a replacement for fmt.Print, which accepts an ansiColor as the first argument.
func Print(col ansiColor, s ...any) {
	fmt.Printf("\033[0;%dm", col.value)
	if col.bold {
		fmt.Printf("\033[1m")
	}
	fmt.Print(s...)
	applySettings()
}

// Println is a replacement for fmt.Println, which accepts an ansiColor as the first argument.
func Println(col ansiColor, s ...any) {
	fmt.Printf("\033[0;%dm", col.value)
	if col.bold {
		fmt.Printf("\033[1m")
	}
	fmt.Println(s...)
	applySettings()
}

// Printf is a replacement for fmt.Printf, which accepts an ansiColor as the first argument.
func Printf(col ansiColor, s string, args ...any) {
	fmt.Printf("\033[0;%dm", col.value)
	if col.bold {
		fmt.Printf("\033[1m")
	}
	fmt.Printf(s, args...)
	applySettings()
}

// Fprint is a replacement for fmt.Fprint, which accepts an ansiColor as the first argument.
func Fprint(output io.Writer, col ansiColor, s ...any) {
	fmt.Fprintf(output, "\033[0;%dm", col.value)
	if col.bold {
		fmt.Fprint(output, "\033[1m")
	}
	fmt.Fprint(output, s...)
	applySettings()
}

// Fprintln is a replacement for fmt.Fprintln, which accepts an ansiColor as the first argument.
func Fprintln(output io.Writer, col ansiColor, s ...any) {
	fmt.Fprintf(output, "\033[0;%dm", col.value)
	if col.bold {
		fmt.Fprintf(output, "\033[1m")
	}
	fmt.Fprintln(output, s...)
	applySettings()
}

// Fprintf is a replacement for fmt.Fprintf, which accepts an ansiColor as the first argument.
func Fprintf(output io.Writer, col ansiColor, s string, args ...any) {
	fmt.Fprintf(output, "\033[0;%dm", col.value)
	if col.bold {
		fmt.Fprintf(output, "\033[1m")
	}
	fmt.Fprintf(output, s, args...)
	applySettings()
}

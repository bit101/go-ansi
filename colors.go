package ansi

import "fmt"

// Colors
type ansiColor struct {
	mod int
	fg  int
	bg  int
}

var (
	Black       ansiColor = ansiColor{0, 30, 40}
	DarkGray    ansiColor = ansiColor{1, 30, 40}
	Red         ansiColor = ansiColor{0, 31, 41}
	LightRed    ansiColor = ansiColor{1, 31, 41}
	Green       ansiColor = ansiColor{0, 32, 42}
	LightGreen  ansiColor = ansiColor{1, 32, 42}
	Brown       ansiColor = ansiColor{0, 33, 43}
	Yellow      ansiColor = ansiColor{1, 33, 43}
	Blue        ansiColor = ansiColor{0, 34, 44}
	LightBlue   ansiColor = ansiColor{1, 34, 44}
	Purple      ansiColor = ansiColor{0, 35, 45}
	LightPurple ansiColor = ansiColor{1, 35, 45}
	Cyan        ansiColor = ansiColor{0, 36, 46}
	LightCyan   ansiColor = ansiColor{1, 36, 46}
	LightGray   ansiColor = ansiColor{0, 37, 47}
	White       ansiColor = ansiColor{1, 37, 47}
	Default     ansiColor = ansiColor{0, 39, 49}

	Bold         ansiColor = ansiColor{0, 1, 1}
	NotBold      ansiColor = ansiColor{0, 22, 22}
	Underline    ansiColor = ansiColor{0, 4, 4}
	NotUnderline ansiColor = ansiColor{0, 24, 24}
	Reversed     ansiColor = ansiColor{0, 7, 7}
	NotReversed  ansiColor = ansiColor{0, 27, 27}
)

var (
	fgColor     ansiColor = Default
	bgColor     ansiColor = Default
	isBold      bool      = false
	isUnderline bool      = false
	isReversed  bool      = false
)

func SetFg(fg ansiColor) {
	fgColor = fg
	applySettings()
}

func SetBg(bg ansiColor) {
	bgColor = bg
	applySettings()
}

func SetBold(bold bool) {
	isBold = bold
	applySettings()
}

func SetUnderline(underline bool) {
	isUnderline = underline
	applySettings()
}

func SetReversed(reversed bool) {
	isReversed = reversed
	applySettings()
}

func SetAll(fg, bg ansiColor, bold, underline, reversed bool) {
	fgColor = fg
	bgColor = bg
	isBold = bold
	isUnderline = underline
	isReversed = reversed
	applySettings()
}

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
	fmt.Printf("\033[38;5;%d;%d;48;5;%d;%d;%d;%d;%dm", fgColor.mod, fgColor.fg, bgColor.mod, bgColor.bg, b, u, r)
}

func Print(col ansiColor, s ...any) {
	fmt.Print("\033[0m")
	fmt.Printf("\033[%d;%dm", col.mod, col.fg)
	fmt.Print(s...)
	applySettings()
}

func Println(col ansiColor, s ...any) {
	fmt.Print("\033[0m")
	fmt.Printf("\033[%d;%dm\033[K", col.mod, col.fg)
	fmt.Println(s...)
	applySettings()
}

func Printf(col ansiColor, s string, args ...any) {
	fmt.Print("\033[0m")
	fmt.Printf("\033[%d;%dm", col.mod, col.fg)
	fmt.Printf(s, args...)
	applySettings()
}

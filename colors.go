package ansi

import "fmt"

// Colors
type ansiColor struct {
	bold  int
	value int
}

var (
	Black      ansiColor = ansiColor{0, 30}
	BoldGray   ansiColor = ansiColor{1, 30}
	Red        ansiColor = ansiColor{0, 31}
	BoldRed    ansiColor = ansiColor{1, 31}
	Green      ansiColor = ansiColor{0, 32}
	BoldGreen  ansiColor = ansiColor{1, 32}
	Yellow     ansiColor = ansiColor{0, 33}
	BoldYellow ansiColor = ansiColor{1, 33}
	Blue       ansiColor = ansiColor{0, 34}
	BoldBlue   ansiColor = ansiColor{1, 34}
	Purple     ansiColor = ansiColor{0, 35}
	BoldPurple ansiColor = ansiColor{1, 35}
	Cyan       ansiColor = ansiColor{0, 36}
	BoldCyan   ansiColor = ansiColor{1, 36}
	White      ansiColor = ansiColor{0, 37}
	BoldWhite  ansiColor = ansiColor{1, 37}
	Default    ansiColor = ansiColor{0, 39}

	Bold         ansiColor = ansiColor{0, 1}
	NotBold      ansiColor = ansiColor{0, 22}
	Underline    ansiColor = ansiColor{0, 4}
	NotUnderline ansiColor = ansiColor{0, 24}
	Reversed     ansiColor = ansiColor{0, 7}
	NotReversed  ansiColor = ansiColor{0, 27}
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
	fmt.Printf("\033[38;5;%d;48;5;%d;%d;%d;%dm", fgColor.value, bgColor.value, b, u, r)
}

func Print(col ansiColor, s ...any) {
	fmt.Print("\033[0m")
	fmt.Printf("\033[%d;%dm", col.bold, col.value)
	fmt.Print(s...)
	applySettings()
}

func Println(col ansiColor, s ...any) {
	fmt.Print("\033[0m")
	fmt.Printf("\033[%d;%dm\033[K", col.bold, col.value)
	fmt.Println(s...)
	applySettings()
}

func Printf(col ansiColor, s string, args ...any) {
	fmt.Print("\033[0m")
	fmt.Printf("\033[%d;%dm", col.bold, col.value)
	fmt.Printf(s, args...)
	applySettings()
}

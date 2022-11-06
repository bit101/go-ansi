package ansi

import "fmt"

// Colors
type ansiColor struct {
	bold  bool
	value int
}

var (
	Black      ansiColor = ansiColor{false, 30}
	BoldBlack  ansiColor = ansiColor{true, 30}
	Red        ansiColor = ansiColor{false, 31}
	BoldRed    ansiColor = ansiColor{true, 31}
	Green      ansiColor = ansiColor{false, 32}
	BoldGreen  ansiColor = ansiColor{true, 32}
	Yellow     ansiColor = ansiColor{false, 33}
	BoldYellow ansiColor = ansiColor{true, 33}
	Blue       ansiColor = ansiColor{false, 34}
	BoldBlue   ansiColor = ansiColor{true, 34}
	Purple     ansiColor = ansiColor{false, 35}
	BoldPurple ansiColor = ansiColor{true, 35}
	Cyan       ansiColor = ansiColor{false, 36}
	BoldCyan   ansiColor = ansiColor{true, 36}
	White      ansiColor = ansiColor{false, 37}
	BoldWhite  ansiColor = ansiColor{true, 37}
	Default    ansiColor = ansiColor{false, 39}

	Bold         ansiColor = ansiColor{false, 1}
	NotBold      ansiColor = ansiColor{false, 22}
	Underline    ansiColor = ansiColor{false, 4}
	NotUnderline ansiColor = ansiColor{false, 24}
	Reversed     ansiColor = ansiColor{false, 7}
	NotReversed  ansiColor = ansiColor{false, 27}
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
	if fg.bold {
		isBold = true
	}
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
	fmt.Printf("\033[0;%d;%d;%d;%d;%dm", fgColor.value, bgColor.value+10, b, u, r)
}

func Print(col ansiColor, s ...any) {
	fmt.Printf("\033[0;%dm", col.value)
	if col.bold {
		fmt.Printf("\033[1m")
	}
	fmt.Print(s...)
	applySettings()
}

func Println(col ansiColor, s ...any) {
	fmt.Printf("\033[0;%dm", col.value)
	if col.bold {
		fmt.Printf("\033[1m")
	}
	fmt.Println(s...)
	applySettings()
}

func Printf(col ansiColor, s string, args ...any) {
	fmt.Printf("\033[0;%dm", col.value)
	if col.bold {
		fmt.Printf("\033[1m")
	}
	fmt.Printf(s, args...)
	applySettings()
}

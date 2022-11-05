package ansi

import "fmt"

// Movement / Clear
func ClearScreen() {
	fmt.Print("\033[2J")
}

func ClearLine() {
	fmt.Print("\033[2K")
}

func CarriageReturn() {
	fmt.Print("\r")
}

func NewLine() {
	fmt.Print("\n")
}

func MoveTo(x, y int) {
	fmt.Printf("\033[%d:%dH", y, x)
}

func MoveUp(n int) {
	fmt.Printf("\033[%dA", n)
}

func MoveDown(n int) {
	fmt.Printf("\033[%dB", n)
}

func MoveRight(n int) {
	fmt.Printf("\033[%dC", n)
}

func MoveLeft(n int) {
	fmt.Printf("\033[%dD", n)
}

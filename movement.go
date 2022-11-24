// Package ansi allows for advanced terminal text manipulation.
package ansi

import "fmt"

// ClearScreen clears the screen.
func ClearScreen() {
	fmt.Print("\033[H\r")
	fmt.Print("\033[2J\r")
}

// ClearLine clears the current line.
func ClearLine() {
	fmt.Print("\033[2K\r")
}

// ClearToEnd clears from the cursor onwards.
func ClearToEnd() {
	fmt.Print("\033[0J")
}

// Backspace deletes the previous character.
func Backspace(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("\b")
	}
}

// Tab inserts a tab character.
func Tab() {
	fmt.Println("\t")
}

// CarriageReturn inserts a carriage return character.
func CarriageReturn() {
	fmt.Print("\r")
}

// NewLine inserts a newline character.
func NewLine() {
	fmt.Print("\n")
}

// MoveTo moves the cursor to the specified position.
func MoveTo(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
}

// MoveUp moves the cursor up the specified number of lines.
func MoveUp(n int) {
	fmt.Printf("\033[%dA", n)
}

// MoveDown moves the cursor down the specified number of lines.
func MoveDown(n int) {
	fmt.Printf("\033[%dB", n)
}

// MoveRight moves the cursor to the right the specified number of spaces.
func MoveRight(n int) {
	fmt.Printf("\033[%dC", n)
}

// MoveLeft moves the cursor to the left the specified number of spaces.
func MoveLeft(n int) {
	fmt.Printf("\033[%dD", n)
}

// GetPos gets the current x, y position of the cursor.
func GetPos() (int, int) {
	var x int
	var y int
	fmt.Print("\033[6n")
	fmt.Scanf("\033[%d;%dR", &y, &x)
	return x, y
}

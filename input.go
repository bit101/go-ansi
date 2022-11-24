// Package ansi allows for advanced terminal text manipulation.
package ansi

import (
	"fmt"
)

// Input represents an input with a prompt and default value.
type Input struct {
	prompt string
	value  string
}

// NewInput creates a new input.
func NewInput(prompt, value string) *Input {
	return &Input{prompt, value}
}

// Run runs the input and returns the input value.
func (i *Input) Run() string {
	i.display()
	i.getInput()
	return i.value
}

func (i *Input) display() {
	ClearLine()
	Printf(Yellow, "%s: ", i.prompt)
	fmt.Print(i.value)
}

func (i *Input) getInput() {
	for {
		c := ReadBytes(1)[0]
		// fmt.Println(c)
		switch c {
		case 127: // backspace
			size := len(i.value)
			if size > 0 {
				i.value = i.value[:size-1]
				i.display()
			}

		case 10: // enter
			fmt.Println()
			return

		case 27:
			ReadBytes(2)

		default:
			if c >= 32 && c < 127 {
				i.value += string(c)
				i.display()
			}
		}
	}
}

func (i *Input) filterByte(c byte) string {
	return ""
}

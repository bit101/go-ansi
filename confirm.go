// Package ansi allows for advanced terminal text manipulation.
package ansi

import "fmt"

// Confirm prompts for a yes/no and returns a bool
type Confirm struct {
	prompt        string
	value         bool
	DefaultChoice bool
	PromptColor   ansiColor
}

// NewConfirm creates a new Confirm.
func NewConfirm(prompt string) *Confirm {
	return &Confirm{prompt, false, false, Yellow}
}

// Run runs the confirm and returns the choice.
func (c *Confirm) Run() bool {
	c.display()
	c.getInput()
	fmt.Println()
	return c.value
}

func (c *Confirm) display() {
	ClearLine()
	yn := "y/N"
	if c.DefaultChoice {
		yn = "Y/n"
	}
	Printf(c.PromptColor, "%s [%s]: ", c.prompt, yn)
}

func (c *Confirm) getInput() {
	for {
		b := ReadBytes(1)
		// fmt.Println(b)
		switch b[0] {
		case 121, 89:
			c.value = true
			c.display()
			fmt.Print("y")
		case 110, 78:
			c.value = false
			c.display()
			fmt.Print("n")
		case 10:
			return
		}
	}
}

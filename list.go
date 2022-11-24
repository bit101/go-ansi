// Package ansi allows for advanced terminal text manipulation.
package ansi

import (
	"fmt"
)

// List respresens a list of options.
type List struct {
	startX, startY int
	list           []string
	index          int
	hasShown       bool
	ShowNumbers    bool
	Prefix         string
	Postfix        string
	DefaultColor   ansiColor
	SelectedColor  ansiColor
}

// NewList creates a new tui list
func NewList(list []string) *List {
	return &List{
		list:          list,
		index:         0,
		ShowNumbers:   true,
		Prefix:        "»",
		Postfix:       "",
		DefaultColor:  Default,
		SelectedColor: Yellow,
	}
}

// Run displays a list of options and gets a result.
func (l *List) Run() (int, string) {
	l.startX, l.startY = GetPos()
	l.display()
	l.getInput()
	fmt.Println()
	return l.index, l.list[l.index]
}

func (l *List) display() {
	if l.hasShown {
		MoveUp(len(l.list))
	}
	l.hasShown = true
	for i, s := range l.list {
		ClearLine()
		col := l.DefaultColor
		pre := " "
		post := ""
		if i == l.index {
			col = l.SelectedColor
			pre = l.Prefix
			post = l.Postfix
		}
		if l.ShowNumbers {
			Printf(col, "%s %d. %s %s\n", pre, i+1, s, post)
		} else {
			Printf(col, "%s %s %s\n", pre, s, post)
		}
	}
	Printf(l.SelectedColor, "Choice: %d", l.index+1)
	MoveLeft(1)
}

func (l *List) getInput() {
	for {
		c := ReadBytes(1)[0]
		// fmt.Println(c)
		switch c {

		case 106: // j or down
			l.index++
			l.correctIndex()
			l.display()

		case 107: // k or up
			l.index--
			l.correctIndex()
			l.display()

		case 27:
			l.handleEscapes()

		case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57: // 0-9
			if l.ShowNumbers && l.checkChoice(c) {
				l.display()
			}

		case 10:
			return
		}
	}
}

func (l *List) handleEscapes() {
	b := string(ReadBytes(2))
	switch b {
	case "[B", "[C":
		l.index++
		l.correctIndex()
		l.display()

	case "[A", "[D":
		l.index--
		l.correctIndex()
		l.display()
	}

}

func (l *List) correctIndex() {
	if l.index < 0 {
		l.index = len(l.list) - 1
	}
	if l.index >= len(l.list) {
		l.index = 0
	}
}

func (l *List) checkChoice(c byte) bool {
	i := int(c - 49)
	if i >= 0 && i < len(l.list) {
		l.index = i
		return true
	}
	return false
}

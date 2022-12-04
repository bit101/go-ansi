// Package main contains some examples of using the functions in the go-ansi module.
package main

import (
	"fmt"
	"time"

	"github.com/bit101/go-ansi"
)

func main() {
	// setFuncs()
	// printFuncs()
	// moveFuncs()
	allColors()
	// mixFuncs()
}

func allColors() {
	// just printing all the colors
	ansi.Println(ansi.Black, "black")
	ansi.Println(ansi.BoldBlack, "black")

	ansi.Println(ansi.Red, "red")
	ansi.Println(ansi.BoldRed, "red")

	ansi.Println(ansi.Green, "green")
	ansi.Println(ansi.BoldGreen, "green")

	ansi.Println(ansi.Yellow, "yellow")
	ansi.Println(ansi.BoldYellow, "yellow")

	ansi.Println(ansi.Blue, "blue")
	ansi.Println(ansi.BoldBlue, "blue")

	ansi.Println(ansi.Purple, "purple")
	ansi.Println(ansi.BoldPurple, "purple")

	ansi.Println(ansi.Cyan, "cyan")
	ansi.Println(ansi.BoldCyan, "cyan")

	ansi.Println(ansi.White, "white")
	ansi.Println(ansi.BoldWhite, "white")
}

func setFuncs() {
	// build up some styles
	ansi.SetFg(ansi.Green)
	fmt.Println("green")
	ansi.SetBold(true)
	fmt.Println("green bold")
	ansi.SetReversed(true)
	fmt.Println("green bold reversed")
	ansi.SetUnderline(true)
	fmt.Println("green bold reversed underlined")

	// reset all styles
	ansi.ResetAll()
	fmt.Println("back to normal")

	// set some more
	ansi.SetBg(ansi.Red)
	ansi.SetFg(ansi.Yellow)
	fmt.Println("yellow on red")
}

func printFuncs() {
	// ansi.Print* functions are one shot. not sticky.
	ansi.Println(ansi.Red, "I am red")
	fmt.Println("I am normal")
	ansi.Println(ansi.Bold, "I am bold")
	fmt.Println("I am normal")
	ansi.Println(ansi.Reversed, "I am reversed")
	ansi.Println(ansi.Default, "I am normal")
}

func moveFuncs() {
	ansi.ClearScreen()
	for i := 1; i <= 10; i++ {
		fmt.Printf("this is line %d\n", i)
	}

	time.Sleep(2 * time.Second)
	ansi.MoveUp(8)
	ansi.ClearLine()

	time.Sleep(1 * time.Second)
	ansi.MoveDown(3)
	ansi.ClearLine()
	ansi.Print(ansi.Red, "changed to red")

	time.Sleep(1 * time.Second)
	ansi.MoveDown(2)
	ansi.CarriageReturn()
	ansi.MoveRight(8)
	ansi.Print(ansi.Green, "line 8")

	time.Sleep(1 * time.Second)
	ansi.MoveDown(3)
	ansi.CarriageReturn()
}

func mixFuncs() {
	ansi.SetFg(ansi.Red)
	ansi.SetBold(true)
	fmt.Println("red bold")

	// ansi.Print* disregards set properties
	ansi.Println(ansi.Green, "green normal")

	// but they are still in play for fmt.Print*
	fmt.Println("still red bold")
}

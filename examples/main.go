package main

import (
	"fmt"
	"time"

	"github.com/bit101/go-ansi"
)

func main() {
	// setFuncs()
	// printFuncs()
	moveFuncs()
}

func setFuncs() {
	ansi.SetFg(ansi.Green)
	fmt.Println("green")
	ansi.SetBold(true)
	fmt.Println("green bold")
	ansi.SetReversed(true)
	fmt.Println("green bold reversed")
	ansi.SetUnderline(true)
	fmt.Println("green bold reversed underlined")
	ansi.ResetAll()
	fmt.Println("back to normal")
	ansi.SetBg(ansi.Red)
	ansi.SetFg(ansi.Yellow)
	fmt.Println("yellow on red")
}

func printFuncs() {
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

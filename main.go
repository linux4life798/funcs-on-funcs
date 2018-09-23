package main

import "fmt"

const Escape = "\033["

const (
	FGColorBlack int = 30 + iota
	FGColorRed
	FGColorGreen
	FGColorYellow
	FGColorBlue
	FGColorPink
	FGColorLightBlue
	FGColorBrown
)

// ConsolePrint is the magic type that is a func
// We will use this func type to hang more function on
type ConsolePrint func(a ...interface{})

func (p ConsolePrint) DisableAll() {
	p(Escape, "0m")
}

func (p ConsolePrint) NewLine() {
	p("\n")
}

func (p ConsolePrint) ColorCode(code int) {
	p(Escape, code, "m")
}

func (p ConsolePrint) Colorln(code int, a ...interface{}) {
	p.ColorCode(code)
	p(a...)
	p.DisableAll()
	p.NewLine()
}

func (p ConsolePrint) Blackln(a ...interface{}) {
	p.Colorln(FGColorBlack, a...)
}
func (p ConsolePrint) Redln(a ...interface{}) {
	p.Colorln(FGColorRed, a...)
}
func (p ConsolePrint) Greenln(a ...interface{}) {
	p.Colorln(FGColorGreen, a...)
}
func (p ConsolePrint) Yellowln(a ...interface{}) {
	p.Colorln(FGColorYellow, a...)
}
func (p ConsolePrint) Blueln(a ...interface{}) {
	p.Colorln(FGColorBlue, a...)
}
func (p ConsolePrint) Pinkln(a ...interface{}) {
	p.Colorln(FGColorPink, a...)
}

func main() {
	var print ConsolePrint = func(a ...interface{}) {
		fmt.Print(a...)
	}

	print("Hello from main print function\n")
	print.Greenln("Hello world from the green func")
	print.Blueln("Hello world from the blue func ")
	print.Yellowln("Hello world from the yellow func")
}

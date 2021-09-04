package main

import "fmt"

const (
	esc   = "\x1b"
	reset = esc + "[0m" // reset; clears all colors and styles (to white on black)

	bold_on      = esc + "[1m" // bold on
	italics_on   = esc + "[3m" // italics on
	underline_on = esc + "[4m" // underline on
	inverse_on   = esc + "[7m" // inverse on; reverses foreground & background colors
	strike_on    = esc + "[9m" // strikethrough on

	bold_off      = esc + "[22m" // bold off
	italics_off   = esc + "[23m" // italics off
	underline_off = esc + "[24m" // underline off
	inverse_off   = esc + "[27m" // inverse off
	strike_off    = esc + "[29m" // strikethrough off

	fg_black   = esc + "[30m" // set foreground color to black
	fg_red     = esc + "[31m" // set foreground color to red
	fg_green   = esc + "[32m" // set foreground color to green
	fg_yellow  = esc + "[33m" // set foreground color to yellow
	fg_blue    = esc + "[34m" // set foreground color to blue
	fg_magenta = esc + "[35m" // set foreground color to magenta (purple)
	fg_cyan    = esc + "[36m" // set foreground color to cyan
	fg_white   = esc + "[37m" // set foreground color to white
	fg_default = esc + "[39m" // set foreground color to default (white)

	bg_black   = esc + "[40m" // set background color to black
	bg_red     = esc + "[41m" // set background color to red
	bg_green   = esc + "[42m" // set background color to green
	bg_yellow  = esc + "[43m" // set background color to yellow
	bg_blue    = esc + "[44m" // set background color to blue
	bg_magenta = esc + "[45m" // set background color to magenta (purple)
	bg_cyan    = esc + "[46m" // set background color to cyan
	bg_white   = esc + "[47m" // set background color to white
	bg_default = esc + "[49m" // set background color to default (black)
)

func main() {
	fmt.Println(fg_green, "Hello", inverse_on, "World!", reset)
}

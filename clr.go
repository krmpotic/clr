package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

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

const (
	default_mark = fg_red
)

func color(re *regexp.Regexp, colorCode string, s string) (t string) {
	i := 0
	for _, m := range re.FindAllStringIndex(s, -1) {
		t += s[i:m[0]] + colorCode + s[m[0]:m[1]] + reset
		i = m[1]
	}
	t += s[i:]
	return
}

type search struct {
	code string // ansi color code
	re   []*regexp.Regexp
}

func parseColorFlag(flag string) (string, bool) {
	switch flag {
	case "-red":
		return fg_red, true
	case "-green":
		return fg_green, true
	case "-yellow":
		return fg_yellow, true
	case "-blue":
		return fg_blue, true
	case "-magenta":
		return fg_magenta, true
	case "-cyan":
		return fg_cyan, true
	default:
		return "", false
	}
}

func main() {
	var searches []search
	for i, a := range os.Args[1:] {
		if clr, ok := parseColorFlag(a); ok {
			var m search
			m.code = clr
			searches = append(searches, m)
			continue
		} else if i == 0 {
			var m search
			m.code = default_mark
			searches = append(searches, m)
		}
		re := regexp.MustCompile(a)
		searches[len(searches)-1].re = append(searches[len(searches)-1].re, re)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l := scanner.Text()
		for _, s := range searches {
			for _, r := range s.re {
				l = color(r, s.code, l)
			}
		}
		fmt.Println(l)
	}
}

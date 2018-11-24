package color

import (
	"fmt"
	"io"
)

type color string

// ANSI color escape sequences
// https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
const (
	Black   color = "0;30"
	Red     color = "0;31"
	Green   color = "0;32"
	Yellow  color = "0;33"
	Blue    color = "0;34"
	Magenta color = "0;35"
	Cyan    color = "0;36"
	White   color = "0;37"
)

const (
	BrightBlack   color = "0;90"
	BrightRed     color = "0;91"
	BrightGreen   color = "0;92"
	BrightYellow  color = "0;93"
	BrightBlue    color = "0;94"
	BrightMagenta color = "0;95"
	BrightCyan    color = "0;96"
	BrightWhite   color = "0;97"
)

// Fprint is like fmt.Fprint but with color.
func Fprint(w io.Writer, c color, a ...interface{}) (n int, err error) {
	format := makeFormat(a)
	return fmt.Fprintf(w, "\x1B[%sm%s\x1B[0m", c, fmt.Sprintf(format, a...))
}

// Fprintf is like fmt.Fprintf but with color.
func Fprintf(w io.Writer, c color, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, "\x1B[%sm%s\x1B[0m", c, fmt.Sprintf(format, a...))
}

// Fprintln is like fmt.Fprintln but with color.
func Fprintln(w io.Writer, c color, a ...interface{}) (n int, err error) {
	format := makeFormat(a)
	return fmt.Fprintf(w, "\x1B[%sm%s\x1B[0m\n", c, fmt.Sprintf(format, a...))
}

// Sprintf is like fmt.Sprintf but with color.
func Sprintf(c color, format string, a ...interface{}) string {
	return fmt.Sprintf("\x1B[%sm%s\x1B[0m", c, fmt.Sprintf(format, a...))
}

// Printf is like fmt.Printf but with color.
func Printf(c color, format string, a ...interface{}) (n int, err error) {
	return fmt.Printf("\x1B[%sm%s\x1B[0m", c, fmt.Sprintf(format, a...))
}

// Println is like fmt.Println but with color.
func Println(c color, a ...interface{}) (n int, err error) {
	format := makeFormat(a)
	return fmt.Printf("\x1B[%sm%s\x1B[0m\n", c, fmt.Sprintf(format, a...))
}

func makeFormat(a ...interface{}) (format string) {
	for i := 0; i < len(a); i++ {
		if i == 0 {
			format += "%v"
		} else {
			format += " %v"
		}
	}
	return format
}

/*
Package color contains some functions for formatting or printing with color.

The following colors are available:
	- Black, Red, Green, Yellow, Blue, Magenta, Cyan, White,
	- BrightBlack, BrightRed, BrightGreen, BrightYellow, BrightBlue, BrightMagenta, BrightCyan, BrightWhite

The functions in this package work in the same way as their counterpart
functions in the 'fmt' package:

Example for formatting:
	// without any formatting
	coloredString := color.Sprintf(color.Yellow, "make this string yellow.")

	// with formatting
	str := "make this string"
	c   := "blue"
	coloredString := color.Sprintf(color.Blue, "%s %s", str, c)

Example for printing:
	var err error
	var s struct{}
	var sl []string

	// fmt.Println with color
	color.Println(color.Green, s, sl)

	// fmt.Printf with color
	color.Printf(color.BrightRed, "some error occurred: %v", err)
*/
package color

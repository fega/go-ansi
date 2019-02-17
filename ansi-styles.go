package ansi

const (
	noColor = "\033[0m"
	// color
	red    = "\033[0;31m"
	black  = "\033[0;30m" // Black
	green  = "\033[0;32m" // Green
	yellow = "\033[0;33m" // Yellow
	blue   = "\033[0;34m" // Blue
	purple = "\033[0;35m" // Purple
	cyan   = "\033[0;36m" // Cyan
	white  = "\033[0;37m" // White

	// Bold
	bblack  = "\033[1;30m" // Black
	bred    = "\033[1;31m" // Red
	bgreen  = "\033[1;32m" // Green
	byellow = "\033[1;33m" // Yellow
	bblue   = "\033[1;34m" // Blue
	bpurple = "\033[1;35m" // Purple
	bcyan   = "\033[1;36m" // Cyan
	bwhite  = "\033[1;37m" // White

	// Underline
	ublack  = "\033[4;30m" // Black
	ured    = "\033[4;31m" // Red
	ugreen  = "\033[4;32m" // Green
	uyellow = "\033[4;33m" // Yellow
	ublue   = "\033[4;34m" // Blue
	upurple = "\033[4;35m" // Purple
	ucyan   = "\033[4;36m" // Cyan
	uwhite  = "\033[4;37m" // White

	// Background
	onBlack  = "\033[40m" // Black
	onRed    = "\033[41m" // Red
	onGreen  = "\033[42m" // Green
	onYellow = "\033[43m" // Yellow
	onBlue   = "\033[44m" // Blue
	onPurple = "\033[45m" // Purple
	onCyan   = "\033[46m" // Cyan
	onWhite  = "\033[47m" // White

	// High Intensity
	iblack  = "\033[0;90m" // Black
	ired    = "\033[0;91m" // Red
	igreen  = "\033[0;92m" // Green
	iyellow = "\033[0;93m" // Yellow
	iblue   = "\033[0;94m" // Blue
	ipurple = "\033[0;95m" // Purple
	icyan   = "\033[0;96m" // Cyan
	iwhite  = "\033[0;97m" // White

	// Bold High Intensity
	biblack  = "\033[1;90m" // Black
	bired    = "\033[1;91m" // Red
	bigreen  = "\033[1;92m" // Green
	biyellow = "\033[1;93m" // Yellow
	biblue   = "\033[1;94m" // Blue
	bipurple = "\033[1;95m" // Purple
	bicyan   = "\033[1;96m" // Cyan
	biwhite  = "\033[1;97m" // White

)

// Red return a Red string
func Red(str string) string {
	var r = red + str + noColor

	return r
}

// Black return a Black string
func Black(str string) string {
	var r = black + str + noColor

	return r
}

// Green return a Green string
func Green(str string) string {
	var r = green + str + noColor

	return r
}

// Yellow return a Yellow string
func Yellow(str string) string {
	var r = yellow + str + noColor

	return r
}

// Blue return a Blue string
func Blue(str string) string {
	var r = blue + str + noColor

	return r
}

// Purple return a Purple string
func Purple(str string) string {
	var r = purple + str + noColor

	return r
}

// Cyan return a Cyan string
func Cyan(str string) string {
	var r = cyan + str + noColor

	return r
}

// White return a White string
func White(str string) string {
	var r = white + str + noColor

	return r
}

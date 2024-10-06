package colorize

const (
	// reset resets the color
	reset = "\033[0m"
	// red represents the red color
	red = "\033[0;31m"
	// green represents the green color
	green = "\033[0;32m"
	// yellow represents the yellow color
	yellow = "\033[0;33m"
	// cyan represents the cyan color
	cyan = "\033[0;36m"
)

// Red returns the string with red color
func Red(s string) string {
	return red + s + reset
}

// Green returns the string with green color
func Green(s string) string {
	return green + s + reset
}

// Yellow returns the string with yellow color
func Yellow(s string) string {
	return yellow + s + reset
}

// CyanColor returns the string with cyan color
func Cyan(s string) string {
	return cyan + s + reset
}

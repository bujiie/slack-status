package color

import (
	"fmt"
)

type Color string

const (
	Reset   Color = "\033[0m"
	Red           = "\033[31m"
	Green         = "\033[32m"
	Yellow        = "\033[33m"
	Blue          = "\033[34m"
	Magenta       = "\033[35m"
	Cyan          = "\033[36m"
	Gray          = "\033[37m"
	White         = "\033[97m"
)

func Colorize(value string, kuler Color) string {
	return fmt.Sprintf("%s%s%s", kuler, value, kuler)
}

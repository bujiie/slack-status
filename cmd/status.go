package main

import (
	"context"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/bujiie/slack-status/internal/color"
	"github.com/bujiie/slack-status/internal/mapping"
	"github.com/bujiie/slack-status/internal/temporal"
	"github.com/bujiie/slack-status/internal/util"
	"os"
	"strconv"
	"time"
)

var statusMapping = mapping.CharMapping{
	"o": ":building:",
	"h": ":house:",
	"p": ":palm-tree:",
	"x": ":x:",
	"v": ":coconut:",
}

func status(ctx context.Context, moment time.Time, args ...string) (*string, error) {
	argWeekOffset := 0
	argStatusPattern := ""

	switch {
	case len(args) > 1:
		n, err := strconv.Atoi(args[0])
		if err != nil {
			return nil, fmt.Errorf(color.Colorize("Error: cannot make status because offset specified does not look like a number (%s).", color.Red), argWeekOffset)
		}
		argWeekOffset = n
		argStatusPattern = args[1]
	case len(args) > 0:
		argStatusPattern = args[0]
	}

	mappedPattern := ""
	for _, char := range argStatusPattern {
		mappedPattern += statusMapping.GetMapping(ctx, string(char))
	}

	weekNumber := temporal.GetWeekNumber(ctx, moment, util.IntToPointer(argWeekOffset))
	return util.StrToPointer(fmt.Sprintf("week %d: %s", weekNumber, mappedPattern)), nil
}

func main() {
	result, err := status(context.Background(), time.Now(), os.Args[1:]...)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = clipboard.WriteAll(*result)
	if err != nil {
		fmt.Println(fmt.Errorf(color.Colorize("Error: could not copy output to clipboard.", color.Red)))
	}

	fmt.Println(*result)
}

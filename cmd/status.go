package main

import (
	"context"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/bujiie/slack-status/internal/color"
	"github.com/bujiie/slack-status/internal/config"
	"github.com/bujiie/slack-status/internal/temporal"
	"github.com/bujiie/slack-status/internal/util"
	"os"
	"strconv"
	"time"
)

func status(ctx context.Context, moment time.Time, args ...string) (*string, error) {
	cfg, err := config.ParseConfig(ctx, "./config_sample.yaml")
	if err != nil {
		return nil, fmt.Errorf("error: cannot parse config file")
	}

	argWeekOffset := 0
	argStatusPattern := ""

	switch {
	case len(args) > 1:
		n, err := strconv.Atoi(args[0])
		if err != nil {
			return nil, fmt.Errorf(color.Colorize("error: cannot make status because offset specified does not look like a number (%s).", color.Red), argWeekOffset)
		}
		argWeekOffset = n
		argStatusPattern = args[1]
	case len(args) > 0:
		argStatusPattern = args[0]
	}

	mappedPattern := ""
	for _, char := range argStatusPattern {
		mappedPattern += cfg.Get(string(char))
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
		fmt.Println(fmt.Errorf("error: could not copy output to clipboard")) // fmt.Errorf(color.Colorize("Error: could not copy output to clipboard", color.Red)))
	}

	fmt.Println(*result)
}

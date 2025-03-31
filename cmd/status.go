package main

import (
	"context"
	"fmt"
	"github.com/bujiie/slack-status/internal/color"
	"github.com/bujiie/slack-status/internal/temporal"
	"github.com/bujiie/slack-status/internal/util"
	"os"
	"strconv"
	"time"
)

func status(ctx context.Context, moment time.Time, args ...string) (*string, error) {
	offset := 0
	if len(args) > 0 {
		n, err := strconv.Atoi(args[0])
		if err != nil {
			return nil, fmt.Errorf(color.Colorize("Error: cannot make status because offset specified does not look like a number (%s).", color.Red), os.Args[1])

		}
		offset = n
	}

	weekNumber := temporal.GetWeekNumber(ctx, moment, util.IntToPointer(offset))
	return util.StrToPointer(fmt.Sprintf("week %d", weekNumber)), nil
}

func main() {
	result, err := status(context.Background(), time.Now(), os.Args[1:]...)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*result)
}

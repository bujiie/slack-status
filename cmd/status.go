package main

import (
	"context"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/bujiie/slack-status/internal/config"
	"github.com/bujiie/slack-status/internal/parse"
	"github.com/bujiie/slack-status/internal/temporal"
	"github.com/bujiie/slack-status/internal/token"
	"os"
	"strconv"
	"time"
)

func AddValues(ctx context.Context, kvPairs map[any]any) context.Context {
	localCtx := ctx
	for key, value := range kvPairs {
		localCtx = context.WithValue(localCtx, key, value)
	}
	return localCtx
}

func getWeekNumberAsString(ctx context.Context) string {
	return strconv.Itoa(temporal.GetWeekNumber(ctx))
}

func getAbbrDayAndMonth(ctx context.Context) string {
	return ctx.Value(config.MomentKey).(time.Time).Format("01 Jan")
}

func getShortMonthAndDay(ctx context.Context) string {
	return temporal.GetStartOfWeek(ctx).Format("1/2")
}

var vpt = token.ValueProviderTable{
	"week_number": getWeekNumberAsString,
	"abbr_dm":     getAbbrDayAndMonth,
	"md_number":   getShortMonthAndDay,
}

func main() {
	args := os.Args
	symbolPattern := ""
	adjustWeeks := 0

	// len(args) > 2 indicates that we have an additional parameter +/-#
	// instructing us to advance the week forward or back specific number of
	// weeks. Note that args[0] always contains the full CLI command.
	if len(args) > 2 {
		n, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		adjustWeeks += n
		symbolPattern = args[2]
	} else {
		symbolPattern = args[1]
	}

	homePath, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	var cfg config.Config
	err = config.ReadYAML(fmt.Sprintf("%s/%s", homePath, config.DefaultFileName), &cfg)
	if err != nil {
		panic(err)
	}

	ctx := AddValues(context.Background(), map[any]any{
		config.MomentKey:      time.Now().AddDate(0, 0, adjustWeeks*7),
		config.StartOfWeekKey: temporal.GetDayOfWeek(cfg.Start),
	})

	var spt token.SymbolProviderTable = cfg.Symbols

	parser := parse.NewParser(&vpt, &spt)
	value := parser.Parse(cfg.Prefix, symbolPattern, ctx)

	err = clipboard.WriteAll(value)
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
}

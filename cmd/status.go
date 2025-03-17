package main

import (
	"context"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/bujiie/slack-status/internal/config"
	"github.com/bujiie/slack-status/internal/temporal"
	"github.com/bujiie/slack-status/internal/token"
	"log"
	"os"
	"strconv"
	"time"
)

var tokenLookup = token.ValueProviderTable{
	"week_number": func(ctx context.Context) string { return strconv.Itoa(temporal.GetWeekNumber(ctx)) },
	"dd_mon":      func(ctx context.Context) string { return ctx.Value(config.MomentKey).(time.Time).Format("01 Jan") },
	"m_d":         func(ctx context.Context) string { return temporal.GetStartOfWeek(ctx).Format("1/2") },
}

func main() {
	args := os.Args
	fwd := 0
	pattern := args[1]
	if len(args) > 2 {
		i, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		fwd += i
		pattern = args[2]
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

	ctx := context.Background()
	ctx = context.WithValue(ctx, config.MomentKey, time.Now().AddDate(0, 0, fwd*7))
	ctx = context.WithValue(ctx, config.StartOfWeekKey, temporal.GetDayOfWeek(cfg.Start))

	spt := make(token.SymbolProviderTable)
	for key, value := range cfg.Symbols {
		spt[key] = value
	}
	prefix := tokenLookup.ResolveString(cfg.Prefix, ctx)
	resolvedPattern := spt.ResolveString(pattern)
	value := fmt.Sprintf("%s%s", prefix, resolvedPattern)

	err = clipboard.WriteAll(value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
}

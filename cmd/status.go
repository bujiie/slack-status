package main

import (
	"context"
	"fmt"
	"github.com/bujiie/slack-status/internal/temporal"
	"time"
)

func main() {
	ctx := context.Background()
	weekNumber := temporal.GetWeekNumber(ctx, time.Now())
	fmt.Printf("week %d\n", weekNumber)
}

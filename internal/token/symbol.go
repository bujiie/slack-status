package token

import (
	"context"
	"github.com/bujiie/slack-status/internal/config"
	"github.com/bujiie/slack-status/internal/temporal"
	"time"
)

type SymbolProviderTable map[string]string

// GetSymbol attempts to find and return the associated symbol for the given
// token. If a symbol is not found, the original token value will be returned.
func (spt SymbolProviderTable) GetSymbol(token string, _ context.Context) string {
	if symbol, exists := spt[token]; exists {
		return symbol
	}
	return token
}

// ResolvePattern attempts to replace each character in the pattern with its
// associated symbol. Note that if no symbol is found, the original character
// will be left in place.
func (spt SymbolProviderTable) ResolvePattern(pattern string, ctx context.Context) string {
	incDayOfWeekPrefix := ctx.Value(config.IncDayOfWeekPrefixKey).(bool)
	startOfWeek := ctx.Value(config.StartOfWeekKey).(time.Weekday)

	result := ""
	weekdayIndex := int(startOfWeek)
	for _, char := range pattern {
		if incDayOfWeekPrefix {
			if weekdayIndex > int(time.Friday) {
				weekdayIndex = int(time.Monday)
			}
			result += temporal.GetDayOfWeekAsString(time.Weekday(weekdayIndex))[0:2]
			weekdayIndex += 1
		}
		result += spt.GetSymbol(string(char), ctx)
		if incDayOfWeekPrefix {
			result += " "
		}
	}
	return result
}

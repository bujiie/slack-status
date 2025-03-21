package token

import (
	"context"
	"fmt"
	"github.com/bujiie/slack-status/internal/config"
	"github.com/bujiie/slack-status/internal/helper"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var symbolTokenA = "a"
var symbolTokenB = "b"

var symbolProviderTable = SymbolProviderTable{
	"a": "1",
	"b": "2",
}

func TestSymbolProviderTable_GetSymbol(t *testing.T) {
	t.Run("should get associated symbol for token", func(t *testing.T) {
		symbol := symbolProviderTable.GetSymbol(symbolTokenA, context.Background())
		assert.Equal(t, "1", symbol)
	})

	t.Run("should return the token when an associated symbol is not found", func(t *testing.T) {
		unknownToken := "unknown"
		symbol := symbolProviderTable.GetSymbol(unknownToken, context.Background())
		assert.Equal(t, unknownToken, symbol)
	})
}

func TestSymbolProviderTable_ResolvePattern(t *testing.T) {
	t.Run("should return string of symbols in order of tokens in the pattern", func(t *testing.T) {
		unknownToken := "r"
		pattern := fmt.Sprintf("%s%s%s%s", symbolTokenA, symbolTokenB, unknownToken, symbolTokenA)
		res := symbolProviderTable.ResolvePattern(pattern, context.Background())
		assert.Equal(t, "12r1", res)
	})

	t.Run("should return string of symbols with day of week character prefixed to each", func(t *testing.T) {
		pattern := fmt.Sprintf("%s%s%s", symbolTokenA, symbolTokenB, symbolTokenA)
		ctx := helper.AddValues(context.Background(), map[any]any{
			config.IncDayOfWeekPrefixKey: true,
			config.StartOfWeekKey:        time.Tuesday,
		})
		res := symbolProviderTable.ResolvePattern(pattern, ctx)
		assert.Equal(t, "Tu1We2Th1", res)
	})
}

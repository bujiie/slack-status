package token

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var symbolTokenA = "a"
var symbolTokenB = "b"

var symbolProviderTable = SymbolProviderTable{
	"a": "1",
	"b": "2",
}

func TestSymbolProviderTable_GetSymbol(t *testing.T) {
	t.Run("should get associated symbol for token", func(t *testing.T) {
		symbol := symbolProviderTable.GetSymbol(symbolTokenA)
		assert.Equal(t, "1", symbol)
	})

	t.Run("should return the token when an associated symbol is not found", func(t *testing.T) {
		unknownToken := "unknown"
		symbol := symbolProviderTable.GetSymbol(unknownToken)
		assert.Equal(t, unknownToken, symbol)
	})
}

func TestSymbolProviderTable_ResolvePattern(t *testing.T) {
	t.Run("should return string of symbols in order of tokens in the pattern", func(t *testing.T) {
		unknownToken := "r"
		pattern := fmt.Sprintf("%s%s%s%s", symbolTokenA, symbolTokenB, unknownToken, symbolTokenA)
		res := symbolProviderTable.ResolvePattern(pattern)
		assert.Equal(t, "12r1", res)
	})
}

package token

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tokenA = "tokenA"
var tokenB = "tokenB"
var valueProviderTable = ValueProviderTable{
	"tokenA": func(ctx context.Context) string { return "valueA" },
	"tokenB": func(ctx context.Context) string { return "valueB" },
}

func TestValueProviderTable_GetProvider(t *testing.T) {
	t.Run("should return the provider associated with the token", func(t *testing.T) {
		ctx := context.Background()

		providerA := valueProviderTable.GetProvider(tokenA)
		providerB := valueProviderTable.GetProvider(tokenB)

		assert.Equal(t, "valueA", providerA(ctx))
		assert.Equal(t, "valueB", providerB(ctx))
	})

	t.Run("should return the default provider when token does not match a provider", func(t *testing.T) {
		ctx := context.Background()

		provider := valueProviderTable.GetProvider("unknownToken")

		assert.Equal(t, "", provider(ctx))
	})
}

func TestValueProviderTable_Resolve(t *testing.T) {
	t.Run("should retrieve provider associated with token and execute it", func(t *testing.T) {
		ctx := context.Background()
		value := valueProviderTable.Resolve(tokenA, ctx)

		assert.Equal(t, "valueA", value)
	})

	t.Run("should retrieve default provider for and unknown token and execute it", func(t *testing.T) {
		ctx := context.Background()
		value := valueProviderTable.Resolve(tokenA, ctx)

		assert.Equal(t, "", value)
	})
}

func TestValueProviderTable_ResolvePattern(t *testing.T) {
	t.Run("should resolve all tokens in a pattern to the values from their respective providers", func(t *testing.T) {
		ctx := context.Background()
		res := valueProviderTable.ResolvePattern("{tokenA}-{tokenB}-{unknown}-{tokenA}", ctx)

		assert.Equal(t, "valueA-valueB--valueA", res)
	})
}

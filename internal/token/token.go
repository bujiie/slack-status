package token

import (
	"context"
	"regexp"
	"strings"
)

type ValueProvider func(context.Context) string

type ValueProviderTable map[string]ValueProvider

// defaultValueProvider does nothing and should always return an empty string.
// This provider should be returned if no other provider can be found for a
// given token.
var defaultValueProvider = func(_ context.Context) string { return "" }

// GetProvider returns the associated provider function from the
// ValueProviderTable if it exists. If a provider cannot be found, the default
// provider will be returned to ensure a provider is always returned.
func (vpt ValueProviderTable) GetProvider(token string) ValueProvider {
	if valueProvider, exists := vpt[token]; exists {
		return valueProvider
	}
	return defaultValueProvider
}

// Resolve finds a provider associated with a token and executes it. Note that
// we expect a provider to always be returned for any token.
func (vpt ValueProviderTable) Resolve(token string, ctx context.Context) string {
	return vpt.GetProvider(token)(ctx)
}

// ResolvePattern extracts (optional) tokens in the pattern parameter and
// attempts to resolve them to associated values. Those values then replace the
// tokens in the pattern.
func (vpt ValueProviderTable) ResolvePattern(pattern string, ctx context.Context) string {
	re := regexp.MustCompile(`{([^}]+)}`)
	matches := re.FindAllStringSubmatch(pattern, -1)

	ret := pattern
	for _, match := range matches {
		tokenWithBraces := match[0]
		token := match[1]

		value := vpt.Resolve(token, ctx)
		// only replace the first occurrence because later occurrences will be
		// handled as a separate, unrelated token.
		ret = strings.Replace(ret, tokenWithBraces, value, 1)
	}
	return ret
}

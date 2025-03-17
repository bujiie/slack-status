package token

import (
	"context"
	"regexp"
	"strings"
)

type ValueProvider func(context.Context) string

type ValueProviderTable map[string]ValueProvider

var defaultValueProvider = func(_ context.Context) string { return "" }

func (vpt ValueProviderTable) GetProvider(token string) ValueProvider {
	if valueProvider, exists := vpt[token]; exists {
		return valueProvider
	}
	return defaultValueProvider
}

func (vpt ValueProviderTable) Resolve(token string, ctx context.Context) string {
	return vpt.GetProvider(token)(ctx)
}

func (vpt ValueProviderTable) ResolveString(subject string, ctx context.Context) string {
	re := regexp.MustCompile(`{([^}]+)}`)
	matches := re.FindAllStringSubmatch(subject, -1)

	result := subject
	for _, match := range matches {
		rawToken := match[0]
		token := match[1]

		value := vpt.GetProvider(token)(ctx)
		result = strings.Replace(result, rawToken, value, 1)
	}
	return result
}

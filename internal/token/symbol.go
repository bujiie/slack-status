package token

type SymbolProviderTable map[string]string

// GetSymbol attempts to find and return the associated symbol for the given
// token. If a symbol is not found, the original token value will be returned.
func (spt SymbolProviderTable) GetSymbol(token string) string {
	if symbol, exists := spt[token]; exists {
		return symbol
	}
	return token
}

// ResolvePattern attempts to replace each character in the pattern with its
// associated symbol. Note that if no symbol is found, the original character
// will be left in place.
func (spt SymbolProviderTable) ResolvePattern(pattern string) string {
	result := ""
	for _, char := range pattern {
		result += spt.GetSymbol(string(char))
	}
	return result
}

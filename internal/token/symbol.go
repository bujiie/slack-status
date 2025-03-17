package token

type SymbolProviderTable map[string]string

func (spt SymbolProviderTable) GetSymbol(token string) string {
	if symbol, exists := spt[token]; exists {
		return symbol
	}
	return token
}

func (spt SymbolProviderTable) ResolveString(subject string) string {
	result := ""
	for _, char := range subject {
		result += spt.GetSymbol(string(char))
	}
	return result
}

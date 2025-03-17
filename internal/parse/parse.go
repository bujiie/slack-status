package parse

import (
	"context"
	"fmt"
	"github.com/bujiie/slack-status/internal/token"
)

type Parser struct {
	Vpt *token.ValueProviderTable
	Spt *token.SymbolProviderTable
}

func NewParser(vpt *token.ValueProviderTable, spt *token.SymbolProviderTable) Parser {
	return Parser{
		Vpt: vpt,
		Spt: spt,
	}
}

func (p *Parser) Parse(prefixPattern string, symbolPattern string, ctx context.Context) string {
	prefix := p.Vpt.ResolvePattern(prefixPattern, ctx)
	resolvedPattern := p.Spt.ResolvePattern(symbolPattern)
	return fmt.Sprintf("%s%s", prefix, resolvedPattern)
}

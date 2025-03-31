package mapping

import "context"

const UnknownChar = "?"

type CharMapping map[string]string

func (cm CharMapping) GetMapping(_ context.Context, char string) string {
	if val, ok := cm[char]; ok {
		return val
	}
	return UnknownChar
}

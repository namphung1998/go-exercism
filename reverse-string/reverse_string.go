package reverse

import "strings"

func Reverse(s string) string {
	var builder strings.Builder

	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		builder.WriteRune(runes[len(runes) - i - 1])
	}

	return builder.String()
}

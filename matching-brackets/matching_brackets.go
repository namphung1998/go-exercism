package brackets

func Bracket(s string) bool {
	runes := []rune(s)
	stack := make([]rune, 0, len(runes))

	mapping := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, c := range runes {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else if c == ')' || c == ']' || c == '}' {
			last := len(stack) - 1
			if len(stack) == 0 || stack[last] != mapping[c] {
				return false
			}

			stack = stack[:last]
		}
	}

	return len(stack) == 0
}

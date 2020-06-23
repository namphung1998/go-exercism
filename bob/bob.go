// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"unicode"
    "strings"
)

func IsYelling(remark string) bool {
	res := false
	for _, r := range remark {
		if unicode.IsNumber(r) {
			continue
		}

		if unicode.IsLower(r) {
			return false
		}
		if unicode.IsUpper(r) {
			res = true
		}

	}

	return res
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	runes := []rune(strings.Trim(remark, " "))
	if len(runes) > 0 {
		if IsYelling(remark) {
			if runes[len(runes)-1] == '?' {
				return "Calm down, I know what I'm doing!"
			}

			return "Whoa, chill out!"
		}
        if runes[len(runes)-1] == '?' {
            return "Sure."
        }
	}
	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		return ""
	}

	if reg.ReplaceAllString(remark, "") == "" {
		return "Fine. Be that way!"
	}

	return "Whatever."

}

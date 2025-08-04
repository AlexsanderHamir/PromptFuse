package promptfuse

import "unicode"

func isWord(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func startsWithSymbol(s string) bool {
	if s == "" {
		return false
	}
	firstRune := []rune(s)[0]
	return !unicode.IsLetter(firstRune) && !unicode.IsNumber(firstRune) && !unicode.IsSpace(firstRune)
}

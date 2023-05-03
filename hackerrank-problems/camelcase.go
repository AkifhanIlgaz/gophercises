package main

import "unicode"

func camelcase(s string) int32 {
	wordCount := 1

	for _, char := range s {
		if unicode.IsUpper(char) {
			wordCount++
		}
	}

	return int32(wordCount)

}

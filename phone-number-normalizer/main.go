package main

import (
	"strings"
)

func main() {

}

func normalize(phone string) string {
	var sb strings.Builder
	// Write benchmark for bytes.Buffer
	for _, char := range phone {
		// unicode.IsDigit(char)
		if char >= '0' && char <= '9' {
			sb.WriteRune(char)
		}

	}
	return sb.String()
}

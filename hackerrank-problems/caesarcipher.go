package main

import "unicode"

func caesarCipher(s string, k int) string {
	encryptedText := []rune{}

	for _, ch := range s {
		encryptedText = append(encryptedText, cipher(ch, k))
	}

	return string(encryptedText)
}

func cipher(char rune, delta int) rune {
	if unicode.IsUpper(char) {
		return rotate(char, 'A', delta)
	} else if unicode.IsLower(char) {
		return rotate(char, 'a', delta)
	}
	return char
}

func rotate(char rune, base, delta int) rune {
	temp := int(char) - base
	temp = (temp + delta) % 26
	return rune(temp + base)
}

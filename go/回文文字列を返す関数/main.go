package main

import (
	"fmt"
	"log"
)

func main() {
	words := []string{"abc", "level", "radar", "hello", "world"}
	fmt.Println(firstPalindromicString(words))
}

func isPalindromic(word string) bool {
	log.Println(len(word))
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}
	return true
}

func firstPalindromicString(words []string) string {
	for _, word := range words {
		if isPalindromic(word) {
			return word
		}
	}
	return ""
}

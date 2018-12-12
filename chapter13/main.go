package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		// Strings

		//Checks if haystack contains needle
		strings.Contains("test", "es"),
		// true

		//counts how many needles are in haystack
		strings.Count("test", "t"),
		// 2

		//checks if string starts with substring
		strings.HasPrefix("test", "te"),
		// true

		//checks if string ends with substring
		strings.HasSuffix("test", "st"),
		// true

		// returns the index of the substring in the string
		strings.Index("test", "e"),
		// 2

		// joins an array of strings on a substring
		strings.Join([]string{"t", "e", "s", "t"}, "-"),
		// "t-e-s-t"

		//returns a string consisting of the given string repeted n times
		strings.Repeat("A", 5),
		// "AAAAA"

		// replaces the first n occurences of substring1 with substring2
		strings.Replace("aaaaaaa", "a", "b", 3),
		// "bbbaaaa"

		// splits a string into an array of substrings on a given substring
		strings.Split("Split this sentance into words", " "),
		// []string{"Split", "this", "sentance", "into", "words"}

		// converts all the letters in a string to lower case
		strings.ToLower("AAAAAAHHHH"),
		// "aaaaaahhhh"

		// converts all the letters in a string to upper case
		strings.ToUpper("aaaaaahhhh"),
		// "AAAAAAHHHH"
	)
}

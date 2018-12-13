package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Strings

	fmt.Println(
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

	// Files & Folders

	//create a file
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	// remember to close it
	defer file.Close()

	//write to a file
	file.WriteString("Hello World")

	// open a file and read it
	// file, err := os.Open("test.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// // remember to close it
	// defer file.Close()

	// // get the file size
	// stat, err := file.Stat()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// size := stat.Size()

	// read the file
	// bs := make([]byte, size)
	// _, err = file.Read(bs)

	//the above can be shortened to:
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	str := string(bs)
	fmt.Println(str)

	// open a directory
	dir, err := os.Open(".")
	if err != nil {
		fmt.Println(err.Error())
	}
	//remember to close it
	defer dir.Close()

	//read the conents of the directory
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	//recursively walk through a directory
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

	// Errors
	err = errors.New("New Error!!!!")
	fmt.Println(err)
}

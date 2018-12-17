package main

import (
	"container/list"
	"crypto/sha1"
	"errors"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
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

	// Containers & Sort
	// List
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	// Sort
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)
	sort.Sort(ByAge(kids))
	fmt.Println(kids)

	// Hashes & Cryptography
	//non cryptogtaphic hash
	crc := crc32.NewIEEE()
	crc.Write([]byte("test."))
	crcSum := crc.Sum32()
	fmt.Println(crcSum)

	//cryptogtaphic hash
	sha := sha1.New()
	sha.Write([]byte("test"))
	shaSum := sha.Sum([]byte{})
	fmt.Println(shaSum)
}

//for Sort
type Person struct {
	Name string
	Age  int
}
type ByName []Person

func (this ByName) Len() int {
	return len(this)
}

func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type ByAge []Person

func (this ByAge) Len() int {
	return len(this)
}

func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}

func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

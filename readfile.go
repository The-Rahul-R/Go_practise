package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// write content to a file
func main() {
	fmt.Println("program to read and write into files")
	info := "Hello gang! How are you doing"

	file1, err := os.Create("./wfile.txt")
	errorcheck(err)
	len, err := io.WriteString(file1, info)
	fmt.Println("the length is", len)
	file1.Close()
	fmt.Println("provide absolute path of file as a string")
	var filepath string
	fmt.Scanln(&filepath)
	readfile(filepath)
}

//read content from a file
func readfile(filename string) {

	readlines, err := ioutil.ReadFile(filename)
	errorcheck(err)
	fmt.Println("the data inside the file is", string(readlines))
}

func errorcheck(err error) {
	if err != nil {
		panic(err)
	}
}

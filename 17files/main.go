package main

import (
	"fmt"
	"io/ioutil"
	"os"
	
)

func main() {
	fmt.Println("Welcome to files in GoLang")
	content := "This should go in a file - LearCodeOnline.in"
	

	file, err := os.Create("./mylcogofile.txt") // os is a package for file handling in go

	checkNilError(err)

	length, err := io.WriteString(file, content) // this return length (read docs)

	checkNilError(err)
	fmt.Println("The length is: ", length)
	defer file.Close() // important to close file, use defer

	readFile("../mylcogofile.txt")

}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename) // the data always read in a byte format
	checkNilError(err)
	fmt.Println("Text from the file \n", string(dataByte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

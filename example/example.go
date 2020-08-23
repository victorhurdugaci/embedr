package main

import (
	"fmt"
	"io/ioutil"

	"github.com/victorhurdugaci/embedr"
)

// Replace "../cmd/embedr" with "embedr" in a real application

//go:generate go run ../cmd/embedr -include *.txt -include **/*.md -package main

func main() {
	embedr.Walk(func(filePath string) error {
		fmt.Printf("File %s:\n", filePath)
		printFileContent(filePath)
		fmt.Println()
		return nil
	})
}

func printFileContent(filePath string) {
	reader, err := embedr.Open(filePath)
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

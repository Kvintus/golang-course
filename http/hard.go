package main

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(filename string) {
	contents, err := os.Open(filename)

	if err != nil {
		fmt.Println("File cannot be opened.")
		os.Exit(1)
	}

	io.Copy(os.Stdout, contents)
}

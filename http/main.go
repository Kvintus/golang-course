package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func mainVideo() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func main() {
	fileName := os.Args[1]
	ReadFile(fileName)
}
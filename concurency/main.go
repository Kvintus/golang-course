package main

import (
	"fmt"
	"net/http"
	"time"
)

func siteStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		return
	}

	fmt.Println(link, "is up")
	c <- link

}

func checkAll() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	c := make(chan string)
	for _, link := range links {
		go siteStatus(link, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(time.Second * 5)
			siteStatus(l, c)
		}(l)
	}
}

func main() {
	timeIt(checkAll)
}

func timeIt(f func()) {
	start := time.Now()
	f()
	fmt.Println((time.Now()).Sub(start))
}

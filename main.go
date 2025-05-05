package main

import (
	"fmt"
	"os"
)

func main() {
	urlArgument := os.Args[1:]

	if len(urlArgument) < 1 {
		fmt.Println("no website provided")
		return
	}
	if len(urlArgument) > 1 {
		fmt.Println("too many arguments provided")
		return
	}
	rawBaseURL := os.Args[1]

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	pages := make(map[string]int)

	crawlPage(rawBaseURL, rawBaseURL, pages)

	for normalizedURL, count := range pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}

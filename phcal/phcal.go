package main

import (
	"flag"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func readWebEvents(url string, foundEvents map[string]string, debug bool) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return
	}

	htmlBody := resp.Body
	defer htmlBody.Close() // close Body when the function returns

	splitIntoTokens := html.NewTokenizer(htmlBody)

	for {
		tempToken := splitIntoTokens.Next()

		switch {
		case tempToken == html.ErrorToken:
			// End of document
			return

		case tempToken == html.StartTagToken:
			currentToken := splitIntoTokens.Token()

			fmt.Println("Start Tag Token:", currentToken.Data)
			for _, a := range currentToken.Attr {
				fmt.Println("Attribute:", a.Key, a.Val)
			}

		case tempToken == html.TextToken:
			currentToken := splitIntoTokens.Token()

			fmt.Println("Text Tag Token:", currentToken.Data)
		}
	}
}

func main() {
	var debug bool

	foundEvents := make(map[string]string)

	urlPtr := flag.String("url", "https://www.phoenixrunning.co.uk/events", "A URL to be examined")
	flag.BoolVar(&debug, "debug", false, "turns print debugging on")

	flag.Parse()

	readWebEvents(*urlPtr, foundEvents, debug)
}

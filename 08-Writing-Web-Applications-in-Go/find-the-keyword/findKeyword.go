// a handy command-line utility that reads a number of web
// pages, which can be found in a text file or read from
// standard input, and returns the number of times a given
// keyword was found in these web pages.
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

// The Data struct type will be used for passing information
// between channels.
type Data struct {
	URL     string
	Keyword string
	Times   int
	Error   error
}

func monitor(values <-chan Data, count int) {
	for i := 0; i < count; i++ {
		x := <-values
		if x.Error == nil {
			fmt.Printf("\t%s\t", x.Keyword)
			fmt.Printf("\t%d\t in\t%s\n", x.Times, x.URL)
		} else {
			fmt.Printf("\t%s\n", x.Error)
		}
	}
}

func processPage(myUrl, keyword string, out chan<- Data) {
	var err error
	times := 0

	URL, err := url.Parse(myUrl)
	if err != nil {
		out <- Data{URL: myUrl, Keyword: keyword, Times: 0, Error: err}
		return
	}

	c := &http.Client{}
	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		out <- Data{URL: myUrl, Keyword: keyword, Times: 0, Error: err}
		return
	}

	httpData, err := c.Do(req)
	if err != nil {
		out <- Data{URL: myUrl, Keyword: keyword, Times: 0, Error: err}
		return
	}

	bodyHTML := ""

	var buffer [1024]byte
	reader := httpData.Body
	for {
		n, err := reader.Read(buffer[:])
		if err != nil {
			break
		}

		bodyHTML += string(buffer[0:n])
	}

	regExp := keyword
	r := regexp.MustCompile(regExp)
	matches := r.FindAllString(bodyHTML, -1)
	times += len(matches)

	newValue := Data{URL: myUrl, Keyword: keyword, Times: times, Error: nil}
	out <- newValue
}

func main() {
	filename := ""
	var f *os.File
	var keyword string

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Not enough arguments!")
		os.Exit(-1)
	}

	if len(arguments) == 2 {
		f = os.Stdin
		keyword = arguments[1]
	} else {
		keyword = arguments[1]
		filename = arguments[2]
		fileHandler, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening %s: %s", filename, err)
			os.Exit(1)
		}
		f = fileHandler
	}
	defer f.Close()

	values := make(chan Data, len(os.Args[1:]))

	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		count += 1
		go func(URL string) {
			processPage(URL, keyword, values)
		}(scanner.Text())
	}

	monitor(values, count)
}

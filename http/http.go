package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

var (
	header string
	method string
	data   string
)

func main() {
	// -hオプション用文言
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage of %s:
   %s [OPTIONS] ARGS...
Options
`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	flag.StringVar(&header, "header", "", "HTTP request header")
	flag.StringVar(&method, "X", "GET", "HTTP request method")
	flag.StringVar(&data, "d", "", "HTTP request payload")
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatal("1 argument should be given as URL.")
	}

	url := format(flag.Args()[0])

	conHTTP(url.String())
}

func format(str string) *url.URL {

	// Add http scheme if it is not included in url.
	// TODO: support https
	if !regexp.MustCompile(`http://*`).MatchString(str) {
		str = "http://" + str
	}

	u, err := url.Parse(str)
	if err != nil {
		log.Fatal(err.Error())
	}
	return u
}

func conHTTP(url string) {

	req, err := http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		log.Fatal(err.Error())
	}

	client := &http.Client{}
	h := regexp.MustCompile(`\s*:\s*`).Split(header, 2)

	if len(h) >= 2 {
		req.Header.Add(h[0], h[1])

		// if request payload exists, set Content-Type "application/x-www-form-urlencoded"
		if data != "" && h[0] != "Content-Type" {
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(body))
}

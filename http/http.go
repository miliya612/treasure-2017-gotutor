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
	fmt.Println(url)
	fmt.Println(url.String())

	conHTTP(url.String())
}

func format(str string) *url.URL {
	u, err := url.Parse(str)
	if err != nil {
		log.Fatal(err.Error())
	}
	if u.IsAbs() {
		u.Scheme = "http"
	}
	return u
}

func conHTTP(url string) {

	req, err := http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		log.Fatal(err.Error())
	}

	// if request payload exists, set Content-Type ""
	if data != "" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	fmt.Println(url)

	client := &http.Client{}
	h := regexp.MustCompile(`\s*:\s*`).Split(header, 2)

	req.Header.Add(h[0], h[1])

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

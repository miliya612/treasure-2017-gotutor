package main

import (
	"bytes"
	"net/http"

	"log"

	"io/ioutil"

	"encoding/json"

	"golang.org/x/net/html"
)

type Page struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u := r.URL.Query().Get("url")
		body := fetchFrom(u)
		p := extractPage(body)

		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		if err := enc.Encode(p); err != nil {
			log.Fatal(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}

func fetchFrom(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	return body
}

func extractPage(h []byte) *Page {
	p := &Page{}
	doc, err := html.Parse(bytes.NewReader(h))
	if err != nil {
		log.Fatal(err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			p.Title = n.FirstChild.Data
		}
		if n.Type == html.ElementNode && n.Data == "meta" {
			if isDescription(n.Attr) {
				for _, a := range n.Attr {
					if a.Key == "content" {
						p.Description = a.Val
						break
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return p
}

func isDescription(s []html.Attribute) bool {
	for _, v := range s {
		if v.Key == "name" && v.Val == "description" {
			return true
		}
	}
	return false
}

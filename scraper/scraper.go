package main

import (
	"bytes"
	"net/http"

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
		if u == "" {
			http.Error(w, "url not specified", http.StatusBadRequest)
			return
		}
		err, body := fetchFrom(u)
		if err != nil {
			http.Error(w, "request failed", http.StatusInternalServerError)
			return
		}
		err, p := extractPage(body)
		if err != nil {
			http.Error(w, "parse failed", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		if err := enc.Encode(p); err != nil {
			http.Error(w, "encoding failed", http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(":8080", nil)
}

func fetchFrom(url string) (error, []byte) {
	resp, err := http.Get(url)
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	return nil, body
}

func extractPage(h []byte) (error, *Page) {
	p := &Page{}
	doc, err := html.Parse(bytes.NewReader(h))
	if err != nil {
		return err, nil
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
	return nil, p
}

func isDescription(s []html.Attribute) bool {
	for _, v := range s {
		if v.Key == "name" && v.Val == "description" {
			return true
		}
	}
	return false
}

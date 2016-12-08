package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func idnesArticles() chan article {
	c := make(chan article)

	go func() {
		defer close(c)

		idnesArticles, err := fetchIdnesArticleLinks()
		if err != nil {
			fmt.Printf("fetching idnes rss feed failed with error: %s", err.Error())
			return
		}

		var wg sync.WaitGroup
		for _, link := range idnesArticles {
			wg.Add(1)
			go func(l string) {
				defer wg.Done()
				a, err := fetchIdnesArticle(l)
				if err != nil {
					fmt.Printf("fetching url %s failed with error: %s", l, err.Error())
					return
				}

				c <- a
			}(link)
		}
		wg.Wait()
	}()

	return c
}

func fetchIdnesArticleLinks() ([]string, error) {
	feed, err := fetchRSSFeed("http://servis.idnes.cz/rss.aspx")
	if err != nil {
		return nil, err
	}

	links := []string{}
	for _, item := range feed.Channel.Items {
		u, err := url.Parse(item.Link)
		if err != nil {
			continue
		}

		u.Fragment = ""
		q := u.Query()
		q.Set("setver", "mauto")
		u.RawQuery = q.Encode()

		links = append(links, u.String())
	}

	return links, nil
}

func fetchIdnesArticle(url string) (article, error) {
	var a article

	a.URL = url
	resp, err := http.Get(url)
	if err != nil {
		return a, err
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return a, err
	}
	defer resp.Body.Close()

	a.Title, a.Text = parseIdnesArticle(doc)
	a.Words = strings.Split(strings.Repeat(a.Title, 5)+a.Text, " ")

	return a, nil
}

func parseIdnesArticle(n *html.Node) (string, string) {
	title, text := "", ""

	var f func(n *html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "h1":
				if n.FirstChild != nil {
					title = strings.Trim(n.FirstChild.Data, " \n")
					break
				}
			case "p":
				if n.FirstChild != nil {
					text += strings.Trim(n.FirstChild.Data, " \n")
					break
				}
			case "div":
				for _, a := range n.Attr {
					if a.Key == "class" && a.Val == "opener" {
						if n.FirstChild != nil {
							text += strings.Trim(n.FirstChild.Data, " \n")
							break
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)

	return title, text
}

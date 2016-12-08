package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func novinkyArticles() chan article {
	c := make(chan article)

	go func() {
		defer close(c)

		novinkyArticles, err := fetchNovinkyArticleLinks()
		if err != nil {
			fmt.Printf("fetching novinky rss feed failed with error: %s", err.Error())
			return
		}

		var wg sync.WaitGroup
		for _, link := range novinkyArticles {
			wg.Add(1)
			go func(l string) {
				defer wg.Done()
				a, err := fetchNovinkyArticle(l)
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

func fetchNovinkyArticleLinks() ([]string, error) {
	feed, err := fetchRSSFeed("https://www.novinky.cz/rss2")
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile("[0-9]+")

	links := []string{}
	for _, i := range feed.Channel.Items {
		ids := re.FindAllString(i.GUID, -1)
		if len(ids) == 1 {
			links = append(links, fmt.Sprintf("https://m.novinky.cz/articleDetails?aId=%s", ids[0]))
		}
	}

	return links, nil
}

func fetchNovinkyArticle(url string) (article, error) {
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

	a.Title, a.Text = parseNovinkyArticle(doc)
	a.Words = strings.Split(strings.Repeat(a.Title, 5)+a.Text, " ")

	return a, nil
}

func parseNovinkyArticle(n *html.Node) (string, string) {
	title, text := "", ""

	var f func(n *html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "h2":
				for _, a := range n.Attr {
					if a.Key == "class" && a.Val == "artCaption" {
						if n.FirstChild != nil && n.FirstChild.NextSibling != nil {
							title = strings.Trim(n.FirstChild.NextSibling.Data, " \n")
							break
						}
					}
				}

			case "span":
				for _, a := range n.Attr {
					if a.Key == "class" && a.Val == "perex" {
						if n.FirstChild != nil && n.FirstChild.NextSibling != nil {
							text += strings.Trim(n.FirstChild.NextSibling.Data, " \n")
							break
						}
					}
				}

			case "p":
				if len(n.Attr) == 0 {
					if n.FirstChild != nil {
						text += strings.Trim(n.FirstChild.Data, " \n")
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)

	text = strings.Replace(text, "\u00a0", " ", -1)

	return title, text
}

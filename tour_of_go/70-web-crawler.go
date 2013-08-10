// Exercise: Web Crawler
//
// In this exercise you'll use Go's concurrency features to parallelize a web crawler.
//
// Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.

package main

import (
	"fmt"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, feed chan []string, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	feed <- urls

	// close(found)
}

func main() {
	new_c := make(chan string)
	uniq_c := make(chan string)
	feed := make(chan []string)
	ex := make(chan int)

	go func() {
		visited := make(map[string]bool)
		for url := range new_c {
			if !visited[url] {
				visited[url] = true
				uniq_c <- url
			} else {
				fmt.Printf("Already visited %s\n", url)
			}
		}
	}()

	go func() {
		for batch := range feed {
			for _, u := range batch {
				new_c <- u
			}
		}
	}()

	go (func() {
		for url := range uniq_c {
			go Crawl(url, feed, fetcher)
		}
	})()

	new_c <- "http://golang.org/"

	<-ex
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		fmt.Printf("Fetching %s\n", url)
		time.Sleep(1001 * time.Millisecond)
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

//Time-stamp: <2017-01-28 02:11:53 hamada>
package exer

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan CrawlResult) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	if depth <= 0 {
		ch <- CrawlResult{}
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		ch <- CrawlResult{Err: err}
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	ch <- CrawlResult{URLs: urls, Crawled: url, Depth: depth}
	return
}

type NextCrawl struct {
	URL   string
	Depth int
}

type CrawlResult struct {
	URLs    []string
	Crawled string
	Depth   int
	Err     error
}

func runCrawl() {
	ch := make(chan CrawlResult)
	nexts := []NextCrawl{NextCrawl{"http://golang.org/", 4}}
	crawleds := make(map[string]int)

	for len(nexts) > 0 {
		goroutines := 0
		for _, n := range nexts {
			if crawleds[n.URL] > 0 {
				continue
			}

			go Crawl(n.URL, n.Depth, fetcher, ch)
			goroutines++
			fmt.Printf("goroutine %d start\n", goroutines)
		}

		nexts = make([]NextCrawl, 0)
		for i := 0; i < goroutines; i++ {
			ret := <-ch
			if ret.Err != nil {
				fmt.Println(ret.Err)
			}
			if ret.Crawled == "" {
				continue
			}

			crawleds[ret.Crawled]++
			for _, url := range ret.URLs {
				nexts = append(nexts, NextCrawl{url, ret.Depth - 1})
			}
		}
	}
}

func Web() {

	if true {
		runCrawl() // ("http://golang.org/", 4, fetcher)
	}

	if false {
		fmt.Printf("v:\t %v\n", fetcher)
		fmt.Printf("T:\t %T\n", fetcher)

		check := func(url string) {
			fmt.Printf("v:\t %v\n", fetcher[url])
			fmt.Printf("T:\t %T\n", fetcher[url])
		}

		check("http://golang.org/")
		check("http://golang.org/pkg/")
		check("http://golang.org/pkg/fmt/")
		check("http://golang.org/pkg/os/")
	}
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
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

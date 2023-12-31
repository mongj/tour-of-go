package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Cache struct {
	mu sync.Mutex
	urls  map[string]int
}

func (c *Cache) Insert(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.urls[key] = 1
}

func (c *Cache) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.urls[key]
}

var wg sync.WaitGroup

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, c Cache) {
	c.Insert(url)
	
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	// simulate slow request
	time.Sleep(time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if c.Value(u) != 1 {
			wg.Add(1)
			go func(u string) {
				defer wg.Done()
				Crawl(u, depth-1, fetcher, c)
			}(u)
		}
	}
	return
}

func main() {
	c := Cache{urls: make(map[string]int)}
	Crawl("https://golang.org/", 4, fetcher, c)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
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
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

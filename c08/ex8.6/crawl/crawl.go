package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"gopl.io/ch5/links"
)

var tokens = make(chan struct{}, 20)

type linkList struct {
	urls  []string
	depth int
}

func main() {
	depth := flag.Int("depth", 3, "-depth=<number>")
	first := flag.String("first", "", "-first=<url to start from>")
	flag.Parse()

	if *first == "" {
		log.Println("Must provide valid starting point")
	}
	startLinks := strings.Split(*first, ",")

	worklist := make(chan linkList)
	var n int

	n++
	go func() { worklist <- linkList{urls: startLinks, depth: 0} }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list.urls {
			if !seen[link] {
				seen[link] = true

				if list.depth <= *depth {
					n++
					go func(link string) {
						worklist <- crawl(link, list.depth)
					}(link)
				}
			}
		}
	}
}

func crawl(url string, depth int) linkList {
	fmt.Println(url, depth)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		fmt.Println(err)
	}
	return linkList{urls: list, depth: depth + 1}
}

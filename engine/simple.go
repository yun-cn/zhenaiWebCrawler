package engine

import (
	"crawler/fetcher"
	"log"
)

// SimpleEngine Type
type SimpleEngine struct{}

// Run Seeds
func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
		log.Printf("Fetching %s\n", r.URL)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		ParseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, ParseResult.Requests...)

		for _, item := range ParseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s\n", r.URL)
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetcher: error fetching URL %s: %v", r.URL, err)
	}
	return r.ParserFunc(body), nil
}

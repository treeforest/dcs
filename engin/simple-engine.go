package engine

import (
	"github.com/treeforest/dcs/fetcher"
	"log"
	)

func Run(seeds ...Request) {
	var requests []Request

	for _, s := range seeds {
		requests = append(requests, s)
	}

	for len(requests) > 0 {
		r := requests[0]

		requests = requests[1:]

		log.Printf("Fetching url: %s", r.Url)
		body, err := fetcher.WebFetch(r.Url)

		if err != nil {
			log.Printf("WebFetch error: %s", r.Url)
		}

		parseResult := r.ParseFunc(body)

		requests = append(requests, parseResult.Reqs...)

		for _, item := range parseResult.Items {
			log.Printf("Got item: %s\n", item)
		}
	}
}

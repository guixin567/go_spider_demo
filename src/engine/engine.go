package engine

import (
	"fmt"
	"log"
	"zhengxin2/src/fetch"
	"zhengxin2/src/model"
)

func Run(seeds ...model.Request) {
	requests := seeds
	count := 0
	for len(requests) > 0 {
		request := requests[0]
		bytes, err := fetch.Fetch(request.Url)
		if err != nil {
			fmt.Printf("fetch error url %s", request.Url)
			continue
		}
		parseResult := request.ParseFunction(bytes)
		for _, item := range parseResult.Items {
			count++
			log.Printf("get index %d item,value %v", count, item)
		}
		requests = requests[1:]
		requests = append(requests, parseResult.Requests...)

	}
}

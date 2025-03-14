package main

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"
)

type MeatCounter map[string]int

type Response struct {
	Beef MeatCounter `json:"beef"`
}

func countMeats(text string) MeatCounter {
	if text == "" {
		return MeatCounter{}
	}

	meats := strings.FieldsFunc(text, func(r rune) bool {
		return r == ' ' || r == ',' || r == '.' || r == '\n'
	})

	numWorkers := 4

	results := make(chan MeatCounter, numWorkers)

	chunkSize := (len(meats) + numWorkers - 1) / numWorkers
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		start := i * chunkSize
		end := min(start+chunkSize, len(meats))

		go func(meatsChunk []string) {
			defer wg.Done()
			localCounter := make(MeatCounter)
			for _, meat := range meatsChunk {
				meat = strings.ToLower(strings.TrimSpace(meat))
				localCounter[meat]++
			}
			results <- localCounter
		}(meats[start:end])
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	finalCounter := make(MeatCounter)
	for localCounter := range results {
		for meat, count := range localCounter {
			finalCounter[meat] += count
		}
	}

	return finalCounter
}

func fetchBacon() (string, error) {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.New("failed to fetch data from Bacon Ipsum API")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("failed to read response body")
	}

	if len(body) == 0 {
		return "", errors.New("API returned empty response")
	}

	return string(body), nil
}

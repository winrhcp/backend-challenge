package main

import (
	"errors"
	"io"
	"net/http"
	"strings"
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

	counter := make(MeatCounter)
	for _, meat := range meats {
		meat = strings.ToLower(strings.TrimSpace(meat))
		counter[meat]++
	}

	return counter
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

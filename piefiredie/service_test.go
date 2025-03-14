package main

import (
	"testing"
)

func TestCountMeats(t *testing.T) {
	text := "bacon bacon beef chicken bacon pork,. pork"
	expected := MeatCounter{
		"bacon":   3,
		"beef":    1,
		"chicken": 1,
		"pork":    2,
	}

	result := countMeats(text)

	for meat, count := range expected {
		if result[meat] != count {
			t.Errorf("Expected %d of %s, but got %d", count, meat, result[meat])
		}
	}
}

func TestCountMeats_EmptyInput(t *testing.T) {
	result := countMeats("")
	if len(result) != 0 {
		t.Errorf("Expected empty map, got %v", result)
	}
}

func TestFetchBacon_EmptyResponse(t *testing.T) {
	fakeFetch := func() (string, error) {
		return "", nil
	}

	result, err := fakeFetch()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "" {
		t.Errorf("Expected empty string, got %s", result)
	}
}

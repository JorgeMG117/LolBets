package test

import (
	"net/http"
	"testing"
)

func TestFetchGophers(t *testing.T) {
	_, err := http.NewRequest("GET", "/gophers", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
}

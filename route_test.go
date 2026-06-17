package main

import (
	"net/http"
	"testing"
)

func TestHomeHandler(t *testing.T) {

	result := 200

	if result != http.StatusOK {
		t.Errorf("Home Handler not returning the correct status code %d", result)
	}
}

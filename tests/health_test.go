//go:build e2e
// +build e2e

package tests

import (
	"github.com/go-resty/resty/v2"
	"testing"
)

func TestHealthCheckEndpoint(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/alive")
	if err != nil {
		t.Errorf("Could not call health endpoint: %q", err)
	}

	want := 200

	if resp.StatusCode() != want {
		t.Errorf("Wanted statuscode %d, got %d", want, resp.StatusCode())
	}
}

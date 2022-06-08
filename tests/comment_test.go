//go:build e2e
// +build e2e

package tests

import (
	"github.com/go-resty/resty/v2"
	"testing"
)

func TestPostComment(t *testing.T) {
	t.Run("can post comment", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.IUYRn6InghAeQrSJgK-s6h8E_EWAbD6KPOywx0FAacc").
			SetBody(`{"slug": "/", "author": "Bob", "body": "hello world"}`).
			Post("http://localhost:8080/api/v1/comment")

		if err != nil {
			t.Errorf("could not post comment: %q", err)
		}

		want := 200

		if resp.StatusCode() != want {
			t.Errorf("wanted responsecode %d, got %d", want, resp.StatusCode())
		}
	})

	t.Run("cannot post comment without authentication", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().SetBody(`{"slug": "/", "author": "Bob", "body": "hello world"}`).Post("http://localhost:8080/api/v1/comment")

		if err != nil {
			t.Errorf("could not post comment: %q", err)
		}

		want := 401

		if resp.StatusCode() != want {
			t.Errorf("wanted responsecode %d, got %d", want, resp.StatusCode())
		}
	})
}

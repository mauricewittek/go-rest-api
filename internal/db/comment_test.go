//go:build integration
// +build integration

package db

import (
	"context"
	"github.com/mauricewittek/go-rest-api/internal/comment"
	"reflect"
	"testing"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		db, err := NewDatabase()
		if err != nil {
			t.Errorf("Could not create the database")
		}

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		if err != nil {
			t.Errorf("Could not post comment")
		}

		newCmt, err := db.GetComment(context.Background(), cmt.ID)
		if err != nil {
			t.Errorf("Could not get comment with id %s", cmt.ID)
		}

		if !reflect.DeepEqual(cmt, newCmt) {
			t.Errorf("got %v, want %v", cmt, newCmt)
		}
	})

	t.Run("test delete comment", func(t *testing.T) {
		db, err := NewDatabase()
		if err != nil {
			t.Errorf("Could not create the database")
		}

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "new-slug",
			Author: "new-author",
			Body:   "new-body",
		})
		if err != nil {
			t.Errorf("Could not post comment")
		}

		err = db.DeleteComment(context.Background(), cmt.ID)
		if err != nil {
			t.Errorf("Could not delete comment with id %s", cmt.ID)
		}

		_, err = db.GetComment(context.Background(), cmt.ID)
		if err == nil {
			t.Errorf("Expected an error getting a comment that should be deleted but didn't get one")
		}
	})
}

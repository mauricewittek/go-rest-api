package main

import (
	"context"
	"fmt"

	"github.com/mauricewittek/go-rest-api/internal/comment"
	"github.com/mauricewittek/go-rest-api/internal/db"
)

// This function is going to be responsible for the instantiation and startup of the application
func Run() error {
	fmt.Println("starting up the application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)

	fmt.Println(cmtService.PostComment(context.Background(), comment.Comment{ID: "0cde9d47-9858-48b1-bc71-65d4a19eb316", Slug: "manual-test", Author: "me", Body: "hello-world"}))
	fmt.Println(cmtService.GetComment(context.Background(), "0cde9d47-9858-48b1-bc71-65d4a19eb316"))

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}

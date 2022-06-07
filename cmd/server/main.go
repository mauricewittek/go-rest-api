package main

import (
	"fmt"

	"github.com/mauricewittek/go-rest-api/internal/comment"
	"github.com/mauricewittek/go-rest-api/internal/db"
	transportHttp "github.com/mauricewittek/go-rest-api/internal/transport/http"
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

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}

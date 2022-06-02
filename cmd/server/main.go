package main

import "fmt"

// This function is going to be responsible for the instantiation and startup of the application
func Run() error {
	fmt.Println("starting up the application")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}

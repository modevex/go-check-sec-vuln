package main

import "fmt"
import "github.com/google/go-github/v62/github"

func main() {
	client := github.NewClient(nil)

	fmt.Printf("%+v\n", client)
}

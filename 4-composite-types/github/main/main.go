package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hanchiang/the-go-programming-language/4-composite-types/github/githubissues"
)

func main() {
	result, err := githubissues.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

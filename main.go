package main

import (
	"context"
	"fmt"

	"github.com/ndaDayo/devdev/api"
)

func main() {
	client := api.NewClient()
	ctx := context.Background()

	owner := "ndaDayo"
	repo := "CompilerInGo"

	commits, resp, err := client.Commits.Get(ctx, owner, repo)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Printf("resp: %+v\n", resp)

	for _, commit := range *commits {
		fmt.Printf("SHA: %s\n", commit.SHA)
	}
}

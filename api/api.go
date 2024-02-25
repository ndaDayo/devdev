package api

import (
	"context"
	"fmt"
)

func GetResource(resource interface{}) (interface{}, error) {
	client := NewClient(WithToken())
	ctx := context.Background()

	switch r := resource.(type) {
	case CommitsParam:
		commits, _, err := client.Commits.Get(ctx, r)
		if err != nil {
			fmt.Println("err", err)
		}
		return commits, nil
	default:
		return nil, nil
	}
}

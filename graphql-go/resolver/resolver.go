package resolver

import (
	"context"
	"fmt"

	"github.com/gopher-dev/experiment/graphql-go/model"

	"github.com/graph-gophers/graphql-go"
)

type Resolver struct{}

func (r *Resolver) FindStoryByID(ctx context.Context, args struct{ ID graphql.ID }) (*model.Story, error) {
	fmt.Println(args.ID)
	return &model.Story{
		Data: struct {
			ID    string
			Title string
		}{
			ID:    "blah",
			Title: "blah",
		},
	}, nil
}

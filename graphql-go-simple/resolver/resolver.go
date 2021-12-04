package resolver

import (
	"context"
	"fmt"
	"time"

	"github.com/gopher-dev/experiment/graphql-go-simple/model"
	"github.com/graph-gophers/graphql-go"
)

type storyInput struct {
	Title       string
	Description string
}

// Resolver :nodoc:
type Resolver struct{}

func (r *Resolver) FindStoryByID(ctx context.Context, args struct{ ID graphql.ID }) (model.Story, error) {
	fmt.Println(args.ID)

	return model.Story{
		Data: model.StoryData{
			ID:     "blahx",
			Title:  "blah",
			Status: "DRAFT",
			Comments: []model.Comment{
				{
					model.CommentData{
						ID:      "comment-1",
						Comment: "Lorem ipsum",
					},
				},
			},
			CreatedAt: time.Now(),
		},
	}, nil
}

func (r *Resolver) CreateStory(ctx context.Context, args struct{ Story storyInput }) (model.Story, error) {
	fmt.Println("title", args.Story.Title)
	fmt.Println("description", args.Story.Description)

	return model.Story{
		Data: model.StoryData{
			ID:    "blah",
			Title: "blah",
		},
	}, nil
}

func (r *Resolver) FindAllStories(args struct{ Size int32 }) ([]model.Story, error) {
	fmt.Println(args.Size)

	return []model.Story{
		{
			Data: model.StoryData{
				ID:    "blah",
				Title: "blah",
			},
		},
		{
			Data: model.StoryData{
				ID:    "blah",
				Title: "blah",
			},
		},
	}, nil
}

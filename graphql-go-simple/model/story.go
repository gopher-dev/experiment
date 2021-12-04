package model

import (
	"time"

	"github.com/graph-gophers/graphql-go"
)

// Story :nodoc:
type Story struct {
	Data StoryData
}

// StoryData :nodoc:
type StoryData struct {
	ID        string
	Title     string
	Status    string
	Comments  []Comment
	CreatedAt time.Time
}

func (s Story) ID() graphql.ID {
	return graphql.ID(s.Data.ID)
}

func (s Story) Title() string {
	return s.Data.Title
}

func (s Story) Status() string {
	return s.Data.Status
}

func (s Story) CreatedAt() string {
	return s.Data.CreatedAt.Format(time.RFC3339)
}

func (s Story) Comments() []Comment {
	return s.Data.Comments
}

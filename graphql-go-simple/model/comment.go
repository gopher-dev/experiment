package model

import (
	"github.com/graph-gophers/graphql-go"
)

// Comment :nodoc:
type Comment struct {
	Data CommentData
}

type CommentData struct {
	ID      string
	Comment string
}

func (c Comment) ID() graphql.ID {
	return graphql.ID(c.Data.ID)
}

func (c Comment) Comment() string {
	return c.Data.Comment
}

package model

import "github.com/graph-gophers/graphql-go"

type Story struct {
	Data struct {
		ID    string
		Title string
	}
}

func (s Story) Id() graphql.ID {
	return graphql.ID(s.Data.ID)
}

func (s Story) Title() string {
	return s.Data.Title
}

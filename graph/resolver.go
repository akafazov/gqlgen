package graph

//go:generate go run github.com/99designs/gqlgen@v0.17.49 generate

import "github.com/akafazov/gqlgen/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

var meetups = []*model.Meetup{
	{
		ID:          "1",
		Name:        "meetup1",
		Description: "description1",
		User:        users[0],
	},
	{
		ID:          "2",
		Name:        "meetup2",
		Description: "description2",
		User:        users[1],
	},
}

var users = []*model.User{
	{
		ID:       "1",
		Username: "user1",
		Email:    "user1@gmail.com",
	},
	{
		ID:       "2",
		Username: "user2",
		Email:    "user2@gmail.com",
	},
}

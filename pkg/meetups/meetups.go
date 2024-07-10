package meetups

import (
	"context"
	"fmt"

	"github.com/akafazov/gqlgen/graph/model"
)

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

func GetMeetups(ctx context.Context) ([]*model.Meetup, error) {
	return meetups, nil
}

func GetUsers(ctx context.Context) ([]*model.User, error) {
	return users, nil
}

func CreateMeetup(input model.NewMeetup) (*model.Meetup, error) {
	m := &model.Meetup{
		ID:          fmt.Sprintf("T%d", len(meetups)+1),
		Name:        input.Name,
		Description: input.Description,
		User:        users[0],
	}
	meetups = append(meetups, m)
	return m, nil
}

func CreateMeetupWithUser(ctx context.Context, obj *model.User) ([]*model.Meetup, error) {
	m := make([]*model.Meetup, 0)
	for _, meetup := range meetups {
		if meetup.User.ID == obj.ID {
			m = append(m, meetup)
		}
	}
	return m, nil
}

func GetUser(ctx context.Context, obj *model.Meetup) (*model.User, error) {
	user := new(model.User)

	for _, u := range users {
		if u.ID == obj.User.ID {
			user = u
			break
		}
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	u := &model.User{
		ID:       fmt.Sprintf("T%d", len(users)+1),
		Username: input.Username,
		Email:    input.Email,
	}
	return u, nil
}

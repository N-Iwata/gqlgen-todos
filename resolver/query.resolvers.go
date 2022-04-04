package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/N-Iwata/gqlgen-todos/graph/generated"
	"github.com/N-Iwata/gqlgen-todos/graph/model"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	for _, v := range r.users {
		todos := []*model.Todo{}
		for _, t := range r.todos {
			if v.ID == t.UserID {
				todos = append(todos, t)
			}
		}
		v.Todos = todos
	}

	return r.users, nil
}

func (r *queryResolver) User(ctx context.Context, userID string) (*model.User, error) {
	user := &model.User{}
	todos := []*model.Todo{}
	for _, v := range r.todos {
		if v.UserID == userID {
			todos = append(todos, v)
		}
	}

	for _, v := range r.users {
		if v.ID == userID {
			user.ID = v.ID
			user.Name = v.Name
			user.Todos = todos
		}
	}

	return user, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }

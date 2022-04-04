package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/N-Iwata/gqlgen-todos/graph/generated"
	"github.com/N-Iwata/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) FinishTodo(ctx context.Context, todoID string) (*model.Todo, error) {
	todo := &model.Todo{}
	for _, v := range r.todos {
		if v.ID == todoID {
			v.Done = true
			todo.ID = v.ID
			todo.Text = v.Text
			todo.UserID = v.UserID
			todo.Done = v.Done
		}
	}
	return todo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, todoID string) (*model.Todo, error) {
	todo := &model.Todo{}
	todos := []*model.Todo{}
	for _, v := range r.todos {
		if v.ID != todoID {
			todos = append(todos, v)
			todo.ID = v.ID
			todo.Text = v.Text
			todo.UserID = v.UserID
			todo.Done = v.Done
		}
	}
	r.todos = todos
	return todo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Name: input.Name,
	}
	r.users = append(r.users, user)
	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

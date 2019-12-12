package go_graphql_poc
import (
	"context"
	"errors"
	"fmt"
	"math/rand"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
type Resolver struct {
	todos []*Todo
	users []*User
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) DeleteTodo(ctx context.Context, input DeleteTodo) (bool, error) {
	var result bool = false

	for index, val := range r.todos {
		if val.ID == input.ID {
			r.todos = append(r.todos[:index], r.todos[index+1:]...)
			result = true
		}
	}

	return result, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*Todo, error) {
	userID := input.UserID

	var foundUser *User

	for _, user := range r.users {
		if user.ID == userID {
			foundUser = user
		}
	}

	if foundUser == nil {
		return nil, errors.New("User not found")
	}

	todo := &Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: foundUser,
	}

	r.todos = append(r.todos, todo)
	return todo, nil
}
func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*User, error) {
	user := &User{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Name: input.Name,
	}

	r.users = append(r.users, user)
	return user, nil
}

type queryResolver struct{ *Resolver }
func (r *queryResolver) Todos(ctx context.Context) ([]*Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	return r.users, nil
}

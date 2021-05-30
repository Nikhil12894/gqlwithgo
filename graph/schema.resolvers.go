package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/Nikhil12894/gqlwithgo/graph/generated"
	"github.com/Nikhil12894/gqlwithgo/graph/model"
)

func (r *mutationResolver) CreateVachil(ctx context.Context, input model.NewVachil) (*model.Vachil, error) {
	video := &model.Vachil{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Title:  input.Title,
		URL:    input.URL,
		Author: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.Vachils = append(r.Vachils, video)
	return video, nil
}

func (r *queryResolver) Vachil(ctx context.Context) ([]*model.Vachil, error) {
	return r.Vachils, nil
}

func (r *subscriptionResolver) VideoAdded(ctx context.Context, repoFullName string) (<-chan *model.Vachil, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

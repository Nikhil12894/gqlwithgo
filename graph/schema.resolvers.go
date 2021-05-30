package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	db "github.com/Nikhil12894/gqlwithgo/dbHandler"
	"github.com/Nikhil12894/gqlwithgo/graph/generated"
	"github.com/Nikhil12894/gqlwithgo/graph/model"
)

func (r *mutationResolver) CreateVachil(ctx context.Context, input model.NewVachil) (*model.Vachil, error) {
	vachil := db.CreateVachil(input)
	return vachil, nil
}

func (r *mutationResolver) CreateBooking(ctx context.Context, input model.NewBooking) (*model.Booking, error) {
	booking := &model.Booking{
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
		User:      &model.User{ID: fmt.Sprintf("%d", input.UserID)},
		Vachil:    &model.Vachil{ID: fmt.Sprintf("%d", input.VachilID)},
		ID:        fmt.Sprintf("%d", rand.Int()),
	}
	r.Bookings = append(r.Bookings, booking)
	return booking, nil
}

func (r *queryResolver) Vachil(ctx context.Context) ([]*model.Vachil, error) {
	return db.FindVachil(), nil
}

func (r *queryResolver) VachilWithType(ctx context.Context, typeArg string) ([]*model.Vachil, error) {
	return db.FindVachilByType(typeArg)
}

func (r *queryResolver) VachilWithTypeCapacity(ctx context.Context, typeArg string, capacity int) ([]*model.Vachil, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AvelebleVachilWithType(ctx context.Context, typeArg string, startDate time.Time, endDate time.Time) ([]*model.Vachil, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AvelableVachilWithTypeCapacity(ctx context.Context, typeArg string, capacity int, startDate time.Time, endDate time.Time) ([]*model.Vachil, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllBooking(ctx context.Context, startDate time.Time, endDate time.Time) ([]*model.Booking, error) {
	return r.Bookings, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

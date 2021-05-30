package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/Nikhil12894/gqlwithgo/graph/generated"
	"github.com/Nikhil12894/gqlwithgo/graph/model"
)

func (r *mutationResolver) CreateVachil(ctx context.Context, input model.NewVachil) (*model.Vachil, error) {
	vachil := &model.Vachil{
		Type:      input.Type,
		Brand:     input.Brand,
		RegNo:     input.RegNo,
		Capacity:  input.Capacity,
		UnitPrice: input.UnitPrice,
	}
	// Create vachil/
	err := r.DB.Create(&vachil).Error
	if err != nil {
		return nil, err
	}
	return vachil, nil
}

func (r *mutationResolver) CreateBooking(ctx context.Context, input model.NewBooking) (*model.Booking, error) {
	booking := &model.Booking{
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
		User:      &model.User{ID: input.UserID},
		Vachil:    &model.Vachil{ID: input.VachilID},
	}
	r.DB.Create(&booking)
	return booking, nil
}

func (r *queryResolver) Vachil(ctx context.Context) ([]*model.Vachil, error) {
	var vachils []*model.Vachil
	r.DB.Find(&vachils)
	return vachils, nil
}

func (r *queryResolver) VachilWithType(ctx context.Context, typeArg string) ([]*model.Vachil, error) {
	var vachils []*model.Vachil
	if err := r.DB.Where("type = ?", typeArg).Find(&vachils).Error; err != nil {
		return nil, err
	}
	return vachils, nil
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
	var bookings []*model.Booking
	r.DB.Find(&bookings)
	return bookings, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

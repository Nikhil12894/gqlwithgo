package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	db "github.com/Nikhil12894/gqlwithgo/dbHandler"
	"github.com/Nikhil12894/gqlwithgo/graph/generated"
	"github.com/Nikhil12894/gqlwithgo/graph/model"
)

func (r *mutationResolver) CreateVachil(ctx context.Context, input model.NewVachil) (*model.Vachil, error) {
	vachil := &model.Vachil{
		Type:      input.Type,
		Brand:     input.Brand,
		Name:      input.Name,
		RegNo:     input.RegNo,
		Capacity:  input.Capacity,
		UnitPrice: input.UnitPrice,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Images:    *input.Images,
		// create method to add images
	}
	// Create vachil/
	err := r.DB.Create(&vachil).Error
	if err != nil {
		return nil, err
	}
	return vachil, nil
}

func (r *mutationResolver) UpdateVachil(ctx context.Context, vachilID int, input model.NewVachil) (*model.Vachil, error) {
	var vachil1 *model.Vachil
	err := r.DB.Where("id = ?", vachilID).First(&vachil1).Error
	if err != nil {
		return nil, err
	}
	vachil := &model.Vachil{
		Type:      input.Type,
		Brand:     input.Brand,
		Name:      input.Name,
		RegNo:     input.RegNo,
		Capacity:  input.Capacity,
		UnitPrice: input.UnitPrice,
		ID:        vachilID,
		UpdatedAt: time.Now(),
		CreatedAt: vachil1.CreatedAt,
		Images:    db.GetImage(*input.Images, vachil1),
	}
	// update vachil/
	err = r.DB.Save(&vachil).Error
	if err != nil {
		return nil, err
	}
	return vachil, nil
}

func (r *mutationResolver) DeleteVachil(ctx context.Context, vachilID int) (bool, error) {
	err := r.DB.Where("id = ?", vachilID).Delete(&model.Vachil{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) CreateBooking(ctx context.Context, input model.NewBooking) (*model.Booking, error) {
	currentTime := time.Now()
	booking := &model.Booking{
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
		UserID:    input.UserID,
		VachilID:  input.VachilID,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
	TotalPriceID := db.CreateTotalPrice(input.StartDate, input.EndDate, input.VachilID)
	if TotalPriceID != 0 {
		booking.TotalPriceID = TotalPriceID
	}
	r.DB.Create(&booking)
	return booking, nil
}

func (r *mutationResolver) UpdateBooking(ctx context.Context, bookingID int, input model.NewBooking) (*model.Booking, error) {
	vachil := &model.Booking{
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
		UserID:    input.UserID,
		VachilID:  input.VachilID,
		UpdatedAt: time.Now(),
		ID:        bookingID,
	}
	TotalPriceID := db.CreateTotalPrice(input.StartDate, input.EndDate, input.VachilID)
	if TotalPriceID != 0 {
		vachil.TotalPriceID = TotalPriceID
	}
	// update vachil/
	err := r.DB.Save(&vachil).Error
	if err != nil {
		return nil, err
	}
	return vachil, nil
}

func (r *mutationResolver) DeleteBooking(ctx context.Context, bookingID int) (bool, error) {
	err := r.DB.Where("id = ?", bookingID).Delete(&model.Booking{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) UpdateUserProfileImage(ctx context.Context, id int, image string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUserPassword(ctx context.Context, id int, password string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input model.NewUserWithType) (*model.User, error) {
	var vachil1 *model.User
	err := r.DB.Where("id = ?", id).First(&vachil1).Error
	if err != nil {
		return nil, err
	}
	vachil := &model.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		UserType:  input.UserType,
		ID:        id,
		UpdatedAt: time.Now(),
		CreatedAt: vachil1.CreatedAt,
		Mobile:    input.Mobile,
		Image:     db.GetUserImage(*input.Image, vachil1),
		Password:  db.GetUserPassword(input.Password, vachil1),
	}
	// update vachil/
	err = r.DB.Save(&vachil).Error
	if err != nil {
		return nil, err
	}
	return vachil, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID int) (bool, error) {
	err := r.DB.Where("id = ?", userID).Delete(&model.User{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *queryResolver) Vachil(ctx context.Context) ([]*model.Vachil, error) {
	var vachils []*model.Vachil
	err := r.DB.Find(&vachils).Error
	if err != nil {
		return nil, err
	}
	return vachils, nil
}

func (r *queryResolver) VachilWithID(ctx context.Context, typeArg int) (*model.Vachil, error) {
	var vachil *model.Vachil
	err := r.DB.Where("id = ?", typeArg).First(&vachil).Error
	if err != nil {
		return nil, err
	}
	return vachil, nil
}

func (r *queryResolver) VachilWithType(ctx context.Context, typeArg string) ([]*model.Vachil, error) {
	var vachils []*model.Vachil
	if err := r.DB.Where("type = ?", typeArg).Find(&vachils).Error; err != nil {
		return nil, err
	}
	return vachils, nil
}

func (r *queryResolver) VachilWithBrand(ctx context.Context, brand string) ([]*model.Vachil, error) {
	var vachils []*model.Vachil
	if err := r.DB.Where("Brand = ?", brand).Find(&vachils).Error; err != nil {
		return nil, err
	}
	return vachils, nil
}

func (r *queryResolver) VachilWithRegNo(ctx context.Context, regNo string) (*model.Vachil, error) {
	var vachils *model.Vachil
	if err := r.DB.Where("RegNo = ?", regNo).First(&vachils).Error; err != nil {
		return nil, err
	}
	return vachils, nil
}

func (r *queryResolver) VachilWithName(ctx context.Context, name string) ([]*model.Vachil, error) {
	var vachils []*model.Vachil
	if err := r.DB.Where("Name = ?", name).First(&vachils).Error; err != nil {
		return nil, err
	}
	return vachils, nil
}

func (r *queryResolver) VachilWithTypeCapacity(ctx context.Context, typeArg string, capacity int) ([]*model.Vachil, error) {
	var vachils []*model.Vachil
	if err := r.DB.Where("type = ? AND capacity =? ", typeArg, capacity).First(&vachils).Error; err != nil {
		return nil, err
	}
	return vachils, nil
}

func (r *queryResolver) AvelebleVachilWithType(ctx context.Context, typeArg string, startDate time.Time, endDate time.Time) ([]*model.Vachil, error) {
	var vachils []*model.Vachil
	if err := r.DB.Where("type = ?", typeArg).First(&vachils).Error; err != nil {
		return nil, err
	}
	return vachils, nil
}

func (r *queryResolver) AvelableVachilWithTypeCapacity(ctx context.Context, typeArg string, capacity int, startDate time.Time, endDate time.Time) ([]*model.Vachil, error) {
	var vachils []*model.Vachil
	if err := r.DB.Where("type = ? AND capacity =? ", typeArg, capacity).First(&vachils).Error; err != nil {
		return nil, err
	}
	return vachils, nil
}

func (r *queryResolver) AllBooking(ctx context.Context, startDate time.Time, endDate time.Time) ([]*model.Booking, error) {
	var bookings []*model.Booking
	err := r.DB.Model(&model.Booking{}).Where("start_date >=? and end_date <= ?", startDate, endDate).Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *queryResolver) AllBookingWithID(ctx context.Context, userID int) ([]*model.Booking, error) {
	var bookings []*model.Booking
	err := r.DB.Find(&bookings).Where("user_id=?", userID).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *queryResolver) AllActiveBookingWithID(ctx context.Context, userID int) ([]*model.Booking, error) {
	var bookings []*model.Booking
	err := r.DB.Find(&bookings).Where("user_id=?", userID).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *queryResolver) UserWithID(ctx context.Context, userID int) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserWithUserName(ctx context.Context, userName string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserPasswordByName(ctx context.Context, email string) (string, error) {
	return "", nil
}

func (r *queryResolver) DistinctVachilImages(ctx context.Context) ([]map[string]interface{}, error) {
	result1 := []map[string]interface{}{}
	err := r.DB.Raw("select distinct images as PreviewImageSrc, images as thumbnailImageSrc,name as Alt, name as Title from vachils", 3).Scan(&result1).Error
	if err != nil {
		return nil, err
	}
	return result1, nil
}

func (r *queryResolver) AllUsers(ctx context.Context, userID int) ([]*model.User, error) {
	var users []*model.User
	err := r.DB.Where("id !=?", userID).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

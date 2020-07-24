package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"BackEnd/graph/model"
	"context"
	"errors"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	user := model.User{
		Email:        input.Email,
		Pass:         input.Pass,
		ProfilePic:   input.ProfilePic,
		Username:     input.Username,
		MembershipID: "",
	}
	_, err := r.DB.Model(&user).Insert()
	if err != nil {
		return nil, errors.New("insert user failed")
	}
	return &user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *model.NewUser) (*model.User, error) {
	var user model.User

	err := r.DB.Model(&user).Where("id=?", id).First()
	if err != nil {
		return nil, errors.New("user not found")
	}
	user.Email = input.Email
	user.Pass = input.Pass
	user.ProfilePic = input.ProfilePic
	user.Username = input.Username
	_, updateErr := r.DB.Model(&user).Where("id=?", id).Update()
	if updateErr != nil {
		return nil, errors.New("update user fialed")
	}
	return &user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	var user model.User

	err := r.DB.Model(&user).Where("id=?", id).First()
	if err != nil {
		return false, errors.New("user not found")
	}
	_, deleteErr := r.DB.Model(&user).Where("id=?", id).Delete()
	if deleteErr != nil {
		return false, errors.New("delete user failed")
	}
	return true, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	err := r.DB.Model(&users).Order("id").Select()

	if err != nil {
		return nil, errors.New("query user failed")
	}
	return users, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.DB.Model(&user).Where("id=?", id).Select()
	if err != nil {
		return nil, errors.New("query user failed")
	}
	return &user, nil
}

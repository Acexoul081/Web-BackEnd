package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"BackEnd/graph/generated"
	"BackEnd/models"
	"context"
	"errors"
	"log"
	"strings"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *models.NewUser) (*models.AuthResponse, error) {
	var user models.User
	err := r.DB.Model(&user).Where("email = ?", input.Email).First()
	if err == nil {
		token, err := user.GenToken()
		if err != nil {
			log.Printf("error while generating token: %v", err)
			return nil, errors.New("something went wrong")
		}
		return &models.AuthResponse{
			AuthToken: token,
			User:      &user,
		}, nil
	}

	err = r.DB.Model(&user).Where("username = ?", input.Username).First()
	if err == nil {
		return nil, errors.New("username already in use")
	}

	newUser := models.User{
		Email:        input.Email,
		ProfilePic:   input.ProfilePic,
		Username:     input.Username,
		MembershipID: "",
	}

	_, err = r.DB.Model(&newUser).Insert()
	if err != nil {
		return nil, errors.New("insert user failed")
	}

	token, err := newUser.GenToken()
	if err != nil {
		log.Printf("error while generating token: %v", err)
		return nil, errors.New("something went wrong")
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      &newUser,
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	var user models.User
	err := r.DB.Model(&user).Where("email=?", input.Email).First()
	if err != nil {
		return nil, errors.New("user not found")
	}

	token, err := user.GenToken()
	if err != nil {
		return nil, errors.New("token error")
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      &user,
	}, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *models.NewUser) (*models.User, error) {
	var user models.User

	err := r.DB.Model(&user).Where("id=?", id).First()
	if err != nil {
		return nil, errors.New("user not found")
	}
	user.Email = input.Email
	user.ProfilePic = input.ProfilePic
	user.Username = input.Username
	_, updateErr := r.DB.Model(&user).Where("id=?", id).Update()
	if updateErr != nil {
		return nil, errors.New("update user fialed")
	}
	return &user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	var user models.User

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

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	var users []*models.User

	err := r.DB.Model(&users).Order("id").Select()

	if err != nil {
		return nil, errors.New("query user failed")
	}
	return users, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := r.DB.Model(&user).Where("id=?", id).Select()
	if err != nil {
		return nil, errors.New("query user failed")
	}
	return &user, nil
}

func (r *queryResolver) GetUserSearch(ctx context.Context, keyword string) ([]*models.User, error) {
	var users []*models.User
	keyword = strings.ToLower(keyword)
	keyword = "%" + keyword + "%"
	err := r.DB.Model(&users).Where("lower(username) LIKE ?", keyword).Order("id").Select()
	if err != nil {
		return nil, errors.New("query user search failed")
	}
	return users, nil
}

func (r *userResolver) Videos(ctx context.Context, obj *models.User) ([]*models.Video, error) {
	var videos []*models.Video
	err := r.DB.Model(&videos).Where("user_id=?", obj.ID).Select()
	if err != nil {
		return nil, errors.New("errors query video from user")
	}
	return videos, nil
}

func (r *userResolver) Subscribers(ctx context.Context, obj *models.User) ([]*models.Abonemen, error) {
	var subscriber []*models.Abonemen
	err := r.DB.Model(&subscriber).Where("user_id=?", obj.ID).Select()
	if err != nil {
		return nil, errors.New("failed query subscriber from user")
	}
	return subscriber, nil
}

func (r *userResolver) Playlists(ctx context.Context, obj *models.User) ([]*models.Playlist, error) {
	var playlist []*models.Playlist
	err := r.DB.Model(&playlist).Where("owner_id=?", obj.ID).Order("id").Select()
	if err != nil {
		return nil, errors.New("failed query playlist from user")
	}
	return playlist, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }

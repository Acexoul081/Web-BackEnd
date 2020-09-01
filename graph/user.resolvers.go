package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"BackEnd/graph/generated"
	"BackEnd/middleware"
	"BackEnd/models"
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	pg "github.com/go-pg/pg/v9"
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
		JoinDate:     time.Now().Format("2006-01-02 15:04:05"),
	}
	if input.Password != nil {
		err = newUser.HashPassword(*input.Password)
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

func (r *mutationResolver) UpdateUser(ctx context.Context, input models.UpdatedUser) (*models.User, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("unauthenticated")
	}

	var user models.User

	err = r.DB.Model(&user).Where("id=?", currentUser.ID).First()
	if err != nil {
		return nil, errors.New("user not found")
	}
	if input.Description != nil {
		user.Description = *input.Description
	}

	if input.Link != nil {
		user.Link = *input.Link
	}

	if input.Banner != nil {
		user.Banner = *input.Banner
	}
	if input.Icon != nil {
		user.ProfilePic = *input.Icon
	}
	_, updateErr := r.DB.Model(&user).Where("id=?", currentUser.ID).Update()
	if updateErr != nil {
		return nil, errors.New("update user failed")
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

func (r *mutationResolver) ChangeUserLocation(ctx context.Context, location string) (*models.User, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("unauthenticated")
	}
	currentUser.Location = location
	_, updateErr := r.DB.Model(currentUser).Where("id=?", currentUser.ID).Update()
	if updateErr != nil {
		return nil, errors.New("update location failed")
	}
	return currentUser, nil
}

func (r *mutationResolver) InsertAbout(ctx context.Context, description string) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAbout(ctx context.Context, description string) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
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

func (r *queryResolver) GetNotification(ctx context.Context, ids []string) (*models.Notification, error) {
	var notification models.Notification

	var videos []*models.Video
	err := r.DB.Model(&videos).Where("user_id in (?)", pg.In(ids)).Order("date_upload DESC").Limit(5).Select()
	if err != nil {
		return nil, errors.New("failed get video for notification")
	}
	notification.Videos = videos

	var posts []*models.Post
	err = r.DB.Model(&posts).Where("channel_id in (?)", pg.In(ids)).Order("post_date DESC").Limit(5).Select()
	if err != nil {
		return nil, errors.New("failed get post for notification")
	}
	notification.Posts = posts

	return &notification, nil
}

func (r *userResolver) Membership(ctx context.Context, obj *models.User) (*models.MembershipDetail, error) {
	var membershipDet models.MembershipDetail
	err := r.DB.Model(&membershipDet).Where("user_id=?", obj.ID).Select()
	if err != nil {
		return nil, nil
	}
	return &membershipDet, nil
}

func (r *userResolver) Videos(ctx context.Context, obj *models.User) ([]*models.Video, error) {
	var videos []*models.Video
	err := r.DB.Model(&videos).Where("user_id=?", obj.ID).Select()
	if err != nil {
		return nil, errors.New("errors query video from user")
	}
	return videos, nil
}

func (r *userResolver) Subscription(ctx context.Context, obj *models.User) ([]*models.Abonemen, error) {
	var subs []*models.Abonemen
	err := r.DB.Model(&subs).Where("subscriber_id=?", obj.ID).Select()
	if err != nil {
		return nil, errors.New("no subscription")
	}
	return subs, nil
}

func (r *userResolver) Subscriber(ctx context.Context, obj *models.User) ([]*models.Abonemen, error) {
	var subs []*models.Abonemen
	err := r.DB.Model(&subs).Where("user_id=?", obj.ID).Select()
	if err != nil {
		return nil, errors.New("no subscriber")
	}
	return subs, nil
}

func (r *userResolver) Playlists(ctx context.Context, obj *models.User) ([]*models.Playlist, error) {
	var playlist []*models.Playlist
	err := r.DB.Model(&playlist).Where("owner_id=?", obj.ID).Order("id").Select()
	if err != nil {
		return nil, errors.New("failed query playlist from user")
	}
	var subs []*models.PlaylistSub
	err = r.DB.Model(&subs).Where("user_id = ?", obj.ID).Select()
	if err != nil {
		return nil, err
	}

	return playlist, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }

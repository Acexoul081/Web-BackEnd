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
	"time"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input models.NewPost) (*models.Post, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("unauthenticated")
	}
	post := models.Post{
		UserID:    currentUser.ID,
		Post:      input.Post,
		PostDate:  time.Now().Format("2006-01-02 15:04:05"),
		ChannelID: input.ChannelID,
		Thumbnail: input.Thumbnail,
	}
	_, err = r.DB.Model(&post).Insert()
	if err != nil {
		return nil, errors.New("failed insert post")
	}
	return &post, nil
}

func (r *mutationResolver) UpdatePostLike(ctx context.Context, id string) (*models.PostLikeDetail, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("unauthenticated")
	}
	var likeDetail models.PostLikeDetail
	err = r.DB.Model(&likeDetail).Where("user_id=? and post_id=?", currentUser.ID, id).Select()
	if err == nil {
		fmt.Println(err)
		likeDetail.Like = true
		likeDetail.UserID = currentUser.ID
		likeDetail.PostID = id
		_, updateErr := r.DB.Model(&likeDetail).Where("user_id=? and post_id =?", currentUser.ID, id).Update()
		if updateErr != nil {
			return nil, errors.New("update like error")
		}
		return &likeDetail, nil
	}
	like := models.PostLikeDetail{
		PostID: id,
		UserID: currentUser.ID,
		Like:   true,
	}

	_, err = r.DB.Model(&like).Insert()
	if err != nil {
		return nil, errors.New("like post failed")
	}
	return &like, nil
}

func (r *mutationResolver) UpdatePostDislike(ctx context.Context, id string) (*models.PostLikeDetail, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("unauthenticated")
	}
	var likeDetail models.PostLikeDetail
	err = r.DB.Model(&likeDetail).Where("user_id=? and post_id=?", currentUser.ID, id).Select()
	if err == nil {
		likeDetail.Like = false
		likeDetail.UserID = currentUser.ID
		likeDetail.PostID = id
		_, updateErr := r.DB.Model(&likeDetail).Where("user_id=? and post_id =?", currentUser.ID, id).Update()
		if updateErr != nil {
			return nil, errors.New("update dislike error")
		}
		return &likeDetail, nil
	}
	dislike := models.PostLikeDetail{
		PostID: id,
		UserID: currentUser.ID,
		Like:   false,
	}

	_, err = r.DB.Model(&dislike).Insert()
	if err != nil {
		return nil, errors.New("dislike post failed")
	}
	return &dislike, nil
}

func (r *postResolver) User(ctx context.Context, obj *models.Post) (*models.User, error) {
	var user models.User
	err := r.DB.Model(&user).Where("id=?", obj.UserID).Select()
	if err != nil {
		return nil, errors.New("failed get user from post")
	}
	return &user, nil
}

func (r *postResolver) Channel(ctx context.Context, obj *models.Post) (*models.User, error) {
	var channel models.User
	err := r.DB.Model(&channel).Where("id=?", obj.ChannelID).Select()
	if err != nil {
		return nil, errors.New("failed get channel from post")
	}
	return &channel, nil
}

func (r *postResolver) Like(ctx context.Context, obj *models.Post) ([]*models.PostLikeDetail, error) {
	var likeDetails []*models.PostLikeDetail
	err := r.DB.Model(&likeDetails).Where("post_id=?", obj.ID).Select()
	if err != nil {
		return nil, errors.New("failed to get like from post")
	}
	return likeDetails, nil
}

func (r *postLikeDetailResolver) Post(ctx context.Context, obj *models.PostLikeDetail) (*models.Post, error) {
	var post models.Post
	err := r.DB.Model(&post).Where("id=?", obj.PostID).Select()
	if err != nil {
		return nil, errors.New("failed get post from detail")
	}
	return &post, nil
}

func (r *postLikeDetailResolver) User(ctx context.Context, obj *models.PostLikeDetail) (*models.User, error) {
	var user models.User
	err := r.DB.Model(&user).Where("id=?", obj.UserID).Select()
	if err != nil {
		return nil, errors.New("failed get user from detail")
	}
	return &user, nil
}

func (r *queryResolver) GetPostsByID(ctx context.Context, id string) ([]*models.Post, error) {
	var posts []*models.Post
	err := r.DB.Model(&posts).Where("user_id=?", id).Select()
	if err != nil {
		return nil, errors.New("failed get post by user id")
	}
	return posts, nil
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// PostLikeDetail returns generated.PostLikeDetailResolver implementation.
func (r *Resolver) PostLikeDetail() generated.PostLikeDetailResolver {
	return &postLikeDetailResolver{r}
}

type postResolver struct{ *Resolver }
type postLikeDetailResolver struct{ *Resolver }

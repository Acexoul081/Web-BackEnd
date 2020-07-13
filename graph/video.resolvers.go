package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"BackEnd/graph/model"
	"context"
	"errors"
	"time"
)

func (r *mutationResolver) CreateVideo(ctx context.Context, input *model.NewVideo) (*model.Video, error) {
	video := model.Video{
		UserID:      input.UserID,
		Title:       input.Title,
		Description: input.Description,
		Label:       input.Label,
		Thumbnail:   input.Thumbnail,
		Category:    input.Category,
		Privacy:     input.Privacy,
		Like: 0,
		Dislike: 0,
		View: 0,
		Link: input.Link,
		DateUpload: time.Now().Format("2006-01-02 15:04:05"),
		DatePublish: "",
	}

	_, err := r.DB.Model(&video).Insert()
	if err != nil {
		return nil, errors.New("insert video failed")
	}
	return &video, nil
}

func (r *mutationResolver) UpdateVideo(ctx context.Context, id string, input *model.NewVideo) (*model.Video, error) {
	var video model.Video

	err := r.DB.Model(&video).Where("id=?", id).First()
	if err != nil {
		return nil, errors.New("video not found")
	}
	video.Title = input.Title
	video.Description = input.Description
	video.Label = input.Label
	video.Thumbnail = input.Thumbnail
	video.Category = input.Category
	video.Privacy = input.Privacy
	_, updateErr := r.DB.Model(&video).Where("id=?", id).Update()
	if updateErr != nil {
		return nil, errors.New("update video failed")
	}
	return &video, nil
}

func (r *mutationResolver) DeleteVideo(ctx context.Context, id string) (bool, error) {
	var video model.Video

	err := r.DB.Model(&video).Where("id=?", id).First()
	if err != nil {
		return false, errors.New("video not found")
	}
	_, deleteErr := r.DB.Model(&video).Where("id=?", id).Delete()

	if deleteErr != nil {
		return false, errors.New("delete video failed")
	}
	return true, nil
}

func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	var videos []*model.Video

	err := r.DB.Model(&videos).Order("id").Select()

	if err != nil {
		return nil, errors.New("query video failed")
	}
	return videos, nil
}

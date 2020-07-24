package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"BackEnd/graph/model"
	"context"
	"errors"
	"math/rand"
	"strings"
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
		Like:        0,
		Dislike:     0,
		View:        0,
		Link:        input.Link,
		DateUpload:  time.Now().Format("2006-01-02 15:04:05"),
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

func (r *mutationResolver) UpdateVideoView(ctx context.Context, id string) (*model.Video, error) {
	var video model.Video
	err := r.DB.Model(&video).Where("id=?", id).First()
	if err != nil {
		return nil, errors.New("video not found")
	}
	video.View += 1
	_, updateErr := r.DB.Model(&video).Where("id=?", id).Update()
	if updateErr != nil {
		return nil, errors.New("update view failed")
	}
	return &video, nil
}

func (r *mutationResolver) UpdateVideoLike(ctx context.Context, id string) (*model.Video, error) {
	var video model.Video
	err := r.DB.Model(&video).Where("id=?", id).First()
	if err != nil {
		return nil, errors.New("video not found")
	}
	video.Like += 1
	_, updateErr := r.DB.Model(&video).Where("id=?", id).Update()
	if updateErr != nil {
		return nil, errors.New("update like failed")
	}
	return &video, nil
}

func (r *mutationResolver) UpdateVideoDislike(ctx context.Context, id string) (*model.Video, error) {
	var video model.Video
	err := r.DB.Model(&video).Where("id=?", id).First()
	if err != nil {
		return nil, errors.New("video not found")
	}
	video.Dislike += 1
	_, updateErr := r.DB.Model(&video).Where("id=?", id).Update()
	if updateErr != nil {
		return nil, errors.New("update dislike failed")
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
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(videos), func(i, j int) {
		videos[i], videos[j] = videos[j], videos[i]
	})
	return videos, nil
}

func (r *queryResolver) GetVideo(ctx context.Context, id string) (*model.Video, error) {
	var video model.Video
	err := r.DB.Model(&video).Where("id=?", id).Select()
	if err != nil {
		return nil, errors.New("query video failed")
	}

	return &video, nil
}

func (r *queryResolver) GetVideosByCategory(ctx context.Context, category int) ([]*model.Video, error) {
	var videos []*model.Video
	err := r.DB.Model(&videos).Where("category=?", category).Order("id").Select()
	if err != nil {
		return nil, errors.New("query video failed")
	}
	return videos, nil
}

func (r *queryResolver) GetVideosByUser(ctx context.Context, userID string) ([]*model.Video, error) {
	var videos []*model.Video
	err := r.DB.Model(&videos).Where("user_id=?", userID).Order("id").Select()
	if err != nil {
		return nil, errors.New("query video failed")
	}
	return videos, nil
}

func (r *queryResolver) GetVideosByUserLimit(ctx context.Context, userID string, limit int) ([]*model.Video, error) {
	var videos []*model.Video
	err := r.DB.Model(&videos).Where("user_id=?", userID).Order("id").Select()
	if err != nil {
		return nil, errors.New("query video with limit failed")
	}
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(videos), func(i, j int) {
		videos[i], videos[j] = videos[j], videos[i]
	})

	return videos[0:limit], nil
}

func (r *queryResolver) GetLastUploadedVideo(ctx context.Context, userID string) (*model.Video, error) {
	var video model.Video

	err := r.DB.Model(&video).Where("user_id=?", userID).Last()
	if err != nil {
		return nil, errors.New("query last video failed")
	}

	return &video, nil
}

func (r *queryResolver) GetVideoSearch(ctx context.Context, keyword string) ([]*model.Video, error) {
	var videos []*model.Video
	keyword = strings.ToLower(keyword)
	keyword = "%" + keyword + "%"
	err := r.DB.Model(&videos).Where("lower(title) LIKE ? or lower(description) LIKE ?", keyword, keyword).Order("title").Select()
	if err != nil {
		return nil, errors.New("query video search failed")
	}
	return videos, nil
}

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"BackEnd/graph/generated"
	"BackEnd/middleware"
	"BackEnd/models"
	"context"
	"errors"
	"math/rand"
	"strings"
	"time"
)

func (r *likeDetailResolver) Video(ctx context.Context, obj *models.LikeDetail) (*models.Video, error) {
	var video models.Video
	err := r.DB.Model(&video).Where("id = ?", obj.VideoID).Select()
	if err != nil {
		return nil, errors.New("query like video failed")
	}
	return &video, nil
}

func (r *likeDetailResolver) User(ctx context.Context, obj *models.LikeDetail) (*models.User, error) {
	var user models.User
	err := r.DB.Model(&user).Where("id = ?", obj.UserID).Select()
	if err != nil {
		return nil, errors.New("query like user failed")
	}
	return &user, nil
}

func (r *mutationResolver) CreateVideo(ctx context.Context, input *models.NewVideo) (*models.Video, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("unauthenticated")
	}
	video := models.Video{
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
		UserID:      currentUser.ID,
		Location:    input.Location,
	}
	if !strings.Contains(*input.DatePublish, "undefined") {
		video.DatePublish = *input.DatePublish
	} else {
		video.DatePublish = time.Now().Format("2006-01-02 15:04:05")
	}

	_, err = r.DB.Model(&video).Insert()
	if err != nil {
		return nil, errors.New("insert video failed")
	}
	return &video, nil
}

func (r *mutationResolver) UpdateVideo(ctx context.Context, id string, input *models.NewVideo) (*models.Video, error) {
	var video models.Video

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

func (r *mutationResolver) UpdateVideoView(ctx context.Context, id string) (*models.Video, error) {
	var video models.Video
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

func (r *mutationResolver) UpdateVideoLike(ctx context.Context, id string) (*models.LikeDetail, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("unauthenticated")
	}
	var videoDetail models.LikeDetail
	err = r.DB.Model(&videoDetail).Where("user_id=? and video_id =?", currentUser.ID, id).First()
	if err != nil {
		videoDetail.Like = true
		videoDetail.UserID = currentUser.ID
		videoDetail.VideoID = id
		_, updateErr := r.DB.Model(&videoDetail).Where("user_id=? and video_id =?", currentUser.ID, id).Update()
		if updateErr != nil {
			return nil, errors.New("update like error")
		}
		return &videoDetail, nil
	}

	video := models.LikeDetail{
		VideoID: id,
		UserID:  currentUser.ID,
		Like:    true,
	}

	_, err = r.DB.Model(&video).Insert()
	if err != nil {
		return nil, errors.New("like video failed")
	}
	return &video, nil
}

func (r *mutationResolver) UpdateVideoDislike(ctx context.Context, id string) (*models.LikeDetail, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("unauthenticated")
	}
	var videoDetail models.LikeDetail
	err = r.DB.Model(&videoDetail).Where("user_id=? and video_id =?", currentUser.ID, id).First()
	if err != nil {
		videoDetail = models.LikeDetail{
			VideoID: id,
			UserID:  currentUser.ID,
			Like:    false,
		}
		_, updateErr := r.DB.Model(&videoDetail).Where("user_id=? and video_id =?", currentUser.ID, id).Update()
		if updateErr != nil {
			return nil, errors.New("update like error")
		}
		return &videoDetail, nil
	}

	video := models.LikeDetail{
		VideoID: id,
		UserID:  currentUser.ID,
		Like:    false,
	}

	_, err = r.DB.Model(&video).Insert()
	if err != nil {
		return nil, errors.New("dislike video failed")
	}
	return &video, nil
}

func (r *mutationResolver) DeleteVideo(ctx context.Context, id string) (bool, error) {
	var video models.Video

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

func (r *queryResolver) Videos(ctx context.Context) ([]*models.Video, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	var videos []*models.Video
	if err != nil || currentUser.MembershipID == "" {
		err = r.DB.Model(&videos).Order("id").Where("NOW() > date_publish ").Select()
		if err != nil {
			return nil, errors.New("query video failed")
		}

	} else {
		err = r.DB.Model(&videos).Order("id").Select()
		if err != nil {
			return nil, errors.New("query video failed")
		}
	}
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(videos), func(i, j int) {
		videos[i], videos[j] = videos[j], videos[i]
	})
	return videos, nil
}

func (r *queryResolver) GetVideo(ctx context.Context, id string) (*models.Video, error) {
	var video models.Video
	err := r.DB.Model(&video).Where("id=?", id).Select()
	if err != nil {
		return nil, errors.New("query video failed")
	}

	return &video, nil
}

func (r *queryResolver) GetVideosByCategory(ctx context.Context, category int) ([]*models.Video, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	var videos []*models.Video
	if err != nil || currentUser.MembershipID == "" {
		err = r.DB.Model(&videos).Order("id").Where("category=? and NOW() > date_publish", category).Select()
		if err != nil {
			return nil, errors.New("query video failed")
		}

	} else {
		err := r.DB.Model(&videos).Where("category=?", category).Order("id").Select()
		if err != nil {
			return nil, errors.New("query video failed")
		}
	}
	return videos, nil
}

func (r *queryResolver) GetVideosByUser(ctx context.Context, userID string) ([]*models.Video, error) {
	var videos []*models.Video
	err := r.DB.Model(&videos).Where("user_id=?", userID).Order("id").Select()
	if err != nil {
		return nil, errors.New("query video failed")
	}
	return videos, nil
}

func (r *queryResolver) GetVideosByUserLimit(ctx context.Context, userID string, limit int) ([]*models.Video, error) {
	var videos []*models.Video
	err := r.DB.Model(&videos).Where("user_id=?", userID).Order("id").Limit(limit).Select()
	if err != nil {
		return nil, errors.New("query video with limit failed")
	}
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(videos), func(i, j int) {
		videos[i], videos[j] = videos[j], videos[i]
	})

	return videos, nil
}

func (r *queryResolver) GetLastUploadedVideo(ctx context.Context, userID string) (*models.Video, error) {
	var video models.Video

	err := r.DB.Model(&video).Where("user_id=?", userID).Last()
	if err != nil {
		return nil, errors.New("query last video failed")
	}

	return &video, nil
}

func (r *queryResolver) GetVideoSearch(ctx context.Context, keyword string) ([]*models.Video, error) {
	var videos []*models.Video
	keyword = strings.ToLower(keyword)
	keyword = "%" + keyword + "%"
	err := r.DB.Model(&videos).Where("lower(title) LIKE ? or lower(description) LIKE ?", keyword, keyword).Order("title").Select()
	if err != nil {
		return nil, errors.New("query video search failed")
	}
	return videos, nil
}

func (r *queryResolver) GetTrendingVideo(ctx context.Context) ([]*models.Video, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	var videos []*models.Video
	if err != nil || currentUser.MembershipID == "" {
		err = r.DB.Model(&videos).Order("id").Where("DATE_PART('day',NOW()-date_upload)<=7 or DATE_PART('month',NOW()-date_upload)<=1 and NOW() > date_publish").Order("view DESC").Limit(20).Select()
		if err != nil {
			return nil, errors.New("query trending failed")
		}

	} else {
		err := r.DB.Model(&videos).Where("DATE_PART('day',NOW()-date_upload)<=7 or DATE_PART('month',NOW()-date_upload)<=1").Order("view DESC").Limit(20).Select()
		if err != nil {
			return nil, errors.New("query trending failed")
		}
	}
	return videos, nil
}

func (r *queryResolver) GetRecommendationVideo(ctx context.Context, category int, videoID string) ([]*models.Video, error) {
	var videos []*models.Video
	err := r.DB.Model(&videos).Where("category=? and id != ?", category, videoID).Order("id").Select()
	if err != nil {
		return nil, errors.New("query video failed")
	}
	return videos, nil
}

func (r *queryResolver) GetVideoByLocation(ctx context.Context, location string) ([]*models.Video, error) {
	var videos []*models.Video
	err := r.DB.Model(&videos).Where("location = ?", location).Order("id").Select()
	if err != nil {
		return nil, errors.New("failed query video by location")
	}
	return videos, nil
}

func (r *queryResolver) GetVideosByTime(ctx context.Context, time string, keyword string) ([]*models.Video, error) {
	var videos []*models.Video
	keyword = strings.ToLower(keyword)
	keyword = "%" + keyword + "%"
	err := r.DB.Model(&videos).Where("date_part(?, NOW()) = extract( ? from date_upload) AND lower(title) LIKE ? ", time, time, keyword).Select()
	if err != nil {
		return nil, errors.New("failed query video by time")
	}
	return videos, nil
}

func (r *videoResolver) User(ctx context.Context, obj *models.Video) (*models.User, error) {
	//var user models.User
	//err := r.DB.Model(&user).Where("id=?", obj.UserID).Select()
	//if err != nil {
	//	return nil, errors.New("query user in video failed")
	//}
	//return &user, nil
	return getUserLoader(ctx).Load(obj.UserID)
}

func (r *videoResolver) Like(ctx context.Context, obj *models.Video) ([]*models.LikeDetail, error) {
	var likes []*models.LikeDetail

	err := r.DB.Model(&likes).Where("video_id = ?", obj.ID).Select()

	if err != nil {
		return nil, errors.New("failed get like")
	}

	return likes, nil
}

func (r *videoResolver) Comments(ctx context.Context, obj *models.Video) ([]*models.Comment, error) {
	var comments []*models.Comment

	err := r.DB.Model(&comments).Where("video_id=?", obj.ID).Order("id").Select()

	if err != nil {
		return nil, errors.New("failed query comments")
	}

	return comments, nil
}

// LikeDetail returns generated.LikeDetailResolver implementation.
func (r *Resolver) LikeDetail() generated.LikeDetailResolver { return &likeDetailResolver{r} }

// Video returns generated.VideoResolver implementation.
func (r *Resolver) Video() generated.VideoResolver { return &videoResolver{r} }

type likeDetailResolver struct{ *Resolver }
type videoResolver struct{ *Resolver }

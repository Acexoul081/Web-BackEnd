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
	"sort"
	"strings"
	"time"
)

func (r *mutationResolver) CreatePlaylist(ctx context.Context, input *models.NewPlaylist) (*models.Playlist, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("unauthenticated")
	}
	playlist := models.Playlist{
		Title:   input.Title,
		Privacy: input.Privacy,
		OwnerID: currentUser.ID,
		Date:    time.Now().Format("2006-01-02 15:04:05"),
	}
	_, err = r.DB.Model(&playlist).Insert()

	if err != nil {
		return nil, errors.New("playlist insert failed")
	}
	return &playlist, nil
}

func (r *mutationResolver) UpdatePlaylist(ctx context.Context, id string, input *models.UpdatedPlaylist) (*models.Playlist, error) {
	var playlist models.Playlist

	err := r.DB.Model(&playlist).Where("id=?", id).First()
	if err != nil {
		return nil, errors.New("playlist not found")
	}
	if input.Title != nil {
		playlist.Title = *input.Title
	}
	if input.Sort != nil {
		playlist.Sort = *input.Sort
	}
	if input.Description != nil {
		playlist.Description = *input.Description
	}
	if input.Privacy != nil {
		playlist.Privacy = *input.Privacy
	}
	_, updateErr := r.DB.Model(&playlist).Where("id=?", id).Update()

	if updateErr != nil {
		return nil, errors.New("update playlist failed")
	}
	return &playlist, nil
}

func (r *mutationResolver) DeletePlaylist(ctx context.Context, id string) (bool, error) {
	var playlist models.Playlist
	err := r.DB.Model(&playlist).Where("id=?", id).First()
	if err != nil {
		return false, errors.New("playlist not found")
	}

	_, deleteErr := r.DB.Model(&playlist).Where("id=?", id).Delete()
	if deleteErr != nil {
		return false, errors.New("delete playlist failed")
	}

	return true, nil
}

func (r *mutationResolver) CreatePlaylistDetail(ctx context.Context, input *models.NewPlaylistDetail) (*models.PlaylistDetail, error) {
	playlistDetail := models.PlaylistDetail{
		PlaylistID: input.PlaylistID,
		VideoID:    input.VideoID,
	}
	_, err := r.DB.Model(&playlistDetail).Insert()

	var playlist models.Playlist
	err = r.DB.Model(&playlist).Where("id=?", input.PlaylistID).First()
	if err != nil {
		return nil, errors.New("playlist not found")
	}
	playlist.Date = time.Now().Format("2006-01-02 15:04:05")
	_, errUpdate := r.DB.Model(&playlist).Where("id=?", input.PlaylistID).Update()
	if errUpdate != nil {
		return nil, errors.New("failed get playlist updated")
	}

	if err != nil {
		return nil, errors.New("playlist detail insert failed")
	}
	return &playlistDetail, nil
}

func (r *mutationResolver) UpdatePlaylistDetail(ctx context.Context, id string, input *models.NewPlaylistDetail) (*models.PlaylistDetail, error) {
	var playlistDetail models.PlaylistDetail

	err := r.DB.Model(&playlistDetail).Where("id=?", id).First()
	if err != nil {
		return nil, errors.New("playlist detail not found")
	}

	_, updateErr := r.DB.Model(&playlistDetail).Where("id=?", id).Update()

	if updateErr != nil {
		return nil, errors.New("playlist detail update failed")
	}
	return &playlistDetail, nil
}

func (r *mutationResolver) DeletePlaylistDetail(ctx context.Context, playlistID string, videoID string) (bool, error) {
	var playlistDetail models.PlaylistDetail

	err := r.DB.Model(&playlistDetail).Where("playlist_id=? and video_id=?", playlistID, videoID).Select()
	if err != nil {
		return false, errors.New("playlist detail not found")
	}

	_, deleteErr := r.DB.Model(&playlistDetail).Where("playlist_id=? and video_id=?", playlistID, videoID).Delete()
	if deleteErr != nil {
		return false, errors.New("delete playlist detail error")
	}
	return true, nil
}

func (r *mutationResolver) CreatePlaylistSub(ctx context.Context, input *models.NewPlaylistSub) (*models.PlaylistSub, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("unauthenticated")
	}
	var sub models.PlaylistSub

	err = r.DB.Model(&sub).Where("playlist_id = ? and user_id=?", input.PlaylistID, currentUser.ID).Select()
	if err == nil {
		return nil, errors.New("already subscribe this playlist")
	}
	newSub := models.PlaylistSub{
		PlaylistID: input.PlaylistID,
		UserID:     currentUser.ID,
	}
	_, err = r.DB.Model(&newSub).Insert()
	if err != nil {
		return nil, errors.New("insert failed")
	}
	return &newSub, nil
}

func (r *mutationResolver) DeletePlaylistSub(ctx context.Context, playlistID string) (bool, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return false, errors.New("unauthenticated")
	}
	var sub models.PlaylistSub
	_, err = r.DB.Model(&sub).Where("user_id = ? and playlist_id = ?", currentUser.ID, playlistID).Delete()
	if err != nil {
		return false, errors.New("failed delete sub")
	}
	return true, nil
}

func (r *mutationResolver) SwapPlaylistDetail(ctx context.Context, newID string, tempID string) (*models.PlaylistDetail, error) {
	var playlistDetail models.PlaylistDetail
	err := r.DB.Model(&playlistDetail).Where("temp_id=?", tempID).Select()

	if err != nil {
		return nil, errors.New("detail playlist not found")
	}

	playlistDetail.VideoID = newID

	_, updateErr := r.DB.Model(&playlistDetail).Where("temp_id=?", tempID).Update()

	if updateErr != nil {
		return nil, errors.New("update playlist swap failed")
	}

	return &playlistDetail, nil
}

func (r *playlistResolver) Owner(ctx context.Context, obj *models.Playlist) (*models.User, error) {
	var owner models.User
	err := r.DB.Model(&owner).Where("id=?", obj.OwnerID).Select()
	if err != nil {
		return nil, errors.New("failed get owner from playlist")
	}
	return &owner, nil
}

func (r *playlistResolver) PlaylistDetail(ctx context.Context, obj *models.Playlist) ([]*models.PlaylistDetail, error) {
	var details []*models.PlaylistDetail
	err := r.DB.Model(&details).Where("playlist_id=?", obj.ID).Select()
	if err != nil {
		return nil, errors.New("failed get detail from playlist")
	}
	return details, nil
}

func (r *playlistResolver) PlaylistSub(ctx context.Context, obj *models.Playlist) ([]*models.PlaylistSub, error) {
	var sub []*models.PlaylistSub
	err := r.DB.Model(&sub).Where("playlist_id=?", obj.ID).Select()
	if err != nil {
		return nil, errors.New("failed get sub from playlist")
	}
	return sub, nil
}

func (r *playlistResolver) Videos(ctx context.Context, obj *models.Playlist) ([]*models.Video, error) {
	var details []*models.PlaylistDetail
	err := r.DB.Model(&details).Where("playlist_id=?", obj.ID).Select()
	if err != nil {
		return nil, errors.New("failed get detail from playlist")
	}
	var videos []*models.Video
	for _, item := range details {
		var video models.Video
		err = r.DB.Model(&video).Where("id=?", item.VideoID).First()
		if err != nil {
			return nil, errors.New("failed query video of playlist best")
		}
		videos = append(videos, &video)
	}
	if obj.Sort == "date add new" {
		sort.SliceStable(videos, func(i, j int) bool {
			return videos[i].DateUpload > videos[j].DateUpload
		})
	} else if obj.Sort == "date add old" {
		sort.SliceStable(videos, func(i, j int) bool {
			return videos[i].DateUpload < videos[j].DateUpload
		})
	} else if obj.Sort == "popular" {
		sort.SliceStable(videos, func(i, j int) bool {
			return videos[i].View > videos[j].View
		})
	} else if obj.Sort == "date publish new" {
		sort.SliceStable(videos, func(i, j int) bool {
			return videos[i].DatePublish > videos[j].DatePublish
		})
	} else if obj.Sort == "date publish old" {
		sort.SliceStable(videos, func(i, j int) bool {
			return videos[i].DatePublish < videos[j].DatePublish
		})
	}
	return videos, nil
}

func (r *playlistDetailResolver) Video(ctx context.Context, obj *models.PlaylistDetail) (*models.Video, error) {
	var video models.Video
	err := r.DB.Model(&video).Where("id=?", obj.VideoID).Select()
	if err != nil {
		return nil, errors.New("failed query video from playlist detail")
	}
	return &video, nil
}

func (r *playlistSubResolver) Playlist(ctx context.Context, obj *models.PlaylistSub) (*models.Playlist, error) {
	var playlist models.Playlist
	err := r.DB.Model(&playlist).Where("id=?", obj.PlaylistID).Select()
	if err != nil {
		return nil, errors.New("playlist not found")
	}
	return &playlist, nil
}

func (r *playlistSubResolver) User(ctx context.Context, obj *models.PlaylistSub) (*models.User, error) {
	var user models.User
	err := r.DB.Model(&user).Where("id=?", obj.UserID).Select()
	if err != nil {
		return nil, errors.New("failed get sub user from playlist")
	}
	return &user, nil
}

func (r *queryResolver) Playlists(ctx context.Context) ([]*models.Playlist, error) {
	var playlists []*models.Playlist

	err := r.DB.Model(&playlists).Order("id").Select()

	if err != nil {
		return nil, errors.New("failed to query playlists")
	}
	return playlists, nil
}

func (r *queryResolver) PlaylistDetails(ctx context.Context) ([]*models.PlaylistDetail, error) {
	var playlistDetails []*models.PlaylistDetail

	err := r.DB.Model(&playlistDetails).Order("video_id").Select()

	if err != nil {
		return nil, errors.New("failed to query playlist details")
	}
	return playlistDetails, nil
}

func (r *queryResolver) GetPlaylistByOwnerID(ctx context.Context, id string, limit *int) ([]*models.Playlist, error) {
	var playlists []*models.Playlist

	query := r.DB.Model(&playlists).Where("owner_id = ?", id).Order("id")
	if limit != nil {
		query.Limit(*limit)
	}
	err := query.Select()
	if err != nil {
		return nil, errors.New("failed to query playlists by owner")
	}

	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(playlists), func(i, j int) {
		playlists[i], playlists[j] = playlists[j], playlists[i]
	})
	return playlists, nil
}

func (r *queryResolver) GetPlaylistByID(ctx context.Context, id string) (*models.Playlist, error) {
	var playlist models.Playlist

	err := r.DB.Model(&playlist).Where("id = ?", id).Select()

	if err != nil {
		return nil, errors.New("failed query playlist by id")
	}

	return &playlist, nil
}

func (r *queryResolver) GetVideoCountOnPlaylist(ctx context.Context, id string) (int, error) {
	var count int
	var playlistDetail models.PlaylistDetail
	count, err := r.DB.Model(&playlistDetail).Where("playlist_id=?", id).Count()
	if err != nil {
		return 0, errors.New("playlist not found")
	}
	return count, nil
}

func (r *queryResolver) GetVideoOfPlaylist(ctx context.Context, id string) ([]*models.Video, error) {
	var playlistDetails []*models.PlaylistDetail

	err := r.DB.Model(&playlistDetails).Where("playlist_id=?", id).Order("video_id").Select()

	if err != nil {
		return nil, errors.New("failed to query playlist detail videos")
	}

	var videos []*models.Video
	for _, item := range playlistDetails {
		var video models.Video
		err = r.DB.Model(&video).Where("id=?", item.VideoID).First()
		if err != nil {
			return nil, errors.New("failed query video of playlist good")
		}
		videos = append(videos, &video)
	}
	sort.SliceStable(videos, func(i, j int) bool {
		return videos[i].Title < videos[j].Title
	})
	return videos, nil
}

func (r *queryResolver) GetFirstVideoOfPlaylist(ctx context.Context, id string) (*models.PlaylistDetail, error) {
	var detail []*models.PlaylistDetail
	err := r.DB.Model(&detail).Where("playlist_id=?", id).Select()
	if err != nil {
		return nil, errors.New("query first video failed")
	}

	return detail[0], nil
}

func (r *queryResolver) GetPlaylistSearch(ctx context.Context, keyword string) ([]*models.Playlist, error) {
	var playlists []*models.Playlist
	keyword = strings.ToLower(keyword)
	keyword = "%" + keyword + "%"
	err := r.DB.Model(&playlists).Where("lower(title) LIKE ?", keyword).Order("title").Select()
	if err != nil {
		return nil, errors.New("query playlist search failed")
	}
	return playlists, nil
}

func (r *queryResolver) GetPlaylistSub(ctx context.Context) ([]*models.PlaylistSub, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("unauthenticated")
	}
	var subs []*models.PlaylistSub
	err = r.DB.Model(&subs).Where("user_id = ?", currentUser.ID).Select()
	if err != nil {
		return nil, errors.New("error get playlist sub from user  id")
	}
	return subs, nil
}

// Playlist returns generated.PlaylistResolver implementation.
func (r *Resolver) Playlist() generated.PlaylistResolver { return &playlistResolver{r} }

// PlaylistDetail returns generated.PlaylistDetailResolver implementation.
func (r *Resolver) PlaylistDetail() generated.PlaylistDetailResolver {
	return &playlistDetailResolver{r}
}

// PlaylistSub returns generated.PlaylistSubResolver implementation.
func (r *Resolver) PlaylistSub() generated.PlaylistSubResolver { return &playlistSubResolver{r} }

type playlistResolver struct{ *Resolver }
type playlistDetailResolver struct{ *Resolver }
type playlistSubResolver struct{ *Resolver }

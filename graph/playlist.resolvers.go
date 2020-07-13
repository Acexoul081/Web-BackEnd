package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"BackEnd/graph/model"
	"context"
	"errors"
)

func (r *mutationResolver) CreatePlaylist(ctx context.Context, input *model.NewPlaylist) (*model.Playlist, error) {
	playlist := model.Playlist{
		OwnerID: input.OwnerID,
		Title:   input.Title,
	}
	_, err := r.DB.Model(&playlist).Insert()

	if err != nil {
		return nil, errors.New("playlist insert failed")
	}
	return &playlist, nil
}

func (r *mutationResolver) UpdatePlaylist(ctx context.Context, id string, input *model.NewPlaylist) (*model.Playlist, error) {
	var playlist model.Playlist

	err := r.DB.Model(&playlist).Where("id=?", id).First()
	if err != nil {
		return nil, errors.New("playlist not found")
	}
	playlist.Title = input.Title
	_, updateErr := r.DB.Model(&playlist).Where("id=?", id).Update()

	if updateErr != nil {
		return nil, errors.New("update playlist failed")
	}
	return &playlist, nil
}

func (r *mutationResolver) DeletePlaylist(ctx context.Context, id string) (bool, error) {
	var playlist model.Playlist
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

func (r *mutationResolver) CreatePlaylistDetail(ctx context.Context, input *model.NewPlaylistDetail) (*model.PlaylistDetail, error) {
	playlistDetail := model.PlaylistDetail{
		PlaylistID: input.PlaylistID,
		VideoID:    input.VideoID,
	}
	_, err := r.DB.Model(&playlistDetail).Insert()

	if err != nil {
		return nil, errors.New("playlist detail insert failed")
	}
	return &playlistDetail, nil
}

func (r *mutationResolver) UpdatePlaylistDetail(ctx context.Context, id string, input *model.NewPlaylistDetail) (*model.PlaylistDetail, error) {
	var playlistDetail model.PlaylistDetail

	err := r.DB.Model(&playlistDetail).Where("id=?", id).First()
	if err != nil {
		return nil, errors.New("playlist detail not found")
	}

	playlistDetail.VideoID = input.VideoID

	_, updateErr := r.DB.Model(&playlistDetail).Where("id=?", id).Update()

	if updateErr != nil {
		return nil, errors.New("playlist detail update failed")
	}
	return &playlistDetail, nil
}

func (r *mutationResolver) DeletePlaylistDetail(ctx context.Context, id string) (bool, error) {
	var playlistDetail model.PlaylistDetail

	err := r.DB.Model(&playlistDetail).Where("id=?", id).First()
	if err != nil {
		return false, errors.New("playlist detail not found")
	}

	_, deleteErr := r.DB.Model(&playlistDetail).Where("id=?", id).Delete()
	if deleteErr != nil {
		return false, errors.New("delete playlist detail error")
	}
	return true, nil
}

func (r *queryResolver) Playlists(ctx context.Context) ([]*model.Playlist, error) {
	var playlists []*model.Playlist

	err := r.DB.Model(&playlists).Order("id").Select()

	if err != nil {
		return nil, errors.New("failed to query playlists")
	}
	return playlists, nil
}

func (r *queryResolver) PlaylistDetails(ctx context.Context) ([]*model.PlaylistDetail, error) {
	var playlistDetails []*model.PlaylistDetail

	err := r.DB.Model(&playlistDetails).Order("video_id").Select()

	if err != nil {
		return nil, errors.New("failed to query playlist details")
	}
	return playlistDetails, nil
}

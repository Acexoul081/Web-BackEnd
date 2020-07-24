package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"BackEnd/graph/model"
	"context"
	"errors"
	"time"
)

func (r *mutationResolver) CreateComment(ctx context.Context, input *model.NewComment) (*model.Comment, error) {
	comment := model.Comment{
		UserID:      input.UserID,
		VideoID:     input.VideoID,
		Comment:     input.Comment,
		CommentDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	_, err := r.DB.Model(&comment).Insert()

	if err != nil {
		return nil, errors.New("insert comment failed")
	}
	return &comment, nil
}

func (r *mutationResolver) UpdateComment(ctx context.Context, id string, input *model.NewComment) (*model.Comment, error) {
	var comment model.Comment

	err := r.DB.Model(&comment).Where("id=?", id).First()
	if err != nil {
		return nil, errors.New("comment not found")
	}

	comment.Comment = input.Comment
	_, updateErr := r.DB.Model(&comment).Where("id=?", id).Update()

	if updateErr != nil {
		return nil, errors.New("update comment failed")
	}
	return &comment, nil
}

func (r *mutationResolver) UpdateCommentLike(ctx context.Context, id string) (*model.Comment, error) {
	var comment model.Comment
	err := r.DB.Model(&comment).Where("id=?", id).First()
	if err != nil {
		return nil, errors.New("comment not found")
	}
	comment.Like += 1
	_, updateErr := r.DB.Model(&comment).Where("id=?", id).Update()
	if updateErr != nil {
		return nil, errors.New("update like failed")
	}
	return &comment, nil
}

func (r *mutationResolver) UpdateCommentDislike(ctx context.Context, id string) (*model.Comment, error) {
	var comment model.Comment
	err := r.DB.Model(&comment).Where("id=?", id).First()
	if err != nil {
		return nil, errors.New("comment not found")
	}
	comment.Dislike += 1
	_, updateErr := r.DB.Model(&comment).Where("id=?", id).Update()
	if updateErr != nil {
		return nil, errors.New("update dislike failed")
	}
	return &comment, nil
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (bool, error) {
	var comment model.Comment

	err := r.DB.Model(&comment).Where("id=?", id).First()
	if err != nil {
		return false, errors.New("comment not found")
	}
	_, deleteErr := r.DB.Model(&comment).Where("id=?", id).Delete()

	if deleteErr != nil {
		return false, errors.New("delete comment failed")
	}
	return true, nil
}

func (r *mutationResolver) CreateReply(ctx context.Context, input *model.NewReply) (*model.Reply, error) {
	reply := model.Reply{
		UserID:    input.UserID,
		CommentID: input.CommentID,
		Reply:     input.Reply,
		ReplyDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	_, err := r.DB.Model(&reply).Insert()

	if err != nil {
		return nil, errors.New("insert reply failed")
	}
	return &reply, nil
}

func (r *mutationResolver) UpdateReply(ctx context.Context, id string, input *model.NewReply) (*model.Reply, error) {
	var reply model.Reply

	err := r.DB.Model(&reply).Where("id=?", id).First()
	if err != nil {
		return nil, errors.New("reply not found")
	}

	reply.Reply = input.Reply

	_, updateErr := r.DB.Model(&reply).Where("id=?", id).Update()

	if updateErr != nil {
		return nil, errors.New("update reply failed")
	}
	return &reply, nil
}

func (r *mutationResolver) UpdateReplyLike(ctx context.Context, id string) (*model.Reply, error) {
	var reply model.Reply
	err := r.DB.Model(&reply).Where("id=?", id).First()
	if err != nil {
		return nil, errors.New("reply not found")
	}
	reply.Like += 1
	_, updateErr := r.DB.Model(&reply).Where("id=?", id).Update()
	if updateErr != nil {
		return nil, errors.New("update like failed")
	}
	return &reply, nil
}

func (r *mutationResolver) UpdateReplyDislike(ctx context.Context, id string) (*model.Reply, error) {
	var reply model.Reply
	err := r.DB.Model(&reply).Where("id=?", id).First()
	if err != nil {
		return nil, errors.New("reply not found")
	}
	reply.Dislike += 1
	_, updateErr := r.DB.Model(&reply).Where("id=?", id).Update()
	if updateErr != nil {
		return nil, errors.New("update like failed")
	}
	return &reply, nil
}

func (r *mutationResolver) DeleteReply(ctx context.Context, id string) (bool, error) {
	var reply model.Reply

	err := r.DB.Model(&reply).Where("id=?", id).First()
	if err != nil {
		return false, errors.New("reply not found")
	}

	_, deleteErr := r.DB.Model(&reply).Where("id=?", id).Delete()

	if deleteErr != nil {
		return false, errors.New("delete reply error")
	}
	return true, nil
}

func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {
	var comments []*model.Comment

	err := r.DB.Model(&comments).Order("id").Select()

	if err != nil {
		return nil, errors.New("failed query comments")
	}

	return comments, nil
}

func (r *queryResolver) Replies(ctx context.Context) ([]*model.Reply, error) {
	var replies []*model.Reply

	err := r.DB.Model(&replies).Order("id").Select()

	if err != nil {
		return nil, errors.New("failed query replies")
	}
	return replies, nil
}

func (r *queryResolver) GetCommentCount(ctx context.Context, id string) (int, error) {
	var count int
	var comment model.Comment
	count, err := r.DB.Model(&comment).Where("video_id=?", id).Count()
	if err != nil {
		return 0, errors.New("comment not found")
	}
	return count, nil
}

func (r *queryResolver) GetCommentsByID(ctx context.Context, id string) ([]*model.Comment, error) {
	var comments []*model.Comment

	err := r.DB.Model(&comments).Where("video_id=?", id).Order("id").Select()

	if err != nil {
		return nil, errors.New("failed query comments")
	}

	return comments, nil
}

func (r *queryResolver) GetRepliesByComment(ctx context.Context, commentID string) ([]*model.Reply, error) {
	var replies []*model.Reply

	err := r.DB.Model(&replies).Where("comment_id=?", commentID).Order("id").Select()

	if err != nil {
		return nil, errors.New("failed query replies")
	}

	return replies, nil
}

func (r *queryResolver) GetReplyCount(ctx context.Context, commentID string) (int, error) {
	var count int
	var reply model.Reply
	count, err := r.DB.Model(&reply).Where("comment_id=?", commentID).Count()
	if err != nil {
		return 0, errors.New("reply not found")
	}
	return count, nil
}

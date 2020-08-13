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

func (r *commentResolver) User(ctx context.Context, obj *models.Comment) (*models.User, error) {
	var user models.User
	err := r.DB.Model(&user).Where("id=?", obj.UserID).Select()
	if err != nil {
		return nil, errors.New("failed query user from comment")
	}
	return &user, nil
}

func (r *commentResolver) Like(ctx context.Context, obj *models.Comment) ([]*models.CommentLikeDetail, error) {
	var comments []*models.CommentLikeDetail
	err := r.DB.Model(&comments).Where("comment_id=?", obj.ID).Select()
	if err != nil {
		return nil, errors.New("failed query comment detail from comment")
	}
	return comments, nil
}

func (r *commentResolver) Replies(ctx context.Context, obj *models.Comment) ([]*models.Reply, error) {
	var replies []*models.Reply
	err := r.DB.Model(&replies).Where("comment_id=?", obj.ID).Select()
	if err != nil {
		return nil, errors.New("failed query reply from comment")
	}
	return replies, nil
}

func (r *commentLikeDetailResolver) Comment(ctx context.Context, obj *models.CommentLikeDetail) (*models.Comment, error) {
	var comment models.Comment
	err := r.DB.Model(&comment).Where("id = ?", obj.CommentID).Select()
	if err != nil {
		return nil, errors.New("failed query comment from comment detail")
	}
	return &comment, nil
}

func (r *commentLikeDetailResolver) User(ctx context.Context, obj *models.CommentLikeDetail) (*models.User, error) {
	var user models.User
	err := r.DB.Model(&user).Where("id=?", obj.UserID).Select()
	if err != nil {
		return nil, errors.New("failed query user from comment detail")
	}
	return &user, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, input *models.NewComment) (*models.Comment, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("unauthenticated")
	} 
	comment := models.Comment{
		VideoID:     input.VideoID,
		Comment:     input.Comment,
		CommentDate: time.Now().Format("2006-01-02 15:04:05"),
		UserID: currentUser.ID,
	}

	_, err = r.DB.Model(&comment).Insert()

	if err != nil {
		return nil, errors.New("insert comment failed")
	}
	return &comment, nil
}

func (r *mutationResolver) UpdateComment(ctx context.Context, id string, input *models.NewComment) (*models.Comment, error) {
	var comment models.Comment

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

func (r *mutationResolver) UpdateCommentLike(ctx context.Context, id string) (*models.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCommentDislike(ctx context.Context, id string) (*models.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (bool, error) {
	var comment models.Comment

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

func (r *mutationResolver) CreateReply(ctx context.Context, input *models.NewReply) (*models.Reply, error) {
	reply := models.Reply{
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

func (r *mutationResolver) UpdateReply(ctx context.Context, id string, input *models.NewReply) (*models.Reply, error) {
	var reply models.Reply

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

func (r *mutationResolver) UpdateReplyLike(ctx context.Context, id string) (*models.Reply, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateReplyDislike(ctx context.Context, id string) (*models.Reply, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteReply(ctx context.Context, id string) (bool, error) {
	var reply models.Reply

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

func (r *queryResolver) Comments(ctx context.Context) ([]*models.Comment, error) {
	var comments []*models.Comment

	err := r.DB.Model(&comments).Order("id").Select()

	if err != nil {
		return nil, errors.New("failed query comments")
	}

	return comments, nil
}

func (r *queryResolver) Replies(ctx context.Context) ([]*models.Reply, error) {
	var replies []*models.Reply

	err := r.DB.Model(&replies).Order("id").Select()

	if err != nil {
		return nil, errors.New("failed query replies")
	}
	return replies, nil
}

func (r *queryResolver) GetCommentCount(ctx context.Context, id string) (int, error) {
	var count int
	var comment models.Comment
	count, err := r.DB.Model(&comment).Where("video_id=?", id).Count()
	if err != nil {
		return 0, errors.New("comment not found")
	}
	return count, nil
}

func (r *queryResolver) GetCommentsByID(ctx context.Context, id string) ([]*models.Comment, error) {
	var comments []*models.Comment

	err := r.DB.Model(&comments).Where("video_id=?", id).Order("id").Select()

	if err != nil {
		return nil, errors.New("failed query comments")
	}

	return comments, nil
}

func (r *queryResolver) GetRepliesByComment(ctx context.Context, commentID string) ([]*models.Reply, error) {
	var replies []*models.Reply

	err := r.DB.Model(&replies).Where("comment_id=?", commentID).Order("id").Select()

	if err != nil {
		return nil, errors.New("failed query replies")
	}

	return replies, nil
}

func (r *queryResolver) GetReplyCount(ctx context.Context, commentID string) (int, error) {
	var count int
	var reply models.Reply
	count, err := r.DB.Model(&reply).Where("comment_id=?", commentID).Count()
	if err != nil {
		return 0, errors.New("reply not found")
	}
	return count, nil
}

func (r *replyResolver) User(ctx context.Context, obj *models.Reply) (*models.User, error) {
	var user models.User
	err := r.DB.Model(&user).Where("id=?", obj.UserID).Select()
	if err != nil {
		return nil, errors.New("query user from reply failed")
	}
	return &user, nil
}

func (r *replyResolver) Like(ctx context.Context, obj *models.Reply) ([]*models.ReplyLikeDetail, error) {
	var replyDetails []*models.ReplyLikeDetail
	err := r.DB.Model(&replyDetails).Where("reply_id=?", obj.ID).Select()
	if err != nil {
		return nil, errors.New("failed query reply detail from reply")
	}
	return replyDetails, nil
}

func (r *replyLikeDetailResolver) Reply(ctx context.Context, obj *models.ReplyLikeDetail) (*models.Reply, error) {
	var reply models.Reply
	err := r.DB.Model(&reply).Where("id=?", obj.ReplyID).Select()
	if err != nil {
		return nil, errors.New("query reply from reply like error")
	}
	return &reply, nil
}

func (r *replyLikeDetailResolver) User(ctx context.Context, obj *models.ReplyLikeDetail) (*models.User, error) {
	var user models.User
	err := r.DB.Model(&user).Where("id=?", obj.UserID).Select()
	if err != nil {
		return nil, errors.New("query user from rely like error")
	}
	return &user, nil
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

// CommentLikeDetail returns generated.CommentLikeDetailResolver implementation.
func (r *Resolver) CommentLikeDetail() generated.CommentLikeDetailResolver {
	return &commentLikeDetailResolver{r}
}

// Reply returns generated.ReplyResolver implementation.
func (r *Resolver) Reply() generated.ReplyResolver { return &replyResolver{r} }

// ReplyLikeDetail returns generated.ReplyLikeDetailResolver implementation.
func (r *Resolver) ReplyLikeDetail() generated.ReplyLikeDetailResolver {
	return &replyLikeDetailResolver{r}
}

type commentResolver struct{ *Resolver }
type commentLikeDetailResolver struct{ *Resolver }
type replyResolver struct{ *Resolver }
type replyLikeDetailResolver struct{ *Resolver }

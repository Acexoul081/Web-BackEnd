package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"BackEnd/graph/generated"
	"BackEnd/middleware"
	"BackEnd/models"
	"context"
	"errors"
)

func (r *abonemenResolver) User(ctx context.Context, obj *models.Abonemen) (*models.User, error) {
	var users models.User
	err := r.DB.Model(&users).Where("id=?", obj.UserID).Select()
	if err != nil {
		return nil, errors.New("failed query user from subscription")
	}
	return &users, nil
}

func (r *abonemenResolver) Subscriber(ctx context.Context, obj *models.Abonemen) (*models.User, error) {
	var users models.User
	err := r.DB.Model(&users).Where("id=?", obj.SubscriberID).Select()
	if err != nil {
		return nil, errors.New("failed query user from subscription")
	}
	return &users, nil
}

func (r *mutationResolver) CreateSubscription(ctx context.Context, input *models.NewSubscription) (*models.Abonemen, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("unauthenticated")
	}
	var sub models.Abonemen
	err = r.DB.Model(&sub).Where("user_id = ? and subscriber_id = ?", input.UserID, currentUser.ID).Select()
	if err == nil {
		return nil, errors.New("already subscribe")
	}
	if currentUser.ID == input.UserID {
		return nil, errors.New("cant subscribe own channel")
	}
	sub = models.Abonemen{
		UserID:       input.UserID,
		SubscriberID: currentUser.ID,
	}
	_, err = r.DB.Model(&sub).Insert()
	if err != nil {
		return nil, errors.New("failed to subscribe")
	}
	return &sub, nil
}

func (r *mutationResolver) UpdateSubscription(ctx context.Context, id string, notif bool) (*models.Abonemen, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("unauthenticated")
	}
	var subscrip models.Abonemen
	err = r.DB.Model(&subscrip).Where("user_id = ? and subscriber_id=?", id, currentUser.ID).Select()
	if err != nil {
		return nil, errors.New("failed get subscription")
	}
	subscrip.Notification = notif
	_, err = r.DB.Model(&subscrip).Where("user_id = ? and subscriber_id=?", id, currentUser.ID).Update()
	if err != nil {
		return nil, errors.New("failed update notification")
	}
	return &subscrip, nil
}

func (r *mutationResolver) DeleteSubscription(ctx context.Context, id string) (bool, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return false, errors.New("unauthenticated")
	}
	var subscription models.Abonemen
	err = r.DB.Model(&subscription).Where("user_id=? and subscriber_id=?", id, currentUser.ID).Select()
	if err != nil {
		return false, errors.New("subscription not found")
	}
	_, deleteErr := r.DB.Model(&subscription).Where("user_id=? and subscriber_id=?", id, currentUser.ID).Delete()
	if deleteErr != nil {
		return false, errors.New("delete subscription failed")
	}
	return true, nil
}

func (r *queryResolver) Subscriptions(ctx context.Context) ([]*models.Abonemen, error) {
	var subscriptions []*models.Abonemen

	err := r.DB.Model(&subscriptions).Order("user_id").Select()

	if err != nil {
		return nil, errors.New("subscriptions query failed")
	}
	return subscriptions, nil
}

func (r *queryResolver) GetSubscriberCount(ctx context.Context, id string) (int, error) {
	var count int
	var subscription models.Abonemen
	count, err := r.DB.Model(&subscription).Where("user_id=?", id).Count()
	if err != nil {
		return 0, errors.New("subscriber not found")
	}
	return count, nil
}

// Abonemen returns generated.AbonemenResolver implementation.
func (r *Resolver) Abonemen() generated.AbonemenResolver { return &abonemenResolver{r} }

type abonemenResolver struct{ *Resolver }

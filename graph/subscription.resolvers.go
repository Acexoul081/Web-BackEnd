package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"BackEnd/graph/model"
	"context"
	"errors"
)

func (r *mutationResolver) CreateSubscription(ctx context.Context, input *model.NewSubscription) (*model.Abonemen, error) {
	subscription := model.Abonemen{
		UserID:       input.UserID,
		SubscriberID: input.SubscriberID,
	}
	_, err := r.DB.Model(&subscription).Insert()

	if err != nil {
		return nil, errors.New("insert subscriber failed")
	}
	return &subscription, nil
}

func (r *mutationResolver) UpdateSubscription(ctx context.Context, id string, input *model.NewSubscription) (*model.Abonemen, error) {
	var subscription model.Abonemen
	err := r.DB.Model(&subscription).Where("id=?", id).First()
	if err != nil {
		return nil, errors.New("subscriber not found")
	}
	subscription.SubscriberID = input.SubscriberID

	_, updateErr := r.DB.Model(&subscription).Where("id=?", id).Update()
	if updateErr != nil {
		return nil, errors.New("update subscription failed")
	}
	return &subscription, nil
}

func (r *mutationResolver) DeleteSubscription(ctx context.Context, id string) (bool, error) {
	var subscription model.Abonemen
	err := r.DB.Model(&subscription).Where("id=?", id).First()
	if err != nil {
		return false, errors.New("subscription not found")
	}
	_, deleteErr := r.DB.Model(&subscription).Where("id=?", id).Delete()
	if deleteErr != nil {
		return false, errors.New("delete subscription failed")
	}
	return true, nil
}

func (r *queryResolver) Subscriptions(ctx context.Context) ([]*model.Abonemen, error) {
	var subscriptions []*model.Abonemen

	err := r.DB.Model(&subscriptions).Order("user_id").Select()

	if err != nil {
		return nil, errors.New("subscriptions query failed")
	}
	return subscriptions, nil
}

func (r *queryResolver) GetSubscriberCount(ctx context.Context, id string) (int, error) {
	var count int
	var subscription model.Abonemen
	count, err := r.DB.Model(&subscription).Where("user_id=?", id).Count()
	if err != nil {
		return 0, errors.New("subscriber not found")
	}
	return count, nil
}

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"BackEnd/graph/generated"
	"BackEnd/models"
	"context"
	"errors"
	"fmt"
)

func (r *abonemenResolver) Subscriber(ctx context.Context, obj *models.Abonemen) (*models.User, error) {
	var users models.User
	err := r.DB.Model(&users).Where("id=?", obj.SubscriberID).Select()
	if err != nil {
		return nil, errors.New("failed query user from subscription")
	}
	return &users, nil
}

func (r *mutationResolver) CreateSubscription(ctx context.Context, input *models.NewSubscription) (*models.Abonemen, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSubscription(ctx context.Context, id string, input *models.NewSubscription) (*models.Abonemen, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteSubscription(ctx context.Context, id string) (bool, error) {
	var subscription models.Abonemen
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

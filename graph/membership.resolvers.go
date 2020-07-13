package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"BackEnd/graph/model"
	"context"
	"errors"
	"time"
)

func (r *mutationResolver) CreateMembershipDetail(ctx context.Context, input *model.NewMembershipDetail) (*model.MembershipDetail, error) {
	membershipDetail := model.MembershipDetail{
		UserID:       input.UserID,
		Bill:         input.Bill,
		MembershipID: input.MembershipID,
		Date:         time.Now().Format("2006-01-02 15:04:05"),
	}

	_, err := r.DB.Model(&membershipDetail).Insert()

	if err != nil {
		return nil, errors.New("insert membership detail failed")
	}
	return &membershipDetail, nil
}

func (r *mutationResolver) UpdateMembershipDetail(ctx context.Context, userID string, input *model.NewMembershipDetail) (*model.MembershipDetail, error) {
	var membershipDetail model.MembershipDetail

	err := r.DB.Model(&membershipDetail).Where("userId=?", userID).First()
	if err != nil {
		return nil, errors.New("membership detail not found")
	}

	membershipDetail.Bill += input.Bill
	_, updateErr := r.DB.Model(&membershipDetail).Where("userId=?", userID).Update()

	if updateErr != nil {
		return nil, errors.New("update membership detail failed")
	}
	return &membershipDetail, nil
}

func (r *mutationResolver) DeleteMembershipDetail(ctx context.Context, userID string) (bool, error) {
	var membershipDetail model.MembershipDetail

	err := r.DB.Model(&membershipDetail).Where("userId=?", userID).First()
	if err != nil {
		return false, errors.New("membership detail not found")
	}
	_, deleteErr := r.DB.Model(&membershipDetail).Where("userId=?", userID).Delete()

	if deleteErr != nil {
		return false, errors.New("delete membership detail failed")
	}

	return true, nil
}

func (r *mutationResolver) CreateMembership(ctx context.Context, input *model.NewMembership) (*model.Membership, error) {
	membership := model.Membership{
		Type: input.Type,
	}

	_, err := r.DB.Model(&membership).Insert()

	if err != nil {
		return nil, errors.New("insert membership failed")
	}

	return &membership, nil
}

func (r *mutationResolver) UpdateMembership(ctx context.Context, id string, input *model.NewMembership) (*model.Membership, error) {
	var membership model.Membership

	err := r.DB.Model(&membership).Where("id=?", id).First()
	if err != nil {
		return nil, errors.New("membership not found")
	}

	membership.Type = input.Type

	_, updateErr := r.DB.Model(&membership).Where("id=?", id).Update()

	if updateErr != nil {
		return nil, errors.New("update membership failed")
	}
	return &membership, nil
}

func (r *mutationResolver) DeleteMembership(ctx context.Context, id string) (bool, error) {
	var membership model.Membership

	err := r.DB.Model(&membership).Where("id=?", id).First()

	if err != nil {
		return false, errors.New("membership not found")
	}

	_, deleteErr := r.DB.Model(&membership).Where("id=?", id).Delete()
	if deleteErr != nil {
		return false, errors.New("delete membership failed")
	}

	return true, nil
}

func (r *queryResolver) MembershipDetails(ctx context.Context) ([]*model.MembershipDetail, error) {
	var membershipDetails []*model.MembershipDetail

	err := r.DB.Model(&membershipDetails).Order("bill").Select()

	if err != nil {
		return nil, errors.New("membership detail query failed")
	}

	return membershipDetails, nil
}

func (r *queryResolver) Memberships(ctx context.Context) ([]*model.Membership, error) {
	var memberships []*model.Membership
	err := r.DB.Model(&memberships).Order("id").Select()

	if err != nil {
		return nil, errors.New("memberships query failed")
	}

	return memberships, nil
}

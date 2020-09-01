package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"BackEnd/graph/generated"
	"BackEnd/middleware"
	"BackEnd/models"
	"context"
	"errors"
	"time"
)

func (r *membershipDetailResolver) Membership(ctx context.Context, obj *models.MembershipDetail) (*models.Membership, error) {
	var membership models.Membership
	err := r.DB.Model(&membership).Where("id=?", obj.MembershipID).Select()
	if err != nil {
		return nil, errors.New("failed get membership from detail")
	}
	return &membership, nil
}

func (r *mutationResolver) CreateMembershipDetail(ctx context.Context, input *models.NewMembershipDetail) (*models.MembershipDetail, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("unauthenticated")
	}
	var membershipDetail models.MembershipDetail

	err = r.DB.Model(&membershipDetail).Where("user_id = ?", currentUser.ID).Select()

	if err != nil {
		newMembershipDetail := models.MembershipDetail{
			Bill:         input.Bill,
			MembershipID: *input.MembershipID,
			Date:         time.Now().Format("2006-01-02 15:04:05"),
			UserID:       currentUser.ID,
		}

		_, err = r.DB.Model(&newMembershipDetail).Insert()

		if err != nil {
			return nil, errors.New("insert membership detail failed")
		}
		return &newMembershipDetail, nil
	}
	addDate, parseErr := time.Parse("2006-01-02 15:04:05", membershipDetail.Date)
	if parseErr != nil {
		return nil, errors.New(addDate.String())
	}
	addDate = addDate.AddDate(0, 1, 0)
	membershipDetail.Date = addDate.Format("2006-01-02 15:04:05")
	_, updateErr := r.DB.Model(&membershipDetail).Where("user_id=?", currentUser.ID).Update()
	if updateErr != nil {
		return nil, errors.New("update membership error")
	}
	return &membershipDetail, nil
}

func (r *mutationResolver) UpdateMembershipDetail(ctx context.Context, userID string, input *models.NewMembershipDetail) (*models.MembershipDetail, error) {
	var membershipDetail models.MembershipDetail

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
	var membershipDetail models.MembershipDetail

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

func (r *mutationResolver) CreateMembership(ctx context.Context, input *models.NewMembership) (*models.Membership, error) {
	membership := models.Membership{
		Type: input.Type,
	}

	_, err := r.DB.Model(&membership).Insert()

	if err != nil {
		return nil, errors.New("insert membership failed")
	}

	return &membership, nil
}

func (r *mutationResolver) UpdateMembership(ctx context.Context, id string, input *models.NewMembership) (*models.Membership, error) {
	var membership models.Membership

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
	var membership models.Membership

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

func (r *queryResolver) MembershipDetails(ctx context.Context) ([]*models.MembershipDetail, error) {
	var membershipDetails []*models.MembershipDetail

	err := r.DB.Model(&membershipDetails).Order("bill").Select()

	if err != nil {
		return nil, errors.New("membership detail query failed")
	}

	return membershipDetails, nil
}

func (r *queryResolver) Memberships(ctx context.Context) ([]*models.Membership, error) {
	var memberships []*models.Membership
	err := r.DB.Model(&memberships).Order("id").Select()

	if err != nil {
		return nil, errors.New("memberships query failed")
	}

	return memberships, nil
}

// MembershipDetail returns generated.MembershipDetailResolver implementation.
func (r *Resolver) MembershipDetail() generated.MembershipDetailResolver {
	return &membershipDetailResolver{r}
}

type membershipDetailResolver struct{ *Resolver }

package rulesuc

import (
	"github.com/ridwanakf/autobase-twitter/internal"
	"github.com/ridwanakf/autobase-twitter/internal/entity"
)

type RulesUsecase struct {
	gw internal.AutobaseGW
}

func NewRulesUsecase(gw internal.AutobaseGW) *RulesUsecase {
	return &RulesUsecase{gw: gw}
}

func (r *RulesUsecase) ResolveRules(target entity.RulesParam) (bool, error) {
	isFollower, err := r.IsFollower(target.UserID)
	if err != nil {
		return false, err
	}

	isFollowing, err:= r.IsFollowing(target.UserID)
	if err != nil{
		return false, err
	}

	followersCount, err := r.FollowersCount(target.UserID)
	if err != nil{
		return false, err
	}

	return (isFollower || !target.MustBeFollower) &&
		(isFollowing || !target.MustBeFollowing) &&
		(followersCount >= target.MinFollowers), nil
}

func (r *RulesUsecase) IsFollower(userID string) (bool, error) {
	panic("implement me")
}

func (r *RulesUsecase) IsFollowing(userID string) (bool, error) {
	panic("implement me")
}

func (r *RulesUsecase) FollowersCount(userID string) (int32, error) {
	panic("implement me")
}

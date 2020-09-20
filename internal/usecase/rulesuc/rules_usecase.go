package rulesuc

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/ridwanakf/autobase-twitter/internal"
	"github.com/ridwanakf/autobase-twitter/internal/entity"
)

type RulesUsecase struct {
	gw internal.AutobaseGW
}

func NewRulesUsecase(gw internal.AutobaseGW) *RulesUsecase {
	return &RulesUsecase{gw: gw}
}

func (r *RulesUsecase) ResolveRules(param entity.RulesParam) (bool, error) {
	relationship, err := r.gw.UsersRelationship(param.UserID, param.TargetID)
	if err != nil {
		return false, err
	}

	isFollower := r.IsFollower(relationship)
	isFollowing := r.IsFollowing(relationship)
	followersCount, err := r.FollowersCount(param.UserID)
	if err != nil {
		return false, err
	}

	return (isFollower || !param.MustBeFollower) &&
		(isFollowing || !param.MustBeFollowing) &&
		(followersCount >= param.MinFollowers), nil
}

// Is Target Follows User
func (r *RulesUsecase) IsFollower(relation twitter.Relationship) bool {
	return relation.Target.Following
}

// Is User Follows Target
func (r *RulesUsecase) IsFollowing(relation twitter.Relationship) bool {
	return relation.Source.Following
}

func (r *RulesUsecase) FollowersCount(userID string) (int32, error) {
	user, err := r.gw.GetUserInfoByID(userID)
	if err != nil{
		return 0, err
	}

	return int32(user.FollowersCount), nil
}

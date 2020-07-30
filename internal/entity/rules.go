package entity

type RulesParam struct {
	UserID       string
	MustBeFollower   bool
	MustBeFollowing  bool
	MinFollowers int32
}

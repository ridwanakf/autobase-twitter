package entity

type RulesParam struct {
	UserID          string
	TargetID        string
	MustBeFollower  bool
	MustBeFollowing bool
	MinFollowers    int32
}

package config

type Rules struct {
	MustFollower  bool  `yaml:"must_follower" env:"RULES_MUST_FOLLOWER"`
	MustFollowing bool  `yaml:"must_following" env:"RULES_MUST_FOLLOWING"`
	MinFollowers  int32 `yaml:"min_followers" env:"RULES_MIN_FOLLOWERS"`
}

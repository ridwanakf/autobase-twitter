package config

type Params struct {
	Keyword         string          `yaml:"keyword" env:"DM_KEYWORD"`
	MessageCount    int             `yaml:"message_count" env:"MESSAGE_COUNT"`
	MessageResponse MessageResponse `yaml:"message_response"`
	DelayDuration   DelayDuration   `yaml:"delay_duration"`
}

type MessageResponse struct {
	Success   string `yaml:"success" env:"MESSAGE_RESPONSE_SUCCESS"`
	Failed    string `yaml:"failed" env:"MESSAGE_RESPONSE_FAILED"`
	Incorrect string `yaml:"incorrect" env:"MESSAGE_RESPONSE_INCORRECT"`
}

type DelayDuration struct {
	Interval  int `yaml:"interval" env:"DELAY_DURATION_INTERVAL"`
	Sleep     int `yaml:"sleep" env:"DELAY_DURATION_SLEEP"`
	RateLimit int `yaml:"rate_limit" env:"DELAY_DURATION_RATELIMIT"`
}

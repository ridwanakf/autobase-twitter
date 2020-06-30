package config

type TwitterKey struct {
	ConsumerKey       string `yaml:"consumer_key" env:"TWITTER_CONSUMER_KEY"`
	ConsumerSecret    string `yaml:"consumer_secret" env:"TWITTER_CONSUMER_SECRET"`
	AccessToken       string `yaml:"access_token" env:"TWITTER_ACCESS_TOKEN"`
	AccessTokenSecret string `yaml:"access_token_secret" env:"TWITTER_ACCESS_TOKEN_SECRET"`
}

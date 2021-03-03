package config

type Config struct {
	Server     Server     `yaml:"server"`
	Params     Params     `yaml:"params"`
	Twitter    TwitterKey `yaml:"twitter_key"`
	UseArchive bool       `yaml:"archive" env:"USE_ARCHIVE"`
	Databases  Database   `yaml:"database"`
	Rules      Rules      `yaml:"rules"`
}

package config

type Config struct {
	Server  Server     `yaml:"server"`
	Params  Params     `yaml:"params"`
	Twitter TwitterKey `yaml:"twitter_key"`
}

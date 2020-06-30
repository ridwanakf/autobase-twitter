package config

type Server struct {
	Port                string `yaml:"port" env:"PORT"`
	Debug               bool   `yaml:"debug" env:"DEBUG"`
	ReadTimeoutSeconds  int    `yaml:"read_timeout_seconds" env:"SERVER_READ_TIMEOUT"`
	WriteTimeoutSeconds int    `yaml:"write_timeout_seconds" env:"SERVER_READ_TIMEOUT"`
}

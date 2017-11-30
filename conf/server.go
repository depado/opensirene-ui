package conf

import "fmt"

// Server is the structure that holds the server configuration
type Server struct {
	Host  string `yaml:"host" env:"SERVER_HOST" default:"127.0.0.1"`
	Port  int    `yaml:"port" env:"SERVER_PORT" default:"8080"`
	Debug bool   `yaml:"debug" env:"SERVER_DEBUG"`
}

// ListenString returns the listen string most routers expect to run
func (s Server) ListenString() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

package configs

import "fmt"

type databaseConfig struct {
	Host            string
	Port            int
	Name            string
	Protocol        string
	User            string
	Password        string
	SslMode         string `mapstructure:"ssl-mode"`
	MaxConntections int    `mapstructure:"max-connections"`
}

func (c databaseConfig) GetDatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", c.User, c.Password, c.Host, c.Port, c.Name)
}

func (c databaseConfig) GetMaxOpenConnections() int {
	return c.MaxConntections
}

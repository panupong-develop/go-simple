package configs

import "time"

type JWTConfig struct {
	Secretkey     string        `mapstructure:"secret-key"`
	AccessExpiry  time.Duration `mapstructure:"access-expiry"`
	RefreshExpiry time.Duration `mapstructure:"refresh-expiry"`
}

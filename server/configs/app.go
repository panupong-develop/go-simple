package configs

import (
	"fmt"
	"strconv"
	"time"
)

type appConfig struct {
	Host           string
	Port           uint16
	Name           string
	Environment    string
	Version        string
	ReadTimeout    time.Duration `mapstructure:"read-timeout"`
	WriteTimeout   time.Duration `mapstructure:"write-timeout"`
	ApiKey         string        `mapstructure:"api-key"`
	AdminKey       string        `mapstructure:"admin-key"`
	BodyLimitBytes int           `mapstructure:"body-limit-bytes"`
	FileLimitBytes int           `mapstructure:"file-limit-bytes"`
	GcpUrl         string        `mapstructure:"gcp-url"`
	FileLogPath    string        `mapstructure:"file-log-path"`
}

func (c appConfig) GetWebAppAddress() string {
	port := strconv.FormatUint(uint64(c.Port), 10)
	return fmt.Sprintf("%s:%s", c.Host, port)
}

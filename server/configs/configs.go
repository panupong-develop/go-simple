package configs

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
)

// singleton
type configInstance struct {
	config *Config
	mutex  sync.Mutex
}

var instance configInstance

// end

type Config struct {
	App      appConfig
	Database databaseConfig
}

func GetConfig() *Config {
	// Prevent race condition of the singleton object
	instance.mutex.Lock()
	defer instance.mutex.Unlock()

	if instance.config == nil {
		log.Fatalln("you must call LoadConfig() first")
	}

	return instance.config
}

func LoadConfig() error {
	// Prevent race condition of the singleton object
	instance.mutex.Lock()
	defer instance.mutex.Unlock()

	// Read from path
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("/workspaces/go-simple/server")

	// Read from environment variables
	viper.AutomaticEnv()

	// Read in config
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	err = viper.Unmarshal(&instance.config)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Load in middleware configs
	// ...

	// print settings
	b, err := json.MarshalIndent(instance.config, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	return nil
}

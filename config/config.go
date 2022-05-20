package leafConfig

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func New(path string, configType string, object interface{}) error {
	// - check file does exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("config file %s does not exists\nerror: %+v", path, err)
	}

	dir := getDirectory(path)
	file, err := getFile(path)

	if err != nil {
		return err
	}

	v := viper.New()
	v.SetConfigName(file)
	v.AddConfigPath(dir)
	v.SetConfigType(configType)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read %s file\nerror: %+v", path, err)
	}

	if err := v.Unmarshal(&object); err != nil {
		return fmt.Errorf("failed to unmarshal config to object\nerror: %+v", err)
	}

	return nil
}

func NewFromEnv(object interface{}) error {
	filename := os.Getenv("CONFIG_FILE")

	if filename == "" {
		filename = ".env"
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", object); err != nil {
			return fmt.Errorf("%s failed to read from env variable", err)
		}
		return nil
	}

	if err := godotenv.Load(filename); err != nil {
		return fmt.Errorf("%s failed to read from .env file", err)
	}

	if err := envconfig.Process("", object); err != nil {
		return fmt.Errorf("%s failed to read from env variable", err)
	}

	return nil
}

func getDirectory(path string) string {
	splits := strings.Split(path, "/")
	return strings.Join(splits[:len(splits)-1], "/")
}

func getFile(path string) (string, error) {
	splits := strings.Split(path, "/")
	last := splits[len(splits)-1]

	return last, nil
}

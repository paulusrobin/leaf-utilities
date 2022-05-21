package leafViper

import (
	"bytes"
	"context"
	"fmt"
	leafFeatureFlag "github.com/paulusrobin/leaf-utilities/featureFlag/featureFlag"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	leafWebClient "github.com/paulusrobin/leaf-utilities/webClient/webClient"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type (
	featureFlag struct {
		viper     viper.Viper
		webClient leafWebClient.Factory
		option    option
		logger    leafLogger.Logger
	}
)

func (f featureFlag) GetKeys() []string {
	return f.viper.AllKeys()
}

func (f featureFlag) GetSettings() map[string]interface{} {
	return f.viper.AllSettings()
}

func (f featureFlag) Get(key string) interface{} {
	return f.viper.Get(key)
}

func New(v viper.Viper, webClient leafWebClient.Factory, logger leafLogger.Logger, opts ...Option) (leafFeatureFlag.FeatureFlag, error) {
	o := defaultOption()
	for _, opt := range opts {
		opt.Apply(&o)
	}

	ff := &featureFlag{
		viper:     v,
		webClient: webClient,
		option:    o,
		logger:    logger,
	}

	var err error
	switch ff.option.featureFlagConfigurationSource {
	case "", "file":
		err = ff.handleFileConfigSource()
		break
	case "custom-http":
		err = ff.handleCustomConfigSource()
		break
	case "etcd", "consul":
		err = ff.handleRemoteConfigSource()
		break
	default:
		err = fmt.Errorf("unsupported feature flag source, only [file, custom-http, etcd, consul] is supported")
		break
	}

	if err != nil {
		return nil, err
	}
	return ff, nil
}

func (f *featureFlag) handleFileConfigSource() error {
	configPath := f.option.featureFlagConfigurationPath
	configType := f.option.featureFlagConfigurationType
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return fmt.Errorf("config file %s does not exists\nerror: %+v", configPath, err)
	}

	dir := getDirectory(configPath)
	file, err := getFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to get %s file\nerror: %+v", file, err)
	}

	f.viper.SetConfigName(file)
	f.viper.AddConfigPath(dir)
	f.viper.SetConfigType(configType)

	if err := f.viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read %s file\nerror: %+v", configPath, err)
	}

	f.viper.WatchConfig()

	return nil
}

func (f *featureFlag) handleCustomConfigSource() error {
	periodicUpdateInterval := f.option.periodicUpdate.interval

	if f.option.featureFlagServer == "" || f.option.featureFlagConfigurationPath == "" {
		return fmt.Errorf("feature flag server and path is required for custom-http configuration source")
	}

	f.viper.SetConfigType(f.option.featureFlagConfigurationType)
	if err := f.retrieveCustomConfig(); err != nil {
		return fmt.Errorf("failed to read custom config\nerror: %+v", err)
	}

	if periodicUpdateInterval > 0 {
		go func() {
			for {
				select {
				case <-time.Tick(periodicUpdateInterval):
					if err := f.retrieveCustomConfig(); err != nil {
						f.logger.Error(leafLogger.BuildMessage(context.Background(), "failed to reload custom remote config",
							leafLogger.WithAttr("err", err.Error())))
					}
				}
			}
		}()
	}

	return nil
}

func (f *featureFlag) retrieveCustomConfig() error {
	webClient := f.webClient.Create(leafWebClient.NewDefaultWebClientOption(f.option.featureFlagTimeout))
	header := http.Header{
		"Content-type": []string{"application/json"},
	}
	response, err := webClient.Get(context.Background(), f.option.featureFlagServer+f.option.featureFlagConfigurationPath, header, nil)
	if err != nil {
		return fmt.Errorf("failed to get feature flag response\nerror: %+v", err)
	}
	bodyBytes, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return fmt.Errorf("failed to read feature flag response body\nerror: %+v", err)
	}

	if err := f.viper.ReadConfig(bytes.NewReader(bodyBytes)); err != nil {
		return fmt.Errorf("failed to set feature flag to runtime\nerror: %+v", err)
	}

	return nil
}

func (f *featureFlag) handleRemoteConfigSource() error {
	configServer := f.option.featureFlagServer
	configPath := f.option.featureFlagConfigurationPath
	configType := f.option.featureFlagConfigurationType
	configSource := f.option.featureFlagConfigurationSource
	periodicUpdateInterval := f.option.periodicUpdate.interval
	if configServer == "" || configPath == "" {
		return fmt.Errorf("feature flag server and path is required for remote configuration source")
	}

	if err := f.viper.AddRemoteProvider(configSource, configServer, configPath); err != nil {
		if err := f.viper.ReadRemoteConfig(); err != nil {
			return fmt.Errorf("failed to add remote provider\nerror: %+v", err)
		}
	}

	f.viper.SetConfigType(configType)
	if err := f.viper.ReadRemoteConfig(); err != nil {
		return fmt.Errorf("failed to add remote provider\nerror: %+v", err)
	}

	if periodicUpdateInterval > 0 {
		go func() {
			for {
				select {
				case <-time.Tick(periodicUpdateInterval):
					err := f.viper.WatchRemoteConfig()
					if err != nil {
						f.logger.Error(leafLogger.BuildMessage(context.Background(), "failed to reload remote config",
							leafLogger.WithAttr("err", err.Error())))
						continue
					}
				}
			}
		}()
	}

	return nil
}

package leafViper

import (
	leafConfigType "github.com/paulusrobin/leaf-utilities/common/constants/configType"
	"strings"
	"time"
)

type (
	ValidatorFunc func(data map[string]interface{}) error
	Option        interface {
		Apply(o *option)
	}

	periodicUpdate struct {
		interval time.Duration
	}
	option struct {
		featureFlagServer              string
		featureFlagConfigurationPath   string
		featureFlagConfigurationType   string
		featureFlagConfigurationSource string
		featureFlagTimeout             time.Duration
		fnValidate                     ValidatorFunc
		periodicUpdate                 periodicUpdate
	}
)

func defaultOption() option {
	return option{
		featureFlagServer:              "",
		featureFlagConfigurationPath:   "./feature-flag.json",
		featureFlagConfigurationType:   leafConfigType.JSON,
		featureFlagConfigurationSource: "",
		featureFlagTimeout:             time.Second,
		fnValidate:                     nil,
		periodicUpdate: periodicUpdate{
			interval: 0,
		},
	}
}

type withFeatureFlagServer string

func (w withFeatureFlagServer) Apply(o *option) {
	o.featureFlagServer = string(w)
}

func WithFeatureFlagServer(server string) Option {
	return withFeatureFlagServer(server)
}

type withFeatureFlagConfigurationPath string

func (w withFeatureFlagConfigurationPath) Apply(o *option) {
	o.featureFlagConfigurationPath = string(w)
}

func WithFeatureFlagConfigurationPath(path string) Option {
	return withFeatureFlagConfigurationPath(path)
}

type withFeatureFlagConfigurationType string

func (w withFeatureFlagConfigurationType) Apply(o *option) {
	o.featureFlagConfigurationType = strings.ToLower(string(w))
}

func WithFeatureFlagConfigurationType(fileType string) Option {
	return withFeatureFlagConfigurationType(fileType)
}

type withFeatureFlagConfigurationSource string

func (w withFeatureFlagConfigurationSource) Apply(o *option) {
	o.featureFlagConfigurationSource = string(w)
}

func WithFeatureFlagConfigurationSource(source string) Option {
	return withFeatureFlagConfigurationSource(source)
}

type withFeatureFlagTimeout time.Duration

func (w withFeatureFlagTimeout) Apply(o *option) {
	o.featureFlagTimeout = time.Duration(w)
}

func WithFeatureFlagTimeout(timeout time.Duration) Option {
	return withFeatureFlagTimeout(timeout)
}

type withValidator ValidatorFunc

func (w withValidator) Apply(o *option) {
	o.fnValidate = ValidatorFunc(w)
}

func WithValidator(fn ValidatorFunc) Option {
	return withValidator(fn)
}

type withPeriodicUpdateInterval time.Duration

func (w withPeriodicUpdateInterval) Apply(o *option) {
	o.periodicUpdate.interval = time.Duration(w)
}

func WithPeriodicUpdateInterval(interval time.Duration) Option {
	return withPeriodicUpdateInterval(interval)
}

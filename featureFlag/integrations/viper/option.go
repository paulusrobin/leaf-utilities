package leafViper

import (
	leafConfigType "github.com/paulusrobin/leaf-utilities/common/constants/configType"
	leafFeatureFlag "github.com/paulusrobin/leaf-utilities/featureFlag/featureFlag"
	"strings"
	"time"
)

type (
	Option interface {
		Apply(o *option)
	}
	periodicUpdate struct {
		interval time.Duration
	}
	BackupServer struct {
		Backup leafFeatureFlag.Backup
	}
	option struct {
		featureFlagServer              string
		featureFlagConfigurationPath   string
		featureFlagConfigurationType   string
		featureFlagConfigurationSource string
		featureFlagTimeout             time.Duration
		periodicUpdate                 periodicUpdate
		backupServer                   BackupServer
	}
)

func defaultOption() option {
	return option{
		featureFlagServer:              "",
		featureFlagConfigurationPath:   "./feature-flag.json",
		featureFlagConfigurationType:   leafConfigType.JSON,
		featureFlagConfigurationSource: "",
		featureFlagTimeout:             time.Second,
		periodicUpdate: periodicUpdate{
			interval: 0,
		},
		backupServer: BackupServer{
			Backup: nil,
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

type withPeriodicUpdateInterval time.Duration

func (w withPeriodicUpdateInterval) Apply(o *option) {
	o.periodicUpdate.interval = time.Duration(w)
}

func WithPeriodicUpdateInterval(interval time.Duration) Option {
	return withPeriodicUpdateInterval(interval)
}

type withBackupServer BackupServer

func (w withBackupServer) Apply(o *option) {
	o.backupServer = BackupServer(w)
}

func WithBackupServer(backupServer BackupServer) Option {
	return withBackupServer(backupServer)
}

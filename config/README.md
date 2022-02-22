# Config

## Config Getter
Config Getter support from .env or environment variable

## Usage
Import the package
```go
import (
    ...
    "github.com/paulusrobin/leaf-utilities/config"
    ...
)
```

Create config struct
```go
type (
    AppConfig struct {
        // - Interface Setting
        HttpEnable      bool `envconfig:"INTERFACE_HTTP_ENABLE" default:"true"`
        MessagingEnable bool `envconfig:"INTERFACE_MESSAGING_ENABLE" default:"true"`
        WorkerEnable    bool `envconfig:"INTERFACE_WORKER_ENABLE" default:"true"`
    }
)
```
Retrieve environtment variable to struct
```go
configuration := AppConfig{}
if err := leafConfig.NewFromEnv(&configuration); err != nil {
	return AppConfig{}, err
}
return configuration, nil
```
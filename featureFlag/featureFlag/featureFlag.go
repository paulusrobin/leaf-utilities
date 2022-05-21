package leafFeatureFlag

type (
	FeatureFlag interface {
		Get(key string) interface{}
		GetKeys() []string
		GetSettings() map[string]interface{}
	}
)

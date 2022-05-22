package leafFeatureFlag

type (
	Backup interface {
		Get(key string) map[string]interface{}
		Set(key string, data map[string]interface{}) error
	}
)

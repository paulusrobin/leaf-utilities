package leafFeatureFlag

type (
	Backup interface {
		Get(key string) (map[string]interface{}, error)
		Set(key string, data map[string]interface{}) error
	}
)

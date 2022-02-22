package leafMandatory

import "github.com/paulusrobin/leaf-utilities/encoding/json"

type Authorization struct {
	authorization string
	token         string
	apiKey        string
	serviceID     string
	serviceSecret string
}

func (a Authorization) Authorization() string {
	return a.authorization
}

func (a Authorization) Token() string {
	return a.token
}

func (a Authorization) ApiKey() string {
	return a.apiKey
}

func (a Authorization) ServiceID() string {
	return a.serviceID
}

func (a Authorization) ServiceSecret() string {
	return a.serviceSecret
}

func (a Authorization) JSON() map[string]interface{} {
	return map[string]interface{}{
		"token":      "***",
		"api_key":    a.ApiKey(),
		"service_id": a.ServiceID(),
	}
}

func (a Authorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.JSON())
}

package model

import (
	"fmt"
	"github.com/paulusrobin/leaf-example/pkg/constant/privilege"
	"strings"
)

type (
	HealthEndpoints []HealthEndpoint
	HealthEndpoint  struct {
		Method string `json:"method"`
		Path   string `json:"path"`
		Verify string `json:"name"`
	}
	HealthRoutesResponse struct {
		Routes []string `json:"routes"`
	}
)

func (e HealthEndpoint) String() string {
	if e.IsDocumentation() {
		e.Verify = privilege.Granted
	}
	return fmt.Sprintf("[%s] %s %s", e.Method, e.Path, e.Verify)
}

func (e HealthEndpoint) IsDocumentation() bool {
	return strings.Contains(e.Path, "/docs")
}

func (e HealthEndpoints) String() []string {
	var eps = make([]string, 0)
	for _, ep := range e {
		if !privilege.Exist(ep.Verify) {
			continue
		}
		eps = append(eps, ep.String())
	}
	return eps
}

func (e HealthEndpoints) Verify() HealthEndpoints {
	var eps = make(HealthEndpoints, 0)
	for _, ep := range e {
		if !privilege.Exist(ep.Verify) {
			continue
		}
		eps = append(eps, ep)
	}

	return eps
}

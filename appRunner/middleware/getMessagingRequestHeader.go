package leafMiddleware

import leafMQ "github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue"

func GetMessagingRequestHeader(message leafMQ.Message, headers ...string) string {
	var value = ""
	for _, header := range headers {
		value = message.Attributes[header]
		if value != "" {
			break
		}
	}
	return value
}

func GetMessagingRequestHeaderWithDefault(message leafMQ.Message, fn func() string, headers ...string) string {
	var value = ""
	for _, header := range headers {
		value = message.Attributes[header]
		if value != "" {
			break
		}
	}
	if value == "" {
		return fn()
	}
	return value
}

package leafLogger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/paulusrobin/leaf-utilities/mandatory"
	"time"
)

type Message map[string]interface{}

func BuildMessage(ctx context.Context, msg string, opts ...MessageOption) Message {
	o := defaultOption()
	for _, opt := range opts {
		opt.Apply(&o)
	}

	var message = map[string]interface{}{
		"timestamp":  time.Now().Format(time.RFC3339),
		"attributes": nil,
	}

	mandatory := leafMandatory.FromContext(ctx)
	messageBody := fmt.Sprintf("%s", msg)
	if mandatory.Valid() {
		message["mandatory"] = mandatory.JSON()
		messageBody = fmt.Sprintf("[%s] %s", mandatory.TraceID(), msg)
	}

	message["message"] = messageBody
	if len(o.json) > 0 {
		var attributes = make(map[string]interface{})
		for key, value := range o.json {
			attributes[key] = value
		}
		message["attributes"] = attributes
	}

	if len(o.masking) > 0 {
		for key, msg := range message {
			message[key] = o.masking.Encode(key, msg)
		}
	}

	return message
}

func (msg Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(msg)
}

func (msg Message) MarshalText() ([]byte, error) {
	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("%v", msg))
	return buffer.Bytes(), nil
}

func (msg Message) String() string {
	return msg["message"].(string)
}

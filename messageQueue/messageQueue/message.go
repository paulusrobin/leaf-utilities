package leafMQ

import "github.com/paulusrobin/leaf-utilities/encoding/json"

type (
	Message struct {
		id         string            `json:"id"`
		Ordering   string            `json:"ordering"`
		Data       []byte            `json:"data"`
		Attributes map[string]string `json:"attributes"`
	}
	publicMessage struct {
		ID         string            `json:"id"`
		Ordering   string            `json:"ordering"`
		Data       []byte            `json:"data"`
		Attributes map[string]string `json:"attributes"`
	}
)

func (m Message) GetID() string {
	return m.id
}

func (m *Message) SetID(ID string) {
	m.id = ID
}

func (m Message) publicMessage() publicMessage {
	return publicMessage{
		ID:         m.id,
		Ordering:   m.Ordering,
		Data:       m.Data,
		Attributes: m.Attributes,
	}
}

func (m Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.publicMessage())
}

func (m *Message) UnmarshalJSON(data []byte) error {
	var publicMessage publicMessage
	if err := json.Unmarshal(data, &publicMessage); err != nil {
		return err
	}
	*m = Message{
		id:         publicMessage.ID,
		Ordering:   publicMessage.Ordering,
		Data:       publicMessage.Data,
		Attributes: publicMessage.Attributes,
	}
	return nil
}

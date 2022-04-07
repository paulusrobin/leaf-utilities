package leafSlack

import (
	"github.com/paulusrobin/leaf-utilities/encoding/json"
)

type (
	Message interface {
		Json() string
	}
	message struct {
		Text   string  `json:"text,omitempty"`
		Blocks []Block `json:"blocks,omitempty"`
	}

	Block struct {
		Type      string     `json:"type,omitempty"`
		BlockID   string     `json:"block_id,omitempty"`
		Text      Text       `json:"text,omitempty"`
		Accessory *Accessory `json:"accessory,omitempty"`
		Fields    []Text     `json:"fields,omitempty"`
		Elements  []Block    `json:"elements,omitempty"`
	}

	Text struct {
		Type string `json:"type,omitempty"`
		Text string `json:"text,omitempty"`
	}

	Accessory struct {
		Type     string `json:"type,omitempty"`
		ImageUrl string `json:"image_url,omitempty"`
		AltText  string `json:"alt_text,omitempty"`
	}
)

func (m message) Json() string {
	jsonByte, _ := json.Marshal(m)
	return string(jsonByte)
}

func NewMessage(options ...MessageOption) (Message, error) {
	o := messageOption{}
	for _, opt := range options {
		opt.Apply(&o)
	}

	if err := o.validate(); err != nil {
		return nil, err
	}

	return &message{
		Text:   o.text,
		Blocks: o.blocks,
	}, nil
}

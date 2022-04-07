package leafSlack

import (
	"fmt"
)

type (
	MessageOption interface {
		Apply(o *messageOption)
	}
	messageOption struct {
		text   string
		blocks []Block
	}
)

type withText string

func (w withText) Apply(o *messageOption) {
	o.text = string(w)
}

func WithText(text string) MessageOption {
	return withText(text)
}

type withBlock Block

func (w withBlock) Apply(o *messageOption) {
	o.blocks = append(o.blocks, Block(w))
}

func WithBlock(block Block) MessageOption {
	return withBlock(block)
}

func (o messageOption) validate() error {
	if o.text == "" && len(o.blocks) == 0 {
		return fmt.Errorf("text or block option must be filled")
	}
	if o.text != "" && len(o.blocks) > 0 {
		return fmt.Errorf("must choose one of text or block option")
	}
	return nil
}

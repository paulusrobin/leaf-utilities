package leafLogger

type (
	messageOptions struct {
		json    map[string]interface{}
		masking MaskedEncoder
	}
	MessageOption interface {
		Apply(o *messageOptions)
	}
)

/*
	======================
		With Attribute
	======================
*/
type withAttribute struct {
	key   string
	value interface{}
}

func (w *withAttribute) Apply(o *messageOptions) {
	if _, found := o.json[w.key]; !found {
		o.json[w.key] = w.value
	}
}

// WithAttr function to add attributes to messages
func WithAttr(key string, value interface{}) MessageOption {
	return &withAttribute{key: key, value: value}
}

/*
	======================
		With Masking
	======================
*/
type withMasking struct {
	key    string
	masked Masked
}

func (w *withMasking) Apply(o *messageOptions) {
	if _, found := o.masking[w.key]; !found {
		o.masking[w.key] = w.masked
	}
}

// WithMasking function to add masking specific key to messages
func WithMasking(key string, masked Masked) MessageOption {
	return &withMasking{key: key, masked: masked}
}

func defaultOption() messageOptions {
	return messageOptions{
		json:    make(map[string]interface{}),
		masking: make(map[string]Masked),
	}
}


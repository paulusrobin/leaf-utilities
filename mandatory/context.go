package leafMandatory

import "context"

const contextKey = `mandatory`

func Context(ctx context.Context, mandatory Mandatory) context.Context {
	ctx = context.WithValue(ctx, contextKey, &mandatory)
	return ctx
}

func FromContext(ctx context.Context) Mandatory {
	if nil == ctx {
		return Mandatory{}
	}

	mandatory, found := ctx.Value(contextKey).(*Mandatory)
	if !found || nil == mandatory {
		return Mandatory{}
	}
	return *mandatory
}

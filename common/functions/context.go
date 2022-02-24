package leafFunctions

import "context"

const skipError = `skip-error`

func DoSkipNoticeError(ctx *context.Context) {
	if nil == ctx {
		return
	}
	*ctx = context.WithValue(*ctx, skipError, true)
}

func DontSkipNoticeError(ctx *context.Context) {
	if nil == ctx {
		return
	}

	skip := (*ctx).Value(skipError)
	if nil == skip {
		return
	}

	*ctx = context.WithValue(*ctx, skipError, false)
}

func SkipNoticeError(ctx context.Context) bool {
	skip := ctx.Value(skipError)
	if nil == skip {
		return false
	}

	if val, ok := skip.(bool); !ok {
		return false
	} else {
		return val
	}
}

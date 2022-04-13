package leafHttpResponse

import (
	"context"
	"github.com/go-playground/validator/v10"
	leafMandatory "github.com/paulusrobin/leaf-utilities/mandatory"
	leafTranslatorValidator "github.com/paulusrobin/leaf-utilities/translator/validator"
	"net/http"
)

type (
	errorResponse struct {
		Error *Error `json:"error,omitempty"`
	}
	Error struct {
		Message string       `json:"message"`
		Code    int          `json:"code"`
		Errors  []ErrorField `json:"errors,omitempty"`
	}
	ErrorField struct {
		Field  string `json:"field"`
		Reason string `json:"reason"`
	}
)

func (e *errorResponse) Val() interface{} {
	return e
}

func newErrorResponse(ctx context.Context, responseCode int, err error) Response {
	var errStruct *Error = nil
	if err != nil {
		eMsg, e := getErrors(ctx, responseCode, err)
		errStruct = &Error{
			Message: eMsg,
			Code:    responseCode,
			Errors:  e,
		}
	}
	return &errorResponse{Error: errStruct}
}

func getErrors(ctx context.Context, responseCode int, err error) (string, []ErrorField) {
	if responseCode != http.StatusBadRequest {
		return err.Error(), nil
	}

	x, ok := err.(validator.ValidationErrors)
	if ok {
		data := buildCustomErrors(ctx, x)
		return data[0].Reason, data
	}
	return err.Error(), nil
}

func buildCustomErrors(ctx context.Context, errs validator.ValidationErrors) []ErrorField {
	mandatory := leafMandatory.FromContext(ctx)
	translator, _ := leafTranslatorValidator.GetTranslator().GetTranslator(mandatory.Language())
	errors := make([]ErrorField, 0)
	for _, err := range errs {
		errors = append(errors, ErrorField{
			Field:  err.Field(),
			Reason: err.Translate(translator),
		})
	}
	return errors
}

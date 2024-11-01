package pkgerr

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/samber/lo"
)

// ValidationError defines a validation specific error structure which conforms with the HTTPError interface
type ValidationError struct {
	namespace      string
	errorCode      string
	httpStatusCode int
	errorBody      map[string]string // useful for multi field validation errors
}

var _ HttpError = (*ValidationError)(nil)

func NewValidationError(
	namespace string,
	errorCode string,
	httpStatusCode int,
	errorBody map[string]string,
) *ValidationError {
	return &ValidationError{
		namespace:      namespace,
		errorCode:      errorCode,
		httpStatusCode: httpStatusCode,
		errorBody:      errorBody,
	}
}

func (e *ValidationError) Error() string {
	es := e.errorBody
	if len(es) == 0 {
		return ""
	}

	keys := make([]string, len(es))
	i := 0
	for key := range es {
		keys[i] = key
		i++
	}
	sort.Strings(keys)

	var s strings.Builder
	for i, key := range keys {
		if i > 0 {
			s.WriteString("; ")
		}
		_, _ = fmt.Fprintf(&s, "%v: (%v)", key, es[key])
	}
	s.WriteString(".")
	return s.String()
}

func (e *ValidationError) HttpStatusCode() int {
	return e.httpStatusCode
}

// ServiceErrorResponseBody defines the response body to send back, this is also used for swagger docs
type ValidationErrorResponseBody struct {
	Namespace string            `json:"namespace,omitempty"`
	Code      string            `json:"code,omitempty"`
	Errors    map[string]string `json:"errors,omitempty"`
}

func (e *ValidationError) ResponseBody() any {
	return ValidationErrorResponseBody{
		Namespace: e.namespace,
		Code:      e.errorCode,
		Errors:    e.errorBody,
	}
}

// WrapValidationError wraps a single field validation error
func WrapValidationError(err error, field string) error {
	if v, ok := err.(validation.Error); ok {
		return NewValidationError("validation", "validation_failed", http.StatusBadRequest, map[string]string{
			field: v.Message(),
		})
	}
	return err
}

// WrapStructValidationError wrapes a multiple field valiadtion error if provided
func WrapStructValidationError(err error) error {
	if v, ok := err.(validation.Errors); ok {
		errMap := lo.MapEntries(v, func(field string, err error) (string, string) {
			return field, err.Error()
		})
		return NewValidationError("validation", "validation_failed", http.StatusBadRequest, errMap)
	}
	return err
}

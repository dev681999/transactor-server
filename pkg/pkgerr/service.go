package pkgerr

// ServiceError defines a generic error structure which conforms with the HTTPError interface
type ServiceError struct {
	namespace      string
	errorCode      string
	httpStatusCode int
	errorBody      string
}

var _ HttpError = (*ServiceError)(nil)

func NewServiceError(
	namespace string,
	errorCode string,
	httpStatusCode int,
	errorBody string,
) *ServiceError {
	return &ServiceError{
		namespace:      namespace,
		errorCode:      errorCode,
		httpStatusCode: httpStatusCode,
		errorBody:      errorBody,
	}
}

func (e *ServiceError) Error() string {
	return e.errorCode
}

func (e *ServiceError) HttpStatusCode() int {
	return e.httpStatusCode
}

// ServiceErrorResponseBody defines the response body to send back, this is also used for swagger docs
type ServiceErrorResponseBody struct {
	Namespace string `json:"namespace,omitempty"`
	Code      string `json:"code,omitempty"`
	Msg       string `json:"msg,omitempty"`
}

func (e *ServiceError) ResponseBody() any {
	return ServiceErrorResponseBody{
		Namespace: e.namespace,
		Code:      e.errorCode,
		Msg:       e.errorBody,
	}
}

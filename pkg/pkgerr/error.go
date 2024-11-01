package pkgerr

// HttpError defines a way for a error handler to send custom errors in response
type HttpError interface {
	// HttpStatusCode defines what status code to send error response with
	HttpStatusCode() int
	// ResponseBody defines what JSON response should be returned
	ResponseBody() any
}

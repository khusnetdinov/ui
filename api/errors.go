package api

type ApiError struct {
	Message string
}

func (e ApiError) Error() string {
	return e.Message
}

func NewApiError(err error) error {
	return &ApiError{
		Message: err.Message,
	}
}

type ApiResponseError struct {
	Code           int
	Message        string
	ResponseParams *ResponseParams
}

func (e ApiResponseError) Error() string {
	return e.Message
}

func NewApiResponseError(response *Response) error {
	return &ApiResponseError{
		Code:           response.ErrorCode,
		Message:        response.Description,
		ResponseParams: response.Parameters,
	}
}

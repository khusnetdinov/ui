package api

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

package api

type ApiResponseError struct {
	Code           int
	Message        string
	ResponseParams *ApiResponseParams
}

func (e ApiResponseError) Error() string {
	return e.Message
}

func NewApiResponseError(response *ApiResponse) error {
	return &ApiResponseError{
		Code:           response.ErrorCode,
		Message:        response.Description,
		ResponseParams: response.Parameters,
	}
}

package api

type Error struct {
	Code           int
	Message        string
	ResponseParams *ResponseParams
}

func (e Error) Error() string {
	return e.Message
}

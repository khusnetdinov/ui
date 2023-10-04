package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const telegramApiUrl = "https://api.telegram.org/bot%s/"
const localApiUrl = "http://localhost:8081/bot%s/"

type Api struct {
	client *http.Client
	Logger *Slog
	url    string
	User   User
}

type ApiResponse struct {
	Ok          bool               `json:"ok"`
	Result      json.RawMessage    `json:"result,omitempty"`
	ErrorCode   int                `json:"error_code,omitempty"`
	Description string             `json:"description,omitempty"`
	Parameters  *ApiResponseParams `json:"parameters,omitempty"`
}

type ApiResponseParams struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int   `json:"retry_after,omitempty"`
}

func NewApi(token string, timeout time.Duration, production bool, debug bool) *Api {
	var url string
	if production {
		url = fmt.Sprintf(telegramApiUrl, token)
	} else {
		url = fmt.Sprintf(localApiUrl, token)
	}

	return &Api{
		Logger: NewSlog(os.Stdout, debug),
		client: &http.Client{Timeout: timeout},
		url:    url,
	}
}

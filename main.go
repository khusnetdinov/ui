package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const token = "6591790550:AAE5s6Mmhs8QsGPDxmTxEB23kvKKg3KrI_w"

// api.go: begin
const telegramApiUrl = "https://api.telegram.org/bot%s/"

const (
	ContentTypeBinary = "application/octet-stream"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeJSON   = "application/json"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
)

const (
	RequestMethodGetMe          = "getMe"          // https://core.telegram.org/bots/api#getme
	RequestMethodGetUpdates     = "getUpdates"     // https://core.telegram.org/bots/api#getupdates
	RequestMethodGetWebHook     = "setWebhook"     // https://core.telegram.org/bots/api#setwebhook
	RequestMethodDeleteWebhook  = "deleteWebhook"  // https://core.telegram.org/bots/api#setwebhook
	RequestMethodGetWebHookInfo = "getWebhookInfo" // https://core.telegram.org/bots/api#getwebhookinfo
)

// const (
// 	RequestMethodSendMessage            = "sendMessage"
// 	RequestMethodForwardMessage         = "forwardMessage"
// 	RequestMethodSendPhoto              = "sendPhoto"
// 	RequestMethodSendAudio              = "sendAudio"
// 	RequestMethodSendDocument           = "sendDocument"
// 	RequestMethodSendSticker            = "sendSticker"
// 	RequestMethodSendVideo              = "sendVideo"
// 	RequestMethodSendVoice              = "sendVoice"
// 	RequestMethodSendVenue              = "sendVenue"
// 	RequestMethodSendContact            = "sendContact"
// 	RequestMethodSendLocation           = "sendLocation"
// 	RequestMethodSendChatAction         = "sendChatAction"
// 	RequestMethodGetUserProfilePhotos   = "getUserProfilePhotos"
// 	RequestMethodGetFile                = "getFile"
// 	RequestMethodKickChatMember         = "kickChatMember"
// 	RequestMethodUnbanChatMember        = "unbanChatMember"
// 	RequestMethodAnswerCallbackQuery    = "answerCallbackQuery"
// 	RequestMethodAnswerInlineQuery      = "answerInlineQuery"
// 	RequestMethodEditMessageText        = "editMessageText"
// 	RequestMethodEditMessageCaption     = "editMessageCaption"
// 	RequestMethodEditMessageReplyMarkup = "editMessageReplyMarkup"
// )

// type Update struct{}        // https://core.telegram.org/bots/api#update
// type WebhookInfo struct{}   // https://core.telegram.org/bots/api#webhookinfo
// type User struct{}          // https://core.telegram.org/bots/api#user
// type Chat struct{}          // https://core.telegram.org/bots/api#chat
// type Message struct{}       // https://core.telegram.org/bots/api#message
// type MessageId struct{}     // https://core.telegram.org/bots/api#messageid
// type MessageEntity struct{} // https://core.telegram.org/bots/api#messageentoty
// type PhotoSize struct{}     // https://core.telegram.org/bots/api#photosize
// type Animation struct{}     // https://core.telegram.org/bots/api#animation
// type Audio struct{}         // https://core.telegram.org/bots/api#audio
// type Document struct{}      // https://core.telegram.org/bots/api#document
// type Story struct{}         // https://core.telegram.org/bots/api#story
// type Video struct{}         // https://core.telegram.org/bots/api#vidoe
// type VideoNote struct{}     // https://core.telegram.org/bots/api#videonote

type Api struct {
	client *http.Client
	token  string
	url    string
}

func NewApi(token string, client *http.Client) *Api {
	return &Api{
		client: client,
		token:  token,
		url:    fmt.Sprintf(telegramApiUrl, token),
	}
}

func (api Api) GetMe() (*http.Response, error) {
	request, err := http.NewRequest(http.MethodPost, api.url+RequestMethodGetMe, nil)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return api.client.Do(request)
}

func (api Api) GetUpdates() (*http.Response, error) {
	return &http.Response{}, nil
}

func (api Api) SetWebHook()     {}
func (api Api) DeleteWebHook()  {}
func (api Api) GetWebHookInfo() {}

// func (api Telegram) SendMessage()          {}
// func (api Telegram) ForwardMessage()       {}
// func (api Telegram) SendPhoto()            {}
// func (api Telegram) SendAudio()            {}
// func (api Telegram) SendDocument()         {}
// func (api Telegram) SendSticker()          {}
// func (api Telegram) SendVideo()            {}
// func (api Telegram) SendVoice()            {}
// func (api Telegram) SendVenue()            {}
// func (api Telegram) SendContact()          {}
// func (api Telegram) SendLocation()         {}
// func (api Telegram) SendChatAction()       {}
// func (api Telegram) GetUserProfilePhotos() {}
// func (api Telegram) GetFile()              {}
// func (api Telegram) AnswerInlineQuery()    {}
// api.go: end

func main() {
	api := NewApi(token, &http.Client{Timeout: time.Duration(1) * time.Second})

	resp, err := api.GetMe()
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", body)
}

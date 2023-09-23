package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	ContentType       = "Content-Type"
	ContentTypeBinary = "application/octet-stream"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeJSON   = "application/json"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
)

const (
	RequestMethodGetUpdates                        = "getUpdates"                        // https://core.telegram.org/bots/api#getupdates
	RequestMethodSetWebHook                        = "setWebhook"                        // https://core.telegram.org/bots/api#setwebhook
	RequestMethodDeleteWebHook                     = "deleteWebhook"                     // https://core.telegram.org/bots/api#deletewebhook
	RequestMethodGetWebHookInfo                    = "getWebhookInfo"                    // https://core.telegram.org/bots/api#getwebhookinfo
	RequestMethodGetMe                             = "getMe"                             // https://core.telegram.org/bots/api#getme
	RequestMethodLogOut                            = "logOut"                            // https://core.telegram.org/bots/api#logout
	RequestMethodClose                             = "close"                             // https://core.telegram.org/bots/api#close
	RequestMethodSendMessage                       = "sendMessage"                       // https://core.telegram.org/bots/api#sendmessage
	// RequestMethodForwardMessage                    = "forwardMessage"                    // https://core.telegram.org/bots/api#forwardmessage
	// RequestMethodCopyMessage                       = "copyMessage"                       // https://core.telegram.org/bots/api#copymessage
	// RequestMethodSendPhoto                         = "sendPhoto"                         // https://core.telegram.org/bots/api#sendphoto
	// RequestMethodSendAudio                         = "sendAudio"                         // https://core.telegram.org/bots/api#sendaudio
	// RequestMethodSendDocument                      = "sendDocument"                      // https://core.telegram.org/bots/api#senddocument
	// RequestMethodSendVideo                         = "sendVideo"                         // https://core.telegram.org/bots/api#sendvideo
	// RequestMethodSendAnimation                     = "sendAnimation"                     // https://core.telegram.org/bots/api#sendanimation
	// RequestMethodSendVoice                         = "sendVoice"                         // https://core.telegram.org/bots/api#sendvoice
	// RequestMethodSendVideoNote                     = "sendVideoNote"                     // https://core.telegram.org/bots/api#sendvideonote
	// RequestMethodSendMediaGroup                    = "sendMediaGroup"                    // https://core.telegram.org/bots/api#sendmediagroup
	// RequestMethodSendLocation                      = "sendLocation"                      // https://core.telegram.org/bots/api#sendlocation
	// RequestMethodSendVenue                         = "sendVenue"                         // https://core.telegram.org/bots/api#sendvenue
	// RequestMethodSendContact                       = "sendContact"                       // https://core.telegram.org/bots/api#sendcontact
	// RequestMethodSendPoll                          = "sendPoll"                          // https://core.telegram.org/bots/api#sendpoll
	// RequestMethodSendDice                          = "sendDice"                          // https://core.telegram.org/bots/api#senddice
	// RequestMethodSendChatAction                    = "sendChatAction"                    // https://core.telegram.org/bots/api#sendchataction
	// RequestMethodGetUserProfilePhotos              = "getUserProfilePhotos"              // https://core.telegram.org/bots/api#getuserprofilephotos
	// RequestMethodGetFile                           = "getFile"                           // https://core.telegram.org/bots/api#getfile
	// RequestMethodBanChatMember                     = "banChatMember"                     // https://core.telegram.org/bots/api#banchatmember
	// RequestMethodUnbanChatMember                   = "unbanChatMember"                   // https://core.telegram.org/bots/api#unbanchatmember
	// RequestMethodRestrictChatMember                = "restrictChatMember"                // https://core.telegram.org/bots/api#restrictchatmember
	// RequestMethodPromoteChatMember                 = "promoteChatMember"                 // https://core.telegram.org/bots/api#promotechatmember
	// RequestMethodSetChatAdministratorCustomTitle   = "setChatAdministratorCustomTitle"   // https://core.telegram.org/bots/api#setchatadministratorcustomtitle
	// RequestMethodBanChatSenderChat                 = "banChatSenderChat"                 // https://core.telegram.org/bots/api#banchatsenderchat
	// RequestMethodUnbanChatSenderChat               = "unbanChatSenderChat"               // https://core.telegram.org/bots/api#unbanchatsenderchat
	// RequestMethodSetChatPermissions                = "setChatPermissions"                // https://core.telegram.org/bots/api#setchatpermissions
	// RequestMethodExportChatInviteLink              = "exportChatInviteLink"              // https://core.telegram.org/bots/api#exportchatinvitelink
	// RequestMethodCreateChatInviteLink              = "createChatInviteLink"              // https://core.telegram.org/bots/api#createchatinvitelink
	// RequestMethodEditChatInviteLink                = "editChatInviteLink"                // https://core.telegram.org/bots/api#editchatinvitelink
	// RequestMethodRevokeChatInviteLink              = "revokeChatInviteLink"              // https://core.telegram.org/bots/api#revokechatinvitelink
	// RequestMethodApproveChatJoinRequest            = "approveChatJoinRequest"            // https://core.telegram.org/bots/api#approvechatjoinrequest
	// RequestMethodDeclineChatJoinRequest            = "declineChatJoinRequest"            // https://core.telegram.org/bots/api#declinechatjoinrequest
	// RequestMethodSetChatPhoto                      = "setChatPhoto"                      // https://core.telegram.org/bots/api#setchatphoto
	// RequestMethodDeleteChatPhoto                   = "deleteChatPhoto"                   // https://core.telegram.org/bots/api#deletechatphoto
	// RequestMethodSetChatTitle                      = "setChatTitle"                      // https://core.telegram.org/bots/api#setchattitle
	// RequestMethodSetChatDescription                = "setChatDescription"                // https://core.telegram.org/bots/api#setchatdescription
	// RequestMethodPinChatMessage                    = "pinChatMessage"                    // https://core.telegram.org/bots/api#pinchatmessage
	// RequestMethodUnpinChatMessage                  = "unpinChatMessage"                  // https://core.telegram.org/bots/api#unpinchatmessage
	// RequestMethodUnpinAllChatMessages              = "unpinAllChatMessages"              // https://core.telegram.org/bots/api#unpinallchatmessages
	// RequestMethodLeaveChat                         = "leaveChat"                         // https://core.telegram.org/bots/api#leavechat
	// RequestMethodGetChat                           = "getChat"                           // https://core.telegram.org/bots/api#getchat
	// RequestMethodGetChatAdministrators             = "getChatAdministrators"             // https://core.telegram.org/bots/api#getchatadministrators
	// RequestMethodGetChatMemberCount                = "getChatMemberCount"                // https://core.telegram.org/bots/api#getchatmembercount
	// RequestMethodGetChatMember                     = "getChatMember"                     // https://core.telegram.org/bots/api#getchatmember
	// RequestMethodSetChatStickerSet                 = "setChatStickerSet"                 // https://core.telegram.org/bots/api#setchatstickerset
	// RequestMethodDeleteChatStickerSet              = "deleteChatStickerSet"              // https://core.telegram.org/bots/api#deletechatstickerset
	// RequestMethodGetForumTopicIconStickers         = "getForumTopicIconStickers"         // https://core.telegram.org/bots/api#getforumtopiciconstickers
	// RequestMethodCreateForumTopic                  = "createForumTopic"                  // https://core.telegram.org/bots/api#createforumtopic
	// RequestMethodEditForumTopic                    = "editForumTopic"                    // https://core.telegram.org/bots/api#editforumtopic
	// RequestMethodCloseForumTopic                   = "closeForumTopic"                   // https://core.telegram.org/bots/api#closeforumtopic
	// RequestMethodReopenForumTopic                  = "reopenForumTopic"                  // https://core.telegram.org/bots/api#reopenforumtopic
	// RequestMethodDeleteForumTopic                  = "deleteForumTopic"                  // https://core.telegram.org/bots/api#deleteforumtopic
	// RequestMethodUnpinAllForumTopicMessages        = "unpinAllForumTopicMessages"        // https://core.telegram.org/bots/api#unpinallforumtopicmessages
	// RequestMethodEditGeneralForumTopic             = "editGeneralForumTopic"             // https://core.telegram.org/bots/api#editgeneralforumtopic
	// RequestMethodCloseGeneralForumTopic            = "closeGeneralForumTopic"            // https://core.telegram.org/bots/api#closegeneralforumtopic
	// RequestMethodReopenGeneralForumTopic           = "reopenGeneralForumTopic"           // https://core.telegram.org/bots/api#reopengeneralforumtopic
	// RequestMethodHideGeneralForumTopic             = "hideGeneralForumTopic"             // https://core.telegram.org/bots/api#hidegeneralforumtopic
	// RequestMethodUnhideGeneralForumTopic           = "unhideGeneralForumTopic"           // https://core.telegram.org/bots/api#unhidegeneralforumtopic
	// RequestMethodUnpinAllGeneralForumTopicMessages = "unpinAllGeneralForumTopicMessages" // https://core.telegram.org/bots/api#unpinallgeneralforumtopicmessages
	// RequestMethodAnswerCallbackQuery               = "answerCallbackQuery"               // https://core.telegram.org/bots/api#answercallbackquery
	// RequestMethodSetMyCommands                     = "setMyCommands"                     // https://core.telegram.org/bots/api#setmycommands
	// RequestMethodDeleteMyCommands                  = "deleteMyCommands"                  // https://core.telegram.org/bots/api#deletemycommands
	// RequestMethodGetMyCommands                     = "getMyCommands"                     // https://core.telegram.org/bots/api#getmycommands
	// RequestMethodSetMyName                         = "setMyName"                         // https://core.telegram.org/bots/api#setmyname
	// RequestMethodGetMyName                         = "getMyName"                         // https://core.telegram.org/bots/api#getmyname
	// RequestMethodSetMyDescription                  = "setMyDescription"                  // https://core.telegram.org/bots/api#setmydescription
	// RequestMethodGetMyDescription                  = "getMyDescription"                  // https://core.telegram.org/bots/api#getmydescription
	// RequestMethodSetMyShortDescription             = "setMyShortDescription"             // https://core.telegram.org/bots/api#setmyshortdescription
	// RequestMethodGetMyShortDescription             = "getMyShortDescription"             // https://core.telegram.org/bots/api#getmyshortdescription
	// RequestMethodSetChatMenuButton                 = "setChatMenuButton"                 // https://core.telegram.org/bots/api#setchatmenubutton
	// RequestMethodGetChatMenuButton                 = "getChatMenuButton"                 // https://core.telegram.org/bots/api#getchatmenubutton
	// RequestMethodSetMyDefaultAdministratorRights   = "setMyDefaultAdministratorRights"   // https://core.telegram.org/bots/api#setmydefaultadministratorrights
	// RequestMethodGetMyDefaultAdministratorRights   = "getMyDefaultAdministratorRights"   // https://core.telegram.org/bots/api#getmydefaultadministratorrights
	// RequestMethodEditMessageText                   = "editMessageText"                   // https://core.telegram.org/bots/api#editmessagetext
	// RequestMethodEditMessageCaption                = "editMessageCaption"                // https://core.telegram.org/bots/api#editmessagecaption
	// RequestMethodEditMessageMedia                  = "editMessageMedia"                  // https://core.telegram.org/bots/api#editmessagemedia
	// RequestMethodEditMessageLiveLocation           = "editMessageLiveLocation"           // https://core.telegram.org/bots/api#editmessagelivelocation
	// RequestMethodStopMessageLiveLocation           = "stopMessageLiveLocation"           // https://core.telegram.org/bots/api#stopmessagelivelocation
	// RequestMethodEditMessageReplyMarkup            = "editMessageReplyMarkup"            // https://core.telegram.org/bots/api#editmessagereplymarkup
	// RequestMethodStopPoll                          = "stopPoll"                          // https://core.telegram.org/bots/api#stoppoll
	// RequestMethodDeleteMessage                     = "deleteMessage"                     // https://core.telegram.org/bots/api#deletemessage
	// RequestMethodSendSticker                       = "sendSticker"                       // https://core.telegram.org/bots/api#sendsticker
	// RequestMethodGetStickerSet                     = "getStickerSet"                     // https://core.telegram.org/bots/api#getstickerset
	// RequestMethodGetCustomEmojiStickers            = "getCustomEmojiStickers"            // https://core.telegram.org/bots/api#getcustomemojistickers
	// RequestMethodUploadStickerFile                 = "uploadStickerFile"                 // https://core.telegram.org/bots/api#uploadstickerfile
	// RequestMethodCreateNewStickerSet               = "createNewStickerSet"               // https://core.telegram.org/bots/api#createnewstickerset
	// RequestMethodAddStickerToSet                   = "addStickerToSet"                   // https://core.telegram.org/bots/api#addstickertoset
	// RequestMethodSetStickerPositionInSet           = "setStickerPositionInSet"           // https://core.telegram.org/bots/api#setstickerpositioninset
	// RequestMethodDeleteStickerFromSet              = "deleteStickerFromSet"              // https://core.telegram.org/bots/api#deletestickerfromset
	// RequestMethodSetStickerEmojiList               = "setStickerEmojiList"               // https://core.telegram.org/bots/api#setstickeremojilist
	// RequestMethodSetStickerKeywords                = "setStickerKeywords"                // https://core.telegram.org/bots/api#setstickerkeywords
	// RequestMethodSetStickerMaskPosition            = "setStickerMaskPosition"            // https://core.telegram.org/bots/api#setstickermaskposition
	// RequestMethodSetStickerSetTitle                = "setStickerSetTitle"                // https://core.telegram.org/bots/api#setstickersettitle
	// RequestMethodSetStickerSetThumbnail            = "setStickerSetThumbnail"            // https://core.telegram.org/bots/api#setstickersetthumbnail
	// RequestMethodSetCustomEmojiStickerSetThumbnail = "setCustomEmojiStickerSetThumbnail" // https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
	// RequestMethodDeleteStickerSet                  = "deleteStickerSet"                  // https://core.telegram.org/bots/api#deletestickerset
	// RequestMethodAnswerInlineQuery                 = "answerInlineQuery"                 // https://core.telegram.org/bots/api#answerinlinequery
	// RequestMethodAnswerWebAppQuery                 = "answerWebAppQuery"                 // https://core.telegram.org/bots/api#answerwebappquery
	// RequestMethodSendInvoice                       = "sendInvoice"                       // https://core.telegram.org/bots/api#sendinvoice
	// RequestMethodCreateInvoiceLink                 = "createInvoiceLink"                 // https://core.telegram.org/bots/api#createinvoicelink
	// RequestMethodAnswerShippingQuery               = "answerShippingQuery"               // https://core.telegram.org/bots/api#answershippingquery
	// RequestMethodAnswerPreCheckoutQuery            = "answerPreCheckoutQuery"            // https://core.telegram.org/bots/api#answerprecheckoutquery
	// RequestMethodSetPassportDataErrors             = "setPassportDataErrors"             // https://core.telegram.org/bots/api#setpassportdataerrors
	// RequestMethodSendGame                          = "sendGame"                          // https://core.telegram.org/bots/api#sendgame
	// RequestMethodSetGameScore                      = "setGameScore"                      // https://core.telegram.org/bots/api#setgamescore
	// RequestMethodGetGameHighScores                 = "getGameHighScores"                 // https://core.telegram.org/bots/api#getgamehighscores
)

func (api Api) GetUpdates(params GetUpdatesParams, result *[]Update) error {
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return err
	}

	requestUrl := api.url + RequestMethodGetUpdates
	request, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewBuffer(jsonParams))
	if err != nil {
		return err
	}
	request.Header.Set(ContentType, ContentTypeJSON)

	response, err := api.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var apiResponse Response
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return err
	}

	if err := json.Unmarshal(apiResponse.Result, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetWebHook(params SetWebHookParams, result *bool) error {
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return err
	}

	requestUrl := api.url + RequestMethodSetWebHook
	request, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewBuffer(jsonParams))
	if err != nil {
		return err
	}
	request.Header.Set(ContentType, ContentTypeJSON)

	response, err := api.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var apiResponse Response
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return err
	}

	if err := json.Unmarshal(apiResponse.Result, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) DeleteWebHook(params DeleteWebHookParams, result *bool) error {
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return err
	}

	requestUrl := api.url + RequestMethodDeleteWebHook
	request, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewBuffer(jsonParams))
	if err != nil {
		return err
	}
	request.Header.Set(ContentType, ContentTypeJSON)

	response, err := api.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var apiResponse Response
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return err
	}

	if err := json.Unmarshal(apiResponse.Result, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetWebHookInfo(params GetWebHookInfoParams, result *WebHookInfo) error {
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return err
	}

	requestUrl := api.url + RequestMethodGetWebHookInfo
	request, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewBuffer(jsonParams))
	if err != nil {
		return err
	}
	request.Header.Set(ContentType, ContentTypeJSON)

	response, err := api.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var apiResponse Response
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return err
	}

	if err := json.Unmarshal(apiResponse.Result, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetMe(params GetMeParams, result *User) error {
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return err
	}

	requestUrl := api.url + RequestMethodGetMe
	request, err := http.NewRequest(http.MethodGet, requestUrl, bytes.NewBuffer(jsonParams))
	if err != nil {
		return err
	}
	request.Header.Set(ContentType, ContentTypeJSON)

	response, err := api.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var apiResponse Response
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return err
	}

	if err := json.Unmarshal(apiResponse.Result, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) LogOut(params LogOutParams, result *bool) error {
	// TODO:

	return nil
}

func (api Api) Close(params CloseParams, result *bool) error {
	// TODO:

	return nil
}

func (api Api) SendMessage(params SendMessageParams, result *Message) error {
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return err
	}

	requestUrl := api.url + RequestMethodSendMessage
	request, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewBuffer(jsonParams))
	if err != nil {
		return err
	}
	request.Header.Set(ContentType, ContentTypeJSON)

	response, err := api.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var apiResponse Response
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return err
	}

	if err := json.Unmarshal(apiResponse.Result, &result); err != nil {
		return err
	}

	return nil
}

// func (api Api) ForwardMessage() error {}
// func (api Api) CopyMessage() error {}
// func (api Api) SendPhoto() error {}
// func (api Api) SendAudio() error {}
// func (api Api) SendDocument() error {}
// func (api Api) SendVideo() error {}
// func (api Api) SendAnimation() error {}
// func (api Api) SendVoice() error {}
// func (api Api) SendVideoNote() error {}
// func (api Api) SendMediaGroup() error {}
// func (api Api) SendLocation() error {}
// func (api Api) SendVenue() error {}
// func (api Api) SendContact() error {}
// func (api Api) SendPoll() error {}
// func (api Api) SendDice() error {}
// func (api Api) SendChatAction() error {}
// func (api Api) GetUserProfilePhotos() error {}
// func (api Api) GetFile() error {}
// func (api Api) BanChatMember() error {}
// func (api Api) UnbanChatMember() error {}
// func (api Api) RestrictChatMember() error {}
// func (api Api) PromoteChatMember() error {}
// func (api Api) SetChatAdministratorCustomTitle() error {}
// func (api Api) BanChatSenderChat() error {}
// func (api Api) UnbanChatSenderChat() error {}
// func (api Api) SetChatPermissions() error {}
// func (api Api) ExportChatInviteLink() error {}
// func (api Api) CreateChatInviteLink() error {}
// func (api Api) EditChatInviteLink() error {}
// func (api Api) RevokeChatInviteLink() error {}
// func (api Api) ApproveChatJoinRequest() error {}
// func (api Api) DeclineChatJoinRequest() error {}
// func (api Api) SetChatPhoto() error {}
// func (api Api) DeleteChatPhoto() error {}
// func (api Api) SetChatTitle() error {}
// func (api Api) SetChatDescription() error {}
// func (api Api) PinChatMessage() error {}
// func (api Api) UnpinChatMessage() error {}
// func (api Api) UnpinAllChatMessages() error {}
// func (api Api) LeaveChat() error {}
// func (api Api) GetChat() error {}
// func (api Api) GetChatAdministrators() error {}
// func (api Api) GetChatMemberCount() error {}
// func (api Api) GetChatMember() error {}
// func (api Api) SetChatStickerSet() error {}
// func (api Api) DeleteChatStickerSet() error {}
// func (api Api) GetForumTopicIconStickers() error {}
// func (api Api) CreateForumTopic() error {}
// func (api Api) EditForumTopic() error {}
// func (api Api) CloseForumTopic() error {}
// func (api Api) ReopenForumTopic() error {}
// func (api Api) DeleteForumTopic() error {}
// func (api Api) UnpinAllForumTopicMessages() error {}
// func (api Api) EditGeneralForumTopic() error {}
// func (api Api) CloseGeneralForumTopic() error {}
// func (api Api) ReopenGeneralForumTopic() error {}
// func (api Api) HideGeneralForumTopic() error {}
// func (api Api) UnhideGeneralForumTopic() error {}
// func (api Api) UnpinAllGeneralForumTopicMessages() error {}
// func (api Api) AnswerCallbackQuery() error {}
// func (api Api) SetMyCommands() error {}
// func (api Api) DeleteMyCommands() error {}
// func (api Api) GetMyCommands() error {}
// func (api Api) SetMyName() error {}
// func (api Api) GetMyName() error {}
// func (api Api) SetMyDescription() error {}
// func (api Api) GetMyDescription() error {}
// func (api Api) SetMyShortDescription() error {}
// func (api Api) GetMyShortDescription() error {}
// func (api Api) SetChatMenuButton() error {}
// func (api Api) GetChatMenuButton() error {}
// func (api Api) SetMyDefaultAdministratorRights() error {}
// func (api Api) GetMyDefaultAdministratorRights() error {}
// func (api Api) EditMessageText() error {}
// func (api Api) EditMessageCaption() error {}
// func (api Api) EditMessageMedia() error {}
// func (api Api) EditMessageLiveLocation() error {}
// func (api Api) StopMessageLiveLocation() error {}
// func (api Api) EditMessageReplyMarkup() error {}
// func (api Api) StopPoll() error {}
// func (api Api) DeleteMessage() error {}
// func (api Api) SendSticker() error {}
// func (api Api) GetStickerSet() error {}
// func (api Api) GetCustomEmojiStickers() error {}
// func (api Api) UploadStickerFile() error {}
// func (api Api) CreateNewStickerSet() error {}
// func (api Api) AddStickerToSet() error {}
// func (api Api) SetStickerPositionInSet() error {}
// func (api Api) DeleteStickerFromSet() error {}
// func (api Api) SetStickerEmojiList() error {}
// func (api Api) SetStickerKeywords() error {}
// func (api Api) SetStickerMaskPosition() error {}
// func (api Api) SetStickerSetTitle() error {}
// func (api Api) SetStickerSetThumbnail() error {}
// func (api Api) SetCustomEmojiStickerSetThumbnail() error {}
// func (api Api) DeleteStickerSet() error {}
// func (api Api) AnswerInlineQuery() error {}
// func (api Api) AnswerWebAppQuery() error {}
// func (api Api) SendInvoice() error {}
// func (api Api) CreateInvoiceLink() error {}
// func (api Api) AnswerShippingQuery() error {}
// func (api Api) AnswerPreCheckoutQuery() error {}
// func (api Api) SetPassportDataErrors() error {}
// func (api Api) SendGame() error {}
// func (api Api) SetGameScore() error {}
// func (api Api) GetGameHighScores() error {}

package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	ContentTypeBinary = "application/octet-stream"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeJSON   = "application/json"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
)

const (
	RequestMethodGetUpdates                        = "getUpdates"                        // https://core.telegram.org/bots/api#getupdates
	// RequestMethodSetWebhook                        = "setWebhook"                        // https://core.telegram.org/bots/api#setwebhook
	// RequestMethodDeleteWebhook                     = "deleteWebhook"                     // https://core.telegram.org/bots/api#deletewebhook
	// RequestMethodGetWebhookInfo                    = "getWebhookInfo"                    // https://core.telegram.org/bots/api#getwebhookinfo
	RequestMethodGetMe = "getMe" // https://core.telegram.org/bots/api#getme
	// RequestMethodLogOut                            = "logOut"                            // https://core.telegram.org/bots/api#logout
	// RequestMethodClose                             = "close"                             // https://core.telegram.org/bots/api#close
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

func (api Api) GetUpdates(params GetUpdatesParams) (Response, error) {
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return Response{}, err
	}

	request, err := http.NewRequest("POST", api.url + RequestMethodGetUpdates, bytes.NewBuffer(jsonParams))
	if err != nil {
		return Response{}, err
	}
	request.Header.Set("Content-Type", ContentTypeJSON)

	response, err := api.client.Do(request)
	if err != nil {
		return Response{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Response{}, err
	}

	var apiResponse Response
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return Response{}, err
	}

	return apiResponse, err
}

// func (api Api) SetWebhook() (Response, error) {}
// func (api Api) DeleteWebhook() (Response, error) {}
// func (api Api) GetWebhookInfo() (Response, error) {}

func (api Api) GetMe() (Response, error) {
	response, err := api.client.Get(api.url + RequestMethodGetMe)
	if err != nil {
		return Response{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Response{}, err
	}

	var apiResponse Response
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return Response{}, err
	}

	return apiResponse, err
}

// func (api Api) LogOut() (Response, error) {}
// func (api Api) Close() (Response, error) {}

func (api Api) SendMessage(params SendMessageParams) (Response, error) {
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return Response{}, err
	}

	request, err := http.NewRequest("POST", api.url + RequestMethodSendMessage, bytes.NewBuffer(jsonParams))
	if err != nil {
		return Response{}, err
	}
	request.Header.Set("Content-Type", ContentTypeJSON)

	response, err := api.client.Do(request)
	if err != nil {
		return Response{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Response{}, err
	}

	var apiResponse Response
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return Response{}, err
	}

	return apiResponse, err
}

// func (api Api) ForwardMessage() (Response, error) {}
// func (api Api) CopyMessage() (Response, error) {}
// func (api Api) SendPhoto() (Response, error) {}
// func (api Api) SendAudio() (Response, error) {}
// func (api Api) SendDocument() (Response, error) {}
// func (api Api) SendVideo() (Response, error) {}
// func (api Api) SendAnimation() (Response, error) {}
// func (api Api) SendVoice() (Response, error) {}
// func (api Api) SendVideoNote() (Response, error) {}
// func (api Api) SendMediaGroup() (Response, error) {}
// func (api Api) SendLocation() (Response, error) {}
// func (api Api) SendVenue() (Response, error) {}
// func (api Api) SendContact() (Response, error) {}
// func (api Api) SendPoll() (Response, error) {}
// func (api Api) SendDice() (Response, error) {}
// func (api Api) SendChatAction() (Response, error) {}
// func (api Api) GetUserProfilePhotos() (Response, error) {}
// func (api Api) GetFile() (Response, error) {}
// func (api Api) BanChatMember() (Response, error) {}
// func (api Api) UnbanChatMember() (Response, error) {}
// func (api Api) RestrictChatMember() (Response, error) {}
// func (api Api) PromoteChatMember() (Response, error) {}
// func (api Api) SetChatAdministratorCustomTitle() (Response, error) {}
// func (api Api) BanChatSenderChat() (Response, error) {}
// func (api Api) UnbanChatSenderChat() (Response, error) {}
// func (api Api) SetChatPermissions() (Response, error) {}
// func (api Api) ExportChatInviteLink() (Response, error) {}
// func (api Api) CreateChatInviteLink() (Response, error) {}
// func (api Api) EditChatInviteLink() (Response, error) {}
// func (api Api) RevokeChatInviteLink() (Response, error) {}
// func (api Api) ApproveChatJoinRequest() (Response, error) {}
// func (api Api) DeclineChatJoinRequest() (Response, error) {}
// func (api Api) SetChatPhoto() (Response, error) {}
// func (api Api) DeleteChatPhoto() (Response, error) {}
// func (api Api) SetChatTitle() (Response, error) {}
// func (api Api) SetChatDescription() (Response, error) {}
// func (api Api) PinChatMessage() (Response, error) {}
// func (api Api) UnpinChatMessage() (Response, error) {}
// func (api Api) UnpinAllChatMessages() (Response, error) {}
// func (api Api) LeaveChat() (Response, error) {}
// func (api Api) GetChat() (Response, error) {}
// func (api Api) GetChatAdministrators() (Response, error) {}
// func (api Api) GetChatMemberCount() (Response, error) {}
// func (api Api) GetChatMember() (Response, error) {}
// func (api Api) SetChatStickerSet() (Response, error) {}
// func (api Api) DeleteChatStickerSet() (Response, error) {}
// func (api Api) GetForumTopicIconStickers() (Response, error) {}
// func (api Api) CreateForumTopic() (Response, error) {}
// func (api Api) EditForumTopic() (Response, error) {}
// func (api Api) CloseForumTopic() (Response, error) {}
// func (api Api) ReopenForumTopic() (Response, error) {}
// func (api Api) DeleteForumTopic() (Response, error) {}
// func (api Api) UnpinAllForumTopicMessages() (Response, error) {}
// func (api Api) EditGeneralForumTopic() (Response, error) {}
// func (api Api) CloseGeneralForumTopic() (Response, error) {}
// func (api Api) ReopenGeneralForumTopic() (Response, error) {}
// func (api Api) HideGeneralForumTopic() (Response, error) {}
// func (api Api) UnhideGeneralForumTopic() (Response, error) {}
// func (api Api) UnpinAllGeneralForumTopicMessages() (Response, error) {}
// func (api Api) AnswerCallbackQuery() (Response, error) {}
// func (api Api) SetMyCommands() (Response, error) {}
// func (api Api) DeleteMyCommands() (Response, error) {}
// func (api Api) GetMyCommands() (Response, error) {}
// func (api Api) SetMyName() (Response, error) {}
// func (api Api) GetMyName() (Response, error) {}
// func (api Api) SetMyDescription() (Response, error) {}
// func (api Api) GetMyDescription() (Response, error) {}
// func (api Api) SetMyShortDescription() (Response, error) {}
// func (api Api) GetMyShortDescription() (Response, error) {}
// func (api Api) SetChatMenuButton() (Response, error) {}
// func (api Api) GetChatMenuButton() (Response, error) {}
// func (api Api) SetMyDefaultAdministratorRights() (Response, error) {}
// func (api Api) GetMyDefaultAdministratorRights() (Response, error) {}
// func (api Api) EditMessageText() (Response, error) {}
// func (api Api) EditMessageCaption() (Response, error) {}
// func (api Api) EditMessageMedia() (Response, error) {}
// func (api Api) EditMessageLiveLocation() (Response, error) {}
// func (api Api) StopMessageLiveLocation() (Response, error) {}
// func (api Api) EditMessageReplyMarkup() (Response, error) {}
// func (api Api) StopPoll() (Response, error) {}
// func (api Api) DeleteMessage() (Response, error) {}
// func (api Api) SendSticker() (Response, error) {}
// func (api Api) GetStickerSet() (Response, error) {}
// func (api Api) GetCustomEmojiStickers() (Response, error) {}
// func (api Api) UploadStickerFile() (Response, error) {}
// func (api Api) CreateNewStickerSet() (Response, error) {}
// func (api Api) AddStickerToSet() (Response, error) {}
// func (api Api) SetStickerPositionInSet() (Response, error) {}
// func (api Api) DeleteStickerFromSet() (Response, error) {}
// func (api Api) SetStickerEmojiList() (Response, error) {}
// func (api Api) SetStickerKeywords() (Response, error) {}
// func (api Api) SetStickerMaskPosition() (Response, error) {}
// func (api Api) SetStickerSetTitle() (Response, error) {}
// func (api Api) SetStickerSetThumbnail() (Response, error) {}
// func (api Api) SetCustomEmojiStickerSetThumbnail() (Response, error) {}
// func (api Api) DeleteStickerSet() (Response, error) {}
// func (api Api) AnswerInlineQuery() (Response, error) {}
// func (api Api) AnswerWebAppQuery() (Response, error) {}
// func (api Api) SendInvoice() (Response, error) {}
// func (api Api) CreateInvoiceLink() (Response, error) {}
// func (api Api) AnswerShippingQuery() (Response, error) {}
// func (api Api) AnswerPreCheckoutQuery() (Response, error) {}
// func (api Api) SetPassportDataErrors() (Response, error) {}
// func (api Api) SendGame() (Response, error) {}
// func (api Api) SetGameScore() (Response, error) {}
// func (api Api) GetGameHighScores() (Response, error) {}

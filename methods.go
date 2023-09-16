package main

import (
	"encoding/json"
	"io/ioutil"
)

const (
	ContentTypeBinary = "application/octet-stream"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeJSON   = "application/json"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
)

const (
	// RequestMethodGetUpdates                        = "getUpdates"                        // https://core.telegram.org/bots/api#getupdates
	// RequestMethodSetWebhook                        = "setWebhook"                        // https://core.telegram.org/bots/api#setwebhook
	// RequestMethodDeleteWebhook                     = "deleteWebhook"                     // https://core.telegram.org/bots/api#deletewebhook
	// RequestMethodGetWebhookInfo                    = "getWebhookInfo"                    // https://core.telegram.org/bots/api#getwebhookinfo
	RequestMethodGetMe = "getMe" // https://core.telegram.org/bots/api#getme
	// RequestMethodLogOut                            = "logOut"                            // https://core.telegram.org/bots/api#logout
	// RequestMethodClose                             = "close"                             // https://core.telegram.org/bots/api#close
	// RequestMethodSendMessage                       = "sendMessage"                       // https://core.telegram.org/bots/api#sendmessage
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

// func (api Api) GetUpdates() (ApiResponse, error) {}
// func (api Api) SetWebhook() (ApiResponse, error) {}
// func (api Api) DeleteWebhook() (ApiResponse, error) {}
// func (api Api) GetWebhookInfo() (ApiResponse, error) {}

func (api Api) GetMe() (ApiResponse, error) {
	response, err := api.client.Get(api.url + RequestMethodGetMe)
	if err != nil {
		return ApiResponse{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ApiResponse{}, err
	}

	var apiResponse ApiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return ApiResponse{}, err
	}

	return apiResponse, err
}

// func (api Api) LogOut() (ApiResponse, error) {}
// func (api Api) Close() (ApiResponse, error) {}
// func (api Api) SendMessage() (ApiResponse, error) {}
// func (api Api) ForwardMessage() (ApiResponse, error) {}
// func (api Api) CopyMessage() (ApiResponse, error) {}
// func (api Api) SendPhoto() (ApiResponse, error) {}
// func (api Api) SendAudio() (ApiResponse, error) {}
// func (api Api) SendDocument() (ApiResponse, error) {}
// func (api Api) SendVideo() (ApiResponse, error) {}
// func (api Api) SendAnimation() (ApiResponse, error) {}
// func (api Api) SendVoice() (ApiResponse, error) {}
// func (api Api) SendVideoNote() (ApiResponse, error) {}
// func (api Api) SendMediaGroup() (ApiResponse, error) {}
// func (api Api) SendLocation() (ApiResponse, error) {}
// func (api Api) SendVenue() (ApiResponse, error) {}
// func (api Api) SendContact() (ApiResponse, error) {}
// func (api Api) SendPoll() (ApiResponse, error) {}
// func (api Api) SendDice() (ApiResponse, error) {}
// func (api Api) SendChatAction() (ApiResponse, error) {}
// func (api Api) GetUserProfilePhotos() (ApiResponse, error) {}
// func (api Api) GetFile() (ApiResponse, error) {}
// func (api Api) BanChatMember() (ApiResponse, error) {}
// func (api Api) UnbanChatMember() (ApiResponse, error) {}
// func (api Api) RestrictChatMember() (ApiResponse, error) {}
// func (api Api) PromoteChatMember() (ApiResponse, error) {}
// func (api Api) SetChatAdministratorCustomTitle() (ApiResponse, error) {}
// func (api Api) BanChatSenderChat() (ApiResponse, error) {}
// func (api Api) UnbanChatSenderChat() (ApiResponse, error) {}
// func (api Api) SetChatPermissions() (ApiResponse, error) {}
// func (api Api) ExportChatInviteLink() (ApiResponse, error) {}
// func (api Api) CreateChatInviteLink() (ApiResponse, error) {}
// func (api Api) EditChatInviteLink() (ApiResponse, error) {}
// func (api Api) RevokeChatInviteLink() (ApiResponse, error) {}
// func (api Api) ApproveChatJoinRequest() (ApiResponse, error) {}
// func (api Api) DeclineChatJoinRequest() (ApiResponse, error) {}
// func (api Api) SetChatPhoto() (ApiResponse, error) {}
// func (api Api) DeleteChatPhoto() (ApiResponse, error) {}
// func (api Api) SetChatTitle() (ApiResponse, error) {}
// func (api Api) SetChatDescription() (ApiResponse, error) {}
// func (api Api) PinChatMessage() (ApiResponse, error) {}
// func (api Api) UnpinChatMessage() (ApiResponse, error) {}
// func (api Api) UnpinAllChatMessages() (ApiResponse, error) {}
// func (api Api) LeaveChat() (ApiResponse, error) {}
// func (api Api) GetChat() (ApiResponse, error) {}
// func (api Api) GetChatAdministrators() (ApiResponse, error) {}
// func (api Api) GetChatMemberCount() (ApiResponse, error) {}
// func (api Api) GetChatMember() (ApiResponse, error) {}
// func (api Api) SetChatStickerSet() (ApiResponse, error) {}
// func (api Api) DeleteChatStickerSet() (ApiResponse, error) {}
// func (api Api) GetForumTopicIconStickers() (ApiResponse, error) {}
// func (api Api) CreateForumTopic() (ApiResponse, error) {}
// func (api Api) EditForumTopic() (ApiResponse, error) {}
// func (api Api) CloseForumTopic() (ApiResponse, error) {}
// func (api Api) ReopenForumTopic() (ApiResponse, error) {}
// func (api Api) DeleteForumTopic() (ApiResponse, error) {}
// func (api Api) UnpinAllForumTopicMessages() (ApiResponse, error) {}
// func (api Api) EditGeneralForumTopic() (ApiResponse, error) {}
// func (api Api) CloseGeneralForumTopic() (ApiResponse, error) {}
// func (api Api) ReopenGeneralForumTopic() (ApiResponse, error) {}
// func (api Api) HideGeneralForumTopic() (ApiResponse, error) {}
// func (api Api) UnhideGeneralForumTopic() (ApiResponse, error) {}
// func (api Api) UnpinAllGeneralForumTopicMessages() (ApiResponse, error) {}
// func (api Api) AnswerCallbackQuery() (ApiResponse, error) {}
// func (api Api) SetMyCommands() (ApiResponse, error) {}
// func (api Api) DeleteMyCommands() (ApiResponse, error) {}
// func (api Api) GetMyCommands() (ApiResponse, error) {}
// func (api Api) SetMyName() (ApiResponse, error) {}
// func (api Api) GetMyName() (ApiResponse, error) {}
// func (api Api) SetMyDescription() (ApiResponse, error) {}
// func (api Api) GetMyDescription() (ApiResponse, error) {}
// func (api Api) SetMyShortDescription() (ApiResponse, error) {}
// func (api Api) GetMyShortDescription() (ApiResponse, error) {}
// func (api Api) SetChatMenuButton() (ApiResponse, error) {}
// func (api Api) GetChatMenuButton() (ApiResponse, error) {}
// func (api Api) SetMyDefaultAdministratorRights() (ApiResponse, error) {}
// func (api Api) GetMyDefaultAdministratorRights() (ApiResponse, error) {}
// func (api Api) EditMessageText() (ApiResponse, error) {}
// func (api Api) EditMessageCaption() (ApiResponse, error) {}
// func (api Api) EditMessageMedia() (ApiResponse, error) {}
// func (api Api) EditMessageLiveLocation() (ApiResponse, error) {}
// func (api Api) StopMessageLiveLocation() (ApiResponse, error) {}
// func (api Api) EditMessageReplyMarkup() (ApiResponse, error) {}
// func (api Api) StopPoll() (ApiResponse, error) {}
// func (api Api) DeleteMessage() (ApiResponse, error) {}
// func (api Api) SendSticker() (ApiResponse, error) {}
// func (api Api) GetStickerSet() (ApiResponse, error) {}
// func (api Api) GetCustomEmojiStickers() (ApiResponse, error) {}
// func (api Api) UploadStickerFile() (ApiResponse, error) {}
// func (api Api) CreateNewStickerSet() (ApiResponse, error) {}
// func (api Api) AddStickerToSet() (ApiResponse, error) {}
// func (api Api) SetStickerPositionInSet() (ApiResponse, error) {}
// func (api Api) DeleteStickerFromSet() (ApiResponse, error) {}
// func (api Api) SetStickerEmojiList() (ApiResponse, error) {}
// func (api Api) SetStickerKeywords() (ApiResponse, error) {}
// func (api Api) SetStickerMaskPosition() (ApiResponse, error) {}
// func (api Api) SetStickerSetTitle() (ApiResponse, error) {}
// func (api Api) SetStickerSetThumbnail() (ApiResponse, error) {}
// func (api Api) SetCustomEmojiStickerSetThumbnail() (ApiResponse, error) {}
// func (api Api) DeleteStickerSet() (ApiResponse, error) {}
// func (api Api) AnswerInlineQuery() (ApiResponse, error) {}
// func (api Api) AnswerWebAppQuery() (ApiResponse, error) {}
// func (api Api) SendInvoice() (ApiResponse, error) {}
// func (api Api) CreateInvoiceLink() (ApiResponse, error) {}
// func (api Api) AnswerShippingQuery() (ApiResponse, error) {}
// func (api Api) AnswerPreCheckoutQuery() (ApiResponse, error) {}
// func (api Api) SetPassportDataErrors() (ApiResponse, error) {}
// func (api Api) SendGame() (ApiResponse, error) {}
// func (api Api) SetGameScore() (ApiResponse, error) {}
// func (api Api) GetGameHighScores() (ApiResponse, error) {}

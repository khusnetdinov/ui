package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/google/uuid"
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
	RequestMethodForwardMessage                    = "forwardMessage"                    // https://core.telegram.org/bots/api#forwardmessage
	RequestMethodCopyMessage                       = "copyMessage"                       // https://core.telegram.org/bots/api#copymessage
	RequestMethodSendPhoto                         = "sendPhoto"                         // https://core.telegram.org/bots/api#sendphoto
	RequestMethodSendAudio                         = "sendAudio"                         // https://core.telegram.org/bots/api#sendaudio
	RequestMethodSendDocument                      = "sendDocument"                      // https://core.telegram.org/bots/api#senddocument
	RequestMethodSendVideo                         = "sendVideo"                         // https://core.telegram.org/bots/api#sendvideo
	RequestMethodSendAnimation                     = "sendAnimation"                     // https://core.telegram.org/bots/api#sendanimation
	RequestMethodSendVoice                         = "sendVoice"                         // https://core.telegram.org/bots/api#sendvoice
	RequestMethodSendVideoNote                     = "sendVideoNote"                     // https://core.telegram.org/bots/api#sendvideonote
	RequestMethodSendMediaGroup                    = "sendMediaGroup"                    // https://core.telegram.org/bots/api#sendmediagroup
	RequestMethodSendLocation                      = "sendLocation"                      // https://core.telegram.org/bots/api#sendlocation
	RequestMethodSendVenue                         = "sendVenue"                         // https://core.telegram.org/bots/api#sendvenue
	RequestMethodSendContact                       = "sendContact"                       // https://core.telegram.org/bots/api#sendcontact
	RequestMethodSendPoll                          = "sendPoll"                          // https://core.telegram.org/bots/api#sendpoll
	RequestMethodSendDice                          = "sendDice"                          // https://core.telegram.org/bots/api#senddice
	RequestMethodSendChatAction                    = "sendChatAction"                    // https://core.telegram.org/bots/api#sendchataction
	RequestMethodGetUserProfilePhotos              = "getUserProfilePhotos"              // https://core.telegram.org/bots/api#getuserprofilephotos
	RequestMethodGetFile                           = "getFile"                           // https://core.telegram.org/bots/api#getfile
	RequestMethodBanChatMember                     = "banChatMember"                     // https://core.telegram.org/bots/api#banchatmember
	RequestMethodUnbanChatMember                   = "unbanChatMember"                   // https://core.telegram.org/bots/api#unbanchatmember
	RequestMethodRestrictChatMember                = "restrictChatMember"                // https://core.telegram.org/bots/api#restrictchatmember
	RequestMethodPromoteChatMember                 = "promoteChatMember"                 // https://core.telegram.org/bots/api#promotechatmember
	RequestMethodSetChatAdministratorCustomTitle   = "setChatAdministratorCustomTitle"   // https://core.telegram.org/bots/api#setchatadministratorcustomtitle
	RequestMethodBanChatSenderChat                 = "banChatSenderChat"                 // https://core.telegram.org/bots/api#banchatsenderchat
	RequestMethodUnbanChatSenderChat               = "unbanChatSenderChat"               // https://core.telegram.org/bots/api#unbanchatsenderchat
	RequestMethodSetChatPermissions                = "setChatPermissions"                // https://core.telegram.org/bots/api#setchatpermissions
	RequestMethodExportChatInviteLink              = "exportChatInviteLink"              // https://core.telegram.org/bots/api#exportchatinvitelink
	RequestMethodCreateChatInviteLink              = "createChatInviteLink"              // https://core.telegram.org/bots/api#createchatinvitelink
	RequestMethodEditChatInviteLink                = "editChatInviteLink"                // https://core.telegram.org/bots/api#editchatinvitelink
	RequestMethodRevokeChatInviteLink              = "revokeChatInviteLink"              // https://core.telegram.org/bots/api#revokechatinvitelink
	RequestMethodApproveChatJoinRequest            = "approveChatJoinRequest"            // https://core.telegram.org/bots/api#approvechatjoinrequest
	RequestMethodDeclineChatJoinRequest            = "declineChatJoinRequest"            // https://core.telegram.org/bots/api#declinechatjoinrequest
	RequestMethodSetChatPhoto                      = "setChatPhoto"                      // https://core.telegram.org/bots/api#setchatphoto
	RequestMethodDeleteChatPhoto                   = "deleteChatPhoto"                   // https://core.telegram.org/bots/api#deletechatphoto
	RequestMethodSetChatTitle                      = "setChatTitle"                      // https://core.telegram.org/bots/api#setchattitle
	RequestMethodSetChatDescription                = "setChatDescription"                // https://core.telegram.org/bots/api#setchatdescription
	RequestMethodPinChatMessage                    = "pinChatMessage"                    // https://core.telegram.org/bots/api#pinchatmessage
	RequestMethodUnpinChatMessage                  = "unpinChatMessage"                  // https://core.telegram.org/bots/api#unpinchatmessage
	RequestMethodUnpinAllChatMessages              = "unpinAllChatMessages"              // https://core.telegram.org/bots/api#unpinallchatmessages
	RequestMethodLeaveChat                         = "leaveChat"                         // https://core.telegram.org/bots/api#leavechat
	RequestMethodGetChat                           = "getChat"                           // https://core.telegram.org/bots/api#getchat
	RequestMethodGetChatAdministrators             = "getChatAdministrators"             // https://core.telegram.org/bots/api#getchatadministrators
	RequestMethodGetChatMemberCount                = "getChatMemberCount"                // https://core.telegram.org/bots/api#getchatmembercount
	RequestMethodGetChatMember                     = "getChatMember"                     // https://core.telegram.org/bots/api#getchatmember
	RequestMethodSetChatStickerSet                 = "setChatStickerSet"                 // https://core.telegram.org/bots/api#setchatstickerset
	RequestMethodDeleteChatStickerSet              = "deleteChatStickerSet"              // https://core.telegram.org/bots/api#deletechatstickerset
	RequestMethodGetForumTopicIconStickers         = "getForumTopicIconStickers"         // https://core.telegram.org/bots/api#getforumtopiciconstickers
	RequestMethodCreateForumTopic                  = "createForumTopic"                  // https://core.telegram.org/bots/api#createforumtopic
	RequestMethodEditForumTopic                    = "editForumTopic"                    // https://core.telegram.org/bots/api#editforumtopic
	RequestMethodCloseForumTopic                   = "closeForumTopic"                   // https://core.telegram.org/bots/api#closeforumtopic
	RequestMethodReopenForumTopic                  = "reopenForumTopic"                  // https://core.telegram.org/bots/api#reopenforumtopic
	RequestMethodDeleteForumTopic                  = "deleteForumTopic"                  // https://core.telegram.org/bots/api#deleteforumtopic
	RequestMethodUnpinAllForumTopicMessages        = "unpinAllForumTopicMessages"        // https://core.telegram.org/bots/api#unpinallforumtopicmessages
	RequestMethodEditGeneralForumTopic             = "editGeneralForumTopic"             // https://core.telegram.org/bots/api#editgeneralforumtopic
	RequestMethodCloseGeneralForumTopic            = "closeGeneralForumTopic"            // https://core.telegram.org/bots/api#closegeneralforumtopic
	RequestMethodReopenGeneralForumTopic           = "reopenGeneralForumTopic"           // https://core.telegram.org/bots/api#reopengeneralforumtopic
	RequestMethodHideGeneralForumTopic             = "hideGeneralForumTopic"             // https://core.telegram.org/bots/api#hidegeneralforumtopic
	RequestMethodUnhideGeneralForumTopic           = "unhideGeneralForumTopic"           // https://core.telegram.org/bots/api#unhidegeneralforumtopic
	RequestMethodUnpinAllGeneralForumTopicMessages = "unpinAllGeneralForumTopicMessages" // https://core.telegram.org/bots/api#unpinallgeneralforumtopicmessages
	RequestMethodAnswerCallbackQuery               = "answerCallbackQuery"               // https://core.telegram.org/bots/api#answercallbackquery
	RequestMethodSetMyCommands                     = "setMyCommands"                     // https://core.telegram.org/bots/api#setmycommands
	RequestMethodDeleteMyCommands                  = "deleteMyCommands"                  // https://core.telegram.org/bots/api#deletemycommands
	RequestMethodGetMyCommands                     = "getMyCommands"                     // https://core.telegram.org/bots/api#getmycommands
	RequestMethodSetMyName                         = "setMyName"                         // https://core.telegram.org/bots/api#setmyname
	RequestMethodGetMyName                         = "getMyName"                         // https://core.telegram.org/bots/api#getmyname
	RequestMethodSetMyDescription                  = "setMyDescription"                  // https://core.telegram.org/bots/api#setmydescription
	RequestMethodGetMyDescription                  = "getMyDescription"                  // https://core.telegram.org/bots/api#getmydescription
	RequestMethodSetMyShortDescription             = "setMyShortDescription"             // https://core.telegram.org/bots/api#setmyshortdescription
	RequestMethodGetMyShortDescription             = "getMyShortDescription"             // https://core.telegram.org/bots/api#getmyshortdescription
	RequestMethodSetChatMenuButton                 = "setChatMenuButton"                 // https://core.telegram.org/bots/api#setchatmenubutton
	RequestMethodGetChatMenuButton                 = "getChatMenuButton"                 // https://core.telegram.org/bots/api#getchatmenubutton
	RequestMethodSetMyDefaultAdministratorRights   = "setMyDefaultAdministratorRights"   // https://core.telegram.org/bots/api#setmydefaultadministratorrights
	RequestMethodGetMyDefaultAdministratorRights   = "getMyDefaultAdministratorRights"   // https://core.telegram.org/bots/api#getmydefaultadministratorrights
	RequestMethodEditMessageText                   = "editMessageText"                   // https://core.telegram.org/bots/api#editmessagetext
	RequestMethodEditMessageCaption                = "editMessageCaption"                // https://core.telegram.org/bots/api#editmessagecaption
	RequestMethodEditMessageMedia                  = "editMessageMedia"                  // https://core.telegram.org/bots/api#editmessagemedia
	RequestMethodEditMessageLiveLocation           = "editMessageLiveLocation"           // https://core.telegram.org/bots/api#editmessagelivelocation
	RequestMethodStopMessageLiveLocation           = "stopMessageLiveLocation"           // https://core.telegram.org/bots/api#stopmessagelivelocation
	RequestMethodEditMessageReplyMarkup            = "editMessageReplyMarkup"            // https://core.telegram.org/bots/api#editmessagereplymarkup
	RequestMethodStopPoll                          = "stopPoll"                          // https://core.telegram.org/bots/api#stoppoll
	RequestMethodDeleteMessage                     = "deleteMessage"                     // https://core.telegram.org/bots/api#deletemessage
	RequestMethodSendSticker                       = "sendSticker"                       // https://core.telegram.org/bots/api#sendsticker
	RequestMethodGetStickerSet                     = "getStickerSet"                     // https://core.telegram.org/bots/api#getstickerset
	RequestMethodGetCustomEmojiStickers            = "getCustomEmojiStickers"            // https://core.telegram.org/bots/api#getcustomemojistickers
	RequestMethodUploadStickerFile                 = "uploadStickerFile"                 // https://core.telegram.org/bots/api#uploadstickerfile
	RequestMethodCreateNewStickerSet               = "createNewStickerSet"               // https://core.telegram.org/bots/api#createnewstickerset
	RequestMethodAddStickerToSet                   = "addStickerToSet"                   // https://core.telegram.org/bots/api#addstickertoset
	RequestMethodSetStickerPositionInSet           = "setStickerPositionInSet"           // https://core.telegram.org/bots/api#setstickerpositioninset
	RequestMethodDeleteStickerFromSet              = "deleteStickerFromSet"              // https://core.telegram.org/bots/api#deletestickerfromset
	RequestMethodSetStickerEmojiList               = "setStickerEmojiList"               // https://core.telegram.org/bots/api#setstickeremojilist
	RequestMethodSetStickerKeywords                = "setStickerKeywords"                // https://core.telegram.org/bots/api#setstickerkeywords
	RequestMethodSetStickerMaskPosition            = "setStickerMaskPosition"            // https://core.telegram.org/bots/api#setstickermaskposition
	RequestMethodSetStickerSetTitle                = "setStickerSetTitle"                // https://core.telegram.org/bots/api#setstickersettitle
	RequestMethodSetStickerSetThumbnail            = "setStickerSetThumbnail"            // https://core.telegram.org/bots/api#setstickersetthumbnail
	RequestMethodSetCustomEmojiStickerSetThumbnail = "setCustomEmojiStickerSetThumbnail" // https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
	RequestMethodDeleteStickerSet                  = "deleteStickerSet"                  // https://core.telegram.org/bots/api#deletestickerset
	RequestMethodAnswerInlineQuery                 = "answerInlineQuery"                 // https://core.telegram.org/bots/api#answerinlinequery
	RequestMethodAnswerWebAppQuery                 = "answerWebAppQuery"                 // https://core.telegram.org/bots/api#answerwebappquery
	RequestMethodSendInvoice                       = "sendInvoice"                       // https://core.telegram.org/bots/api#sendinvoice
	RequestMethodCreateInvoiceLink                 = "createInvoiceLink"                 // https://core.telegram.org/bots/api#createinvoicelink
	RequestMethodAnswerShippingQuery               = "answerShippingQuery"               // https://core.telegram.org/bots/api#answershippingquery
	RequestMethodAnswerPreCheckoutQuery            = "answerPreCheckoutQuery"            // https://core.telegram.org/bots/api#answerprecheckoutquery
	RequestMethodSetPassportDataErrors             = "setPassportDataErrors"             // https://core.telegram.org/bots/api#setpassportdataerrors
	RequestMethodSendGame                          = "sendGame"                          // https://core.telegram.org/bots/api#sendgame
	RequestMethodSetGameScore                      = "setGameScore"                      // https://core.telegram.org/bots/api#setgamescore
	RequestMethodGetGameHighScores                 = "getGameHighScores"                 // https://core.telegram.org/bots/api#getgamehighscores
)

func (api Api) request(url string, params interface{}, result interface{}) error {
	uuid := fmt.Sprintf("%s", uuid.New())
	api.Logger.Debug(uuid, "url", url, "params", params)

	jsonParams, err := json.Marshal(params)
	if err != nil {
		return err
	}

	requestUrl := api.url + url
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

	var apiResponse ApiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return err
	}
	api.Logger.Debug(uuid, "apiResponse", apiResponse)

	if !apiResponse.Ok {
		return NewApiResponseError(&apiResponse)
	}

	if err := json.Unmarshal(apiResponse.Result, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetUpdates(params RequestParamsGetUpdates, result *[]Update) error {
	if err := api.request(RequestMethodGetUpdates, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetWebHook(params RequestParamsSetWebHook, result *bool) error {
	if err := api.request(RequestMethodSetWebHook, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) DeleteWebHook(params RequestParamsDeleteWebHook, result *bool) error {
	if err := api.request(RequestMethodDeleteWebHook, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetWebHookInfo(params RequestParamsGetWebHookInfo, result *WebHookInfo) error {
	if err := api.request(RequestMethodGetWebHookInfo, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetMe(params RequestParamsGetMe, result *User) error {
	if err := api.request(RequestMethodGetMe, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) LogOut(params RequestParamsLogOut, result *bool) error {
	if err := api.request(RequestMethodLogOut, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) Close(params RequestParamsClose, result *bool) error {
	if err := api.request(RequestMethodClose, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SendMessage(params RequestParamsSendMessage, result *Message) error {
	if err := api.request(RequestMethodSendMessage, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) ForwardMessage(params RequestParamsForwardMessage, result *Message) error {
	if err := api.request(RequestMethodForwardMessage, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) CopyMessage(params RequestParamsCopyMessage, result *int64) error {
	if err := api.request(RequestMethodCopyMessage, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SendPhoto(params RequestParamsSendPhoto, result *Message) error {
	if err := api.request(RequestMethodSendPhoto, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SendAudio(params RequestParamsSendAudio, result *Message) error {
	if err := api.request(RequestMethodSendAudio, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SendDocument(params RequestParamsSendDocument, result *Message) error {
	if err := api.request(RequestMethodSendDocument, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SendVideo(params RequestParamsSendVideo, result *Message) error {
	if err := api.request(RequestMethodSendVideo, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SendAnimation(params RequestParamsSendAnimation, result *Message) error {
	if err := api.request(RequestMethodSendAnimation, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SendVoice(params RequestParamsSendVoice, result *Message) error {
	if err := api.request(RequestMethodSendVoice, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SendVideoNote(params RequestParamsSendVideoNote, result *Message) error {
	if err := api.request(RequestMethodSendVideoNote, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SendMediaGroup(params RequestParamsSendMediaGroup, result *Message) error {
	if err := api.request(RequestMethodSendMediaGroup, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SendLocation(params RequestParamsSendLocation, result *Message) error {
	if err := api.request(RequestMethodSendLocation, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SendVenue(params RequestParamsSendVenue, result *Message) error {
	if err := api.request(RequestMethodSendVenue, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SendContact(params RequestParamsSendContact, result *Message) error {
	if err := api.request(RequestMethodSendContact, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SendPoll(params RequestParamsSendPoll, result *Message) error {
	if err := api.request(RequestMethodSendPoll, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SendDice(params RequestParamsSendDice, result *Message) error {
	if err := api.request(RequestMethodSendDice, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SendChatAction(params RequestParamsSendChatAction, result *bool) error {
	if err := api.request(RequestMethodSendChatAction, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetUserProfilePhotos(params RequestParamsGetUserProfilePhotos, result *UserProfilePhotos) error {
	if err := api.request(RequestMethodGetUserProfilePhotos, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetFile(params RequestParamsGetFile, result *File) error {
	if err := api.request(RequestMethodGetFile, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) BanChatMember(params RequestParamsBanChatMember, result *bool) error {
	if err := api.request(RequestMethodBanChatMember, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) UnbanChatMember(params RequestParamsUnbanChatMember, result *bool) error {
	if err := api.request(RequestMethodUnbanChatMember, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) RestrictChatMember(params RequestParamsRestrictChatMember, result *bool) error {
	if err := api.request(RequestMethodRestrictChatMember, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) PromoteChatMember(params RequestParamsPromoteChatMember, result *bool) error {
	if err := api.request(RequestMethodPromoteChatMember, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetChatAdministratorCustomTitle(params RequestParamsSetChatAdministratorCustomTitle, result *bool) error {
	if err := api.request(RequestMethodSetChatAdministratorCustomTitle, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) BanChatSenderChat(params RequestParamsBanChatSenderChat, result *bool) error {
	if err := api.request(RequestMethodBanChatSenderChat, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) UnbanChatSenderChat(params RequestParamsUnbanChatSenderChat, result *bool) error {
	if err := api.request(RequestMethodUnbanChatSenderChat, params, &result); err != nil {
		return err
	 }

	return nil
}

func (api Api) SetChatPermissions(params RequestParamsSetChatPermissions, result *bool) error {
	if err := api.request(RequestMethodSetChatPermissions, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) ExportChatInviteLink(params RequestParamsExportChatInviteLink, result *string) error {
	if err := api.request(RequestMethodExportChatInviteLink, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) CreateChatInviteLink(params RequestParamsCreateChatInviteLink, result *ChatInviteLink) error {
	if err := api.request(RequestMethodCreateChatInviteLink, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) EditChatInviteLink(params RequestParamsEditChatInviteLink, result *ChatInviteLink) error {
	if err := api.request(RequestMethodEditChatInviteLink, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) RevokeChatInviteLink(params RequestParamsRevokeChatInviteLink, result *ChatInviteLink) error {
	if err := api.request(RequestMethodRevokeChatInviteLink, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) ApproveChatJoinRequest(params RequestParamsApproveChatJoinRequest, result *bool) error {
	if err := api.request(RequestMethodApproveChatJoinRequest, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) DeclineChatJoinRequest(params RequestParamsDeclineChatJoinRequest, result *bool) error {
	if err := api.request(RequestMethodDeclineChatJoinRequest, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetChatPhoto(params RequestParamsSetChatPhoto, result *bool) error {
	if err := api.request(RequestMethodSetChatPhoto, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) DeleteChatPhoto(params RequestParamsDeleteChatPhoto, result *bool) error {
	if err := api.request(RequestMethodDeleteChatPhoto, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetChatTitle(params RequestParamsSetChatTitle, result *bool) error {
	if err := api.request(RequestMethodSetChatTitle, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetChatDescription(params RequestParamsSetChatDescription, result *bool) error {
	if err := api.request(RequestMethodSetChatDescription, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) PinChatMessage(params RequestParamsPinChatMessage, result *bool) error {
	if err := api.request(RequestMethodPinChatMessage, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) UnpinChatMessage(params RequestParamsUnpinChatMessage, result *bool) error {
	if err := api.request(RequestMethodUnpinChatMessage, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) UnpinAllChatMessages(params RequestParamsUnpinAllChatMessages, result *bool) error {
	if err := api.request(RequestMethodUnpinAllChatMessages, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) LeaveChat(params RequestParamsLeaveChat, result *bool) error {
	if err := api.request(RequestMethodLeaveChat, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetChat(params RequestParamsGetChat, result *Chat) error {
	if err := api.request(RequestMethodGetChat, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetChatAdministrators(params RequestParamsGetChatAdministrators, result *[]ChatMember) error {
	if err := api.request(RequestMethodGetChatAdministrators, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetChatMemberCount(params RequestParamsGetChatMemberCount, result *int64) error {
	if err := api.request(RequestMethodGetChatMemberCount, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetChatMember(params RequestParamsGetChatMember, result *ChatMember) error {
	if err := api.request(RequestMethodGetChatMember, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetChatStickerSet(params RequestParamsSetChatStickerSet, result *bool) error {
	if err := api.request(RequestMethodSetChatStickerSet, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) DeleteChatStickerSet(params RequestParamsDeleteChatStickerSet, result *bool) error {
	if err := api.request(RequestMethodDeleteChatStickerSet, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetForumTopicIconStickers(params RequestParamsGetForumTopicIconStickers, result *[]Sticker) error {
	if err := api.request(RequestMethodGetForumTopicIconStickers, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) CreateForumTopic(params RequestParamsCreateForumTopic, result *ForumTopic) error {
	if err := api.request(RequestMethodCreateForumTopic, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) EditForumTopic(params RequestParamsEditForumTopic, result *bool) error {
	if err := api.request(RequestMethodEditForumTopic, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) CloseForumTopic(params RequestParamsCloseForumTopic, result *bool) error {
	if err := api.request(RequestMethodCloseForumTopic, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) ReopenForumTopic(params RequestParamsReopenForumTopic, result *bool) error {
	if err := api.request(RequestMethodReopenForumTopic, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) DeleteForumTopic(params RequestParamsDeleteForumTopic, result *bool) error {
	if err := api.request(RequestMethodDeleteForumTopic, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) UnpinAllForumTopicMessages(params RequestParamsUnpinAllForumTopicMessages, result *bool) error {
	if err := api.request(RequestMethodUnpinAllForumTopicMessages, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) EditGeneralForumTopic(params RequestParamsEditGeneralForumTopic, result *bool) error {
	if err := api.request(RequestMethodEditGeneralForumTopic, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) CloseGeneralForumTopic(params RequestParamsCloseGeneralForumTopic, result *bool) error {
	if err := api.request(RequestMethodCloseGeneralForumTopic, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) ReopenGeneralForumTopic(params RequestParamsReopenGeneralForumTopic, result *bool) error {
	if err := api.request(RequestMethodReopenGeneralForumTopic, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) HideGeneralForumTopic(params RequestParamsHideGeneralForumTopic, result *bool) error {
	if err := api.request(RequestMethodHideGeneralForumTopic, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) UnhideGeneralForumTopic(params RequestParamsUnhideGeneralForumTopic, result *bool) error {
	if err := api.request(RequestMethodUnhideGeneralForumTopic, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) UnpinAllGeneralForumTopicMessages(params RequestParamsUnpinAllGeneralForumTopicMessages, result *bool) error {
	if err := api.request(RequestMethodUnpinAllGeneralForumTopicMessages, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) AnswerCallbackQuery(params RequestParamsAnswerCallbackQuery, result *bool) error {
	if err := api.request(RequestMethodAnswerCallbackQuery, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetMyCommands(params RequestParamsSetMyCommands, result *bool) error {
	if err := api.request(RequestMethodSetMyCommands, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) DeleteMyCommands(params RequestParamsDeleteMyCommands, result *bool) error {
	if err := api.request(RequestMethodDeleteMyCommands, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetMyCommands(params RequestParamsGetMyCommands, result *[]BotCommand) error {
	if err := api.request(RequestMethodGetMyCommands, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetMyName(params RequestParamsSetMyName, result *bool) error {
	if err := api.request(RequestMethodSetMyName, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetMyName(params RequestParamsGetMyName, result *BotName) error {
	if err := api.request(RequestMethodGetMyName, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetMyDescription(params RequestParamsSetMyDescription, result *bool) error {
	if err := api.request(RequestMethodSetMyDescription, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetMyDescription(params RequestParamsGetMyDescription, result *BotDescription) error {
	if err := api.request(RequestMethodGetMyDescription, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetMyShortDescription(params RequestParamsSetMyShortDescription, result *bool) error {
	if err := api.request(RequestMethodSetMyShortDescription, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetMyShortDescription(params RequestParamsGetMyShortDescription, result *BotShortDescription) error {
	if err := api.request(RequestMethodGetMyShortDescription, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetChatMenuButton(params RequestParamsSetChatMenuButton, result *bool) error {
	if err := api.request(RequestMethodSetChatMenuButton, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetChatMenuButton(params RequestParamsGetChatMenuButton, result *MenuButton) error {
	if err := api.request(RequestMethodGetChatMenuButton, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetMyDefaultAdministratorRights(params RequestParamsSetMyDefaultAdministratorRights, result *bool) error {
	if err := api.request(RequestMethodSetMyDefaultAdministratorRights, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetMyDefaultAdministratorRights(params RequestParamsGetMyDefaultAdministratorRights, result *ChatAdministratorRights) error {
	if err := api.request(RequestMethodGetMyDefaultAdministratorRights, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) EditMessageText(params RequestParamsEditMessageText, result *Message) error {
	if err := api.request(RequestMethodEditMessageText, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) EditInlineMessageText(params RequestParamsEditMessageText, result *bool) error {
	if err := api.request(RequestMethodEditMessageText, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) EditMessageCaption(params RequestParamsEditMessageCaption, result *Message) error {
	if err := api.request(RequestMethodEditMessageCaption, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) EditInlineMessageCaption(params RequestParamsEditMessageCaption, result **bool) error {
	if err := api.request(RequestMethodEditMessageCaption, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) EditMessageMedia(params RequestParamsEditMessageMedia, result *Message) error {
	if err := api.request(RequestMethodEditMessageMedia, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) EditInlineMessageMedia(params RequestParamsEditMessageMedia, result *bool) error {
	if err := api.request(RequestMethodEditMessageMedia, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) EditMessageLiveLocation(params RequestParamsEditMessageLiveLocation, result *Message) error {
	if err := api.request(RequestMethodEditMessageLiveLocation, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) EditInlineMessageLiveLocation(params RequestParamsEditMessageLiveLocation, result *bool) error {
	if err := api.request(RequestMethodEditMessageLiveLocation, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) StopMessageLiveLocation(params RequestParamsStopMessageLiveLocation, result *Message) error {
	if err := api.request(RequestMethodStopMessageLiveLocation, params, &result); err != nil {
		return err
	}

	return nil
}


func (api Api) StopInlineMessageLiveLocation(params RequestParamsStopMessageLiveLocation, result *bool) error {
	if err := api.request(RequestMethodStopMessageLiveLocation, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) EditMessageReplyMarkup(params RequestParamsEditMessageReplyMarkup, result *Message) error {
	if err := api.request(RequestMethodEditMessageReplyMarkup, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) EditInlineMessageReplyMarkup(params RequestParamsEditMessageReplyMarkup, result *bool) error {
	if err := api.request(RequestMethodEditMessageReplyMarkup, params, &result); err != nil {
		return err
	}

	return nil
}


func (api Api) DeleteMessage(params RequestParamsDeleteMessage, result *bool) error {
	if err := api.request(RequestMethodDeleteMessage, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SendSticker(params RequestParamsSendSticker, result *Message) error {
	if err := api.request(RequestMethodSendSticker, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetStickerSet(params RequestParamsGetStickerSet, result *StickerSet) error {
	if err := api.request(RequestMethodGetStickerSet, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetCustomEmojiStickers(params RequestParamsGetCustomEmojiStickers, result *Sticker) error {
	if err := api.request(RequestMethodGetCustomEmojiStickers, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) UploadStickerFile(params RequestParamsUploadStickerFile, result *File) error {
	if err := api.request(RequestMethodUploadStickerFile, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) CreateNewStickerSet(params RequestParamsCreateNewStickerSet, result *bool) error {
	if err := api.request(RequestMethodCreateNewStickerSet, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) AddStickerToSet(params RequestParamsAddStickerToSet, result *bool) error {
	if err := api.request(RequestMethodAddStickerToSet, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetStickerPositionInSet(params RequestParamsSetStickerPositionInSet, result *bool) error {
	if err := api.request(RequestMethodSetStickerPositionInSet, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) DeleteStickerFromSet(params RequestParamsDeleteStickerFromSet, result *bool) error {
	if err := api.request(RequestMethodDeleteStickerFromSet, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetStickerEmojiList(params RequestParamsSetStickerEmojiList, result *bool) error {
	if err := api.request(RequestMethodSetStickerEmojiList, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetStickerKeywords(params RequestParamsSetStickerKeywords, result *bool) error {
	if err := api.request(RequestMethodSetStickerKeywords, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetStickerMaskPosition(params RequestParamsSetStickerMaskPosition, result *bool) error {
	if err := api.request(RequestMethodSetStickerMaskPosition, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetStickerSetTitle(params RequestParamsSetStickerSetTitle, result *bool) error {
	if err := api.request(RequestMethodSetStickerSetTitle, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetStickerSetThumbnail(params RequestParamsSetStickerSetThumbnail, result *bool) error {
	if err := api.request(RequestMethodSetStickerSetThumbnail, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetCustomEmojiStickerSetThumbnail(params RequestParamsSetCustomEmojiStickerSetThumbnail, result *bool) error {
	if err := api.request(RequestMethodSetCustomEmojiStickerSetThumbnail, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) DeleteStickerSet(params RequestParamsDeleteStickerSet, result *bool) error {
	if err := api.request(RequestMethodDeleteStickerSet, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) AnswerInlineQuery(params RequestParamsAnswerInlineQuery, result *bool) error {
	if err := api.request(RequestMethodAnswerInlineQuery, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) AnswerWebAppQuery(params RequestParamsAnswerWebAppQuery, result *SentWebAppMessage) error {
	if err := api.request(RequestMethodAnswerWebAppQuery, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SendInvoice(params RequestParamsSendInvoice, result *Message) error {
	if err := api.request(RequestMethodSendInvoice, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) CreateInvoiceLink(params RequestParamsCreateInvoiceLink, result *string) error {
	if err := api.request(RequestMethodCreateInvoiceLink, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) AnswerShippingQuery(params RequestParamsAnswerShippingQuery, result *bool) error {
	if err := api.request(RequestMethodAnswerShippingQuery, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) AnswerPreCheckoutQuery(params RequestParamsAnswerPreCheckoutQuery, result *bool) error {
	if err := api.request(RequestMethodAnswerPreCheckoutQuery, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetPassportDataErrors(params RequestParamsSetPassportDataErrors, result *bool) error {
	if err := api.request(RequestMethodSetPassportDataErrors, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SendGame(params RequestParamsSendGame, result *Message) error {
	if err := api.request(RequestMethodSendGame, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetGameScore(params RequestParamsSetGameScore, result *Message) error {
	if err := api.request(RequestMethodSetGameScore, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) SetInlineGameScore(params RequestParamsSetGameScore, result *bool) error {
	if err := api.request(RequestMethodSetGameScore, params, &result); err != nil {
		return err
	}

	return nil
}

func (api Api) GetGameHighScores(params RequestParamsGetGameHighScores, result *[]GameHighScore) error {
	if err := api.request(RequestMethodGetGameHighScores, params, &result); err != nil {
		return err
	}

	return nil
}

package api

// https://core.telegram.org/bots/api#getupdates
type RequestParamsGetUpdates struct {
	Offset         int64    `json:"offset,omitempty"`
	Limit          int64    `json:"limit,omitempty"`
	Timeout        int64    `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// https://core.telegram.org/bots/api#setwebhook
type RequestParamsSetWebHook struct {
	Url                string     `json:"url"`
	Certificate        *InputFile `json:"certificate,omitempty"`
	IpAddress          int64      `json:"ip_address,omitempty"`
	MaxConnections     int64      `json:"max_connections,omitempty"`
	AllowedUpdates     []string   `json:"allowed_updates,omitempty"`
	DropPendingUpdates bool       `json:"drop_pending_updates,omitempty"`
	SecretToken        string     `json:"secret_token,omitempty"`
}

// https://core.telegram.org/bots/api#deletewebhook
type RequestParamsDeleteWebHook struct {
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

// https://core.telegram.org/bots/api#getwebhookinfo
type RequestParamsGetWebHookInfo struct{}

// https://core.telegram.org/bots/api#getme
type RequestParamsGetMe struct{}

// https://core.telegram.org/bots/api#logout
type RequestParamsLogOut struct {}

// https://core.telegram.org/bots/api#close
type RequestParamsClose struct {}

// https://core.telegram.org/bots/api#sendmessage
type RequestParamsSendMessage struct {
	ChatId                   int64           `json:"chat_id"`
	MessageThreadId          int64           `json:"message_thread_id,omitempty"`
	Text                     string          `json:"text"`
	ParseMode                string          `json:"parse_mode,omitempty"`
	Entities                 []MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview    bool            `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool            `json:"disable_notification,omitempty"`
	ProtectContent           bool            `json:"protect_content,omitempty"`
	ReplyToMessageId         int64           `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply`json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#forwardmessage
type RequestParamsForwardMessage struct {
	ChatId              int64 `json:"chat_id"`
	MessageThreadId     int64 `json:"message_thread_id,omitempty"`
	FromChatId          int64 `json:"from_chat_id"`
	DisableNotification bool  `json:"disable_notification,omitempty"`
	ProtectContent      bool  `json:"protect_content,omitempty"`
	MessageId           int64 `json:"message_id"`
}

// https://core.telegram.org/bots/api#copymessage
type RequestParamsCopyMessage struct {
	ChatId                   int64           `json:"chat_id"`
	MessageThreadId          int64           `json:"message_thread_id,omitempty"`
	FromChatId               int64           `json:"from_chat_id"`
	MessageId                int64           `json:"message_id"`
	Caption                  string          `json:"caption,omitempty"`
	ParseMode                string          `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity `json:"caption_entities,omitempty"`
	DisableNotification      bool            `json:"disable_notification,omitempty"`
	ProtectContent           bool            `json:"protect_content,omitempty"`
	ReplyToMessageId         int64           `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply`json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendphoto
type RequestParamsSendPhoto struct {
	ChatId                   int64           `json:"chat_id"`
	MessageThreadId          int64           `json:"message_thread_id,omitempty"`
	Photo                    InputFile       `json:"photo"`
	Caption                  string          `json:"caption,omitempty"`
	ParseMode                string          `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler               bool            `json:"has_spoiler,omitempty"`
	DisableNotification      bool            `json:"disable_notification,omitempty"`
	ProtectContent           bool            `json:"protect_content,omitempty"`
	ReplyToMessageId         int64           `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply`json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendaudio
type RequestParamsSendAudio struct {
	ChatId                   int64           `json:"chat_id"`
	MessageThreadId          int64           `json:"message_thread_id,omitempty"`
	Audio                    InputFile       `json:"photo"`
	Caption                  string          `json:"caption,omitempty"`
	ParseMode                string          `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler               bool            `json:"has_spoiler,omitempty"`
	DisableNotification      bool            `json:"disable_notification,omitempty"`
	ProtectContent           bool            `json:"protect_content,omitempty"`
	ReplyToMessageId         int64           `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply`json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#senddocument
type RequestParamsSendDocument struct {
	ChatId          int64 `json:"chat_id"`
	MessageThreadId int64 `json:"message_thread_id,omitempty"`
	// Document InputFile or String`json:"document"`
	// Thumbnail InputFile or String`json:"thumbnail,omitempty"`
	Caption                     string          `json:"caption,omitempty"`
	ParseMode                   string          `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool            `json:"disable_content_type_detection,omitempty"`
	DisableNotification         bool            `json:"disable_notification,omitempty"`
	ProtectContent              bool            `json:"protect_content,omitempty"`
	ReplyToMessageId            int64           `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply    bool            `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply`json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendvideo
type RequestParamsSendVideo struct {
	ChatId          int64 `json:"chat_id"`
	MessageThreadId int64 `json:"message_thread_id,omitempty"`
	// Video InputFile or String`json:"video"`
	Duration int64 `json:"duration,omitempty"`
	Width    int64 `json:"width,omitempty"`
	Height   int64 `json:"height,omitempty"`
	// Thumbnail InputFile or String`json:"thumbnail,omitempty"`
	Caption                  string          `json:"caption,omitempty"`
	ParseMode                string          `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler               bool            `json:"has_spoiler,omitempty"`
	SupportsStreaming        bool            `json:"supports_streaming,omitempty"`
	DisableNotification      bool            `json:"disable_notification,omitempty"`
	ProtectContent           bool            `json:"protect_content,omitempty"`
	ReplyToMessageId         int64           `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply`json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendanimation
type RequestParamsSendAnimation struct {
	ChatId          int64 `json:"chat_id"`
	MessageThreadId int64 `json:"message_thread_id,omitempty"`
	// Animation InputFile or String`json:"animation"`
	Duration int64 `json:"duration,omitempty"`
	Width    int64 `json:"width,omitempty"`
	Height   int64 `json:"height,omitempty"`
	// Thumbnail InputFile or String`json:"thumbnail,omitempty"`
	Caption                  string          `json:"caption,omitempty"`
	ParseMode                string          `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler               bool            `json:"has_spoiler,omitempty"`
	DisableNotification      bool            `json:"disable_notification,omitempty"`
	ProtectContent           bool            `json:"protect_content,omitempty"`
	ReplyToMessageId         int64           `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply`json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendvoice
type RequestParamsSendVoice struct {
	ChatId          int64 `json:"chat_id"`
	MessageThreadId int64 `json:"message_thread_id,omitempty"`
	// Voice InputFile or String`json:"voice"`
	Caption                  string          `json:"caption,omitempty"`
	ParseMode                string          `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity `json:"caption_entities,omitempty"`
	Duration                 int64           `json:"duration,omitempty"`
	DisableNotification      bool            `json:"disable_notification,omitempty"`
	ProtectContent           bool            `json:"protect_content,omitempty"`
	ReplyToMessageId         int64           `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply`json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendvideonote
type RequestParamsSendVideoNote struct {
	ChatId          int64 `json:"chat_id"`
	MessageThreadId int64 `json:"message_thread_id,omitempty"`
	// VideoNote InputFile or String`json:"video_note"`
	Duration int64 `json:"duration,omitempty"`
	Length   int64 `json:"length,omitempty"`
	// Thumbnail InputFile or String`json:"thumbnail,omitempty"`
	DisableNotification      bool  `json:"disable_notification,omitempty"`
	ProtectContent           bool  `json:"protect_content,omitempty"`
	ReplyToMessageId         int64 `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool  `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply`json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendmediagroup
type RequestParamsSendMediaGroup struct {
	ChatId          int64 `json:"chat_id"`
	MessageThreadId int64 `json:"message_thread_id,omitempty"`
	// Media []InputMediaAudio, InputMediaDocument, InputMediaPhoto and InputMediaVideo`json:"media"`
	DisableNotification      bool  `json:"disable_notification,omitempty"`
	ProtectContent           bool  `json:"protect_content,omitempty"`
	ReplyToMessageId         int64 `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool  `json:"allow_sending_without_reply,omitempty"`
}

// https://core.telegram.org/bots/api#sendlocation
type RequestParamsSendLocation struct {
	ChatId                   int64   `json:"chat_id"`
	MessageThreadId          int64   `json:"message_thread_id,omitempty"`
	Latitude                 float64 `json:"latitude"`
	Longitude                float64 `json:"longitude"`
	HorizontalAccuracy       float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod               int64   `json:"live_period,omitempty"`
	Heading                  int64   `json:"heading,omitempty"`
	ProximityAlertRadius     int64   `json:"proximity_alert_radius,omitempty"`
	DisableNotification      bool    `json:"disable_notification,omitempty"`
	ProtectContent           bool    `json:"protect_content,omitempty"`
	ReplyToMessageId         int64   `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool    `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply`json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendvenue
type RequestParamsSendVenue struct {
	ChatId                   int64   `json:"chat_id"`
	MessageThreadId          int64   `json:"message_thread_id,omitempty"`
	Latitude                 float64 `json:"latitude"`
	Longitude                float64 `json:"longitude"`
	Title                    string  `json:"title"`
	Address                  string  `json:"address"`
	FoursquareId             string  `json:"foursquare_id,omitempty"`
	FoursquareType           string  `json:"foursquare_type,omitempty"`
	GooglePlaceId            string  `json:"google_place_id,omitempty"`
	GooglePlaceType          string  `json:"google_place_type,omitempty"`
	DisableNotification      bool    `json:"disable_notification,omitempty"`
	ProtectContent           bool    `json:"protect_content,omitempty"`
	ReplyToMessageId         int64   `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool    `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply`json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendcontact
type RequestParamsSendContact struct {
	ChatId                   int64  `json:"chat_id"`
	MessageThreadId          int64  `json:"message_thread_id,omitempty"`
	PhoneNumber              string `json:"phone_number"`
	FirstName                string `json:"first_name"`
	LastName                 string `json:"last_name,omitempty"`
	Vcard                    string `json:"vcard,omitempty"`
	DisableNotification      bool   `json:"disable_notification,omitempty"`
	ProtectContent           bool   `json:"protect_content,omitempty"`
	ReplyToMessageId         int64  `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool   `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply`json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendpoll
type RequestParamsSendPoll struct {
	ChatId                   int64           `json:"chat_id"`
	MessageThreadId          int64           `json:"message_thread_id,omitempty"`
	Question                 string          `json:"question"`
	Options                  []string        `json:"options"`
	IsAnonymous              bool            `json:"is_anonymous,omitempty"`
	Type                     string          `json:"type,omitempty"`
	AllowsMultipleAnswers    bool            `json:"allows_multiple_answers,omitempty"`
	CorrectOptionId          int64           `json:"correct_option_id,omitempty"`
	Explanation              string          `json:"explanation,omitempty"`
	ExplanationParseMode     string          `json:"explanation_parse_mode,omitempty"`
	ExplanationEntities      []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod               int64           `json:"open_period,omitempty"`
	CloseDate                int64           `json:"close_date,omitempty"`
	IsClosed                 bool            `json:"is_closed,omitempty"`
	DisableNotification      bool            `json:"disable_notification,omitempty"`
	ProtectContent           bool            `json:"protect_content,omitempty"`
	ReplyToMessageId         int64           `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply`json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#senddice
type RequestParamsSendDice struct {
	ChatId                   int64  `json:"chat_id"`
	MessageThreadId          int64  `json:"message_thread_id,omitempty"`
	Emoji                    string `json:"emoji,omitempty"`
	DisableNotification      bool   `json:"disable_notification,omitempty"`
	ProtectContent           bool   `json:"protect_content,omitempty"`
	ReplyToMessageId         int64  `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool   `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply`json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendchataction
type RequestParamsSendChatAction struct {
	ChatId                   int64  `json:"chat_id"`
	MessageThreadId          int64  `json:"message_thread_id,omitempty"`
	action                   string `json:"action,omitempty"`
}

// https://core.telegram.org/bots/api#getuserprofilephotos
type RequestParamsGetUserProfilePhotos struct {
	UserId int64 `json:"user_id"`
	Offset int64 `json:"offset,omitempty"`
	Limit  int64 `json:"limit,omitempty"`
}

// https://core.telegram.org/bots/api#getfile
type RequestParamsGetFile struct {
	FileId string `json:"file_id"`
}

// https://core.telegram.org/bots/api#banchatmember
type RequestParamsBanChatMember struct {
	ChatId         int64 `json:"chat_id"`
	UserId         int64 `json:"user_id"`
	UntilDate      int64 `json:"until_date,omitempty"`
	RevokeMessages bool  `json:"revoke_messages,omitempty"`
}

// https://core.telegram.org/bots/api#unbanchatmember
type RequestParamsUnbanChatMember struct {
	ChatId       int64 `json:"chat_id"`
	UserId       int64 `json:"user_id"`
	OnlyIfBanned bool  `json:"only_if_banned,omitempty"`
}

// https://core.telegram.org/bots/api#restrictchatmember
type RequestParamsRestrictChatMember struct {
	ChatId                        int64            `json:"chat_id"`
	UserId                        int64            `json:"user_id"`
	Permissions                   *ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions bool             `json:"use_independent_chat_permissions,omitempty"`
	UntilDate                     int64            `json:"until_date,omitempty"`
}

// https://core.telegram.org/bots/api#promotechatmember
type RequestParamsPromoteChatMember struct {
	ChatId              int64 `json:"chat_id"`
	UserId              int64 `json:"user_id"`
	IsAnonymous         bool  `json:"is_anonymous,omitempty"`
	CanManageChat       bool  `json:"can_manage_chat,omitempty"`
	CanPostMessages     bool  `json:"can_post_messages,omitempty"`
	CanEditMessages     bool  `json:"can_edit_messages,omitempty"`
	CanDeleteMessages   bool  `json:"can_delete_messages,omitempty"`
	CanManageVideoChats bool  `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers  bool  `json:"can_restrict_members,omitempty"`
	CanPromoteMembers   bool  `json:"can_promote_members,omitempty"`
	CanChangeInfo       bool  `json:"can_change_info,omitempty"`
	CanInviteUsers      bool  `json:"can_invite_users,omitempty"`
	CanPinMessages      bool  `json:"can_pin_messages,omitempty"`
	CanManageTopics     bool  `json:"can_manage_topics,omitempty"`
}

// https://core.telegram.org/bots/api#setchatadministratorcustomtitle
type RequestParamsSetChatAdministratorCustomTitle struct {
	ChatId      int64  `json:"chat_id"`
	UserId      int64  `json:"user_id"`
	CustomTitle string `json:"custom_title"`
}

// https://core.telegram.org/bots/api#banchatsenderchat
type RequestParamsBanChatSenderChat struct {
	ChatId       int64 `json:"chat_id"`
	SenderChatId int64 `json:"sender_chat_id"`
}

// https://core.telegram.org/bots/api#unbanchatsenderchat
type RequestParamsUnbanChatSenderChat struct {
	ChatId       int64 `json:"chat_id"`
	SenderChatId int64 `json:"sender_chat_id"`
}

// https://core.telegram.org/bots/api#setchatpermissions
type RequestParamsSetChatPermissions struct {
	ChatId                        int64            `json:"chat_id"`
	Permissions                   *ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions bool             `json:"use_independent_chat_permissions,omitempty"`
}

// https://core.telegram.org/bots/api#exportchatinvitelink
type RequestParamsExportChatInviteLink struct {
	ChatId int64 `json:"chat_id"`
}

// https://core.telegram.org/bots/api#createchatinvitelink
type RequestParamsCreateChatInviteLink struct {
	ChatId             int64  `json:"chat_id"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int64  `json:"expire_date,omitempty"`
	MemberLimit        int64  `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

// https://core.telegram.org/bots/api#editchatinvitelink
type RequestParamsEditChatInviteLink struct {
	ChatId             int64  `json:"chat_id"`
	InviteLink         string `json:"invite_link"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int64  `json:"expire_date,omitempty"`
	MemberLimit        int64  `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

// https://core.telegram.org/bots/api#revokechatinvitelink
type RequestParamsRevokeChatInviteLink struct {
	ChatId     int64  `json:"chat_id"`
	InviteLink string `json:"invite_link"`
}

// https://core.telegram.org/bots/api#approvechatjoinrequest
type RequestParamsApproveChatJoinRequest struct {
	ChatId int64 `json:"chat_id"`
	UserId int64 `json:"user_id"`
}

// https://core.telegram.org/bots/api#declinechatjoinrequest
type RequestParamsDeclineChatJoinRequest struct {
	ChatId int64 `json:"chat_id"`
	UserId int64 `json:"user_id"`
}

// https://core.telegram.org/bots/api#setchatphoto
type RequestParamsSetChatPhoto struct {
	ChatId int64      `json:"chat_id"`
	Photo  *InputFile `json:"photo"`
}

// https://core.telegram.org/bots/api#deletechatphoto
type RequestParamsDeleteChatPhoto struct {
	ChatId int64 `json:"chat_id"`
}

// https://core.telegram.org/bots/api#setchattitle
type RequestParamsSetChatTitle struct {
	ChatId int64  `json:"chat_id"`
	Title  string `json:"title"`
}

// https://core.telegram.org/bots/api#setchatdescription
type RequestParamsSetChatDescription struct {
	ChatId      int64  `json:"chat_id"`
	Description string `json:"description,omitempty"`
}

// https://core.telegram.org/bots/api#pinchatmessage
type RequestParamsPinChatMessage struct {
	ChatId              int64 `json:"chat_id"`
	MessageId           int64 `json:"message_id"`
	DisableNotification bool  `json:"disable_notification,omitempty"`
}

// https://core.telegram.org/bots/api#unpinchatmessage
type RequestParamsUnpinChatMessage struct {
	ChatId    int64 `json:"chat_id"`
	MessageId int64 `json:"message_id,omitempty"`
}

// https://core.telegram.org/bots/api#unpinallchatmessages
type RequestParamsUnpinAllChatMessages struct {
	ChatId int64 `json:"chat_id"`
}

// https://core.telegram.org/bots/api#leavechat
type RequestParamsLeaveChat struct {
	ChatId int64 `json:"chat_id"`
}

// https://core.telegram.org/bots/api#getchat
type RequestParamsGetChat struct {
	ChatId int64 `json:"chat_id"`
}

// https://core.telegram.org/bots/api#getchatadministrators
type RequestParamsGetChatAdministrators struct {
	ChatId int64 `json:"chat_id"`
}

// https://core.telegram.org/bots/api#getchatmembercount
type RequestParamsGetChatMemberCount struct {
	ChatId int64 `json:"chat_id"`
}

// https://core.telegram.org/bots/api#getchatmember
type RequestParamsGetChatMember struct {
	ChatId int64 `json:"chat_id"`
	UserId int64 `json:"user_id"`
}

// https://core.telegram.org/bots/api#setchatstickerset
type RequestParamsSetChatStickerSet struct {
	ChatId         int64  `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
}

// https://core.telegram.org/bots/api#deletechatstickerset
type RequestParamsDeleteChatStickerSet struct {
	ChatId int64 `json:"chat_id"`
}

// https://core.telegram.org/bots/api#getforumtopiciconstickers
type RequestParamsGetForumTopicIconStickers struct {
}

// https://core.telegram.org/bots/api#createforumtopic
type RequestParamsCreateForumTopic struct {
	ChatId            int64  `json:"chat_id"`
	Name              string `json:"name"`
	IconColor         int64  `json:"icon_color,omitempty"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

// https://core.telegram.org/bots/api#editforumtopic
type RequestParamsEditForumTopic struct {
	ChatId            int64  `json:"chat_id"`
	MessageThreadId   int64  `json:"message_thread_id"`
	Name              string `json:"name,omitempty"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

// https://core.telegram.org/bots/api#closeforumtopic
type RequestParamsCloseForumTopic struct {
	ChatId          int64 `json:"chat_id"`
	MessageThreadId int64 `json:"message_thread_id"`
}

// https://core.telegram.org/bots/api#reopenforumtopic
type RequestParamsReopenForumTopic struct {
	ChatId          int64 `json:"chat_id"`
	MessageThreadId int64 `json:"message_thread_id"`
}

// https://core.telegram.org/bots/api#deleteforumtopic
type RequestParamsDeleteForumTopic struct {
	ChatId          int64 `json:"chat_id"`
	MessageThreadId int64 `json:"message_thread_id"`
}

// https://core.telegram.org/bots/api#unpinallforumtopicmessages
type RequestParamsUnpinAllForumTopicMessages struct {
	ChatId          int64 `json:"chat_id"`
	MessageThreadId int64 `json:"message_thread_id"`
}

// https://core.telegram.org/bots/api#editgeneralforumtopic
type RequestParamsEditGeneralForumTopic struct {
	ChatId int64  `json:"chat_id"`
	Name   string `json:"name"`
}

// https://core.telegram.org/bots/api#closegeneralforumtopic
type RequestParamsCloseGeneralForumTopic struct {
	ChatId int64 `json:"chat_id"`
}

// https://core.telegram.org/bots/api#reopengeneralforumtopic
type RequestParamsReopenGeneralForumTopic struct {
	ChatId int64 `json:"chat_id"`
}

// https://core.telegram.org/bots/api#hidegeneralforumtopic
type RequestParamsHideGeneralForumTopic struct {
	ChatId int64 `json:"chat_id"`
}

// https://core.telegram.org/bots/api#unhidegeneralforumtopic
type RequestParamsUnhideGeneralForumTopic struct {
	ChatId int64 `json:"chat_id"`
}

// https://core.telegram.org/bots/api#unpinallgeneralforumtopicmessages
type RequestParamsUnpinAllGeneralForumTopicMessages struct {
	ChatId int64 `json:"chat_id"`
}

// https://core.telegram.org/bots/api#answercallbackquery
type RequestParamsAnswerCallbackQuery struct {
}

// https://core.telegram.org/bots/api#setmycommands
type RequestParamsSetMyCommands struct {
	Commands     []BotCommand     `json:"commands"`
	Scope        *BotCommandScope `json:"scope,omitempty"`
	LanguageCode string           `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#deletemycommands
type RequestParamsDeleteMyCommands struct {
	Scope        *BotCommandScope `json:"scope,omitempty"`
	LanguageCode string           `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#getmycommands
type RequestParamsGetMyCommands struct {
	Scope        *BotCommandScope `json:"scope,omitempty"`
	LanguageCode string           `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#setmyname
type RequestParamsSetMyName struct {
	Name         string `json:"name,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#getmyname
type RequestParamsGetMyName struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#setmydescription
type RequestParamsSetMyDescription struct {
	Description  string `json:"description,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#getmydescription
type RequestParamsGetMyDescription struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#setmyshortdescription
type RequestParamsSetMyShortDescription struct {
	ShortDescription string `json:"short_description,omitempty"`
	LanguageCode     string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#getmyshortdescription
type RequestParamsGetMyShortDescription struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#setchatmenubutton
type RequestParamsSetChatMenuButton struct {
	ChatId     int64       `json:"chat_id,omitempty"`
	MenuButton *MenuButton `json:"menu_button,omitempty"`
}

// https://core.telegram.org/bots/api#getchatmenubutton
type RequestParamsGetChatMenuButton struct {
	ChatId int64 `json:"chat_id,omitempty"`
}

// https://core.telegram.org/bots/api#setmydefaultadministratorrights
type RequestParamsSetMyDefaultAdministratorRights struct {
	Rights      *ChatAdministratorRights `json:"rights,omitempty"`
	ForChannels bool                     `json:"for_channels,omitempty"`
}

// https://core.telegram.org/bots/api#getmydefaultadministratorrights
type RequestParamsGetMyDefaultAdministratorRights struct {
	ForChannels bool `json:"for_channels,omitempty"`
}

// https://core.telegram.org/bots/api#editmessagetext
type RequestParamsEditMessageText struct {
	ChatId                int64                 `json:"chat_id,omitempty"`
	MessageId             int64                 `json:"message_id,omitempty"`
	InlineMessageId       string                `json:"inline_message_id,omitempty"`
	Text                  string                `json:"text"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	Entities              []MessageEntity       `json:"entities,omitempty"`
	DisableWebPagePreview bool                  `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#editmessagecaption
type RequestParamsEditMessageCaption struct {
	ChatId          int64                 `json:"chat_id,omitempty"`
	MessageId       int64                 `json:"message_id,omitempty"`
	InlineMessageId string                `json:"inline_message_id,omitempty"`
	Caption         string                `json:"caption,omitempty"`
	ParseMode       string                `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#editmessagemedia
type RequestParamsEditMessageMedia struct {
	ChatId          int64                 `json:"chat_id,omitempty"`
	MessageId       int64                 `json:"message_id,omitempty"`
	InlineMessageId string                `json:"inline_message_id,omitempty"`
	Media           *InputMedia           `json:"media"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#editmessagelivelocation
type RequestParamsEditMessageLiveLocation struct {
	ChatId               int64                 `json:"chat_id,omitempty"`
	MessageId            int64                 `json:"message_id,omitempty"`
	InlineMessageId      string                `json:"inline_message_id,omitempty"`
	Latitude             float64               `json:"latitude"`
	Longitude            float64               `json:"longitude"`
	HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`
	Heading              int64                 `json:"heading,omitempty"`
	ProximityAlertRadius int64                 `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#stopmessagelivelocation
type RequestParamsStopMessageLiveLocation struct {
	ChatId          int64                 `json:"chat_id,omitempty"`
	MessageId       int64                 `json:"message_id,omitempty"`
	InlineMessageId string                `json:"inline_message_id,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#editmessagereplymarkup
type RequestParamsEditMessageReplyMarkup struct {
	ChatId          int64                 `json:"chat_id,omitempty"`
	MessageId       int64                 `json:"message_id,omitempty"`
	InlineMessageId string                `json:"inline_message_id,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#stoppoll
type RequestParamsStopPoll struct {
	ChatId      int64                 `json:"chat_id"`
	MessageId   int64                 `json:"message_id"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#deletemessage
type RequestParamsDeleteMessage struct {
	ChatId    int64 `json:"chat_id"`
	MessageId int64 `json:"message_id"`
}

// https://core.telegram.org/bots/api#sendsticker
type RequestParamsSendSticker struct {
	ChatId          int64 `json:"chat_id"`
	MessageThreadId int64 `json:"message_thread_id,omitempty"`
	// Sticker InputFile or String`json:"sticker"`
	Emoji                    string `json:"emoji,omitempty"`
	DisableNotification      bool   `json:"disable_notification,omitempty"`
	ProtectContent           bool   `json:"protect_content,omitempty"`
	ReplyToMessageId         int64  `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool   `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply`json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#getstickerset
type RequestParamsGetStickerSet struct {
	Name string `json:"name"`
}

// https://core.telegram.org/bots/api#getcustomemojistickers
type RequestParamsGetCustomEmojiStickers struct {
	CustomEmojiIds []string `json:"custom_emoji_ids"`
}

// https://core.telegram.org/bots/api#uploadstickerfile
type RequestParamsUploadStickerFile struct {
	UserId        int64      `json:"user_id"`
	Sticker       *InputFile `json:"sticker"`
	StickerFormat string     `json:"sticker_format"`
}

// https://core.telegram.org/bots/api#createnewstickerset
type RequestParamsCreateNewStickerSet struct {
	UserId          int64          `json:"user_id"`
	Name            string         `json:"name"`
	Title           string         `json:"title"`
	Stickers        []InputSticker `json:"stickers"`
	StickerFormat   string         `json:"sticker_format"`
	StickerType     string         `json:"sticker_type,omitempty"`
	NeedsRepainting bool           `json:"needs_repainting,omitempty"`
}

// https://core.telegram.org/bots/api#addstickertoset
type RequestParamsAddStickerToSet struct {
	UserId  int64         `json:"user_id"`
	Name    string        `json:"name"`
	Sticker *InputSticker `json:"sticker"`
}

// https://core.telegram.org/bots/api#setstickerpositioninset
type RequestParamsSetStickerPositionInSet struct {
	Sticker  string `json:"sticker"`
	Position int64  `json:"position"`
}

// https://core.telegram.org/bots/api#deletestickerfromset
type RequestParamsDeleteStickerFromSet struct {
	Sticker string `json:"sticker"`
}

// https://core.telegram.org/bots/api#setstickeremojilist
type RequestParamsSetStickerEmojiList struct {
	Sticker   string   `json:"sticker"`
	EmojiList []string `json:"emoji_list"`
}

// https://core.telegram.org/bots/api#setstickerkeywords
type RequestParamsSetStickerKeywords struct {
	Sticker  string   `json:"sticker"`
	Keywords []string `json:"keywords,omitempty"`
}

// https://core.telegram.org/bots/api#setstickermaskposition
type RequestParamsSetStickerMaskPosition struct {
	Sticker      string        `json:"sticker"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}

// https://core.telegram.org/bots/api#setstickersettitle
type RequestParamsSetStickerSetTitle struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

// https://core.telegram.org/bots/api#setstickersetthumbnail
type RequestParamsSetStickerSetThumbnail struct {
	Name   string `json:"name"`
	UserId int64  `json:"user_id"`
	// Thumbnail InputFile or String`json:"thumbnail,omitempty"`
}

// https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
type RequestParamsSetCustomEmojiStickerSetThumbnail struct {
	Name          string `json:"name"`
	CustomEmojiId string `json:"custom_emoji_id,omitempty"`
}

// https://core.telegram.org/bots/api#deletestickerset
type RequestParamsDeleteStickerSet struct {
	Name string `json:"name"`
}

// https://core.telegram.org/bots/api#answerinlinequery
type RequestParamsAnswerInlineQuery struct {
	InlineQueryId string                    `json:"inline_query_id"`
	Results       []InlineQueryResult       `json:"results"`
	CacheTime     int64                     `json:"cache_time,omitempty"`
	IsPersonal    bool                      `json:"is_personal,omitempty"`
	NextOffset    string                    `json:"next_offset,omitempty"`
	Button        *InlineQueryResultsButton `json:"button,omitempty"`
}

// https://core.telegram.org/bots/api#answerwebappquery
type RequestParamsAnswerWebAppQuery struct {
	WebAppQueryId string             `json:"web_app_query_id"`
	Result        *InlineQueryResult `json:"result"`
}

// https://core.telegram.org/bots/api#sendinvoice
type RequestParamsSendInvoice struct {
	ChatId                    int64                 `json:"chat_id"`
	MessageThreadId           int64                 `json:"message_thread_id,omitempty"`
	Title                     string                `json:"title"`
	Description               string                `json:"description"`
	Payload                   string                `json:"payload"`
	ProviderToken             string                `json:"provider_token"`
	Currency                  string                `json:"currency"`
	Prices                    []LabeledPrice        `json:"prices"`
	MaxTipAmount              int64                 `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int                 `json:"suggested_tip_amounts,omitempty"`
	StartParameter            string                `json:"start_parameter,omitempty"`
	ProviderData              string                `json:"provider_data,omitempty"`
	PhotoUrl                  string                `json:"photo_url,omitempty"`
	PhotoSize                 int64                 `json:"photo_size,omitempty"`
	PhotoWidth                int64                 `json:"photo_width,omitempty"`
	PhotoHeight               int64                 `json:"photo_height,omitempty"`
	NeedName                  bool                  `json:"need_name,omitempty"`
	NeedPhoneNumber           bool                  `json:"need_phone_number,omitempty"`
	NeedEmail                 bool                  `json:"need_email,omitempty"`
	NeedShippingAddress       bool                  `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool                  `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool                  `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool                  `json:"is_flexible,omitempty"`
	DisableNotification       bool                  `json:"disable_notification,omitempty"`
	ProtectContent            bool                  `json:"protect_content,omitempty"`
	ReplyToMessageId          int64                 `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply  bool                  `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup               *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#createinvoicelink
type RequestParamsCreateInvoiceLink struct {
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             string         `json:"provider_token"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	MaxTipAmount              int64          `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int          `json:"suggested_tip_amounts,omitempty"`
	ProviderData              string         `json:"provider_data,omitempty"`
	PhotoUrl                  string         `json:"photo_url,omitempty"`
	PhotoSize                 int64          `json:"photo_size,omitempty"`
	PhotoWidth                int64          `json:"photo_width,omitempty"`
	PhotoHeight               int64          `json:"photo_height,omitempty"`
	NeedName                  bool           `json:"need_name,omitempty"`
	NeedPhoneNumber           bool           `json:"need_phone_number,omitempty"`
	NeedEmail                 bool           `json:"need_email,omitempty"`
	NeedShippingAddress       bool           `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool           `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool           `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool           `json:"is_flexible,omitempty"`
}

// https://core.telegram.org/bots/api#answershippingquery
type RequestParamsAnswerShippingQuery struct {
	ShippingQueryId string           `json:"shipping_query_id"`
	Ok              bool             `json:"ok"`
	ShippingOptions []ShippingOption `json:"shipping_options,omitempty"`
	ErrorMessage    string           `json:"error_message,omitempty"`
}

// https://core.telegram.org/bots/api#answerprecheckoutquery
type RequestParamsAnswerPreCheckoutQuery struct {
	PreCheckoutQueryId string `json:"pre_checkout_query_id"`
	Ok                 bool   `json:"ok"`
	ErrorMessage       string `json:"error_message,omitempty"`
}

// https://core.telegram.org/bots/api#setpassportdataerrors
type RequestParamsSetPassportDataErrors struct {
}

// https://core.telegram.org/bots/api#sendgame
type RequestParamsSendGame struct {
	ChatId                   int64                 `json:"chat_id"`
	MessageThreadId          int64                 `json:"message_thread_id,omitempty"`
	GameShortName            string                `json:"game_short_name"`
	DisableNotification      bool                  `json:"disable_notification,omitempty"`
	ProtectContent           bool                  `json:"protect_content,omitempty"`
	ReplyToMessageId         int64                 `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                  `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#setgamescore
type RequestParamsSetGameScore struct {
	UserId             int64  `json:"user_id"`
	Score              int64  `json:"score"`
	Force              bool   `json:"force,omitempty"`
	DisableEditMessage bool   `json:"disable_edit_message,omitempty"`
	ChatId             int64  `json:"chat_id,omitempty"`
	MessageId          int64  `json:"message_id,omitempty"`
	InlineMessageId    string `json:"inline_message_id,omitempty"`
}

// https://core.telegram.org/bots/api#getgamehighscores
type RequestParamsGetGameHighScores struct {
	UserId             int64  `json:"user_id"`
	ChatId             int64  `json:"chat_id,omitempty"`
	MessageId          int64  `json:"message_id,omitempty"`
	InlineMessageId    string `json:"inline_message_id,omitempty"`
}

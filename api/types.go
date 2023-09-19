package api

// https://core.telegram.org/bots/api#update
type Update struct {
	UpdateId           int64               `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
	PollAnswer         *PollAnswer         `json:"poll_answer,omitempty"`
	MyChatMember       *ChatMemberUpdated  `json:"my_chat_member,omitempty"`
	ChatMember         *ChatMemberUpdated  `json:"chat_member,omitempty"`
	ChatJoinRequest    *ChatJoinRequest    `json:"chat_join_request,omitempty"`
}

// // https://core.telegram.org/bots/api#webhookinfo
// type WebhookInfo struct {
// 	Url                          string   `json:"url"`
// 	HasCustomCertificate         bool     `json:"has_custom_certificate"`
// 	PendingUpdateCount           int64    `json:"pending_update_count"`
// 	IpAddress                    string   `json:"ip_address,omitempty"`
// 	LastErrorDate                int64    `json:"last_error_date,omitempty"`
// 	LastErrorMessage             string   `json:"last_error_message,omitempty"`
// 	LastSynchronizationErrorDate int64    `json:"last_synchronization_error_date,omitempty"`
// 	MaxConnections               int64    `json:"max_connections,omitempty"`
// 	AllowedUpdates               []string `json:"allowed_updates,omitempty"`
// }

// https://core.telegram.org/bots/api#user
type User struct {
	Id                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name,omitempty"`
	UserName                string `json:"username,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	IsPremium               bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
}

// https://core.telegram.org/bots/api#chat
type Chat struct {
	Id                                 int64            `json:"id"`
	Type                               string           `json:"type"`
	Title                              string           `json:"title,omitempty"`
	UserName                           string           `json:"username,omitempty"`
	FirstName                          string           `json:"first_name,omitempty"`
	LastName                           string           `json:"last_name,omitempty"`
	IsForum                            bool             `json:"is_forum,omitempty"`
	Photo                              *ChatPhoto       `json:"photo,omitempty"`
	ActiveUserNames                    []string         `json:"active_usernames,omitempty"`
	EmojiStatusCustomEmojiId           string           `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate          int64            `json:"emoji_status_expiration_date,omitempty"`
	Bio                                string           `json:"bio,omitempty"`
	HasPrivateForwards                 bool             `json:"has_private_forwards,omitempty"`
	HasRestrictedVoiceAndVideoMessages bool             `json:"has_restricted_voice_and_video_messages,omitempty"`
	JoinToSendMessages                 bool             `json:"join_to_send_messages,omitempty"`
	JoinByRequest                      bool             `json:"join_by_request,omitempty"`
	Description                        string           `json:"description,omitempty"`
	InviteLink                         string           `json:"invite_link,omitempty"`
	PinnedMessage                      *Message         `json:"pinned_message,omitempty"`
	Permissions                        *ChatPermissions `json:"permissions,omitempty"`
	SlowModeDelay                      int64            `json:"slow_mode_delay,omitempty"`
	MessageAutoDeleteTime              int64            `json:"message_auto_delete_time,omitempty"`
	HasAggressiveAntiSpamEnabled       bool             `json:"has_aggressive_anti_spam_enabled,omitempty"`
	HasHiddenMembers                   bool             `json:"has_hidden_members,omitempty"`
	HasProtectedContent                bool             `json:"has_protected_content,omitempty"`
	StickerSetName                     string           `json:"sticker_set_name,omitempty"`
	CanSetStickerSet                   bool             `json:"can_set_sticker_set,omitempty"`
	LinkedChatId                       int64            `json:"linked_chat_id,omitempty"`
	Location                           *ChatLocation    `json:"location,omitempty"`
}

// https://core.telegram.org/bots/api#message
type Message struct {
	MessageId                     int64                          `json:"message_id"`
	MessageThreadId               int64                          `json:"message_thread_id,omitempty"`
	From                          *User                          `json:"from,omitempty"`
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`
	Date                          int64                          `json:"date"`
	Chat                          *Chat                          `json:"chat"`
	ForwardFrom                   *User                          `json:"forward_from,omitempty"`
	ForwardFromChat               *Chat                          `json:"forward_from_chat,omitempty"`
	ForwardFromMessageId          int64                          `json:"forward_from_message_id,omitempty"`
	ForwardSignature              string                         `json:"forward_signature,omitempty"`
	ForwardSenderName             string                         `json:"forward_sender_name,omitempty"`
	ForwardDate                   int64                          `json:"forward_date,omitempty"`
	IsTopicMessage                bool                           `json:"is_topic_message,omitempty"`
	IsAutomaticForward            bool                           `json:"is_automatic_forward,omitempty"`
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`
	ViaBot                        *User                          `json:"via_bot,omitempty"`
	EditDate                      int64                          `json:"edit_date,omitempty"`
	HasProtectedContent           bool                           `json:"has_protected_content,omitempty"`
	MediaGroupId                  string                         `json:"media_group_id,omitempty"`
	AuthorSignature               string                         `json:"author_signature,omitempty"`
	Text                          string                         `json:"text,omitempty"`
	Entities                      []MessageEntity                `json:"entities,omitempty"`
	Animation                     *Animation                     `json:"animation,omitempty"`
	Audio                         *Audio                         `json:"audio,omitempty"`
	Document                      *Document                      `json:"document,omitempty"`
	Photo                         []PhotoSize                    `json:"photo,omitempty"`
	Sticker                       *Sticker                       `json:"sticker,omitempty"`
	Story                         *Story                         `json:"story,omitempty"`
	Video                         *Video                         `json:"video,omitempty"`
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`
	Voice                         *Voice                         `json:"voice,omitempty"`
	Caption                       string                         `json:"caption,omitempty"`
	CaptionEntities               []MessageEntity                `json:"caption_entities,omitempty"`
	HasMediaSpoiler               bool                           `json:"has_media_spoiler,omitempty"`
	Contact                       *Contact                       `json:"contact,omitempty"`
	Dice                          *Dice                          `json:"dice,omitempty"`
	Game                          *Game                          `json:"game,omitempty"`
	Poll                          *Poll                          `json:"poll,omitempty"`
	Venue                         *Venue                         `json:"venue,omitempty"`
	Location                      *Location                      `json:"location,omitempty"`
	NewChatMembers                []User                         `json:"new_chat_members,omitempty"`
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`
	NewChatTitle                  string                         `json:"new_chat_title,omitempty"`
	NewChatPhoto                  []PhotoSize                    `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               bool                           `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              bool                           `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            bool                           `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatId               int64                          `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatId             int64                          `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                 *Message                       `json:"pinned_message,omitempty"`
	Invoice                       *Invoice                       `json:"invoice,omitempty"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`
	UserShared                    *UserShared                    `json:"user_shared,omitempty"`
	ChatShared                    *ChatShared                    `json:"chat_shared,omitempty"`
	ConnectedWebsite              string                         `json:"connected_website,omitempty"`
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed,omitempty"`
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created,omitempty"`
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited,omitempty"`
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed,omitempty"`
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`
	WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`
}

// // https://core.telegram.org/bots/api#messageid
// type MessageId struct {
// 	MessageId int64 `json:"message_id"`
// }

// https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type          string `json:"type"`
	Offset        int64  `json:"offset"`
	Length        int64  `json:"length"`
	Url           string `json:"url,omitempty"`
	User          *User  `json:"user,omitempty"`
	Language      string `json:"language,omitempty"`
	CustomEmojiId string `json:"custom_emoji_id,omitempty"`
}

// https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
	FileSize     int64  `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#animation
type Animation struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Width        int64      `json:"width"`
	Height       int64      `json:"height"`
	Duration     int64      `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#audio
type Audio struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Duration     int64      `json:"duration"`
	Performer    string     `json:"performer,omitempty"`
	Title        string     `json:"title,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
}

// https://core.telegram.org/bots/api#document
type Document struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#story
type Story struct {
}

// https://core.telegram.org/bots/api#video
type Video struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Width        int64      `json:"width"`
	Height       int64      `json:"height"`
	Duration     int64      `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Length       int64      `json:"length"`
	Duration     int64      `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#voice
type Voice struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Duration     int64  `json:"duration"`
	MimeType     string `json:"mime_type,omitempty"`
	FileSize     int64  `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#contact
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserId      int64  `json:"user_id,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}

// https://core.telegram.org/bots/api#dice
type Dice struct {
	Emoji string `json:"emoji"`
	Value int64  `json:"value"`
}

// https://core.telegram.org/bots/api#polloption
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int64  `json:"voter_count"`
}

// https://core.telegram.org/bots/api#pollanswer
type PollAnswer struct {
	PollId    string  `json:"poll_id"`
	VoterChat *Chat   `json:"voter_chat,omitempty"`
	User      *User   `json:"user,omitempty"`
	OptionIds []int64 `json:"option_ids"`
}

// https://core.telegram.org/bots/api#poll
type Poll struct {
	Id                    string          `json:"id"`
	Question              string          `json:"question"`
	Options               []PollOption    `json:"options"`
	TotalVoterCount       int64           `json:"total_voter_count"`
	IsClosed              bool            `json:"is_closed"`
	IsAnonymous           bool            `json:"is_anonymous"`
	Type                  string          `json:"type"`
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`
	CorrectOptionId       int64           `json:"correct_option_id,omitempty"`
	Explanation           string          `json:"explanation,omitempty"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            int64           `json:"open_period,omitempty"`
	CloseDate             int64           `json:"close_date,omitempty"`
}

// https://core.telegram.org/bots/api#location
type Location struct {
	Longitude            float64 `json:"longitude"`
	Latitude             float64 `json:"latitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int64   `json:"live_period,omitempty"`
	Heading              int64   `json:"heading,omitempty"`
	ProximityAlertRadius int64   `json:"proximity_alert_radius,omitempty"`
}

// https://core.telegram.org/bots/api#venue
type Venue struct {
	Location        *Location `json:"location"`
	Title           string    `json:"title"`
	Address         string    `json:"address"`
	FoursquareId    string    `json:"foursquare_id,omitempty"`
	FoursquareType  string    `json:"foursquare_type,omitempty"`
	GooglePlaceId   string    `json:"google_place_id,omitempty"`
	GooglePlaceType string    `json:"google_place_type,omitempty"`
}

// https://core.telegram.org/bots/api#webappdata
type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

// https://core.telegram.org/bots/api#proximityalerttriggered
type ProximityAlertTriggered struct {
	Traveler *User `json:"traveler"`
	Watcher  *User `json:"watcher"`
	Distance int64 `json:"distance"`
}

// https://core.telegram.org/bots/api#messageautodeletetimerchanged
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int64 `json:"message_auto_delete_time"`
}

// https://core.telegram.org/bots/api#forumtopiccreated
type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int64  `json:"icon_color"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

// https://core.telegram.org/bots/api#forumtopicclosed
type ForumTopicClosed struct {
}

// https://core.telegram.org/bots/api#forumtopicedited
type ForumTopicEdited struct {
	Name              string `json:"name,omitempty"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

// https://core.telegram.org/bots/api#forumtopicreopened
type ForumTopicReopened struct {
}

// https://core.telegram.org/bots/api#generalforumtopichidden
type GeneralForumTopicHidden struct {
}

// https://core.telegram.org/bots/api#generalforumtopicunhidden
type GeneralForumTopicUnhidden struct {
}

// https://core.telegram.org/bots/api#usershared
type UserShared struct {
	RequestId int64 `json:"request_id"`
	UserId    int64 `json:"user_id"`
}

// https://core.telegram.org/bots/api#chatshared
type ChatShared struct {
	RequestId int64 `json:"request_id"`
	ChatId    int64 `json:"chat_id"`
}

// https://core.telegram.org/bots/api#writeaccessallowed
type WriteAccessAllowed struct {
	WebAppName string `json:"web_app_name,omitempty"`
}

// https://core.telegram.org/bots/api#videochatscheduled
type VideoChatScheduled struct {
	StartDate int64 `json:"start_date"`
}

// https://core.telegram.org/bots/api#videochatstarted
type VideoChatStarted struct {
}

// https://core.telegram.org/bots/api#videochatended
type VideoChatEnded struct {
	Duration int64 `json:"duration"`
}

// https://core.telegram.org/bots/api#videochatparticipantsinvited
type VideoChatParticipantsInvited struct {
	Users []User `json:"users"`
}

// // https://core.telegram.org/bots/api#userprofilephotos
// type UserProfilePhotos struct {
// 	TotalCount int64       `json:"total_count"`
// 	Photos     []PhotoSize `json:"photos"`
// }

// https://core.telegram.org/bots/api#file
type File struct {
}

// https://core.telegram.org/bots/api#webappinfo
type WebAppInfo struct {
	Url string `json:"url"`
}

// // https://core.telegram.org/bots/api#replykeyboardmarkup
// type ReplyKeyboardMarkup struct {
// 	Keyboard              []KeyboardButton `json:"keyboard"`
// 	IsPersistent          bool             `json:"is_persistent,omitempty"`
// 	ResizeKeyboard        bool             `json:"resize_keyboard,omitempty"`
// 	OneTimeKeyboard       bool             `json:"one_time_keyboard,omitempty"`
// 	InputFieldPlaceholder string           `json:"input_field_placeholder,omitempty"`
// 	Selective             bool             `json:"selective,omitempty"`
// }

// // https://core.telegram.org/bots/api#keyboardbutton
// type KeyboardButton struct {
// 	Text            string                     `json:"text"`
// 	RequestUser     *KeyboardButtonRequestUser `json:"request_user,omitempty"`
// 	RequestChat     *KeyboardButtonRequestChat `json:"request_chat,omitempty"`
// 	RequestContact  bool                       `json:"request_contact,omitempty"`
// 	RequestLocation bool                       `json:"request_location,omitempty"`
// 	RequestPoll     *KeyboardButtonPollType    `json:"request_poll,omitempty"`
// 	WebApp          *WebAppInfo                `json:"web_app,omitempty"`
// }

// // https://core.telegram.org/bots/api#keyboardbuttonrequestuser
// type KeyboardButtonRequestUser struct {
// 	RequestId     int64 `json:"request_id"`
// 	UserIsBot     bool  `json:"user_is_bot,omitempty"`
// 	UserIsPremium bool  `json:"user_is_premium,omitempty"`
// }

// // https://core.telegram.org/bots/api#keyboardbuttonrequestchat
// type KeyboardButtonRequestChat struct {
// 	RequestId               int64                    `json:"request_id"`
// 	ChatIsChannel           bool                     `json:"chat_is_channel"`
// 	ChatIsForum             bool                     `json:"chat_is_forum,omitempty"`
// 	ChatHasUserName         bool                     `json:"chat_has_username,omitempty"`
// 	ChatIsCreated           bool                     `json:"chat_is_created,omitempty"`
// 	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"`
// 	BotAdministratorRights  *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`
// 	BotIsMember             bool                     `json:"bot_is_member,omitempty"`
// }

// // https://core.telegram.org/bots/api#keyboardbuttonpolltype
// type KeyboardButtonPollType struct {
// 	Type string `json:"type,omitempty"`
// }

// // https://core.telegram.org/bots/api#replykeyboardremove
// type ReplyKeyboardRemove struct {
// 	RemoveKeyboard bool `json:"remove_keyboard"`
// 	Selective      bool `json:"selective,omitempty"`
// }

// https://core.telegram.org/bots/api#inlinekeyboardmarkup
type InlineKeyboardMarkup struct {
	InlineKeyboard []InlineKeyboardButton `json:"inline_keyboard"`
}

// https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`
	Url                          string                       `json:"url,omitempty"`
	CallbackData                 string                       `json:"callback_data,omitempty"`
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`
	LoginUrl                     *LoginUrl                    `json:"login_url,omitempty"`
	SwitchInlineQuery            string                       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string                       `json:"switch_inline_query_current_chat,omitempty"`
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`
	Pay                          bool                         `json:"pay,omitempty"`
}

// https://core.telegram.org/bots/api#loginurl
type LoginUrl struct {
}

// https://core.telegram.org/bots/api#switchinlinequerychosenchat
type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query,omitempty"`
	AllowUserChats    bool   `json:"allow_user_chats,omitempty"`
	AllowBotChats     bool   `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   bool   `json:"allow_group_chats,omitempty"`
	AllowChannelChats bool   `json:"allow_channel_chats,omitempty"`
}

// https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	Id              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageId string   `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data,omitempty"`
	GameShortName   string   `json:"game_short_name,omitempty"`
}

// // https://core.telegram.org/bots/api#forcereply
// type ForceReply struct {
// 	ForceReply            bool   `json:"force_reply"`
// 	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
// 	Selective             bool   `json:"selective,omitempty"`
// }

// https://core.telegram.org/bots/api#chatphoto
type ChatPhoto struct {
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
	BigFileId         string `json:"big_file_id"`
	BigFileUniqueId   string `json:"big_file_unique_id"`
}

// https://core.telegram.org/bots/api#chatinvitelink
type ChatInviteLink struct {
	InviteLink              string `json:"invite_link"`
	Creator                 *User  `json:"creator"`
	CreatesJoinRequest      bool   `json:"creates_join_request"`
	IsPrimary               bool   `json:"is_primary"`
	IsRevoked               bool   `json:"is_revoked"`
	Name                    string `json:"name,omitempty"`
	ExpireDate              int64  `json:"expire_date,omitempty"`
	MemberLimit             int64  `json:"member_limit,omitempty"`
	PendingJoinRequestCount int64  `json:"pending_join_request_count,omitempty"`
}

// // https://core.telegram.org/bots/api#chatadministratorrights
// type ChatAdministratorRights struct {
// 	IsAnonymous         bool `json:"is_anonymous"`
// 	CanManageChat       bool `json:"can_manage_chat"`
// 	CanDeleteMessages   bool `json:"can_delete_messages"`
// 	CanManageVideoChats bool `json:"can_manage_video_chats"`
// 	CanRestrictMembers  bool `json:"can_restrict_members"`
// 	CanPromoteMembers   bool `json:"can_promote_members"`
// 	CanChangeInfo       bool `json:"can_change_info"`
// 	CanInviteUsers      bool `json:"can_invite_users"`
// 	CanPostMessages     bool `json:"can_post_messages,omitempty"`
// 	CanEditMessages     bool `json:"can_edit_messages,omitempty"`
// 	CanPinMessages      bool `json:"can_pin_messages,omitempty"`
// 	CanManageTopics     bool `json:"can_manage_topics,omitempty"`
// }

// https://core.telegram.org/bots/api#chatmember
type ChatMember struct {
}

// // https://core.telegram.org/bots/api#chatmemberowner
// type ChatMemberOwner struct {
// 	Status      string `json:"status"`
// 	User        *User  `json:"user"`
// 	IsAnonymous bool   `json:"is_anonymous"`
// 	CustomTitle string `json:"custom_title,omitempty"`
// }

// // https://core.telegram.org/bots/api#chatmemberadministrator
// type ChatMemberAdministrator struct {
// 	Status              string `json:"status"`
// 	User                *User  `json:"user"`
// 	CanBeEdited         bool   `json:"can_be_edited"`
// 	IsAnonymous         bool   `json:"is_anonymous"`
// 	CanManageChat       bool   `json:"can_manage_chat"`
// 	CanDeleteMessages   bool   `json:"can_delete_messages"`
// 	CanManageVideoChats bool   `json:"can_manage_video_chats"`
// 	CanRestrictMembers  bool   `json:"can_restrict_members"`
// 	CanPromoteMembers   bool   `json:"can_promote_members"`
// 	CanChangeInfo       bool   `json:"can_change_info"`
// 	CanInviteUsers      bool   `json:"can_invite_users"`
// 	CanPostMessages     bool   `json:"can_post_messages,omitempty"`
// 	CanEditMessages     bool   `json:"can_edit_messages,omitempty"`
// 	CanPinMessages      bool   `json:"can_pin_messages,omitempty"`
// 	CanManageTopics     bool   `json:"can_manage_topics,omitempty"`
// 	CustomTitle         string `json:"custom_title,omitempty"`
// }

// // https://core.telegram.org/bots/api#chatmembermember
// type ChatMemberMember struct {
// 	Status string `json:"status"`
// 	User   *User  `json:"user"`
// }

// // https://core.telegram.org/bots/api#chatmemberrestricted
// type ChatMemberRestricted struct {
// 	Status                string `json:"status"`
// 	User                  *User  `json:"user"`
// 	IsMember              bool   `json:"is_member"`
// 	CanSendMessages       bool   `json:"can_send_messages"`
// 	CanSendAudios         bool   `json:"can_send_audios"`
// 	CanSendDocuments      bool   `json:"can_send_documents"`
// 	CanSendPhotos         bool   `json:"can_send_photos"`
// 	CanSendVideos         bool   `json:"can_send_videos"`
// 	CanSendVideoNotes     bool   `json:"can_send_video_notes"`
// 	CanSendVoiceNotes     bool   `json:"can_send_voice_notes"`
// 	CanSendPolls          bool   `json:"can_send_polls"`
// 	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
// 	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
// 	CanChangeInfo         bool   `json:"can_change_info"`
// 	CanInviteUsers        bool   `json:"can_invite_users"`
// 	CanPinMessages        bool   `json:"can_pin_messages"`
// 	CanManageTopics       bool   `json:"can_manage_topics"`
// 	UntilDate             int64  `json:"until_date"`
// }

// // https://core.telegram.org/bots/api#chatmemberleft
// type ChatMemberLeft struct {
// 	Status string `json:"status"`
// 	User   *User  `json:"user"`
// }

// // https://core.telegram.org/bots/api#chatmemberbanned
// type ChatMemberBanned struct {
// 	Status    string `json:"status"`
// 	User      *User  `json:"user"`
// 	UntilDate int64  `json:"until_date"`
// }

// https://core.telegram.org/bots/api#chatmemberupdated
type ChatMemberUpdated struct {
	Chat                    *Chat           `json:"chat"`
	From                    *User           `json:"from"`
	Date                    int64           `json:"date"`
	OldChatMember           *ChatMember     `json:"old_chat_member"`
	NewChatMember           *ChatMember     `json:"new_chat_member"`
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`
	ViaChatFolderInviteLink bool            `json:"via_chat_folder_invite_link,omitempty"`
}

// https://core.telegram.org/bots/api#chatjoinrequest
type ChatJoinRequest struct {
	Chat       *Chat           `json:"chat"`
	From       *User           `json:"from"`
	UserChatId *int64          `json:"user_chat_id"`
	Date       int64           `json:"date"`
	Bio        string          `json:"bio,omitempty"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

// https://core.telegram.org/bots/api#chatpermissions
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendAudios         bool `json:"can_send_audios,omitempty"`
	CanSendDocuments      bool `json:"can_send_documents,omitempty"`
	CanSendPhotos         bool `json:"can_send_photos,omitempty"`
	CanSendVideos         bool `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     bool `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     bool `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`
	CanManageTopics       bool `json:"can_manage_topics,omitempty"`
}

// https://core.telegram.org/bots/api#chatlocation
type ChatLocation struct {
	Location *Location `json:"location"`
	Address  string    `json:"address"`
}

// // https://core.telegram.org/bots/api#forumtopic
// type ForumTopic struct {
// 	MessageThreadId   int64  `json:"message_thread_id"`
// 	Name              string `json:"name"`
// 	IconColor         int64  `json:"icon_color"`
// 	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
// }

// // https://core.telegram.org/bots/api#botcommand
// type BotCommand struct {
// 	Command     string `json:"command"`
// 	Description string `json:"description"`
// }

// // https://core.telegram.org/bots/api#botcommandscope
// type BotCommandScope struct {
// }

// // https://core.telegram.org/bots/api#determining-list-of-commands
// type DeterminingListOfCommands struct {
// }

// // https://core.telegram.org/bots/api#botcommandscopedefault
// type BotCommandScopeDefault struct {
// 	Type string `json:"type"`
// }

// // https://core.telegram.org/bots/api#botcommandscopeallprivatechats
// type BotCommandScopeAllPrivateChats struct {
// 	Type string `json:"type"`
// }

// // https://core.telegram.org/bots/api#botcommandscopeallgroupchats
// type BotCommandScopeAllGroupChats struct {
// 	Type string `json:"type"`
// }

// // https://core.telegram.org/bots/api#botcommandscopeallchatadministrators
// type BotCommandScopeAllChatAdministrators struct {
// 	Type string `json:"type"`
// }

// // https://core.telegram.org/bots/api#botcommandscopechat
// type BotCommandScopeChat struct {
// 	Type   string `json:"type"`
// 	ChatId int64  `json:"chat_id"`
// }

// // https://core.telegram.org/bots/api#botcommandscopechatadministrators
// type BotCommandScopeChatAdministrators struct {
// 	Type   string `json:"type"`
// 	ChatId int64  `json:"chat_id"`
// }

// // https://core.telegram.org/bots/api#botcommandscopechatmember
// type BotCommandScopeChatMember struct {
// 	Type   string `json:"type"`
// 	ChatId int64  `json:"chat_id"`
// 	UserId int64  `json:"user_id"`
// }

// // https://core.telegram.org/bots/api#botname
// type BotName struct {
// 	Name string `json:"name"`
// }

// // https://core.telegram.org/bots/api#botdescription
// type BotDescription struct {
// 	Description string `json:"description"`
// }

// // https://core.telegram.org/bots/api#botshortdescription
// type BotShortDescription struct {
// 	ShortDescription string `json:"short_description"`
// }

// // https://core.telegram.org/bots/api#menubutton
// type MenuButton struct {
// }

// // https://core.telegram.org/bots/api#menubuttoncommands
// type MenuButtonCommands struct {
// 	Type string `json:"type"`
// }

// // https://core.telegram.org/bots/api#menubuttonwebapp
// type MenuButtonWebApp struct {
// 	Type   string      `json:"type"`
// 	Text   string      `json:"text"`
// 	WebApp *WebAppInfo `json:"web_app"`
// }

// // https://core.telegram.org/bots/api#menubuttondefault
// type MenuButtonDefault struct {
// 	Type string `json:"type"`
// }

// // https://core.telegram.org/bots/api#responseparameters
// type ResponseParameters struct {
// 	MigrateToChatId int64 `json:"migrate_to_chat_id,omitempty"`
// 	RetryAfter      int64 `json:"retry_after,omitempty"`
// }

// // https://core.telegram.org/bots/api#inputmedia
// type InputMedia struct {
// }

// // https://core.telegram.org/bots/api#inputmediaphoto
// type InputMediaPhoto struct {
// 	Type            string          `json:"type"`
// 	Media           string          `json:"media"`
// 	Caption         string          `json:"caption,omitempty"`
// 	ParseMode       string          `json:"parse_mode,omitempty"`
// 	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
// 	HasSpoiler      bool            `json:"has_spoiler,omitempty"`
// }

// // https://core.telegram.org/bots/api#inputmediavideo
// type InputMediaVideo struct {
// 	Type              string          `json:"type"`
// 	Media             string          `json:"media"`
// 	Thumbnail         *InputFile      `json:"thumbnail,omitempty"`
// 	Caption           string          `json:"caption,omitempty"`
// 	ParseMode         string          `json:"parse_mode,omitempty"`
// 	CaptionEntities   []MessageEntity `json:"caption_entities,omitempty"`
// 	Width             int64           `json:"width,omitempty"`
// 	Height            int64           `json:"height,omitempty"`
// 	Duration          int64           `json:"duration,omitempty"`
// 	SupportsStreaming bool            `json:"supports_streaming,omitempty"`
// 	HasSpoiler        bool            `json:"has_spoiler,omitempty"`
// }

// // https://core.telegram.org/bots/api#inputmediaanimation
// type InputMediaAnimation struct {
// 	Type            string          `json:"type"`
// 	Media           string          `json:"media"`
// 	Thumbnail       *InputFile      `json:"thumbnail,omitempty"`
// 	Caption         string          `json:"caption,omitempty"`
// 	ParseMode       string          `json:"parse_mode,omitempty"`
// 	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
// 	Width           int64           `json:"width,omitempty"`
// 	Height          int64           `json:"height,omitempty"`
// 	Duration        int64           `json:"duration,omitempty"`
// 	HasSpoiler      bool            `json:"has_spoiler,omitempty"`
// }

// // https://core.telegram.org/bots/api#inputmediaaudio
// type InputMediaAudio struct {
// 	Type            string          `json:"type"`
// 	Media           string          `json:"media"`
// 	Thumbnail       *InputFile      `json:"thumbnail,omitempty"`
// 	Caption         string          `json:"caption,omitempty"`
// 	ParseMode       string          `json:"parse_mode,omitempty"`
// 	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
// 	Duration        int64           `json:"duration,omitempty"`
// 	Performer       string          `json:"performer,omitempty"`
// 	Title           string          `json:"title,omitempty"`
// }

// // https://core.telegram.org/bots/api#inputmediadocument
// type InputMediaDocument struct {
// 	Type                        string          `json:"type"`
// 	Media                       string          `json:"media"`
// 	Thumbnail                   *InputFile      `json:"thumbnail,omitempty"`
// 	Caption                     string          `json:"caption,omitempty"`
// 	ParseMode                   string          `json:"parse_mode,omitempty"`
// 	CaptionEntities             []MessageEntity `json:"caption_entities,omitempty"`
// 	DisableContentTypeDetection bool            `json:"disable_content_type_detection,omitempty"`
// }

// // https://core.telegram.org/bots/api#inputfile
// type InputFile struct {
// }

// https://core.telegram.org/bots/api#sticker
type Sticker struct {
	FileId           string        `json:"file_id"`
	FileUniqueId     string        `json:"file_unique_id"`
	Type             string        `json:"type"`
	Width            int64         `json:"width"`
	Height           int64         `json:"height"`
	IsAnimated       bool          `json:"is_animated"`
	IsVideo          bool          `json:"is_video"`
	Thumbnail        *PhotoSize    `json:"thumbnail,omitempty"`
	Emoji            string        `json:"emoji,omitempty"`
	SetName          string        `json:"set_name,omitempty"`
	PremiumAnimation *File         `json:"premium_animation,omitempty"`
	MaskPosition     *MaskPosition `json:"mask_position,omitempty"`
	CustomEmojiId    string        `json:"custom_emoji_id,omitempty"`
	NeedsRepainting  bool          `json:"needs_repainting,omitempty"`
	FileSize         int64         `json:"file_size,omitempty"`
}

// // https://core.telegram.org/bots/api#stickerset
// type StickerSet struct {
// 	Name        string     `json:"name"`
// 	Title       string     `json:"title"`
// 	StickerType string     `json:"sticker_type"`
// 	IsAnimated  bool       `json:"is_animated"`
// 	IsVideo     bool       `json:"is_video"`
// 	Stickers    []Sticker  `json:"stickers"`
// 	Thumbnail   *PhotoSize `json:"thumbnail,omitempty"`
// }

// https://core.telegram.org/bots/api#maskposition
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

// // https://core.telegram.org/bots/api#inputsticker
// type InputSticker struct {
// 	Sticker      *InputFile    `json:"sticker"`
// 	EmojiList    []string      `json:"emoji_list"`
// 	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
// 	Keywords     []string      `json:"keywords,omitempty"`
// }

// https://core.telegram.org/bots/api#inlinequery
type InlineQuery struct {
	Id       string    `json:"id"`
	From     *User     `json:"from"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	ChatType string    `json:"chat_type,omitempty"`
	Location *Location `json:"location,omitempty"`
}

// // https://core.telegram.org/bots/api#inlinequeryresultsbutton
// type InlineQueryResultsButton struct {
// 	Text           string      `json:"text"`
// 	WebApp         *WebAppInfo `json:"web_app,omitempty"`
// 	StartParameter string      `json:"start_parameter,omitempty"`
// }

// // https://core.telegram.org/bots/api#inlinequeryresult
// type InlineQueryResult struct {
// }

// // https://core.telegram.org/bots/api#inlinequeryresultarticle
// type InlineQueryResultArticle struct {
// 	Type                string                `json:"type"`
// 	Id                  string                `json:"id"`
// 	Title               string                `json:"title"`
// 	InputMessageContent *InputMessageContent  `json:"input_message_content"`
// 	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
// 	Url                 string                `json:"url,omitempty"`
// 	HideUrl             bool                  `json:"hide_url,omitempty"`
// 	Description         string                `json:"description,omitempty"`
// 	ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`
// 	ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`
// 	ThumbnailHeight     int64                 `json:"thumbnail_height,omitempty"`
// }

// // https://core.telegram.org/bots/api#inlinequeryresultphoto
// type InlineQueryResultPhoto struct {
// 	Type                string                `json:"type"`
// 	Id                  string                `json:"id"`
// 	PhotoUrl            string                `json:"photo_url"`
// 	ThumbnailUrl        string                `json:"thumbnail_url"`
// 	PhotoWidth          int64                 `json:"photo_width,omitempty"`
// 	PhotoHeight         int64                 `json:"photo_height,omitempty"`
// 	Title               string                `json:"title,omitempty"`
// 	Description         string                `json:"description,omitempty"`
// 	Caption             string                `json:"caption,omitempty"`
// 	ParseMode           string                `json:"parse_mode,omitempty"`
// 	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
// 	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
// 	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
// }

// // https://core.telegram.org/bots/api#inlinequeryresultgif
// type InlineQueryResultGif struct {
// 	Type                string                `json:"type"`
// 	Id                  string                `json:"id"`
// 	GifUrl              string                `json:"gif_url"`
// 	GifWidth            int64                 `json:"gif_width,omitempty"`
// 	GifHeight           int64                 `json:"gif_height,omitempty"`
// 	GifDuration         int64                 `json:"gif_duration,omitempty"`
// 	ThumbnailUrl        string                `json:"thumbnail_url"`
// 	ThumbnailMimeType   string                `json:"thumbnail_mime_type,omitempty"`
// 	Title               string                `json:"title,omitempty"`
// 	Caption             string                `json:"caption,omitempty"`
// 	ParseMode           string                `json:"parse_mode,omitempty"`
// 	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
// 	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
// 	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
// }

// // https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
// type InlineQueryResultMpeg4Gif struct {
// 	Type                string                `json:"type"`
// 	Id                  string                `json:"id"`
// 	Mpeg4Url            string                `json:"mpeg4_url"`
// 	Mpeg4Width          int64                 `json:"mpeg4_width,omitempty"`
// 	Mpeg4Height         int64                 `json:"mpeg4_height,omitempty"`
// 	Mpeg4Duration       int64                 `json:"mpeg4_duration,omitempty"`
// 	ThumbnailUrl        string                `json:"thumbnail_url"`
// 	ThumbnailMimeType   string                `json:"thumbnail_mime_type,omitempty"`
// 	Title               string                `json:"title,omitempty"`
// 	Caption             string                `json:"caption,omitempty"`
// 	ParseMode           string                `json:"parse_mode,omitempty"`
// 	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
// 	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
// 	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
// }

// // https://core.telegram.org/bots/api#inlinequeryresultvideo
// type InlineQueryResultVideo struct {
// }

// // https://core.telegram.org/bots/api#inlinequeryresultaudio
// type InlineQueryResultAudio struct {
// 	Type                string                `json:"type"`
// 	Id                  string                `json:"id"`
// 	AudioUrl            string                `json:"audio_url"`
// 	Title               string                `json:"title"`
// 	Caption             string                `json:"caption,omitempty"`
// 	ParseMode           string                `json:"parse_mode,omitempty"`
// 	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
// 	Performer           string                `json:"performer,omitempty"`
// 	AudioDuration       int64                 `json:"audio_duration,omitempty"`
// 	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
// 	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
// }

// // https://core.telegram.org/bots/api#inlinequeryresultvoice
// type InlineQueryResultVoice struct {
// 	Type                string                `json:"type"`
// 	Id                  string                `json:"id"`
// 	VoiceUrl            string                `json:"voice_url"`
// 	Title               string                `json:"title"`
// 	Caption             string                `json:"caption,omitempty"`
// 	ParseMode           string                `json:"parse_mode,omitempty"`
// 	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
// 	VoiceDuration       int64                 `json:"voice_duration,omitempty"`
// 	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
// 	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
// }

// // https://core.telegram.org/bots/api#inlinequeryresultdocument
// type InlineQueryResultDocument struct {
// 	Type                string                `json:"type"`
// 	Id                  string                `json:"id"`
// 	Title               string                `json:"title"`
// 	Caption             string                `json:"caption,omitempty"`
// 	ParseMode           string                `json:"parse_mode,omitempty"`
// 	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
// 	DocumentUrl         string                `json:"document_url"`
// 	MimeType            string                `json:"mime_type"`
// 	Description         string                `json:"description,omitempty"`
// 	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
// 	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
// 	ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`
// 	ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`
// 	ThumbnailHeight     int64                 `json:"thumbnail_height,omitempty"`
// }

// // https://core.telegram.org/bots/api#inlinequeryresultlocation
// type InlineQueryResultLocation struct {
// 	Type                 string                `json:"type"`
// 	Id                   string                `json:"id"`
// 	Latitude             float64               `json:"latitude"`
// 	Longitude            float64               `json:"longitude"`
// 	Title                string                `json:"title"`
// 	HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`
// 	LivePeriod           int64                 `json:"live_period,omitempty"`
// 	Heading              int64                 `json:"heading,omitempty"`
// 	ProximityAlertRadius int64                 `json:"proximity_alert_radius,omitempty"`
// 	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
// 	InputMessageContent  *InputMessageContent  `json:"input_message_content,omitempty"`
// 	ThumbnailUrl         string                `json:"thumbnail_url,omitempty"`
// 	ThumbnailWidth       int64                 `json:"thumbnail_width,omitempty"`
// 	ThumbnailHeight      int64                 `json:"thumbnail_height,omitempty"`
// }

// // https://core.telegram.org/bots/api#inlinequeryresultvenue
// type InlineQueryResultVenue struct {
// 	Type                string                `json:"type"`
// 	Id                  string                `json:"id"`
// 	Latitude            float64               `json:"latitude"`
// 	Longitude           float64               `json:"longitude"`
// 	Title               string                `json:"title"`
// 	Address             string                `json:"address"`
// 	FoursquareId        string                `json:"foursquare_id,omitempty"`
// 	FoursquareType      string                `json:"foursquare_type,omitempty"`
// 	GooglePlaceId       string                `json:"google_place_id,omitempty"`
// 	GooglePlaceType     string                `json:"google_place_type,omitempty"`
// 	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
// 	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
// 	ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`
// 	ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`
// 	ThumbnailHeight     int64                 `json:"thumbnail_height,omitempty"`
// }

// // https://core.telegram.org/bots/api#inlinequeryresultcontact
// type InlineQueryResultContact struct {
// 	Type                string                `json:"type"`
// 	Id                  string                `json:"id"`
// 	PhoneNumber         string                `json:"phone_number"`
// 	FirstName           string                `json:"first_name"`
// 	LastName            string                `json:"last_name,omitempty"`
// 	Vcard               string                `json:"vcard,omitempty"`
// 	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
// 	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
// 	ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`
// 	ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`
// 	ThumbnailHeight     int64                 `json:"thumbnail_height,omitempty"`
// }

// // https://core.telegram.org/bots/api#inlinequeryresultgame
// type InlineQueryResultGame struct {
// 	Type          string                `json:"type"`
// 	Id            string                `json:"id"`
// 	GameShortName string                `json:"game_short_name"`
// 	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
// }

// // https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
// type InlineQueryResultCachedPhoto struct {
// 	Type                string                `json:"type"`
// 	Id                  string                `json:"id"`
// 	PhotoFileId         string                `json:"photo_file_id"`
// 	Title               string                `json:"title,omitempty"`
// 	Description         string                `json:"description,omitempty"`
// 	Caption             string                `json:"caption,omitempty"`
// 	ParseMode           string                `json:"parse_mode,omitempty"`
// 	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
// 	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
// 	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
// }

// // https://core.telegram.org/bots/api#inlinequeryresultcachedgif
// type InlineQueryResultCachedGif struct {
// 	Type                string                `json:"type"`
// 	Id                  string                `json:"id"`
// 	GifFileId           string                `json:"gif_file_id"`
// 	Title               string                `json:"title,omitempty"`
// 	Caption             string                `json:"caption,omitempty"`
// 	ParseMode           string                `json:"parse_mode,omitempty"`
// 	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
// 	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
// 	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
// }

// // https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
// type InlineQueryResultCachedMpeg4Gif struct {
// 	Type                string                `json:"type"`
// 	Id                  string                `json:"id"`
// 	Mpeg4FileId         string                `json:"mpeg4_file_id"`
// 	Title               string                `json:"title,omitempty"`
// 	Caption             string                `json:"caption,omitempty"`
// 	ParseMode           string                `json:"parse_mode,omitempty"`
// 	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
// 	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
// 	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
// }

// // https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
// type InlineQueryResultCachedSticker struct {
// 	Type                string                `json:"type"`
// 	Id                  string                `json:"id"`
// 	StickerFileId       string                `json:"sticker_file_id"`
// 	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
// 	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
// }

// // https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
// type InlineQueryResultCachedDocument struct {
// 	Type                string                `json:"type"`
// 	Id                  string                `json:"id"`
// 	Title               string                `json:"title"`
// 	DocumentFileId      string                `json:"document_file_id"`
// 	Description         string                `json:"description,omitempty"`
// 	Caption             string                `json:"caption,omitempty"`
// 	ParseMode           string                `json:"parse_mode,omitempty"`
// 	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
// 	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
// 	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
// }

// // https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
// type InlineQueryResultCachedVideo struct {
// 	Type                string                `json:"type"`
// 	Id                  string                `json:"id"`
// 	VideoFileId         string                `json:"video_file_id"`
// 	Title               string                `json:"title"`
// 	Description         string                `json:"description,omitempty"`
// 	Caption             string                `json:"caption,omitempty"`
// 	ParseMode           string                `json:"parse_mode,omitempty"`
// 	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
// 	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
// 	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
// }

// // https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
// type InlineQueryResultCachedVoice struct {
// 	Type                string                `json:"type"`
// 	Id                  string                `json:"id"`
// 	VoiceFileId         string                `json:"voice_file_id"`
// 	Title               string                `json:"title"`
// 	Caption             string                `json:"caption,omitempty"`
// 	ParseMode           string                `json:"parse_mode,omitempty"`
// 	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
// 	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
// 	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
// }

// // https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
// type InlineQueryResultCachedAudio struct {
// 	Type                string                `json:"type"`
// 	Id                  string                `json:"id"`
// 	AudioFileId         string                `json:"audio_file_id"`
// 	Caption             string                `json:"caption,omitempty"`
// 	ParseMode           string                `json:"parse_mode,omitempty"`
// 	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
// 	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
// 	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
// }

// // https://core.telegram.org/bots/api#inputmessagecontent
// type InputMessageContent struct {
// }

// // https://core.telegram.org/bots/api#inputtextmessagecontent
// type InputTextMessageContent struct {
// 	MessageText           string          `json:"message_text"`
// 	ParseMode             string          `json:"parse_mode,omitempty"`
// 	Entities              []MessageEntity `json:"entities,omitempty"`
// 	DisableWebPagePreview bool            `json:"disable_web_page_preview,omitempty"`
// }

// // https://core.telegram.org/bots/api#inputlocationmessagecontent
// type InputLocationMessageContent struct {
// 	Latitude             float64 `json:"latitude"`
// 	Longitude            float64 `json:"longitude"`
// 	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
// 	LivePeriod           int64   `json:"live_period,omitempty"`
// 	Heading              int64   `json:"heading,omitempty"`
// 	ProximityAlertRadius int64   `json:"proximity_alert_radius,omitempty"`
// }

// // https://core.telegram.org/bots/api#inputvenuemessagecontent
// type InputVenueMessageContent struct {
// 	Latitude        float64 `json:"latitude"`
// 	Longitude       float64 `json:"longitude"`
// 	Title           string  `json:"title"`
// 	Address         string  `json:"address"`
// 	FoursquareId    string  `json:"foursquare_id,omitempty"`
// 	FoursquareType  string  `json:"foursquare_type,omitempty"`
// 	GooglePlaceId   string  `json:"google_place_id,omitempty"`
// 	GooglePlaceType string  `json:"google_place_type,omitempty"`
// }

// // https://core.telegram.org/bots/api#inputcontactmessagecontent
// type InputContactMessageContent struct {
// 	PhoneNumber string `json:"phone_number"`
// 	FirstName   string `json:"first_name"`
// 	LastName    string `json:"last_name,omitempty"`
// 	Vcard       string `json:"vcard,omitempty"`
// }

// // https://core.telegram.org/bots/api#inputinvoicemessagecontent
// type InputInvoiceMessageContent struct {
// 	Title                     string         `json:"title"`
// 	Description               string         `json:"description"`
// 	Payload                   string         `json:"payload"`
// 	ProviderToken             string         `json:"provider_token"`
// 	Currency                  string         `json:"currency"`
// 	Prices                    []LabeledPrice `json:"prices"`
// 	MaxTipAmount              int64          `json:"max_tip_amount,omitempty"`
// 	SuggestedTipAmounts       []int64        `json:"suggested_tip_amounts,omitempty"`
// 	ProviderData              string         `json:"provider_data,omitempty"`
// 	PhotoUrl                  string         `json:"photo_url,omitempty"`
// 	PhotoSize                 int64          `json:"photo_size,omitempty"`
// 	PhotoWidth                int64          `json:"photo_width,omitempty"`
// 	PhotoHeight               int64          `json:"photo_height,omitempty"`
// 	NeedName                  bool           `json:"need_name,omitempty"`
// 	NeedPhoneNumber           bool           `json:"need_phone_number,omitempty"`
// 	NeedEmail                 bool           `json:"need_email,omitempty"`
// 	NeedShippingAddress       bool           `json:"need_shipping_address,omitempty"`
// 	SendPhoneNumberToProvider bool           `json:"send_phone_number_to_provider,omitempty"`
// 	SendEmailToProvider       bool           `json:"send_email_to_provider,omitempty"`
// 	IsFlexible                bool           `json:"is_flexible,omitempty"`
// }

// https://core.telegram.org/bots/api#choseninlineresult
type ChosenInlineResult struct {
	ResultId        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageId string    `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}

// // https://core.telegram.org/bots/api#sentwebappmessage
// type SentWebAppMessage struct {
// 	InlineMessageId string `json:"inline_message_id,omitempty"`
// }

// // https://core.telegram.org/bots/api#labeledprice
// type LabeledPrice struct {
// 	Label  string `json:"label"`
// 	Amount int64  `json:"amount"`
// }

// https://core.telegram.org/bots/api#invoice
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int64  `json:"total_amount"`
}

// https://core.telegram.org/bots/api#shippingaddress
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

// https://core.telegram.org/bots/api#orderinfo
type OrderInfo struct {
	Name            string           `json:"name,omitempty"`
	PhoneNumber     string           `json:"phone_number,omitempty"`
	Email           string           `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// // https://core.telegram.org/bots/api#shippingoption
// type ShippingOption struct {
// 	Id     string         `json:"id"`
// 	Title  string         `json:"title"`
// 	Prices []LabeledPrice `json:"prices"`
// }

// https://core.telegram.org/bots/api#successfulpayment
type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	TotalAmount             int64      `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	ShippingOptionId        string     `json:"shipping_option_id,omitempty"`
	OrderInfo               *OrderInfo `json:"order_info,omitempty"`
	TelegramPaymentChargeId string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeId string     `json:"provider_payment_charge_id"`
}

// https://core.telegram.org/bots/api#shippingquery
type ShippingQuery struct {
	Id              string           `json:"id"`
	From            *User            `json:"from"`
	InvoicePayload  string           `json:"invoice_payload"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// https://core.telegram.org/bots/api#precheckoutquery
type PreCheckoutQuery struct {
	Id               string     `json:"id"`
	From             *User      `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int64      `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionId string     `json:"shipping_option_id,omitempty"`
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`
}

// https://core.telegram.org/bots/api#passportdata
type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials *EncryptedCredentials      `json:"credentials"`
}

// https://core.telegram.org/bots/api#passportfile
type PassportFile struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	FileDate     int64  `json:"file_date"`
}

// https://core.telegram.org/bots/api#encryptedpassportelement
type EncryptedPassportElement struct {
	Type        string         `json:"type"`
	Data        string         `json:"data,omitempty"`
	PhoneNumber string         `json:"phone_number,omitempty"`
	Email       string         `json:"email,omitempty"`
	Files       []PassportFile `json:"files,omitempty"`
	FrontSide   *PassportFile  `json:"front_side,omitempty"`
	ReverseSide *PassportFile  `json:"reverse_side,omitempty"`
	Selfie      *PassportFile  `json:"selfie,omitempty"`
	Translation []PassportFile `json:"translation,omitempty"`
	Hash        string         `json:"hash"`
}

// https://core.telegram.org/bots/api#encryptedcredentials
type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

// // https://core.telegram.org/bots/api#passportelementerror
// type PassportElementError struct {
// }

// // https://core.telegram.org/bots/api#passportelementerrordatafield
// type PassportElementErrorDataField struct {
// 	Source    string `json:"source"`
// 	Type      string `json:"type"`
// 	FieldName string `json:"field_name"`
// 	DataHash  string `json:"data_hash"`
// 	Message   string `json:"message"`
// }

// // https://core.telegram.org/bots/api#passportelementerrorfrontside
// type PassportElementErrorFrontSide struct {
// 	Source   string `json:"source"`
// 	Type     string `json:"type"`
// 	FileHash string `json:"file_hash"`
// 	Message  string `json:"message"`
// }

// // https://core.telegram.org/bots/api#passportelementerrorreverseside
// type PassportElementErrorReverseSide struct {
// 	Source   string `json:"source"`
// 	Type     string `json:"type"`
// 	FileHash string `json:"file_hash"`
// 	Message  string `json:"message"`
// }

// // https://core.telegram.org/bots/api#passportelementerrorselfie
// type PassportElementErrorSelfie struct {
// 	Source   string `json:"source"`
// 	Type     string `json:"type"`
// 	FileHash string `json:"file_hash"`
// 	Message  string `json:"message"`
// }

// // https://core.telegram.org/bots/api#passportelementerrorfile
// type PassportElementErrorFile struct {
// 	Source   string `json:"source"`
// 	Type     string `json:"type"`
// 	FileHash string `json:"file_hash"`
// 	Message  string `json:"message"`
// }

// // https://core.telegram.org/bots/api#passportelementerrorfiles
// type PassportElementErrorFiles struct {
// 	Source     string   `json:"source"`
// 	Type       string   `json:"type"`
// 	FileHashes []string `json:"file_hashes"`
// 	Message    string   `json:"message"`
// }

// // https://core.telegram.org/bots/api#passportelementerrortranslationfile
// type PassportElementErrorTranslationFile struct {
// 	Source   string `json:"source"`
// 	Type     string `json:"type"`
// 	FileHash string `json:"file_hash"`
// 	Message  string `json:"message"`
// }

// // https://core.telegram.org/bots/api#passportelementerrortranslationfiles
// type PassportElementErrorTranslationFiles struct {
// 	Source     string   `json:"source"`
// 	Type       string   `json:"type"`
// 	FileHashes []string `json:"file_hashes"`
// 	Message    string   `json:"message"`
// }

// // https://core.telegram.org/bots/api#passportelementerrorunspecified
// type PassportElementErrorUnspecified struct {
// 	Source      string `json:"source"`
// 	Type        string `json:"type"`
// 	ElementHash string `json:"element_hash"`
// 	Message     string `json:"message"`
// }

// https://core.telegram.org/bots/api#game
type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	Animation    *Animation      `json:"animation,omitempty"`
}

// https://core.telegram.org/bots/api#callbackgame
type CallbackGame struct {
}

// // https://core.telegram.org/bots/api#gamehighscore
// type GameHighScore struct {
// 	Position int64 `json:"position"`
// 	User     *User `json:"user"`
// 	Score    int64 `json:"score"`
// }

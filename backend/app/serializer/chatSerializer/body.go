package chatSerializer

import (
	"angrymiao-ai/app/tools/openAiSdk"
	"mime/multipart"

	"github.com/sashabaranov/go-openai"
)

type ChatBodyRequest struct {
	Message    *string               `form:"message" json:"message" binding:"required"`
	DeviceCode string                `form:"device_code" json:"device_code" binding:"required"`
	Image      *multipart.FileHeader `form:"image" binding:"omitempty"`
	//Image      *string `json:"image" binding:"omitempty"`
}

type VoiceChatBodyRequest struct {
	DeviceCode string                `form:"device_code" binding:"required"`
	AudioFile  *multipart.FileHeader `form:"audio" binding:"required"`
}

type HistoryChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Image   string `json:"Image"`
}

type ChatReplyResp struct {
	*ChatReply
	Error string `json:"error"`
}

type ChatReply struct {
	ReplyMsg string `json:"reply_msg"`
	FuncName string `json:"func_name"`
	//IsChangeConf bool   `json:"is_change_conf"`
	//ChangeJson   any    `json:"change_json"`
}

type GlassesConfigRequest struct {
	ForwardMode   string `json:"forward_mode" binding:"required,oneof=all media"`
	DiscordUserID string `json:"discord_user_id" binding:"omitempty"`
	DeviceCode    string `json:"device_code" binding:"required"`
}

type ChatToolCall struct {
	Msg      openai.ChatCompletionMessage
	History  []HistoryChatMessage
	Client   *openAiSdk.Client
	FuncName string
	FuncArgs string
}

package chat

import (
	"angrymiao-ai/app/exception"
	"angrymiao-ai/app/serializer/chatSerializer"
	serializerResponse "angrymiao-ai/app/serializer/response"
	"angrymiao-ai/app/service/chatService"
	"angrymiao-ai/app/validate"
	"github.com/gin-gonic/gin"
)

// VoiceChat
//
//	@Summary	AI 语音聊天接口
//	@Tags		AI
//	@Accept		application/json
//	@Produce	application/json
//	@Security	ApiKeyAuth
//	@Param		data	body		chatSerializer.VoiceChatBodyRequest	true	"data"
//	@Success	200		{object}	chatSerializer.ChatReply
//	@Router		/chat [post]
func (h *Handler) VoiceChat(c *gin.Context) {
	var body *chatSerializer.VoiceChatBodyRequest
	if !validate.CheckFormData(c, &body) {
		return
	}

	ttsFilePath, err := new(chatService.Service).HandleVoiceChat(body, c)
	if err != nil {
		exception.BadRequestException(c, err.Error())
		return
	}

	serializerResponse.Create(c, map[string]string{
		"ttsFilePath": ttsFilePath,
	})
	return
}

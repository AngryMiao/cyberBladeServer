package chat

import (
	"angrymiao-ai/app/exception"
	"angrymiao-ai/app/serializer/chatSerializer"
	serializerResponse "angrymiao-ai/app/serializer/response"
	"angrymiao-ai/app/service/chatService"
	"angrymiao-ai/app/service/forwardService"
	"angrymiao-ai/app/tools/pool"
	"angrymiao-ai/app/validate"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func simulateWork(duration time.Duration) (interface{}, error) {
	time.Sleep(duration)
	// 随机生成成功或失败
	if rand.Float32() < 0.1 { // 10% 失败率
		return nil, fmt.Errorf("random error occurred")
	}
	return fmt.Sprintf("work completed after %v", duration), nil
}

// Chat
//
//	@Summary	AI 文字聊天接口
//	@Tags		AI
//	@Accept		application/json
//	@Produce	application/json
//	@Security	ApiKeyAuth
//	@Param		data	body		chatSerializer.ChatBodyRequest	true	"data"
//	@Success	200		{object}	chatSerializer.ChatReply
//	@Router		/chat [post]
func (h *Handler) Chat(c *gin.Context) {
	var body *chatSerializer.ChatBodyRequest
	if !validate.CheckFormData(c, &body) {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resultChan := make(chan *chatSerializer.ChatReplyResp, 1)
	err := pool.Submit(func() {
		response, err := chatService.GetChatService().Answer(ctx, body)
		if err != nil {
			resultChan <- &chatSerializer.ChatReplyResp{
				Error: err.Error(),
			}
			return
		}
		resultChan <- &chatSerializer.ChatReplyResp{
			ChatReply: response,
		}
	})
	if err != nil {
		exception.BadRequestException(c, err.Error())
		return
	}

	select {
	case result := <-resultChan:
		if result.Error != "" {
			exception.BadRequestException(c, result.Error)
			return
		}
		serializerResponse.OK(c, result.ChatReply)
	case <-ctx.Done():
		exception.TimeOutRequestException(c, "")
		return
	}

}

// EarphoneConfig
//
//	@Summary	AI 耳机配置接口
//	@Tags		AI
//	@Accept		application/json
//	@Produce	application/json
//	@Security	ApiKeyAuth
//	@Param		data	body		chatSerializer.GlassesConfigRequest	true	"data"
//	@Success	200		{object}	ent.ChatConfig
//	@Router		/chat/earphone-config [post]
func (h *Handler) EarphoneConfig(c *gin.Context) {
	var body *chatSerializer.GlassesConfigRequest
	if !validate.CheckBody(c, &body) {
		return
	}

	resp, err := new(forwardService.Service).GlassesConfigSave(body)
	if err != nil {
		exception.BadRequestException(c, err.Error())
		return
	}

	serializerResponse.Create(c, resp)
	return
}

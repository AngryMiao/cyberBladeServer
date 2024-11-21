package customer

import (
	"angrymiao-ai/app/handler"
	"angrymiao-ai/app/middleware"
	"github.com/gin-gonic/gin"
)

func ChatRouter(g *gin.RouterGroup, handler *handler.Handler) {
	chatGroup := g.Group("chat").Use(middleware.JWTAuthMiddleware())
	{
		chatGroup.POST("", handler.Customer.Chat.Chat)
		chatGroup.POST("earphone-config", handler.Customer.Chat.EarphoneConfig)
		chatGroup.POST("voice", handler.Customer.Chat.VoiceChat)
	}
}

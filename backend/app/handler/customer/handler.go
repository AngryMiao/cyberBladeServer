package customer

import (
	"angrymiao-ai/app/handler/customer/chat"
)

/*
统一的handler入口，小接口细分到各自的package去做处理，然后聚合。
*/

type Handler struct {
	Chat *chat.Handler
}

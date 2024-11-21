package chat

import (
	"angrymiao-ai/app/cache"
)

const (
	keyChatHistory         = "chat_history"
	TTLKeyTaskCreateSecond = 3600
)

type Cache struct {
}

func (*Cache) KeyPrefixCartItem(key string) string {
	return cache.KeyGenerate(keyChatHistory, key)
}

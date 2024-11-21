package chatService

import (
	"angrymiao-ai/app/cache"
	"angrymiao-ai/app/cache/chat"
	"angrymiao-ai/app/serializer/chatSerializer"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/sashabaranov/go-openai"
)

const (
	maxHistoryLength = 20
	keyChatHistory   = "chat_history"
)

func (s *Service) saveMessage(deviceCode string, message chatSerializer.HistoryChatMessage) error {
	hashKey := new(chat.Cache).KeyPrefixCartItem(deviceCode)
	// 获取当前历史记录
	history, err := s.getHistory(deviceCode)
	if err != nil {
		return err
	}

	// 添加新消息
	history = append(history, message)

	if len(history) > maxHistoryLength {
		history = history[len(history)-maxHistoryLength:]
	}

	historyJSON, err := json.Marshal(history)
	if err != nil {
		return err
	}

	// save Redis
	_, err = cache.HSet(keyChatHistory, hashKey, string(historyJSON))
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) getHistory(deviceCode string) ([]chatSerializer.HistoryChatMessage, error) {
	var (
		history = make([]chatSerializer.HistoryChatMessage, 0)
		hashKey = new(chat.Cache).KeyPrefixCartItem(deviceCode)
	)

	// 从 Redis 获取历史记录
	historyJSON, err := cache.HGet(keyChatHistory, hashKey)
	if err == redis.Nil {
		return history, nil
	} else if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(historyJSON), &history)
	if err != nil {
		return nil, err
	}

	return history, nil
}

func (s *Service) convertToOpenAIMessages(history []chatSerializer.HistoryChatMessage) []openai.ChatCompletionMessage {
	messages := make([]openai.ChatCompletionMessage, len(history))
	for i, msg := range history {
		var (
			content     = ""
			messagePart []openai.ChatMessagePart
		)
		if msg.Image != "" {
			messagePart = []openai.ChatMessagePart{
				{
					Type: openai.ChatMessagePartTypeText,
					Text: msg.Content,
				},
				{
					Type: openai.ChatMessagePartTypeImageURL,
					ImageURL: &openai.ChatMessageImageURL{
						URL:    msg.Image,
						Detail: openai.ImageURLDetailLow,
					},
				},
			}
		} else {
			content = msg.Content
		}
		messages[i] = openai.ChatCompletionMessage{
			Role:         msg.Role,
			Content:      content,
			MultiContent: messagePart,
		}
	}
	return messages
}

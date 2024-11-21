package chatService

import (
	"angrymiao-ai/app/serializer/chatSerializer"
	"angrymiao-ai/app/tools/openAiSdk"
	"angrymiao-ai/pkg"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
)

func (s *Service) Answer(ctx context.Context, body *chatSerializer.ChatBodyRequest) (*chatSerializer.ChatReply, error) {
	var (
		client = openAiSdk.NewClient()
	)

	chatData := chatSerializer.HistoryChatMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: *body.Message,
	}

	if body.Image != nil {
		// 图片先由后端上传服务器，后面改用前端上传处理响应更快。
		//chatData.Image = body.Image
		imageUrl, err := s.AnswerImageDeal(body.Image)
		if err != nil {
			return nil, err
		}
		chatData.Image = imageUrl
	}

	err := s.saveMessage(body.DeviceCode, chatData)

	if err != nil {
		return nil, err
	}

	// 获取并打印完整的对话历史
	history, err := s.getHistory(body.DeviceCode)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.OpenAiClient.CreateChatCompletion(ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT4oMini,
			Messages: s.convertToOpenAIMessages(history),
			Tools: []openai.Tool{
				s.forwardFunc(),
				//s.resetKeyboardFunc(),
			},
		},
	)
	//fmt.Println("resp", resp)

	if err != nil || len(resp.Choices) != 1 {
		return nil, errors.New(fmt.Sprintf("Completion error: err:%v len(choices):%v\n", err, len(resp.Choices)))
	}

	var (
		msg      = resp.Choices[0].Message
		funcName = ""
	)

	if msg.Content != "" {
		err = s.saveMessage(body.DeviceCode, chatSerializer.HistoryChatMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: msg.Content,
		})

		if err != nil {
			return nil, err
		}
	}

	if len(msg.ToolCalls) >= 1 {
		//fmt.Println("msg.ToolCalls", msg.ToolCalls)
		funcName = msg.ToolCalls[0].Function.Name
		funcArgs := msg.ToolCalls[0].Function.Arguments

		callResp, err := s.chatToolCalls(chatSerializer.ChatToolCall{
			Client:   client,
			Msg:      msg,
			History:  history,
			FuncName: funcName,
			FuncArgs: funcArgs,
		}, body)

		if err != nil {
			return nil, err
		}
		msg.Content = callResp
	}

	return &chatSerializer.ChatReply{
		ReplyMsg: msg.Content,
		FuncName: funcName,
	}, nil
}

func (s *Service) chatToolCalls(callArgs chatSerializer.ChatToolCall, body *chatSerializer.ChatBodyRequest) (string, error) {
	funcMap := s.CallFunc()
	replyMsg := ""
	var callData any

	switch callArgs.FuncName {
	case "forwardMsg":
		fm := funcMap[callArgs.FuncName].(func(map[string]string, string) any)

		var args = map[string]string{}
		json.Unmarshal([]byte(callArgs.FuncArgs), &args)

		callData = fm(args, body.DeviceCode)
		if pkg.IsErrorType(callData) {
			return "", callData.(error)
		}
		replyMsg = callData.(string)
	}

	//callArgs.History = append(callArgs.History, chatSerializer.HistoryChatMessage{
	//	Role:    openai.ChatMessageRoleUser,
	//	Content: replyMsg,
	//})
	//msg, err := s.callFuncChat(callArgs.Client.OpenAiClient, s.convertToOpenAIMessages(callArgs.History), callArgs.Msg, replyMsg)

	//fmt.Println("msgmsgmsgmsg", msg)
	//fmt.Println("msg.Content", msg.Content)
	if replyMsg != "" {
		err := s.saveMessage(body.DeviceCode, chatSerializer.HistoryChatMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: replyMsg,
		})

		if err != nil {
			return "", err
		}
	}

	return replyMsg, nil
}

// function calling 答案回复给open ai
func (s *Service) callFuncChat(client *openai.Client, dialogue []openai.ChatCompletionMessage, msg openai.ChatCompletionMessage, funcResp string) (openai.ChatCompletionMessage, error) {
	dialogue = append(dialogue, openai.ChatCompletionMessage{
		Role:       openai.ChatMessageRoleTool,
		Content:    funcResp,
		Name:       msg.ToolCalls[0].Function.Name,
		ToolCallID: msg.ToolCalls[0].ID,
	})

	fmt.Println("dialogue", dialogue)
	resp, err := client.CreateChatCompletion(context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT4oMini,
			Messages: dialogue,
		},
	)

	if err != nil || len(resp.Choices) != 1 {
		return msg, errors.New(fmt.Sprintf("2nd completion error: err:%v len(choices):%v\n", err, len(resp.Choices)))
	}

	return resp.Choices[0].Message, nil
}

func (s *Service) OpenAiText2Embedding(text []string) ([]openai.Embedding, error) {
	var (
		ctx    = context.Background()
		client = openAiSdk.NewClient()
	)
	queryReq := openai.EmbeddingRequest{
		Input: text,
		Model: openai.LargeEmbedding3,
	}

	resp, err := client.OpenAiClient.CreateEmbeddings(ctx, queryReq)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error creating query embedding:%s", err))
	}

	return resp.Data, nil
}

func (s *Service) CallFunc() map[string]any {
	ForwardFunc := s.Forward
	return map[string]any{
		"forwardMsg": ForwardFunc,
	}
}

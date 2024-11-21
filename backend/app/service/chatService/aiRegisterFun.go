package chatService

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

func (s *Service) forwardFunc() openai.Tool {
	params := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"allReturn": {
				Type:        jsonschema.String,
				Description: "根据最新的消息，需要把当前AI回答的消息转发给手机或者Discord里面",
			},
			"preReturn": {
				Type:        jsonschema.String,
				Description: "根据最新的消息，需要把当前AI回答的上一条消息转发给手机或者Discord里面",
			},
		},
		Required: []string{"forward"},
	}

	f := openai.FunctionDefinition{
		Name:        "forwardMsg",
		Description: "AI根据最新的消息回复判断是否需要转发到手机或者Discord里面",
		Parameters:  params,
	}

	return openai.Tool{
		Type:     openai.ToolTypeFunction,
		Function: &f,
	}
}

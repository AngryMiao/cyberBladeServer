package openAiSdk

import (
	"angrymiao-ai/config"
	"github.com/sashabaranov/go-openai"
)

var client *Client

type Client struct {
	OpenAiClient *openai.Client
	config       *config.OpenAiConfig
}

func initClient(c *Client) {
	c.config = config.Conf.OpenAi

	openAiConfig := openai.DefaultConfig(c.config.Key)
	openAiConfig.BaseURL = c.config.BaseUrl

	//if config.Conf.Mode == "local" {
	//	proxyUrl, err := url.Parse("http://localhost:7890")
	//	if err != nil {
	//		panic(err)
	//	}
	//	transport := &http.Transport{
	//		Proxy: http.ProxyURL(proxyUrl),
	//	}
	//	openAiConfig.HTTPClient = &http.Client{
	//		Transport: transport,
	//	}
	//}

	c.OpenAiClient = openai.NewClientWithConfig(openAiConfig)

	//c.OpenAiClient = openai.NewClient(c.config.Key)
}

func NewClient() *Client {
	if client == nil {
		client = new(Client)
		initClient(client)
	}

	return client
}

package discordSdk

import (
	"angrymiao-ai/config"
	"github.com/angrymiao/discordgo"
)

var client *Client

type Client struct {
	DiscordClient *discordgo.Session
	config        *config.DiscordConfig
}

func initClient(c *Client) error {
	c.config = config.Conf.Discord

	dClient, err := discordgo.New("Bot " + c.config.Token)
	c.DiscordClient = dClient
	if err != nil {
		return err
	}
	return nil
}

func NewClient() (*Client, error) {
	var err error
	if client == nil {
		client = new(Client)
		err = initClient(client)
	}

	return client, err
}

package chatService

import (
	"angrymiao-ai/app/ent"
	"angrymiao-ai/app/ent/chatconfig"
	"angrymiao-ai/app/model"
	"angrymiao-ai/app/tools/discordSdk"
	"context"
	"github.com/angrymiao/discordgo"
)

func (s *Service) Forward(args map[string]string, deviceCode string) any {
	var err error
	var pushMsg string
	if allReturn, isOk := args["allReturn"]; isOk {
		pushMsg, err = s.forwardLogic(allReturn, deviceCode)
	}

	if preReturn, isOk := args["preReturn"]; isOk {
		pushMsg, err = s.forwardLogic(preReturn, deviceCode)
	}

	if err != nil {
		return err
	}

	return pushMsg
}

func (s *Service) forwardLogic(msg, deviceCode string) (string, error) {
	var (
		client = model.DB
		ctx    = context.Background()
	)

	configObj, err := client.ChatConfig.Query().Where(chatconfig.DeviceCode(deviceCode)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return "", err
	}

	if configObj == nil {
		return "", nil
	}

	if configObj.DiscordUserID != nil {
		// Discord推送消息
		discordClient, err := discordSdk.NewClient()
		if err != nil {
			return "", err
		}

		dClient := discordClient.DiscordClient
		dClient.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildMembers
		channel, err := dClient.UserChannelCreate(*configObj.DiscordUserID)
		if err != nil {
			return "", err
		}

		go dClient.ChannelMessageSend(channel.ID, msg)
		return "消息已发送到Discord私信", err
	}

	return "没有可以发送的社交平台", nil
}

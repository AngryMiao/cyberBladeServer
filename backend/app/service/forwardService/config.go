package forwardService

import (
	"angrymiao-ai/app/ent"
	"angrymiao-ai/app/ent/chatconfig"
	"angrymiao-ai/app/model"
	"angrymiao-ai/app/serializer/chatSerializer"
	"context"
)

func (s *Service) GlassesConfigSave(args *chatSerializer.GlassesConfigRequest) (*ent.ChatConfig, error) {
	var (
		client = model.DB
		ctx    = context.Background()
	)

	configObj, err := client.ChatConfig.Query().Where(chatconfig.DeviceCode(args.DeviceCode)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if configObj == nil {
		configObj, err = client.ChatConfig.Create().
			SetDeviceCode(args.DeviceCode).
			SetForwardMode(chatconfig.ForwardMode(args.ForwardMode)).
			SetDiscordUserID(args.DiscordUserID).
			Save(ctx)
	} else {
		configObj, err = client.ChatConfig.UpdateOne(configObj).
			SetForwardMode(chatconfig.ForwardMode(args.ForwardMode)).
			SetDiscordUserID(args.DiscordUserID).
			Save(ctx)
	}

	if err != nil {
		return nil, err
	}
	return configObj, nil
}

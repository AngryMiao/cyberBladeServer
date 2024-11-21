package cmd

import (
	"angrymiao-ai/app/model"
	"angrymiao-ai/app/model/hook"
	"angrymiao-ai/config"
)

func InitModel(c *config.Config) {
	model.Init(c)
	hook.AddHook(model.DB)

}

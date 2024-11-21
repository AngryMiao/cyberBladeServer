package hook

import (
	"angrymiao-ai/app/ent"
)

func EntHook(client *ent.Client) {

}

func AddHook(client *ent.Client) {
	EntHook(client)
}

package schema

import (
	chatMapping "angrymiao-ai/app/mapping/chat"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// ChatConfig holds the schema definition for the ChatConfig entity.
type ChatConfig struct {
	ent.Schema
}

func (ChatConfig) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "chat_config"},
	}
}

func (ChatConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the ChatConfig.
func (ChatConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("device_code").Unique().NotEmpty(),
		field.String("discord_user_id").Optional().Nillable(),
		field.Enum("forward_mode").Values(chatMapping.ForwardMode...).Default(chatMapping.ForwardModeAll).Optional(),
	}
}

// Edges of the ChatConfig.
func (ChatConfig) Edges() []ent.Edge {
	return nil
}

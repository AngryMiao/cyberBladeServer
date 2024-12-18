// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChatConfigColumns holds the columns for the "chat_config" table.
	ChatConfigColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "device_code", Type: field.TypeString, Unique: true},
		{Name: "discord_user_id", Type: field.TypeString, Nullable: true},
		{Name: "forward_mode", Type: field.TypeEnum, Nullable: true, Enums: []string{"all", "media"}, Default: "all"},
	}
	// ChatConfigTable holds the schema information for the "chat_config" table.
	ChatConfigTable = &schema.Table{
		Name:       "chat_config",
		Columns:    ChatConfigColumns,
		PrimaryKey: []*schema.Column{ChatConfigColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChatConfigTable,
	}
)

func init() {
	ChatConfigTable.Annotation = &entsql.Annotation{
		Table: "chat_config",
	}
}

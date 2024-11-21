package model

import (
	"angrymiao-ai/app/ent"
	entMigrate "angrymiao-ai/app/ent/migrate"
	"angrymiao-ai/app/log"
	"angrymiao-ai/app/mapping"
	"angrymiao-ai/config"
	"context"
	_ "github.com/lib/pq"
)

// DB global db
var DB *ent.Client

func Init(c *config.Config) {
	DB = newDB(c.DB)

	if c.Mode == mapping.DevMode || c.Mode == mapping.DevLocal {
		//DB = DB.Debug()
	}

	migrate()

}

func newDB(c *config.DBConfig) *ent.Client {
	db, err := ent.Open(c.Dialect, c.DSN)

	if err != nil {
		log.Log.Panic("connect to db fail, err=%v", err)
	}

	//_, err = db.ExecContext(context.Background(), "CREATE EXTENSION IF NOT EXISTS vector")
	//if err != nil {
	//	panic(err)
	//}

	return db
}

func migrate() {
	// Run the auto migration tool.
	if err := DB.Schema.Create(
		context.Background(),
		entMigrate.WithDropIndex(true),
		entMigrate.WithDropColumn(true),
	); err != nil {
		log.Log.Fatalf("failed creating schema resources: %v", err)
	}
}

package main

import (
	"angrymiao-ai/app/auth"
	"angrymiao-ai/app/cache"
	"angrymiao-ai/app/log"
	"angrymiao-ai/app/router"
	"angrymiao-ai/app/router/customer"
	"angrymiao-ai/app/tools/pool"
	"angrymiao-ai/cmd"
	"angrymiao-ai/config"
	"flag"
)

//	@BasePath	/api

// @title						AngryMiao Customer API
// @version					0.0.1
// @description				angrymiao C端 后端接口.
// @description				<br />
// @description				auth 格式: Bearer TOKEN
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	load()
	run()
}

func load() {
	// 解析配置项
	flag.Parse()

	// log
	log.Init()

	// config
	config.InitCustomerConfig()

	// init and migrate db
	cmd.InitModel(config.Conf)

	// init cache
	cache.Init(config.Conf)

	// jwt
	auth.Init(config.Conf)
}

func run() {
	f := router.GetInstance()
	if err := f.Init(); err != nil {
		log.Log.Fatal("Framework initialization failed:", err)
	}

	// router
	r := customer.NewRouter(f)
	r.Init(f.GetEngine(), config.Conf.Mode)

	// pool
	if err := pool.InitPool(100); err != nil {
		log.Log.Fatalf("Failed to init pool: %v", err)
	}
	defer pool.Release()

	if err := f.Run(); err != nil {
		log.Log.Panic("Server running failed:, err=%v", err)
	}
}

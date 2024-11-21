package config

/*
	配置模板, 用来做统一配置的入口,
	将代码和配置分离，
*/

// InitCustomerConfig 配置c端
func InitCustomerConfig() {
	if confPath == "" {
		confPath = "./config/files/local/customer.yaml"
	}

	Init(confPath)
}

package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

func Init(confPath string) {
	v := readViperConfig(confPath)

	if err := v.Unmarshal(&Conf); err != nil {
		panic(fmt.Sprintf("Unmarshal config file fail, error:%s\n", err))
	}
}

func readViperConfig(confPath string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(confPath)

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return v
}


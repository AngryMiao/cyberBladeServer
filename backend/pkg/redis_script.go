package pkg

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"io/ioutil"
	"log"
	"os"
)

func LoadLuaScript(fileName string, client *redis.Client) string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("open %v %v.", fileName, err.Error()))
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err.Error())
	}

	script := redis.NewScript(string(bytes))

	sha, err := script.Load(client.Context(), client).Result()
	if err != nil {
		log.Fatalln(err)
	}
	return sha
}

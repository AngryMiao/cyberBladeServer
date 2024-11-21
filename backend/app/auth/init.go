package auth

import (
	"angrymiao-ai/app/log"
	"angrymiao-ai/config"
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
)

var (
	verifyKey *rsa.PublicKey

	issuer string
)

// Init 加载配置文件
func Init(c *config.Config) {

	issuer = c.JWT.Issuer

	verifyBytes, err := ioutil.ReadFile(c.JWT.PublicKeyPath)
	if err != nil {
		log.Log.Fatal(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Log.Fatal(err)
	}

}

package pkg

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/pbkdf2"
	"log"
	"strconv"
	"strings"
)

func HashAndSaltPassword(password string) string {
	// sha 值
	randByte := make([]byte, 8)
	_, err := rand.Read(randByte)
	if err != nil {
		log.Println(err)
	}
	base64RandByte := base64.StdEncoding.EncodeToString(randByte)
	salt := []byte(base64RandByte)

	//迭代次数
	iterTimes := 100000

	dk := pbkdf2.Key([]byte(password), salt, iterTimes, 32, sha256.New)

	hashedPassword := "pbkdf2_sha256$100000$" + string(salt) + "$" + base64.StdEncoding.EncodeToString(dk)
	return hashedPassword
}

func GenerateRandomPassword() string {
	result, _ := password.Generate(16, 6, 4, false, false)
	return result
}

func ComparePassword(HashedPassword string, password string) bool {

	s := strings.Split(HashedPassword, "$")

	salt := []byte(s[2])

	// saved password iteration value should be converted to int
	iter, _ := strconv.Atoi(s[1])

	dk := pbkdf2.Key([]byte(password), salt, iter, 32, sha256.New)

	hashedPW := "pbkdf2_sha256$100000$" + s[2] + "$" + base64.StdEncoding.EncodeToString(dk)

	if subtle.ConstantTimeCompare([]byte(HashedPassword), []byte(hashedPW)) == 0 {
		return false
	}

	return true
}

package pkg

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"math/rand"
	"time"
)

var capitalLetters = []rune("ABCDEFGHJKLMNPQRSTUVWXYZ")
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numbers = []rune("0123456789")
var lettersNumbers = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

var EasyAlpha = []rune("ABCDEFGHJKLMNPQRSTUVWXYZ0123456789")

func randomCode(n int, keys []rune) string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	keySizes := len(keys)
	b := make([]rune, n)

	for i := range b {
		b[i] = keys[random.Intn(keySizes)]
	}
	return string(b)
}

func GenerateRandomNumber(n int) string {
	return randomCode(n, numbers)
}

func GenerateRandomKey(n int) string {
	return randomCode(n, lettersNumbers)
}

func GenerateRandomCapitalLetter(n int) string {
	return randomCode(n, capitalLetters)
}

func GenerateRandomKeyWithNanoid(n int) string {
	code, err := gonanoid.New(n)
	if err != nil {
		code = GenerateRandomKey(n)
	}
	return code
}

func GenerateAlphaRandomKeyWithNanoid(n int, keys []rune) string {
	code, err := gonanoid.Generate(string(keys), n)
	if err != nil {
		code = randomCode(n, keys)
	}
	return code
}

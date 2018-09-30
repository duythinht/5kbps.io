package vault

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestEncryptText(t *testing.T) {
	k := GenerateEncryptionKey()
	for i := 10; i < 60; i++ {
		text := randString(i)
		ct, err := EncryptText(text, k)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		if !strings.HasPrefix(ct, CipherTextBegin) || !strings.HasSuffix(ct, TextEnd) {
			t.Log(ct)
			t.Fail()
		}

		decrypted, err := DecryptText(ct, k)

		if err != nil {
			t.Log(err)
			t.Fail()
		}
		if decrypted != text {
			t.Log(text, decrypted)
			t.Fail()
		}
	}
}

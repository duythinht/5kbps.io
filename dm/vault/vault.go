package vault

import (
	"encoding/base64"
	"io"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

const (
	CipherTextBegin = "dm::cipher::Text["
	PlainTextBegin  = "dm::plain::Text["
	TextEnd         = "]"
)

func EncryptText(t string, k string) (string, error) {

	key, err := base64.StdEncoding.DecodeString(k)

	if err != nil {
		return "", nil
	}

	ciphertext, err := encrypt([]byte(t), key)

	if err != nil {
		return "", err
	}

	return CipherTextBegin + base64.StdEncoding.EncodeToString(ciphertext) + TextEnd, nil
}

func DecryptText(t string, k string) (string, error) {
	key, err := base64.StdEncoding.DecodeString(k)

	if err != nil {
		return "", nil
	}

	ciphertext, err := base64.StdEncoding.DecodeString(t[len(CipherTextBegin) : len(t)-len(TextEnd)])

	if err != nil {
		return "", err
	}

	plaintext, err := decrypt(ciphertext, key)

	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func EncryptFile(f string, k string) string {
	return ""
}

func DecryptFile(f io.Reader, k string) (string, error) {
	buff, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	data := make(map[interface{}]interface{}, 0)
	yaml.Unmarshal(buff, data)

	decrypted, err := inspectDecrypt(data, k)

	if err != nil {
		return "", err
	}

	out, err := yaml.Marshal(decrypted)
	return string(out), err
}

func GenerateEncryptionKey() string {
	key := newEncryptionKey()
	return base64.StdEncoding.EncodeToString(key)
}

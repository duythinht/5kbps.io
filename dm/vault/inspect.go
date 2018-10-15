package vault

import (
	"strings"
)

func inspectDecrypt(m interface{}, key string) (m1 interface{}, err error) {
	switch mt := m.(type) {
	case map[interface{}]interface{}:
		for k, v := range mt {
			mt[k], err = inspectDecrypt(v, key)
			if err != nil {
				return nil, err
			}
		}
		return mt, nil
	case []interface{}:
		for i, v := range mt {
			mt[i], err = inspectDecrypt(v, key)
			if err != nil {
				return nil, err
			}
		}
		return mt, nil
	case string:
		s := strings.TrimSpace(mt)
		if strings.HasPrefix(s, CipherTextBegin) && strings.HasSuffix(s, TextEnd) {
			return DecryptText(mt, key)
		}
	}

	return m, nil
}

func inspectEncrypt(m interface{}, key string) (m1 interface{}, err error) {
	switch mt := m.(type) {
	case map[interface{}]interface{}:
		for k, v := range mt {
			mt[k], err = inspectEncrypt(v, key)
			if err != nil {
				return nil, err
			}
		}
		return mt, nil
	case []interface{}:
		for i, v := range mt {
			mt[i], err = inspectEncrypt(v, key)
			if err != nil {
				return nil, err
			}
		}
		return mt, nil
	case string:
		s := strings.TrimSpace(mt)
		if strings.HasPrefix(s, PlainTextBegin) && strings.HasSuffix(s, TextEnd) {
			return EncryptText(mt, key)
		}
	}

	return m, nil
}

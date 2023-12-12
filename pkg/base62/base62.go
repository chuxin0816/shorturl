package base62

import (
	"bytes"
	"errors"
)

const base62Code = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var errInvalidChar = errors.New("invalid base62 char")

func Encode(num uint64) string {
	if num == 0 {
		return string(base62Code[0])
	}
	var buf bytes.Buffer
	for num > 0 {
		buf.WriteByte(base62Code[num%62])
		num /= 62
	}

	// reverse buf
	b := buf.Bytes()
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}

	return string(b)
}

func Decode(str string) (uint64, error) {
	var res uint64
	for _, c := range str {
		num, err := base62Index(c)
		if err != nil {
			return 0, err
		}
		res = res*62 + num
	}
	return res, nil
}

func base62Index(r rune) (uint64, error) {
	switch {
	case '0' <= r && r <= '9':
		return uint64(r - '0'), nil
	case 'a' <= r && r <= 'z':
		return uint64(r - 'a' + 10), nil
	case 'A' <= r && r <= 'Z':
		return uint64(r - 'A' + 36), nil
	default:
		return 0, errInvalidChar
	}
}

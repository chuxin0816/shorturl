package base62

import (
	"bytes"
	"errors"
)

const (
	base62Code   = "n1F6U5r48dgAEOWeCwTbjhSI3aBzHscmiRGQ0luX9pvJy2kDKfoMNPq7YxVLtZ"
	base62System = 62
)

var (
	errInvalidChar = errors.New("invalid base62 char")
	base62Map      = initBase62Map()
)

func Encode(num uint64) string {
	if num == 0 {
		return string(base62Code[0])
	}
	var buf bytes.Buffer
	for num > 0 {
		buf.WriteByte(base62Code[num%62])
		num /= base62System
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

func initBase62Map() map[rune]uint64 {
	m := make(map[rune]uint64)
	for i, c := range base62Code {
		m[c] = uint64(i)
	}
	return m
}

func base62Index(r rune) (uint64, error) {
	if v, ok := base62Map[r]; ok {
		return v, nil
	}
	return 0, errInvalidChar
}

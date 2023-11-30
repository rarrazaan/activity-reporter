package helper

import (
	"math/rand"
	"mini-socmed/internal/cons"
	"strings"
	"time"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

type (
	RandomString interface {
		RandStringBytesMaskImprSrcSB() string
	}
	randomString struct{}
)

var src = rand.NewSource(time.Now().UnixNano())

func (s *randomString) RandStringBytesMaskImprSrcSB() string {
	sb := strings.Builder{}
	sb.Grow(cons.PostIDStringLength)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := cons.PostIDStringLength-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func NewRandomString() RandomString {
	return &randomString{}
}

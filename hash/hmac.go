package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"
)

// NewHMAC returns an new hash-based message authentication code
func NewHMAC(key string) HMAC {
	h := hmac.New(sha256.New, []byte(key))
	return HMAC{
		hmac: h,
	}
}

// HMAC holds the hash-bashed message data
type HMAC struct {
	hmac hash.Hash
}

// Hash generates a HMAC
func (h HMAC) Hash(input string) string {
	h.hmac.Reset()
	err := h.hmac.Write([]byte(input))
	if err != nil {
		fmt.Printf("ErrorWithHmac %v", err)
	}
	b := h.hmac.Sum(nil)
	return base64.URLEncoding.EncodeToString(b)
}

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func hmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	// yorktodo為什麼要h.sum
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	fmt.Println(hmacSha256("test", "thisissecrety"))
}

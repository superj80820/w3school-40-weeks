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
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	sharedSecret := "小明與早餐店阿姨的共同鑰匙"
	badGuySecret := "壞人的鑰匙"
	meals := "小明餐點: 大冰紅"

	// 壞人利用自己的鑰匙產生HMAC
	badGuyHMAC := hmacSha256(meals, badGuySecret)

	// 早餐店阿姨利用與小明共同的鑰匙產生HMAC
	trueHMAC := hmacSha256(meals, sharedSecret)

	// 早餐店阿姨比對此兩個HMAC，發現不同，故此訊息不是小明傳送的
	if badGuyHMAC != trueHMAC {
		fmt.Println("Two HMACs are not the same.")
		return
	}
}

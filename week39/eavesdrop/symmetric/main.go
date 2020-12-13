// 大量參考: https://gist.github.com/yingray/57fdc3264b1927ef0f984b533d63abab
package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

func Ecrypt(plaintext string, key []byte, iv []byte, blockSize int) string {
	pad := func(ciphertext []byte, blockSize int, after int) []byte {
		padding := (blockSize - len(ciphertext)%blockSize)
		padtext := bytes.Repeat([]byte{byte(padding)}, padding)
		return append(ciphertext, padtext...)
	}

	bPlaintext := pad([]byte(plaintext), blockSize, len(plaintext))
	block, _ := aes.NewCipher(key)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext)
}

func Decrypt(ciphertext string, key []byte, iv []byte) string {
	unpad := func(ciphertext []byte) []byte {
		length := len(ciphertext)
		unpadding := int(ciphertext[length-1])
		return ciphertext[:(length - unpadding)]
	}

	decodeData, _ := hex.DecodeString(ciphertext)
	block, _ := aes.NewCipher(key)
	blockMode := cipher.NewCBCDecrypter(block, iv)
	originData := make([]byte, len(decodeData))
	blockMode.CryptBlocks(originData, decodeData)
	return string(unpad(originData))
}

func main() {
	// 小明與早餐店阿姨共同的鑰匙
	key := []byte("di93bi39a^*(2i$2ajg9^ha9fj@hswe(")
	// 這邊比較特別一點，由於是使用CBC演算法，所以在加密與解密時會多一個隨機數iv，
	// 這可以讓「相同的明文加密後，會產生不同的加密訊息」，以避免壞人透過相同的加密訊息來推斷資訊
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		log.Fatal(err)
	}

	// 小明透過鑰匙加密訊息
	plaintext := "小明的付款密碼: 12345"
	ecryptMsg := Ecrypt(plaintext, key, iv, aes.BlockSize)

	// 早餐店阿姨透過鑰匙解密訊息
	decryptMsg := Decrypt(ecryptMsg, key, iv)
	fmt.Println("早餐店阿姨", decryptMsg)
}

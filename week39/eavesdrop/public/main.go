// 大量參考: https://gist.github.com/mfridman/c0c5ece512f63d429c4589196a1d4242
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

// LoadFile load the file to bytes
func LoadFile(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

// BytesToPrivateKey bytes to private key
func BytesToPrivateKey(priv []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		log.Fatal(err)
	}
	return key
}

// BytesToPublicKey bytes to public key
func BytesToPublicKey(pub []byte) *rsa.PublicKey {
	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		log.Fatal(err)
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		log.Fatal("not ok")
	}
	return key
}

// EncryptWithPublicKey encrypts data with public key
func EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) []byte {
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		log.Fatal(err)
	}
	return ciphertext
}

// DecryptWithPrivateKey decrypts data with private key
func DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) []byte {
	hash := sha512.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
	if err != nil {
		log.Fatal(err)
	}
	return plaintext
}

func main() {
	// 早餐店阿姨產生公私鑰
	privateKey := BytesToPrivateKey(LoadFile("./key"))
	publicKey := &privateKey.PublicKey

	// 小明獲得早餐店阿姨的公鑰，利用此公鑰加密
	encryptedMsg := EncryptWithPublicKey([]byte("小明的付款密碼: 12345"), publicKey)

	// 小明將 encryptedMsg 傳送給早餐店阿姨

	// 早餐店阿姨使用私鑰解開此訊息
	msg := DecryptWithPrivateKey(encryptedMsg, privateKey)
	fmt.Println(string(msg))
}

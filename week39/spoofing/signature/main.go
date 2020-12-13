// 大量參考: https://gist.github.com/mfridman/c0c5ece512f63d429c4589196a1d4242
package main

import (
	"crypto"
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

func main() {
	// 壞人的私鑰
	badGuyPrivateKey := BytesToPrivateKey(LoadFile("./badGuyKey"))
	// 小明的公鑰
	goodGuyPublicKey := BytesToPrivateKey(LoadFile("./goodGuyKey")).PublicKey

	// 壞人開始偽造小明的訊息
	messageBytes := []byte("小明餐點: 大冰紅")
	hash := sha512.New()
	hash.Write(messageBytes)
	hashed := hash.Sum(nil)

	// 壞人用自己的私鑰簽名，並非小明的
	signature, err := rsa.SignPKCS1v15(rand.Reader, badGuyPrivateKey, crypto.SHA512, hashed)
	if err != nil {
		panic(err)
	}

	// 早餐店阿姨取得小明的公鑰，利用此公鑰驗證之後發現不是小明傳的訊息
	err = rsa.VerifyPKCS1v15(&goodGuyPublicKey, crypto.SHA512, hashed, signature)
	if err != nil {
		fmt.Println("Two signatures are not the same. Error: ", err)
		return
	}
}

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
	goodGuyPrivateKey := BytesToPrivateKey(LoadFile("./goodGuyKey"))
	// 小明的公鑰，公鑰可以透過私要來取得，所以這邊就不在載入公鑰檔案了
	goodGuyPublicKey := goodGuyPrivateKey.PublicKey

	// 小明用自己的私鑰對訊息簽章
	messageBytes := []byte("小明餐點: 大冰奶")
	hash := sha512.New()
	hash.Write(messageBytes)
	hashed := hash.Sum(nil)

	// 小明用自己的私鑰簽名
	signature, err := rsa.SignPKCS1v15(rand.Reader, goodGuyPrivateKey, crypto.SHA512, hashed)
	if err != nil {
		panic(err)
	}

	// 早餐店阿姨取得小明的公鑰，驗證後發現的確是小明，傳送餐點回去給小明
	err = rsa.VerifyPKCS1v15(&goodGuyPublicKey, crypto.SHA512, hashed, signature)
	if err != nil {
		fmt.Println("Two signatures are not the same. Error: ", err)
		return
	}
	fmt.Println("Verify the signature is correct")

	// 小明獲得餐點，並且吃完後開市賴帳

	// 早餐店阿姨說明當初`rsa.VerifyPKCS1v15`利用小明的公鑰驗證後的確是小明用私鑰簽章的，故證明小明確實有點過餐
}

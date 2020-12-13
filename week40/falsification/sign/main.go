package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"week39/utils"
)

func main() {
	// privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	// if err != nil {
	// 	fmt.Printf("rsa.GenerateKey: %v\n", err)
	// 	return
	// }

	// message := "Hello World!"
	// messageBytes := bytes.NewBufferString(message)
	// hash := sha512.New()
	// hash.Write(messageBytes.Bytes())
	// digest := hash.Sum(nil)

	// fmt.Printf("messageBytes: %v\n", messageBytes)
	// fmt.Printf("hash: %V\n", hash)
	// fmt.Printf("digest: %v\n", digest)

	// signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA512, digest)
	// if err != nil {
	// 	fmt.Printf("rsa.SignPKCS1v15 error: %v\n", err)
	// 	return
	// }

	// fmt.Printf("signature: %v\n", signature)

	// err = rsa.VerifyPKCS1v15(&privateKey.PublicKey, crypto.SHA512, digest, signature)
	// if err != nil {
	// 	fmt.Printf("rsa.VerifyPKCS1v15 error: %V\n", err)
	// }

	// fmt.Println("Signature good!")

	privateKey := utils.BytesToPrivateKey(utils.LoadPrivateKey("../test"))

	messageBytes := []byte("Hello World!")
	hash := sha256.New()
	hash.Write(messageBytes)
	// yorktodo為什麼要sum
	hashed := hash.Sum(nil)
	// message := []byte("test")
	// // yorktodo: hashed 需查證
	// hashed := sha256.Sum256(message)

	// signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		panic(err)
	}
	fmt.Println(signature)

	// york todo 256/512的差別是？
	// err = rsa.VerifyPKCS1v15(&privateKey.PublicKey, crypto.SHA512, hashed[:], signature)
	err = rsa.VerifyPKCS1v15(&privateKey.PublicKey, crypto.SHA256, hashed, signature)
	if err != nil {
		panic(err)
	}

	fmt.Println("Signature good!")
}

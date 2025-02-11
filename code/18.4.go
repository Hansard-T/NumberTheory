package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
)

func generateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	pubKey := &privKey.PublicKey
	return privKey, pubKey, nil
}

func encryptMessage(pubKey *rsa.PublicKey, message string) (string, error) {
	messageBytes := []byte(message)

	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, messageBytes, nil)
	if err != nil {
		return "", err
	}

	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)
	return encodedCiphertext, nil
}

func decryptMessage(privKey *rsa.PrivateKey, encryptedMessage string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedMessage)
	if err != nil {
		return "", err
	}

	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privKey, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func main() {
	privKey, pubKey, err := generateRSAKeyPair(2048)
	if err != nil {
		log.Fatalf("密钥生成失败: %v", err)
	}

	originalMessage := "Hello world!"

	encryptedMessage, err := encryptMessage(pubKey, originalMessage)
	if err != nil {
		log.Fatalf("加密失败: %v", err)
	}
	fmt.Println("加密后的消息: ")
	fmt.Println(encryptedMessage)
	decryptedMessage, err := decryptMessage(privKey, encryptedMessage)
	if err != nil {
		log.Fatalf("解密失败: %v", err)
	}
	fmt.Println("解密后的消息: ")
	fmt.Println(decryptedMessage)
}
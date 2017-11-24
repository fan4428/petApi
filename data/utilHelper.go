package data

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

//Encrypt AesCBC 加密
func Encrypt(plantText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key) //选择加密算法
	if err != nil {
		return nil, err
	}
	fmt.Println("block:", block)
	plantText = PKCS7Padding(plantText, block.BlockSize())

	blockModel := cipher.NewCBCEncrypter(block, key)

	ciphertext := make([]byte, len(plantText))

	blockModel.CryptBlocks(ciphertext, plantText)
	return ciphertext, nil
}

//Dncrypt AesCBC 解密
func Dncrypt(plantText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key) //选择加密算法
	if err != nil {
		return nil, err
	}
	fmt.Println("block:", block)
	// plantText = PKCS7Padding(plantText, block.BlockSize())
	blockSize := block.BlockSize()

	blockModel := cipher.NewCBCDecrypter(block, key[:blockSize])

	ciphertext := make([]byte, len(plantText))

	blockModel.CryptBlocks(ciphertext, plantText)
	return ciphertext, nil
}

//PKCS7Padding PKCS7规范
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

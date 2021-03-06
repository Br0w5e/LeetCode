package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"
)

func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}

func encryptAES(src []byte, key []byte)([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	src = padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src, nil
}

func decryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = unpadding(src)
	return src, nil
}

func unpadding(src []byte) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	return src[:n-unPadNum]
}

func main()  {
	d := []byte("hello, aes")
	key := []byte("hgfedcba87654321")
	fmt.Println("加密前:", string(d))
	x1, err := encryptAES(d, key)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("加密后：", string(x1))
	x2, err := decryptAES(x1, key)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("解密后:", string(x2))
}

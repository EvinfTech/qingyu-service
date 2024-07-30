package util

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"errors"
)

func padding(src []byte) []byte {
	//填充个数
	padding := aes.BlockSize - len(src)%aes.BlockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, paddingText...)
}

func unPadding(src []byte) []byte {
	size := len(src)
	if (size - int(src[size-1])) > 0 {
		return src[:(size - int(src[size-1]))]
	}
	return src
}

// Encrypt 加密
func Encrypt(key []byte, src []byte) ([]byte, error) {
	//key只能是 16 24 32长度
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//padding
	src = padding(src)
	//返回加密结果
	encryptData := make([]byte, len(src))
	//存储每次加密的数据
	tmpData := make([]byte, 16)

	//分组分块加密
	for index := 0; index < len(src); index += 16 {
		block.Encrypt(tmpData, src[index:index+16])
		copy(encryptData[index:index+16], tmpData)
	}
	return encryptData, nil
}

// Decrypt 解密
func Decrypt(key []byte, src []byte) ([]byte, error) {
	//key只能是 16 24 32长度
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(src)%aes.BlockSize != 0 {
		return nil, errors.New("AES decrypt data length error")
	}
	//返回解密结果
	decryptData := make([]byte, len(src))
	//存储每次解密的数据
	tmpData := make([]byte, 16)

	//分组分块解密
	for index := 0; index < len(src); index += 16 {
		block.Decrypt(tmpData, src[index:index+16])
		copy(decryptData[index:index+16], tmpData)
	}
	resData := unPadding(decryptData)
	return resData, nil
}

// EncryptString 加密字符串
func EncryptString(key string, src string) string {
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return ""
	}
	b, err := Encrypt(keyBytes, []byte(src))
	if err != nil {
		return ""
	}
	bhex := hex.EncodeToString(b)
	return bhex
}

// DecryptString 解密字符串
func DecryptString(key string, src string) (string, error) {
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}
	decodeBytes, err := hex.DecodeString(src)
	if err != nil {
		return "", err
	}

	dataBytes, err := Decrypt(keyBytes, decodeBytes)
	return string(dataBytes), err
}

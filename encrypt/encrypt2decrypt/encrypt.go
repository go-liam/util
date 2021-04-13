package encrypt2decrypt

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"errors"
	"strconv"
)

/*
	key := []byte("2fa6c1e9")
	str :="1234567890"
	strEncrypted, err := Encrypt(str, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encrypted:", strEncrypted)
	strDecrypted, err := Decrypt(strEncrypted, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decrypted:", strDecrypted)
*/



const defaultKey  =  "9fa6c8e0"
//var key = []byte(defaultKey)

func DecryptedID(st string,key string) string {
	if key ==""{
		key = defaultKey
	}
	key1 := []byte(key)
	strDecrypted, err := Decrypt(st, key1)
	if err != nil {
		return "0"
	}
	return strDecrypted
}

func EncryptedID(id int64,key string) string {
	if key ==""{
		key = defaultKey
	}
	key1 := []byte(key)
	str := strconv.FormatInt(id, 10)
	strEncrypted, err := Encrypt(str, key1)
	if err != nil {
		return ""
	}
	return strEncrypted
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

func Encrypt(text string, key []byte) (string, error) {
	src := []byte(text)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

func Decrypt(decrypted string, key []byte) (string, error) {
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out), nil
}

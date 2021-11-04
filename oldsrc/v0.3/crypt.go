package crypt

import (
	"crypto/aes"
	"crypto/des"
	"encoding/base64"
	"fmt"
)

func EncryptData(data []byte, algo string, encode string, key []byte) []byte {
	// Encrypts the data using the key
	// encodes if encode == base64 or == hex
	// returns the encrypted data
	// algo is the algorithm to use
	// output is the output file if specified
	// key is the key to use
	var output []byte
	if encode == "base64" {
		output = base64Encode(data)
	} else if encode == "hex" {
		output = hexEncode(data)
	} else {
		output = data
	}

	if algo == "aes" {
		output = aesEncrypt(data, []byte(key))
	} else if algo == "des" {
		output = desEncrypt(output, []byte(key))
	} else {
		output = []byte("Error: Invalid algorithm")
	}

	return output

}

func aesEncrypt(data []byte, key []byte) []byte {
	c, err := aes.NewCipher(key)
	CheckError(err)
	out := make([]byte, len(data))
	c.Encrypt(out, []byte(data))
	return out
}

func aesDecrypt(ct []byte, key []byte) []byte {
	c, err := aes.NewCipher(key)
	CheckError(err)
	out := make([]byte, len(ct))
	c.Decrypt(out, ct)
	return out
}

func desEncrypt(data []byte, key []byte) []byte {
	c, err := des.NewCipher(key)
	CheckError(err)
	out := make([]byte, len(data))
	c.Encrypt(out, []byte(data))
	return out
}

func desDecrypt(ct []byte, key []byte) []byte {
	c, err := des.NewCipher(key)
	CheckError(err)
	out := make([]byte, len(ct))
	c.Decrypt(out, ct)
	return out
}

func base64Encode(data []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(data))
}

func hexEncode(data []byte) []byte {
	return []byte(fmt.Sprintf("%x", data))
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		//os.Exit(1)
	}
}

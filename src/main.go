package main

import (
	"crypto/aes"
	"crypto/des"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	if argExists("--help") || argExists("-h") || argExists("-?") {
		println(getHelp())
		os.Exit(0)
	}

	if argExists("--encrypt") {
		encrypt := []byte(handleArg("--encrypt"))
		algo := handleArg("-t")
		encode := handleArg("--encode")
		length, err := strconv.Atoi(handleArg("-l"))
		CheckError(err)
		key := generateKey(length)
		if encrypt == nil {
			println("Error: No data provided")
			println(getHelp())
			os.Exit(1)
		}
		if algo == "" {
			println("Error: No encryption algorithm provided")
			println(getHelp())
			os.Exit(1)
		}
		if encode == "" {
			encode = "false"
		}
		encryptedData := encryptData(encrypt, algo, encode, length, key)
		fmt.Println(string(encryptedData) + " " + string(key))
	}

	if argExists("--decrypt") { // NOT IMPLEMENTED YET
		decrypt := []byte(handleArg("--decrypt"))
		key := []byte(handleArg("--key"))
		if decrypt == nil {
			println("Error: No data provided")
			println(getHelp())
			os.Exit(1)
		}
		if key == nil {
			println("Error: No key provided")
			println(getHelp())
			os.Exit(1)
		}

	}
}

func argExists(arg string) bool {
	args := os.Args
	for i := 0; i < len(args); i++ {
		if strings.Contains(args[i], arg) {
			return true
		}
	}
	return false
}

func handleArg(arg string) string {
	args := os.Args
	var data string = ""
	for i := 0; i < len(args); i++ {
		if strings.Contains(args[i], arg) {
			data = args[i+1]
		}
	}
	return data
}

func getHelp() string {
	return "Usage: ./sed <command>\n" +
		"Commands:\n" +
		"\tCryptography Utility:\n" +
		"\t\t--encrypt <data>\n" +
		"\t\t\tEncrypts data and outputs keys\n" +
		"\t\t-t <encryption algorithm>\n" +
		"\t\t\tSpecify encryption algorithm\n" +
		"\t\t\t\taes / des\n" +
		"\t\t-l <len>\n" +
		"\t\t\tSets the key length for generating the key\n" +
		"\t\t--encode <hex/base64>\n" +
		"\t\t\tEncodes the data with base64 (Use if data is too large)\n" +
		"\t\t--decrypt <data>\n" +
		"\t\t\tDecrypts data with RSA (Requires private key)\n" +
		"\t\t--key <key>\n" +
		"\t\t\tSets the private key for crypting\n" +
		"\t\t--decode\n" +
		"\t\t\tDecodes the data with base64 (Use if data is encoded)\n" +
		"\t\t-o <file/dir>\n" +
		"\t\t\tSets the output file or directory for the data and keys\n"

}

/************************
 * Cryptography Utility	*
 ************************/

func encryptData(data []byte, algo string, encode string, length int, key []byte) []byte {
	// Encrypts the data using the key
	// encodes if encode == base64 or == hex
	// returns the encrypted data
	// algo is the algorithm to use
	// output is the output file if specified
	// length is the length of the key to generate
	// if key == nil then it will generate a key
	if algo == "aes" {
		if key == nil {
			key = generateKey(length)
		}
		if encode == "base64" {
			return base64Encode(aesEncrypt(data, key))
		} else if encode == "hex" {
			return hexEncode(aesEncrypt(data, key))
		} else {
			return aesEncrypt(data, key)
		}
	}
	if algo == "des" {
		if key == nil {
			key = generateKey(length)
		}
		if encode == "base64" {
			return base64Encode(desEncrypt(data, key))
		} else if encode == "hex" {
			return hexEncode(desEncrypt(data, key))
		} else {
			return desEncrypt(data, key)
		}
	}
	return nil
}

func decryptData(data []byte, algo string, encode string, key []byte) []byte {
	// Decrypts the data using the key
	// encodes if encode == base64 or == hex
	// returns the decrypted data
	// algo is the algorithm to use
	// output is the output file if specified
	// key is the key to use
	if algo == "aes" {
		if encode == "base64" {
			return aesDecrypt(base64Decode(data), key)
		} else if encode == "hex" {
			return aesDecrypt(hexDecode(data), key)
		} else {
			return aesDecrypt(data, key)
		}
	}
	if algo == "des" {
		if encode == "base64" {
			return desDecrypt(base64Decode(data), key)
		} else if encode == "hex" {
			return desDecrypt(hexDecode(data), key)
		} else {
			return desDecrypt(data, key)
		}
	}
	return nil
}

func generateKey(length int) []byte {
	key := make([]byte, length)
	return key
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

/***********************
 * Encoding / Decoding *
 ***********************/

func base64Encode(data []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(data))
}

func base64Decode(data []byte) []byte {
	decoded, err := base64.StdEncoding.DecodeString(string(data))
	CheckError(err)
	return decoded
}

func hexEncode(data []byte) []byte {
	return []byte(fmt.Sprintf("%x", data))
}

func hexDecode(data []byte) []byte {
	return []byte(fmt.Sprintf("%s", data))
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		//os.Exit(1)
	}
}

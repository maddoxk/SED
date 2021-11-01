package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {

	if handleArg("-h") != "Argument not found" {
		fmt.Println(getHelp())
		os.Exit(0)
	}

	for _, arg := range os.Args[1:] {
		if arg == "-d" {
			if handleArg("-o") != "Argument not found" {
				ioutil.WriteFile(handleArg("-o"), []byte(durl(handleArg("-d"))), 0644)
			} else {
				fmt.Println(durl(handleArg("-d")))
			}
		} else if arg == "--encrypt" {
			data, err := ioutil.ReadFile(handleArg("--encrypt"))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			var keyLen int
			keyLen, err = strconv.Atoi(handleArg("-l"))
			if err != nil {
				fmt.Println(err)
			}
			privateKey, publicKey := generateKeys(keyLen)
			if handleArg("--encode") != "Argument not found" {
				data = []byte(base64.StdEncoding.EncodeToString(data))
			}
			cipheredData := encrypt(string(data), publicKey)
			if handleArg("-o") != "Argument not found" {
				ioutil.WriteFile(handleArg("-o"), []byte(cipheredData), 0644)
				ioutil.WriteFile(handleArg("-o")+".pub", []byte(base64.StdEncoding.EncodeToString(publicKey.N.Bytes())), 0644)
				ioutil.WriteFile(handleArg("-o")+".pri", []byte(base64.StdEncoding.EncodeToString(privateKey.D.Bytes())), 0644)
			} else {
				fmt.Println(cipheredData)
			}
		} else if arg == "--decrypt" {
			rawData, err := ioutil.ReadFile(handleArg("--decrypt"))
			data := string(rawData)
			decodedData, err := base64.StdEncoding.DecodeString(data)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			block, _ := pem.Decode([]byte(handleArg("--key")))
			privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
			decryptedData := decrypt(decodedData, privateKey)
			if handleArg("--decode") != "Argument not found" {
				decryptedData, err := base64.StdEncoding.DecodeString(decryptedData)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				if handleArg("-o") != "Argument not found" {
					ioutil.WriteFile(handleArg("-o"), decryptedData, 0644)
				} else {
					fmt.Println(decryptedData)
				}
			} else {
				if handleArg("-o") != "Argument not found" {
					ioutil.WriteFile(handleArg("-o"), []byte(decryptedData), 0644)
				} else {
					fmt.Println(decryptedData)
				}
			}
		}
	}
}

func handleArg(arg string) string {
	args := os.Args
	for i := 1; i < len(args); i++ {
		if args[i] == arg {
			if i+1 < len(args) {
				return args[i+1]
			} else {
				return "Invalid number of arguments"
			}
		}
	}
	return "Argument not found"
}

func durl(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "Error"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error"
	}
	//Encrypt the body with RSA
	data := string(body)
	return data

	//generateKey(1024))
}

func generateKeys(bitLen int) (*rsa.PrivateKey, *rsa.PublicKey) {
	//Generate a RSA key
	privatekey, Error := rsa.GenerateKey(rand.Reader, bitLen)
	if Error != nil {
		fmt.Println("Error generating key")
		os.Exit(1)
	}
	publickey := &privatekey.PublicKey
	return privatekey, publickey
}

func generatePublicKey(bitLen int) *rsa.PublicKey {
	//Generate a RSA key
	privatekey, Error := rsa.GenerateKey(rand.Reader, bitLen)
	if Error != nil {
		fmt.Println("Error generating key")
		os.Exit(1)
	}
	publickey := &privatekey.PublicKey
	return publickey
}

func encrypt(data string, publicKey *rsa.PublicKey) string {
	//Encrypt the body with RSA
	ciphertext, Error := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(data))
	if Error != nil {
		fmt.Println("Error encrypting data" + Error.Error())
		os.Exit(1)
	}
	return string(ciphertext)
}

func decrypt(data []byte, privateKey *rsa.PrivateKey) string {
	decryptedData, Error := rsa.DecryptPKCS1v15(rand.Reader, privateKey, data)
	if Error != nil {
		fmt.Println("Error decrypting data" + Error.Error())
		os.Exit(1)
	}
	return string(decryptedData)
}

func getHelp() string {
	return "Usage: ./eurl <command>\n" +
		"Commands:\n" +
		"\tStandard HTTP Utility:\n" +
		"\t\t-d <url>\n" +
		"\t\t\tDownloads the url and returns the body\n" +
		"\tCryptography Utility:\n" +
		"\t\t--encrypt <data>\n" +
		"\t\t\tEncrypts data with RSA and outputs keys\n" +
		"\t\t-l <len>\n" +
		"\t\t\tSets the key length for generating the key\n" +
		"\t\t--dynamic\n" +
		"\t\t\tGenerates key size basic on length of data\n" +
		"\t\t--encode\n" +
		"\t\t\tEncodes the data with base64 (Use if data is too large)\n" +
		"\t\t--decrypt <data>\n" +
		"\t\t\tDecrypts data with RSA (Requires private key)\n" +
		"\t\t--key <key>\n" +
		"\t\t\tSets the private key for decrypting\n" +
		"\t\t--decode\n" +
		"\t\t\tDecodes the data with base64 (Use if data is encoded)\n" +
		"\t\t-o <file/dir>\n" +
		"\t\t\tSets the output file or directory for the data and keys\n"

}

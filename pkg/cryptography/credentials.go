package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/hoto/jenkins-credentials-decryptor/pkg/xml"
	"log"
	"regexp"
	"strings"
)

func DecryptCredentials(credentials *[]xml.Credential, secret []byte) ([]xml.Credential, error) {
	decryptedCredentials := make([]xml.Credential, len(*credentials))
	copy(decryptedCredentials, *credentials)

	for i, credential := range *credentials {
		for key, value := range credential.Tags {
			if isEncrypted(value) {
				decoded := base64Decode(value)
				decrypted := decrypt(decoded, secret)
				decryptedCredentials[i].Tags[key] = decrypted
			}
		}
		printFields(i, credential) // TODO: move to main

	}
	return decryptedCredentials, nil
}

// TODO check for first and last char and in-between
func isEncrypted(text string) bool {
	return strings.Contains(text, "{") && strings.Contains(text, "}")
}

func base64Decode(text string) []byte {
	encoded := regexp.MustCompile("{(.*?)}").FindStringSubmatch(text)[1]
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	check(err)
	return decoded
}

func decrypt(decoded []byte, secret []byte) string {
	if decoded[0] == 1 { // TODO handle the other case
		/*
		  p = p[1:] #Strip the version
		  p = p[4:] # Strip the iv length
		  p = p[4:] # Strip the data length
		  iv_length = 16
		  iv = p[:iv_length]
		  p = p[iv_length:]
		  o = AES.new(secret, AES.MODE_CBC, iv)
		  decrypted_p = o.decrypt(p)
		*/
		cipherText := decoded[1:]   // strip version
		cipherText = cipherText[4:] // strip iv length
		cipherText = cipherText[4:] // strip data length
		ivLength := 16              // TODO calculate this
		iv := cipherText[:ivLength]
		cipherText = cipherText[ivLength:] //strip iv

		block, err := aes.NewCipher(secret)
		check(err)
		mode := cipher.NewCBCDecrypter(block, iv)
		mode.CryptBlocks(cipherText, cipherText)

		return strings.TrimSpace(string(cipherText)) // TODO can a password have space at the front or end?
	}
	return "ERROR_UNKNOWN_VERSION"
}

func printFields(i int, credential xml.Credential) {
	fmt.Println(i)
	for k, v := range credential.Tags {
		fmt.Printf("\t%s=%s\n", k, v)
	}
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

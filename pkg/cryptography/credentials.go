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

func DecryptCredentials(credentials *[]xml.Credential, secret []byte) {
	for i, credential := range *credentials {
		decodeBase64Fields(credential, secret)
		printFields(i, credential)
	}
}

func decodeBase64Fields(credential xml.Credential, secret []byte) {
	for k, v := range credential.Tags {
		if strings.Contains(v, "{") {
			encoded := regexp.MustCompile("{(.*?)}").FindStringSubmatch(v)[1]
			decoded, err := base64.StdEncoding.DecodeString(encoded)
			check(err)
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
				cipherText = cipherText[ivLength:]

				block, err := aes.NewCipher(secret)
				check(err)
				mode := cipher.NewCBCDecrypter(block, iv)
				mode.CryptBlocks(cipherText, cipherText)

				credential.Tags[k] = strings.TrimSpace(string(cipherText))
			}
		}
	}
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

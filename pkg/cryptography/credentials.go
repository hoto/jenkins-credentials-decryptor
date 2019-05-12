package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
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
			if isBase64EncodedSecret(value) {
				decoded := base64Decode(value)
				decrypted := decrypt(decoded, secret)
				decryptedCredentials[i].Tags[key] = decrypted
			}
		}

	}
	return decryptedCredentials, nil
}

func isBase64EncodedSecret(text string) bool {
	if strings.HasPrefix(text, "{") && strings.HasSuffix(text, "}") {
		encoded := regexp.MustCompile("{(.*?)}").FindStringSubmatch(text)[1]
		_, err := base64.StdEncoding.DecodeString(encoded)
		if err == nil {
			return true
		}
	}
	if strings.HasSuffix(text, "=") {
		_, err := base64.StdEncoding.DecodeString(text)
		if err == nil {
			return true
		}
	}
	return false
}

func base64Decode(text string) []byte {
	encoded := text
	if strings.HasPrefix(text, "{") && strings.HasSuffix(text, "}") {
		encoded = regexp.MustCompile("{(.*?)}").FindStringSubmatch(text)[1]
	}
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	check(err)
	return decoded
}

func decrypt(decoded []byte, secret []byte) string {
	if decoded[0] == 1 { // TODO handle the other case
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

		trimmed := strings.TrimSpace(string(cipherText))
		withoutPadding := strings.Replace(string(trimmed), string('\x05'), "", -1)
		withoutPadding = strings.Replace(string(withoutPadding), string('\x06'), "", -1)
		return string(withoutPadding)
	}
	return string(decoded)
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

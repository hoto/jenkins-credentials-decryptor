package cryptography

import (
	"crypto/aes"
	cipherLib "crypto/cipher"
	"encoding/base64"
	"github.com/hoto/jenkins-credentials-decryptor/pkg/xml"
	"log"
	"regexp"
	"strings"
)

/*
  This is some next level reverse engineering.
  Kudos to http://xn--thibaud-dya.fr/jenkins_credentials.html
*/
func DecryptCredentials(credentials *[]xml.Credential, secret []byte) ([]xml.Credential, error) {
	decryptedCredentials := make([]xml.Credential, len(*credentials))
	copy(decryptedCredentials, *credentials)

	for i, credential := range *credentials {
		for key, value := range credential.Tags {
			if isBase64EncodedSecret(value) {
				encodedCipher := stripBrackets(value)
				cipher := base64Decode(encodedCipher)
				decrypted := decrypt(cipher, secret)
				decryptedCredentials[i].Tags[key] = decrypted
			}
		}

	}
	return decryptedCredentials, nil
}

/*
  New format of declaring a field to be a "base64 decoded secret" is by using {} brackets.
  Example:

    <password>{AQAAABAAAAAgPT7JbBVgyWiivobt0CJEduLyP0lB3uyTj+D5WBvVk6jyG6BQFPYGN4Z3VJN2JLDm}</password>

  Old format does not use the {} brackets.
  Instead jenkins seems to be usually suffixing the encoding with '=' sign.
  Example:

     <password>B+4pJjkJXD+pzyT9lcq8M8vF+p5YU4HmWy+MWldEdG4=</password>

  I'm not sure how to distinguish other encoded secrets from the "old days of jenkins".
  I don't want to comprehend Jenkins code from 4 years ago just to handle some edge cases.
  I can't try to decode all values because there are some phrases which
  would be false positive e.g. "root" (which is a valid base64 encoding)
*/
func isBase64EncodedSecret(text string) bool {
	if isBracketed(text) {
		encoded := textBetweenBrackets(text)
		return isBase64Encoded(encoded)
	}
	if strings.HasSuffix(text, "=") {
		return isBase64Encoded(text)
	}
	return false
}

func isBase64Encoded(text string) bool {
	_, err := base64.StdEncoding.DecodeString(text)
	if err == nil {
		return true
	}
	return false
}

func stripBrackets(text string) string {
	if isBracketed(text) {
		return textBetweenBrackets(text)
	}
	return text
}

func isBracketed(text string) bool {
	return strings.HasPrefix(text, "{") && strings.HasSuffix(text, "}")
}

func base64Decode(encoded string) []byte {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	check(err)
	return decoded
}

func textBetweenBrackets(text string) string {
	return regexp.MustCompile("{(.*?)}").FindStringSubmatch(text)[1]
}

func decrypt(cipher []byte, secret []byte) string {
	if cipher[0] == 1 { // you've gotta love jenkins
		return decryptNewFormatCredentials(cipher, secret)
	} else {
		return decryptOldFormatCredentials(cipher, secret)
	}
}

func decryptNewFormatCredentials(cipher []byte, secret []byte) string {
	cipher = cipher[1:] // strip version
	cipher = cipher[4:] // strip iv length
	cipher = cipher[4:] // strip data length
	ivLength := 16      // TODO calculate this
	iv := cipher[:ivLength]
	cipher = cipher[ivLength:] //strip iv
	block, err := aes.NewCipher(secret)
	check(err)
	mode := cipherLib.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipher, cipher)
	trimmed := strings.TrimSpace(string(cipher))

	// TODO strip PKCS7 padding with math not by strings.Replace()
	withoutPadding := strings.Replace(string(trimmed), string('\x05'), "", -1)
	withoutPadding = strings.Replace(string(withoutPadding), string('\x06'), "", -1)
	return withoutPadding
}

func decryptOldFormatCredentials(decoded []byte, secret []byte) string {
	decrypted := string(decryptAes128Ecb(decoded, secret))
	return strings.Replace(decrypted, "::::MAGIC::::", "", -1)
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

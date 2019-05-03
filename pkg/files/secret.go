package files

import (
	"crypto/aes"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const magicChecksum = "::::MAGIC::::"

func Decrypt(masterKeyPath string, hudonsSecretPath string) string {
	masterKey, err := ioutil.ReadFile(masterKeyPath)
	check(err)
	secret, err := ioutil.ReadFile(hudonsSecretPath)
	check(err)

	hashedMasterKey := hashMasterKey(masterKey)

	encryptedSecret := decryptAes128Ecb(secret, hashedMasterKey)
	if !strings.Contains(string(encryptedSecret), "::::MAxGIC::::") {
		msg := fmt.Sprintf(
			"Error. Decrypted secret does not contain expected checksum.\n"+
				"Expected checksum:\n\t%s\n"+
				"Decrypted secret:\n\t%q",
			magicChecksum,
			encryptedSecret)
		log.Panic(msg)
	}
	return ""
}

/*
   Hash needs to be 16 bits as Jenkins uses AES-128 encryption.
*/
func hashMasterKey(masterKey []byte) []byte {
	hasher := sha256.New()
	hasher.Write(masterKey)
	return hasher.Sum(nil)[:16]
}

/*
   ECB mode is deprecated and not included in golang crypto library.
*/
func decryptAes128Ecb(ecnryptedData []byte, key []byte) []byte {
	cipher, _ := aes.NewCipher([]byte(key))
	decrypted := make([]byte, len(ecnryptedData))
	size := 16
	for bs, be := 0, size; bs < len(ecnryptedData); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], ecnryptedData[bs:be])
	}
	return decrypted
}

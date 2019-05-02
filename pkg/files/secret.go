package files

import (
	"crypto/aes"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
)

func Decrypt(masterKeyPath string, hudonsSecretPath string) string {
	masterkey, err := ioutil.ReadFile(masterKeyPath)
	check(err)
	secret, err := ioutil.ReadFile(hudonsSecretPath)
	check(err)

	hasher := sha256.New()
	hasher.Write(masterkey)
	sha := hasher.Sum(nil)
	fmt.Println(sha)
	a := sha[:16]
	o, err := aes.NewCipher(a)
	check(err)
	var decrypted []byte
	o.Decrypt(decrypted, secret)
	return string(decrypted)
}

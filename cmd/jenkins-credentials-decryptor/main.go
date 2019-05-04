package main

import (
	"github.com/hoto/jenkins-credentials-decryptor/pkg/config"
	"github.com/hoto/jenkins-credentials-decryptor/pkg/cryptography"
	"github.com/hoto/jenkins-credentials-decryptor/pkg/xml"
	"io/ioutil"
	"log"
)

func main() {
	config.ParseFlags()

	credentialsXml := readFile(config.CredentialsXmlPath)
	masterKey := readFile(config.MasterKeyPath)
	encryptedHudsonSecret := readFile(config.HudsonSecretPath)

	secret, err := cryptography.DecryptHudsonSecret(masterKey, encryptedHudsonSecret)
	check(err)

	credentials := xml.ParseCredentialsXml(credentialsXml)
	cryptography.DecryptCredentials(credentials, secret)
}

func readFile(path string) []byte {
	content, err := ioutil.ReadFile(path)
	check(err)
	return content
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

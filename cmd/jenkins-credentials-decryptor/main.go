package main

import (
	"io/ioutil"
	"log"

	"github.com/hoto/jenkins-credentials-decryptor/pkg/config"
	"github.com/hoto/jenkins-credentials-decryptor/pkg/cryptography"
	"github.com/hoto/jenkins-credentials-decryptor/pkg/printer"
	"github.com/hoto/jenkins-credentials-decryptor/pkg/xml"
)

func main() {
	config.ParseFlags()

	credentialsXml := readFile(config.CredentialsXmlPath)
	masterKey := readFile(config.MasterKeyPath)
	encryptedHudsonSecret := readFile(config.HudsonSecretPath)

	credentials, err := xml.ParseCredentialsXml(credentialsXml)
	check(err)
	secret, err := cryptography.DecryptHudsonSecret(masterKey, encryptedHudsonSecret)
	check(err)

	decryptedCredentials, _ := cryptography.DecryptCredentials(credentials, secret)

	printer.Print(decryptedCredentials, config.OutputFormat)
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

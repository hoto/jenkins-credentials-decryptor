package main

import (
	"fmt"
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

	credentials, err := xml.ParseCredentialsXml(credentialsXml)
	check(err)
	secret, err := cryptography.DecryptHudsonSecret(masterKey, encryptedHudsonSecret)
	check(err)

	decryptedCredentials, _ := cryptography.DecryptCredentials(credentials, secret)
	print(decryptedCredentials)
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

func print(decryptedCredentials []xml.Credential) {
	for i, credential := range decryptedCredentials {
		fmt.Println(i)
		for k, v := range credential.Tags {
			fmt.Printf("\t%s=%s\n", k, v)
		}
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hoto/jenkins-credentials-decryptor/pkg/config"
	"github.com/hoto/jenkins-credentials-decryptor/pkg/cryptography"
	"github.com/hoto/jenkins-credentials-decryptor/pkg/xml"
	"io/ioutil"
	"log"
	"strings"
)

type Credential struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	Passphrase  string `json:"passphrase"`
	Password    string `json:"password"`
	Path        string `json:"path"`
	Scope       string `json:"scope"`
	Secret      string `json:"secret"`
	Username    string `json:"username"`
}

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
	if config.JsonOutput {
		printJson(decryptedCredentials)
	} else {
		print(decryptedCredentials)
	}
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
			fmt.Printf("\t%s: %s\n", k, v)
		}
	}
}

func printJson(decryptedCredentials []xml.Credential) {
	trimCutSet := "\u0001\u0004\u0010\u000e\u000f\u0007\b"
	credentials := make([]Credential, 0, len(decryptedCredentials))
	for _, credential := range decryptedCredentials {
		t := credential.Tags
		temp := Credential{
			Description: t["description"],
			Id:          t["id"],
			Passphrase:  strings.Trim(t["passphrase"], trimCutSet),
			Password:    strings.Trim(t["password"], trimCutSet),
			Path:        t["path"],
			Scope:       t["scope"],
			Secret:      strings.Trim(t["secret"], trimCutSet),
			Username:    t["username"],
		}
		credentials = append(credentials, temp)
	}
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	err := enc.Encode(&credentials)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf.String())
}
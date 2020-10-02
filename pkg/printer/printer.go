package printer

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/hoto/jenkins-credentials-decryptor/pkg/xml"
)

type Credential struct {
	AccessKey          string `json:"accessKey,omitempty"`
	Description        string `json:"description,omitempty"`
	FileName           string `json:"fileName,omitempty"`
	IamMfaSerialNumber string `json:"iamMfaSerialNumber,omitempty"`
	IamRoleArn         string `json:"iamRoleArn,omitempty"`
	Id                 string `json:"id,omitempty"`
	Passphrase         string `json:"passphrase,omitempty"`
	Password           string `json:"password,omitempty"`
	Path               string `json:"path,omitempty"`
	PrivateKey         string `json:"privateKey,omitempty"`
	PrivateKeySource   string `json:"privateKeySource,omitempty"`
	RoleId             string `json:"roleId,omitempty"`
	Scope              string `json:"scope,omitempty"`
	Secret             string `json:"secret,omitempty"`
	SecretBytes        string `json:"secretBytes,omitempty"`
	SecretId           string `json:"secretId,omitempty"`
	SecretKey          string `json:"secretKey,omitempty"`
	Username           string `json:"username,omitempty"`
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func Print(decryptedCredentials []xml.Credential, outputFormat string) {
	if outputFormat == "json" {
		printJson(decryptedCredentials)
	} else {
		printText(decryptedCredentials)
	}
}

func printText(decryptedCredentials []xml.Credential) {
	for i, credential := range decryptedCredentials {
		fmt.Println(i)
		for k, v := range credential.Tags {
			fmt.Printf("\t%s: %s\n", k, v)
		}
	}
}

func printJson(decryptedCredentials []xml.Credential) {
	trimCutSet := "\u0001\u0004\u0007\u000e\u000f\u0010\b"
	credentials := make([]Credential, 0, len(decryptedCredentials))
	for _, credential := range decryptedCredentials {
		t := credential.Tags
		//secretBytes, err := base64.StdEncoding.DecodeString(t["secretBytes"])
		secretBytes, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(t["secretBytes"])
		check(err)
		temp := Credential{
			AccessKey:          t["accessKey"],
			Description:        t["description"],
			FileName:           t["fileName"],
			IamMfaSerialNumber: t["iamMfaSerialNumber"],
			IamRoleArn:         t["iamRoleArn"],
			Id:                 t["id"],
			Passphrase:         strings.Trim(t["passphrase"], trimCutSet),
			Password:           strings.Trim(t["password"], trimCutSet),
			Path:               t["path"],
			PrivateKey:         t["privateKey"],
			PrivateKeySource:   t["privateKeySource"],
			RoleId:             t["roleId"],
			Scope:              t["scope"],
			Secret:             strings.Trim(t["secret"], trimCutSet),
			SecretBytes:        base64.StdEncoding.EncodeToString(secretBytes),
			SecretId:           t["secretId"],
			SecretKey:          t["secretKey"],
			Username:           t["username"],
		}
		credentials = append(credentials, temp)
	}
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	err := enc.Encode(&credentials)
	check(err)
	fmt.Println(buf.String())
}

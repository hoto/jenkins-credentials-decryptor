package main

import (
	"fmt"
	"github.com/beevik/etree"
	"io/ioutil"
	"regexp"
)

const (
	credentialsXpath = "//java.util.concurrent.CopyOnWriteArrayList/*"
)

type Credential struct {
	tags map[string]string
}

func readCredentials() *[]Credential {
	credentials := make([]Credential, 0)
	for i, credentialNode := range readCredentialsXml().FindElements(credentialsXpath) {
		fmt.Println(i)
		credential := &Credential{
			tags: map[string]string{},
		}
		for _, field := range credentialNode.ChildElements() {
			fmt.Printf("\t%s=%s\n", field.Tag, field.Text())
			credential.tags[field.Tag] = field.Text()
		}
		credentials = append(credentials, *credential)
	}
	return &credentials
}

/*
 HACK ALERT:
 Stripping xml version line as current native and third party xml decoders
 refuses to parse xml version 1.0+
 Jenkins uses xml version 1.1+ so this may blow up.
*/
func readCredentialsXml() *etree.Document {
	credentials, err := ioutil.ReadFile("test/resources/credentials.xml")
	check(err)
	sanitizedCredentials := regexp.
		MustCompile("(?m)^.*<?xml.*$").
		ReplaceAllString(string(credentials), "")
	document := etree.NewDocument()
	err = document.ReadFromString(sanitizedCredentials)
	check(err)
	return document
}

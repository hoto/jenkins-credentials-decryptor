package main

import (
	"github.com/beevik/etree"
	"io/ioutil"
	"regexp"
	"strings"
)

const (
	credentialsXpath = "//java.util.concurrent.CopyOnWriteArrayList/*"
)

type Credential struct {
	tags map[string]string
}

func ReadCredentials(path string) *[]Credential {
	credentials := make([]Credential, 0)
	for _, credentialNode := range parseCredentialsXml(path).FindElements(credentialsXpath) {
		credential := &Credential{
			tags: map[string]string{},
		}
		for _, child := range credentialNode.ChildElements() {
			reduceFields(child, credential)
		}
		credentials = append(credentials, *credential)
	}
	return &credentials
}

/*
  There is a possibility that a field will get overridden but I haven't seen an example like that.
*/
func reduceFields(node *etree.Element, credential *Credential) {
	credential.tags[node.Tag] = strings.TrimSpace(node.Text())
	for _, child := range node.ChildElements() {
		credential.tags[child.Tag] = strings.TrimSpace(child.Text())
		reduceFields(child, credential)
	}
}

/*
 HACK ALERT:
 Stripping xml version line as current native and third party xml decoders
 refuses to parse xml version 1.0+
 Jenkins uses xml version 1.1+ so this may blow up.
*/
func parseCredentialsXml(path string) *etree.Document {
	credentials, err := ioutil.ReadFile(path)
	check(err)
	sanitizedCredentials := regexp.
		MustCompile("(?m)^.*<?xml.*$").
		ReplaceAllString(string(credentials), "")
	document := etree.NewDocument()
	err = document.ReadFromString(sanitizedCredentials)
	check(err)
	return document
}

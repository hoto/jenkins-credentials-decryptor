package cryptography

import (
	"github.com/beevik/etree"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

const (
	credentialsXpath = "//java.util.concurrent.CopyOnWriteArrayList/*"
)

type Credential struct {
	Tags map[string]string
}

func ReadCredentials(path string) *[]Credential {
	credentials := make([]Credential, 0)
	for _, credentialNode := range parseCredentialsXml(path).FindElements(credentialsXpath) {
		credential := &Credential{
			Tags: map[string]string{},
		}
		for _, child := range credentialNode.ChildElements() {
			reduceFields(child, credential)
		}
		credentials = append(credentials, *credential)
	}
	return &credentials
}

/*
  There is a possibility that a field could get overridden but I haven't seen an example of that yet.
*/
func reduceFields(node *etree.Element, credential *Credential) {
	credential.Tags[node.Tag] = strings.TrimSpace(node.Text())
	for _, child := range node.ChildElements() {
		credential.Tags[child.Tag] = strings.TrimSpace(child.Text())
		reduceFields(child, credential)
	}
}

func parseCredentialsXml(path string) *etree.Document {
	credentialsXml, err := ioutil.ReadFile(path)
	check(err)
	document := etree.NewDocument()
	err = document.ReadFromString(stripXmlVersion(credentialsXml))
	check(err)
	return document
}

/*
 HACK ALERT:
 Stripping xml version because I could not find any decoder which would parse xml version 1.0+
 Jenkins uses xml version 1.1+ so this may blow up.
*/
func stripXmlVersion(credentials []byte) string {
	return regexp.
		MustCompile("(?m)^.*<?xml.*$").
		ReplaceAllString(string(credentials), "")
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

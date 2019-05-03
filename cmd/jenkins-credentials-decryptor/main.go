package main

import (
	"fmt"
	"github.com/hoto/jenkins-credentials-decryptor/pkg/xml"
	"io/ioutil"
	"log"
)

func main() {
	credentialsXml, err := ioutil.ReadFile("test/resources/credentials.xml")
	check(err)

	credentials := xml.ParseCredentialsXml(credentialsXml)
	for i, credential := range *credentials {
		fmt.Println(i)
		for k, v := range credential.Tags {
			fmt.Printf("\t%s=%s\n", k, v)
		}
	}
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

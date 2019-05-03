package main

import (
	"fmt"
	"github.com/hoto/jenkins-credentials-decryptor/pkg/cryptography"
)

func main() {
	credentials := cryptography.ReadCredentials("test/resources/credentials.xml")
	for i, credential := range *credentials {
		fmt.Println(i)
		for k, v := range credential.Tags {
			fmt.Printf("\t%s=%s\n", k, v)
		}
	}
}

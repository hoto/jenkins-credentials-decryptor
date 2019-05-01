package main

import (
	"fmt"
	"log"
)

func main() {
	credentials := ReadCredentials()
	for i, credential := range *credentials {
		fmt.Println(i)
		for k, v := range credential.tags {
			fmt.Printf("\t%s=%s\n", k, v)
		}
	}
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

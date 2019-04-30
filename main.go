package main

import (
	"fmt"
	"github.com/beevik/etree"
	"io/ioutil"
	"log"
	"regexp"
)

func main() {
	credentials, err := ioutil.ReadFile("test/resources/credentials.xml")
	check(err)
	/*
	 HACK ALERT:
	 Stripping xml version line as current native and third party xml decoders
	 refuses to parse xml version 1.0+
	 Jenkins uses xml version 1.1+ so this may blow up.
	 That line looks like this:
	 <?xml version='1.1' encoding='UTF-8'?>
	*/
	sanitizedCredentials := regexp.
		MustCompile("(?m)^.*<?xml.*$").
		ReplaceAllString(string(credentials), "")

	doc := etree.NewDocument()
	if err := doc.ReadFromString(sanitizedCredentials); err != nil {
		panic(err)
	}
	for _, node := range doc.FindElements("//java.util.concurrent.CopyOnWriteArrayList/*") {
		children := node.ChildElements()
		fmt.Println(node.Tag)
		for _, child := range children {
			fmt.Printf("\t%s=%s\n", child.Tag, child.Text())
		}
	}
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

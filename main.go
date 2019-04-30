package main

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
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

	decoder := xml.NewDecoder(strings.NewReader(sanitizedCredentials))
	readToken(decoder)
}

func readToken(decoder *xml.Decoder) {
	var nodeName string
	decoder.Token()
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			log.Print("Finished reading whole file.")
			break
		}
		if err != nil {
			log.Println("Error when tokenizing:", err)
		}
		if token == nil {
			log.Print("No token")
			break
		}
		switch node := token.(type) {
		case xml.StartElement:
			nodeName = node.Name.Local
			println("Node:", nodeName)
		default:
		}
	}
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

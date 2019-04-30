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
	parseXml(decoder)
}

func parseXml(decoder *xml.Decoder) {
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		check(err)
		switch node := token.(type) {
		case xml.StartElement:
			println("StartElement=", node.Name.Local)
		case xml.CharData:
			println("CharData=", string(node))
		default:
		}
	}
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

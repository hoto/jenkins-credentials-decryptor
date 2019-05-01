package main

import "log"

//const (
//	credentialsXpath = "//java.util.concurrent.CopyOnWriteArrayList/*"
//)
//
//func main() {
//	for i, credential := range readCredentialsXml().FindElements(credentialsXpath) {
//		fmt.Println(i)
//		for _, field := range credential.ChildElements() {
//			fmt.Printf("\t%s=%s\n", field.Tag, field.Text())
//		}
//	}
//}
//
///*
// HACK ALERT:
// Stripping xml version line as current native and third party xml decoders
// refuses to parse xml version 1.0+
// Jenkins uses xml version 1.1+ so this may blow up.
//*/
//func readCredentialsXml() *etree.Document {
//	credentials, err := ioutil.ReadFile("test/resources/credentials.xml")
//	check(err)
//	sanitizedCredentials := regexp.
//		MustCompile("(?m)^.*<?xml.*$").
//		ReplaceAllString(string(credentials), "")
//	document := etree.NewDocument()
//	err = document.ReadFromString(sanitizedCredentials)
//	check(err)
//	return document
//}
//
func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

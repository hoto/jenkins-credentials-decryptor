package config

import (
	"flag"
	"fmt"
	"os"
)

var (
	MasterKeyPath      string
	HudsonSecretPath   string
	CredentialsXmlPath string
)

const (
	empty           = ""
	masterKeyDesc   = "(required) master.key file location"
	secretDesc      = "(required) hudson.util.Secret file location"
	credentialsDesc = "(required) credentials.xml file location"
	usage           = `Usage:

  jenkins-credentials-decryptor \
    -m master.key \
    -s hudson.util.Secret \
    -c credentials.xml

Flags:

`
)

func ParseFlags() {
	flag.Usage = overrideUsage()

	flag.StringVar(&MasterKeyPath, "m", empty, masterKeyDesc)
	flag.StringVar(&HudsonSecretPath, "s", empty, secretDesc)
	flag.StringVar(&CredentialsXmlPath, "c", empty, credentialsDesc)

	flag.Parse()

	if isEmpty(MasterKeyPath) || isEmpty(HudsonSecretPath) || isEmpty(CredentialsXmlPath) {
		printUsageAndExit()
	}
}

func overrideUsage() func() {
	return func() {
		_, _ = fmt.Fprintf(os.Stdout, usage)
		flag.PrintDefaults()
	}
}

func isEmpty(text string) bool {
	return text == empty
}

func printUsageAndExit() {
	fmt.Printf("Please provide all required flags.\n\n")
	flag.Usage()
	os.Exit(1)
}

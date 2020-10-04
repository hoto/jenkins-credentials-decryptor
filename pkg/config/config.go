package config

import (
	"flag"
	"fmt"
	"os"
)

var (
	Version     string
	ShortCommit string
	BuildDate   string

	MasterKeyPath      string
	HudsonSecretPath   string
	CredentialsXmlPath string

	OutputFormat string
)

const (
	empty           = ""
	showVersionDesc = "(optional) show version"
	masterKeyDesc   = "(required) master.key file location"
	secretDesc      = "(required) hudson.util.Secret file location"
	credentialsDesc = "(required) credentials.xml file location"
	outputFormat    = "(optional) output format [json|text]"
	usage           = `Usage:

  jenkins-credentials-decryptor \
    -m master.key \
    -s hudson.util.Secret \
    -c credentials.xml \
    -o json

Flags:

`
)

func ParseFlags() {
	flag.Usage = overrideUsage()

	showVersion := flag.Bool("version", false, showVersionDesc)

	flag.StringVar(&MasterKeyPath, "m", empty, masterKeyDesc)
	flag.StringVar(&HudsonSecretPath, "s", empty, secretDesc)
	flag.StringVar(&CredentialsXmlPath, "c", empty, credentialsDesc)
	flag.StringVar(&OutputFormat, "o", "json", outputFormat)

	flag.Parse()

	if *showVersion {
		fmt.Printf("jenkins-credentials-decryptor version %s, commit %s, build %s\n",
			Version, ShortCommit, BuildDate)
		os.Exit(0)
	}

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

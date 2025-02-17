package printer

import (
	"bytes"
	"embed"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/hoto/jenkins-credentials-decryptor/pkg/xml"
)

//go:embed k8sSecret.tmpl
var k8sSecretTemplate embed.FS

type Credential struct {
	AccessKey             string `json:"accessKey,omitempty"`
	AccessToken           string `json:"accessToken,omitempty"`
	ApiToken              string `json:"apiToken,omitempty"`
	AppId                 string `json:"appID,omitempty"`
	ClientCertificate     string `json:"clientCertificate,omitempty"`
	ClientKeySecret       string `json:"clientKey,omitempty"`
	Description           string `json:"description,omitempty"`
	FileName              string `json:"fileName,omitempty"`
	IamMfaSerialNumber    string `json:"iamMfaSerialNumber,omitempty"`
	IamRoleArn            string `json:"iamRoleArn,omitempty"`
	Id                    string `json:"id,omitempty"`
	MountPath             string `json:"mountPath,omitempty"`
	Namespace             string `json:"namespace,omitempty"`
	Owner                 string `json:"owner,omitempty"`
	Passphrase            string `json:"passphrase,omitempty"`
	Password              string `json:"password,omitempty"`
	Path                  string `json:"path,omitempty"`
	PrivateKey            string `json:"privateKey,omitempty"`
	PrivateKeySource      string `json:"privateKeySource,omitempty"`
	RoleId                string `json:"roleId,omitempty"`
	Scope                 string `json:"scope,omitempty"`
	Secret                string `json:"secret,omitempty"`
	SecretBytes           string `json:"secretBytes,omitempty"`
	SecretId              string `json:"secretId,omitempty"`
	SecretKey             string `json:"secretKey,omitempty"`
	SecretType            string `json:"secretType,omitempty"`
	ServerCaCertificate   string `json:"serverCaCertificate,omitempty"`
	Token                 string `json:"token,omitempty"`
	UsePolicies           string `json:"usePolicies,omitempty"`
	Username              string `json:"username,omitempty"`
	UploadedKeystoreBytes string `json:"uploadedKeyStoreBytes,omitempty"`
}

type KeyStoreSource struct {
	UploadedKeystoreBytes string `json:"uploadedKeystoreBytes,omitempty"`
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func Print(decryptedCredentials []xml.Credential, outputFormat string) {
	if outputFormat == "json" {
		printJson(decryptedCredentials)
	} else if outputFormat == "k8secret" {
		printK8sSecrets(decryptedCredentials)
	} else {
		printText(decryptedCredentials)
	}
}

func printText(decryptedCredentials []xml.Credential) {
	for i, credential := range decryptedCredentials {
		fmt.Println(i)
		for k, v := range credential.Tags {
			fmt.Printf("\t%s: %s\n", k, v)
		}
	}
}

func credentialsStore(decryptedCredentials []xml.Credential) []Credential {
	trimCutSet := "\u0001\u0004\u0007\u000e\u000f\u0010\b"
	credentials := make([]Credential, 0, len(decryptedCredentials))
	for _, credential := range decryptedCredentials {
		t := credential.Tags
		temp := Credential{
			AccessKey:             t["accessKey"],
			AccessToken:           t["accessToken"],
			ApiToken:              t["apiToken"],
			AppId:                 t["appID"],
			ClientCertificate:     t["clientCertificate"],
			ClientKeySecret:       t["clientKey"],
			Description:           t["description"],
			FileName:              t["fileName"],
			IamMfaSerialNumber:    t["iamMfaSerialNumber"],
			IamRoleArn:            t["iamRoleArn"],
			Id:                    t["id"],
			MountPath:             t["mountPath"],
			Namespace:             t["namespace"],
			Owner:                 t["owner"],
			Passphrase:            strings.Trim(t["passphrase"], trimCutSet),
			Password:              strings.Trim(t["password"], trimCutSet),
			Path:                  t["path"],
			PrivateKey:            t["privateKey"],
			PrivateKeySource:      t["privateKeySource"],
			ServerCaCertificate:   t["serverCaCertificate"],
			RoleId:                t["roleId"],
			Scope:                 t["scope"],
			Secret:                strings.Trim(t["secret"], trimCutSet),
			SecretBytes:           b64.StdEncoding.EncodeToString([]byte(t["secretBytes"])),
			SecretId:              t["secretId"],
			SecretKey:             t["secretKey"],
			SecretType:            getType(t["secretType"]),
			Token:                 t["token"],
			UploadedKeystoreBytes: b64.StdEncoding.EncodeToString([]byte(t["uploadedKeystoreBytes"])),
			UsePolicies:           t["usePolicies"],
			Username:              t["username"],
		}
		credentials = append(credentials, temp)
	}
	return credentials

}

func printJson(decryptedCredentials []xml.Credential) {
	credentials := credentialsStore(decryptedCredentials)
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	err := enc.Encode(&credentials)
	check(err)
	fmt.Println(buf.String())
}

func getType(key string) string {

	credentialsType := map[string]string{
		"com.dabsquared.gitlabjenkins.connection.GitLabApiTokenImpl":                "gitlabToken",
		"io.jenkins.plugins.gitlabserverconfig.credentials.PersonalAccessTokenImpl": "gitlabToken",
		"com.cloudbees.plugins.credentials.impl.UsernamePasswordCredentialsImpl":    "usernamePassword",
		"org.jenkinsci.plugins.plaincredentials.impl.StringCredentialsImpl":         "secretText",
		"org.jenkinsci.plugins.plaincredentials.impl.FileCredentialsImpl":           "secretFile",
		"com.cloudbees.plugins.credentials.impl.CertificateCredentialsImpl":         "certificate",
		"com.cloudbees.jenkins.plugins.sshcredentials.impl.BasicSSHUserPrivateKey":  "basicSSHUserPrivateKey",
		"com.cloudbees.jenkins.plugins.awscredentials.AWSCredentialsImpl":           "aws",
		//"org.jenkinsci.plugins.openstack.compute.auth.OpenstackCredentialv3":        "openstack",
		"org.jenkinsci.plugins.github_branch_source.GitHubAppCredentials":          "gitHubApp",
		"org.jenkinsci.plugins.github__branch__source.GitHubAppCredentials":        "gitHubApp",
		"com.datapipe.jenkins.vault.credentials.VaultAppRoleCredential":            "vaultAppRole",
		"com.datapipe.jenkins.vault.credentials.VaultTokenCredential":              "vaultToken",
		"com.datapipe.jenkins.vault.credentials.VaultGithubTokenCredential":        "VaultGitHubToken",
		"org.jenkinsci.plugins.docker.commons.credentials.DockerServerCredentials": "x509ClientCert",
	}
	if value, ok := credentialsType[key]; ok {
		return value
	} else {
		return key
	}

}
func printK8sSecrets(decryptedCredentials []xml.Credential) {
	fmt.Println("Parsing the secrets into a k8sSecrets file...")
	credentials := credentialsStore(decryptedCredentials)

	funcMap := template.FuncMap{
		"toLowerAndRemoveUnderscores": func(s string) string {
			lower := strings.ToLower(s)
			return strings.ReplaceAll(lower, "_", "")
		},
		"noSpaces": func(s string) string {

			noNewLines := strings.ReplaceAll(s, "\n", "")
			return noNewLines
		},
		"indent": func(spaces int, s string) string {
			padding := strings.Repeat(" ", spaces)
			return strings.ReplaceAll(s, "\n", "\n"+padding)
		},
	}

	templ, err := template.New("k8sSecret.tmpl").Funcs(funcMap).ParseFS(k8sSecretTemplate, "k8sSecret.tmpl")

	if err != nil {
		fmt.Printf("Error parsing embedded template: %v\n", err)
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	k8sSecretsFile := dir + "/k8sSecret.yaml"
	file, err := os.Create(k8sSecretsFile)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	err = templ.Execute(file, credentials)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Secrets parsed into k8sSecrets file: ", k8sSecretsFile)
}

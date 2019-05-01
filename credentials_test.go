package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	gitlab = &Credential{
		tags: map[string]string{
			"scope":       "GLOBAL",
			"id":          "gitlab",
			"description": "Gitlab admin user",
			"username":    "gitlabadmin",
			"password":    "{AQAAABAAAAAgPT7JbBVgyWiivobt0CJEduLyP0lB3uyTj+D5WBvVk6jyG6BQFPYGN4Z3VJN2JLDm}",
		},
	}
	bastion = &Credential{
		tags: map[string]string{
			"scope":            "GLOBAL",
			"id":               "production-bastion",
			"description":      "Production bastion ssh key",
			"username":         "root",
			"passphrase":       "{AQAAABAAAAAQmEZaw8Ev9tClXWVQye1TR2KgF3p/wGoYs/TEQCmsxCk=}",
			"privateKeySource": "\n            ",
		},
	}
)

func Test_reads_credentials_from_xml_file(t *testing.T) {
	expected := []Credential{*gitlab, *bastion}

	credentials := ReadCredentials()

	assert.Equal(t, expected, *credentials)
}

package xml

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var (
	gitlab = &Credential{
		Tags: map[string]string{
			"scope":       "GLOBAL",
			"id":          "gitlab",
			"description": "Gitlab admin user",
			"username":    "gitlabadmin",
			"password":    "{AQAAABAAAAAgPT7JbBVgyWiivobt0CJEduLyP0lB3uyTj+D5WBvVk6jyG6BQFPYGN4Z3VJN2JLDm}",
		},
	}
	bastion = &Credential{
		Tags: map[string]string{
			"scope":            "GLOBAL",
			"id":               "production-bastion",
			"description":      "Production bastion ssh key",
			"username":         "root",
			"passphrase":       "{AQAAABAAAAAQmEZaw8Ev9tClXWVQye1TR2KgF3p/wGoYs/TEQCmsxCk=}",
			"privateKeySource": "",
			"privateKey":       "{AQAAABAAAAOATgRVSIzrinNAlvbf/h2yDhN/yvJXpb/KcZCKKmyQqRolE65dHfnhdO1zLEv5Ek3uCXHWgNxSC7j9uk1/y/ckt+SCg8nPD9aOfLhSDpuuw0DO7IUvouPB9pSiW76c5QFzIOFb6Qx0ZPU7LO0VyK83yXnOpSQdbSOMnC9iFtUxSxIF70mSBxVY60SD87xXA4aT3wRtf2YoyAwjf2Y1x/T0GJhvQHBKAqlDRJ7Y9SzI4tLVvVyGidhuz6m2/KIioUVvRHI7kQ5azr+mp7sAXF1O4PGXz6BnvGSwFmmY/ggNnHvTzhoaVDWQm5PKVEp9fB9uX2SuRVYvdJ5/2hSx/woWAsQZezoOzW5d5UBh9NMP6HRCbJnR68/P0yeAw0Kbrb0U/wF5baZ/mwU5AZS4Uy93S5BSxaXmQBNUKu+HNPV3vG9kCX2tYzasuoNX1DTmpGPRc1Osqf1s9PEWzH90wFwQGDrALGvmUdkshqldyDm+RRZbJRF6puBTjBOXyzrFDvpn41xjpT3SabPn7xv45du5qW+95yP/yhjD7jTthsZRgQwkduGv1XLFDZ3Mbh8hqjObLwmcl0QyyDvrLrWA/Kyw1NpZ5U0A/rA46Eh/OZIiQVEk4JAyNewqWGCR+OTvyTK0EYqIZWgpbiSgRmVoUN6gmUdLJCjCPAvJZAZ3DlEzRpnabF6zIKsCsucDhvHn4F3JXpytLz8roVK/pX3VCm5VwkcSogcfKKfRpy+SMVY0LVcRkJCI32LhXduTxYQuzSEUqHz7tz3TP8DzVQeYUWdfMtg+kxQQd466CDGdzvINA+spuwtwQ9ogjZWdKhNr+Kk7fpCMI8SY1b2+ATYKXE/mY/gkNp+rIj9EMpjNWP81rZ7pcEIxycBOzADpPadgocri4cmSYDYmbf1/QYrTcp9jr6OfD6dvLjkJkQ+pKKF522swTxTIYxNitaRLKT7WpynAwBnXJLzwzsKAGw3UBuWuKPI5kcszAVodThAZqODMHLjB8LElWyoZ4+nqmlC+zLh8dJha3I0JEVL3hWp8sOr1S6kFPjahy/vxPYpLARlqBO3ZVLgM01aueMyaAagRWrZC6OgZw0NXW002FC5Vrfm2QaxLYz4JuuUZYsXySNwAaqf5FKGgO9dsm90EboP6dyp9BRi7h5Z8nWaGP9yJp92u6KvUe93UmSfX2ZPHCScRnBNFWQE8nDgCKsnKWenZmOkL}",
		},
	}
)

func Test_reads_credentials_from_xml_file(t *testing.T) {
	expectedCredentials := []Credential{*gitlab, *bastion}
	credentialsXml, _ := ioutil.ReadFile("../../test/resources/jenkins_2.141/credentials.xml")

	actualCredentials, _ := ParseCredentialsXml(credentialsXml)

	assert.Equal(t, expectedCredentials, *actualCredentials)
}

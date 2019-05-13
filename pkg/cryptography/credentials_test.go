package cryptography

import (
	"github.com/hoto/jenkins-credentials-decryptor/pkg/xml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var (
	oldFormatEncryptedCredentials = []xml.Credential{
		{
			Tags: map[string]string{
				"username": "user_1",
				"password": "B+4pJjkJXD+pzyT9lcq8M8vF+p5YU4HmWy+MWldEdG4=",
			},
		},
		{
			Tags: map[string]string{
				"username":   "user_1",
				"privateKey": "LINE_1\nLINE_2\nLINE_2",
				"passphrase": "7fPxAjao0hmb9ggFCSl8WsvF+p5YU4HmWy+MWldEdG4=",
			},
		},
	}

	oldFormatDecryptedCredentials = []xml.Credential{
		{
			Tags: map[string]string{
				"username": "user_1",
				"password": "password_1\t\t\t\t\t\t\t\t\t"},
		},
		{
			Tags: map[string]string{
				"passphrase": "passphrase\t\t\t\t\t\t\t\t\t",
				"username":   "user_1",
				"privateKey": "LINE_1\nLINE_2\nLINE_2"},
		},
	}

	newFormatEncryptedCredentials = []xml.Credential{
		{
			Tags: map[string]string{
				"username": "xfireadmin",
				"password": "AQAAABAAAAAQtnCexFYLFtmTQCL0x3wnirMnXVA7aZy+lfrfso+SjHI=",
			},
		},
		{
			Tags: map[string]string{
				"username": "gitlabadmin",
				"password": "{AQAAABAAAAAgPT7JbBVgyWiivobt0CJEduLyP0lB3uyTj+D5WBvVk6jyG6BQFPYGN4Z3VJN2JLDm}",
			},
		},
		{
			Tags: map[string]string{
				"username":   "root",
				"passphrase": "{AQAAABAAAAAQmEZaw8Ev9tClXWVQye1TR2KgF3p/wGoYs/TEQCmsxCk=}",
				"privateKey": "{AQAAABAAAAOATgRVSIzrinNAlvbf/h2yDhN/yvJXpb/KcZCKKmyQqRolE65dHfnhdO1zLEv5Ek3uCXHWgNxSC7j9uk1/y/ckt+SCg8nPD9aOfLhSDpuuw0DO7IUvouPB9pSiW76c5QFzIOFb6Qx0ZPU7LO0VyK83yXnOpSQdbSOMnC9iFtUxSxIF70mSBxVY60SD87xXA4aT3wRtf2YoyAwjf2Y1x/T0GJhvQHBKAqlDRJ7Y9SzI4tLVvVyGidhuz6m2/KIioUVvRHI7kQ5azr+mp7sAXF1O4PGXz6BnvGSwFmmY/ggNnHvTzhoaVDWQm5PKVEp9fB9uX2SuRVYvdJ5/2hSx/woWAsQZezoOzW5d5UBh9NMP6HRCbJnR68/P0yeAw0Kbrb0U/wF5baZ/mwU5AZS4Uy93S5BSxaXmQBNUKu+HNPV3vG9kCX2tYzasuoNX1DTmpGPRc1Osqf1s9PEWzH90wFwQGDrALGvmUdkshqldyDm+RRZbJRF6puBTjBOXyzrFDvpn41xjpT3SabPn7xv45du5qW+95yP/yhjD7jTthsZRgQwkduGv1XLFDZ3Mbh8hqjObLwmcl0QyyDvrLrWA/Kyw1NpZ5U0A/rA46Eh/OZIiQVEk4JAyNewqWGCR+OTvyTK0EYqIZWgpbiSgRmVoUN6gmUdLJCjCPAvJZAZ3DlEzRpnabF6zIKsCsucDhvHn4F3JXpytLz8roVK/pX3VCm5VwkcSogcfKKfRpy+SMVY0LVcRkJCI32LhXduTxYQuzSEUqHz7tz3TP8DzVQeYUWdfMtg+kxQQd466CDGdzvINA+spuwtwQ9ogjZWdKhNr+Kk7fpCMI8SY1b2+ATYKXE/mY/gkNp+rIj9EMpjNWP81rZ7pcEIxycBOzADpPadgocri4cmSYDYmbf1/QYrTcp9jr6OfD6dvLjkJkQ+pKKF522swTxTIYxNitaRLKT7WpynAwBnXJLzwzsKAGw3UBuWuKPI5kcszAVodThAZqODMHLjB8LElWyoZ4+nqmlC+zLh8dJha3I0JEVL3hWp8sOr1S6kFPjahy/vxPYpLARlqBO3ZVLgM01aueMyaAagRWrZC6OgZw0NXW002FC5Vrfm2QaxLYz4JuuUZYsXySNwAaqf5FKGgO9dsm90EboP6dyp9BRi7h5Z8nWaGP9yJp92u6KvUe93UmSfX2ZPHCScRnBNFWQE8nDgCKsnKWenZmOkL}",
			},
		},
	}
	newFormatDecryptedCredentials = []xml.Credential{
		{
			Tags: map[string]string{
				"username": "xfireadmin",
				"password": "ilovexfire",
			},
		},
		{
			Tags: map[string]string{
				"username": "gitlabadmin",
				"password": "Drmhze6EPcv0fN_81Bj",
			},
		},
		{
			Tags: map[string]string{
				"username":   "root",
				"passphrase": "IEPkO5nYhEG",
				"privateKey": `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCwgxo7cl2RajAWFseL0JAIBJbZ6dFWBGcq7+TMkP8viDwfLj4u
iYqERw+Y/lW0VZxuQuVMBfcCCINTG0S3W+MYPKiHKSaQWV53oOyPUCWaU1WjMHG4
Y3DFeE8NomOqLEOjCHDAIkZDzeEO14S8OW0fEycvR7Opo8lI6TJ4xAi9gwIDAQAB
AoGAJDu1TdCrLmd62X3xllzIxCyU/sSFiT+8Ic8+y1NUXuB7XvcyIoFvYrnnlMNY
unz8cJHg2ds7mjo/IvctAuqk0gQ5cMUgxf5QzP0ZlIcHyq8lB0YsRki0aZHSuv7r
N2KKUrayPTAJrA08GPYA+koLnY1R/yNiauo3D0ERe7KdCUECQQDolVFRXQ3KAlsu
ZGpMXbjm3H+V610D1/F4xg8qae5tKZTpbPwSnhHvCYLfNE6CSGf4JE/nhhWQk5th
j8oB4A0hAkEAwkiWXgckr+w2q+nApKDu5co27vdOPmC0q+8gk1+hCLs6Tn5V0dVx
z3aqUdQpiIYOrd5FSfXa7YNnO8VgmoayIwJAK/XFB+7ho1PsrgkWulZgk2oLx2dU
DlzrbBtrVGXvRby9Q51wy4gK9bZDgTKewCs1U4Zxf94tB0WO8dK+qLoTYQJBALcQ
4KcfAgHGiUl6C+zUO+dIoHSRkSeTxgpQW5iiPkHU8b7uqfz7q676OMi8Kpqa/w/z
5cQoJq8w50BZ3oocq5MCQBew/PwOfusahnBiUoFY0CfWTR4HZ86Uo1zgtPKoLCUG
hDA6SHkmIEPkO5nYhEGMryddRI7rsB4EKJaQ8AnJ7r4=
-----END RSA PRIVATE KEY-----`,
			},
		},
	}
)

func Test_decrypts_old_format_credentials(t *testing.T) {
	secret, _ := ioutil.ReadFile("../../test/resources/jenkins_1.625.1/decrypted/hudson.util.Secret")

	credentials, _ := DecryptCredentials(&oldFormatEncryptedCredentials, secret)

	assert.Equal(t, credentials, oldFormatDecryptedCredentials)
}

func Test_decrypts_new_format_credentials(t *testing.T) {
	secret, _ := ioutil.ReadFile("../../test/resources/jenkins_2.141/decrypted/hudson.util.Secret")

	credentials, _ := DecryptCredentials(&newFormatEncryptedCredentials, secret)

	assert.Equal(t, credentials, newFormatDecryptedCredentials)
}

package cryptography

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

const (
	// hudson.util.Secret contains '::::MAGIC::::' checksum when it's created by Jenkins
	encryptedHudsonSecretWithNoMagic = "UKFS\xe6\xedv\xf8\x987xo^\x83\xf2*\xb4\x03\x97.\xe7\xd2\xde\x14\xa3\xb6\xcfF\x9c\xa3^)q\xa5\xa0\x85h\xb1'\xaa\xb4\xad\xdb\xf0\x8dKe\x06\xa0k˪ˣX\x8c6\xe6\x14V\x86:\xeb\x1dq/\xfa\xaf?\xf5®>\xec\x83HC\x83\xf9\xc2\xf1qo\x87\x9f\xef(\xed\x06\xb7\x93`\xf2fC\xccy\xe6\xe0Bۙ\xcc\x1e5[\x9c\x9b\xa0K\x9e\xab\xb09\xecA\x1d19HS\xb3<\xa7\xa4\xec\xce\xf3\xb7\xe5\xde\x10J\x06\xdeK9zj\x85\x95\xcee\x19=}W-h\xfb\xb9\x121V\xb1F\xeeK\xf1\t\xe8\x87\xf6d\xe1\xb8\xfd%\xca:a\xdcnH\xdf\xfc\xd2\xc9[\xf8e-΅\xab\xbc\x04\xdfK'1j%\xbe\x93\x12\xfb\x00\x8a\x89\x84\xc1\x1f`\x9bڏy\xedMc\xfcGrh\xcf\x1e\xef!~\xec\xbd\xf5\xba\x97]u\xffr\t\xf7\x19X9\xcfo\xce\x15}l\xbaM\x89~\xe5s\xed\xd8:\xb6ᓋRX\x84#\xabu[\x07\xf8\xde\x1awH\xc2;b \x04\xc3"
)

func Test_return_error_when_encrypted_secret_does_not_contain_checksum(t *testing.T) {
	masterKey, _ := ioutil.ReadFile("../../test/resources/jenkins_2.141/master.key")

	_, err := DecryptHudsonSecret(masterKey, []byte(encryptedHudsonSecretWithNoMagic))

	assert.Contains(t, err.Error(), "Error. Decrypted hudson secret does not contain expected checksum.")
	assert.Contains(t, err.Error(), "::::MAGIC::::")
}

func Test_decrypts_secret(t *testing.T) {
	masterKey, _ := ioutil.ReadFile("../../test/resources/jenkins_2.141/master.key")
	encryptedSecret, _ := ioutil.ReadFile("../../test/resources/jenkins_2.141/hudson.util.Secret")
	expectedDecryptedHudsonSecret, _ := ioutil.ReadFile("../../test/resources/jenkins_2.141/decrypted/hudson.util.Secret")

	actualDecryptedHudsonSecret, _ := DecryptHudsonSecret(masterKey, encryptedSecret)

	assert.Equal(t, actualDecryptedHudsonSecret, expectedDecryptedHudsonSecret)
	assert.True(t, len(actualDecryptedHudsonSecret) > 1)
}

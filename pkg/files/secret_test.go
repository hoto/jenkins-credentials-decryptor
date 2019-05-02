package files

import (
	"testing"
)

func Test_something(t *testing.T) {
	Decrypt(
		"../../test/resources/master.key",
		"../../test/resources/hudson.util.Secret")
}

package files

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_something(t *testing.T) {
	secret, err := Decrypt(
		"../../test/resources/master.key",
		"../../test/resources/hudson.util.Secret")
	check(err)
	fmt.Printf("Secret:%q", secret)
	assert.True(t, true)
}

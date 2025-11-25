package testutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertNilError(t *testing.T, err error) {
	var msg string
	if err != nil {
		msg = err.Error()
	}
	assert.Nil(t, err, msg)
}

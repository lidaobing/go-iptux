package errors

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	assert.Equal(t, NewNotImplementedError().Error(), "libiptux.Error: [0] not implemented")
	assert.Equal(t, NewErrorWithCausedBy(NotImplementd, "hello world", NewNotImplementedError()).Error(),
		`
libiptux.Error: [0] hello world
  Caused By: libiptux.Error: [0] not implemented`[1:])
}
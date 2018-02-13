package libiptux

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewShareFile(t *testing.T) {
	shareFile := NewShareFile(nil)
	assert.NotNil(t, shareFile)
}
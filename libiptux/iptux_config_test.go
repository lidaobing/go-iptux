package libiptux_test

import (
	"testing"
	"github.com/lidaobing/go-iptux/libiptux"
	"github.com/stretchr/testify/assert"
)

func TestNewIptuxConfig(t *testing.T) {
	_, err := libiptux.NewIptuxConfig("foobar")
	assert.Nil(t, err)
	assert.NoError(t, err)
}

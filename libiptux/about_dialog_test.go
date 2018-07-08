package libiptux_test

import (
	"testing"
	"github.com/gotk3/gotk3/gtk"
	"github.com/stretchr/testify/assert"
	"github.com/lidaobing/go-iptux/libiptux"
)

func TestNewAboutDialog(t *testing.T) {
	gtk.Init(nil)
	assert.NotNil(t, libiptux.NewAboutDialog())
}

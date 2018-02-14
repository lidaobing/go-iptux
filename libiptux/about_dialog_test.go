package libiptux_test

import (
	"testing"
	"github.com/mattn/go-gtk/gtk"
	"github.com/stretchr/testify/assert"
	"github.com/lidaobing/go-iptux/libiptux"
)

func TestNewAboutDialog(t *testing.T) {
	gtk.Init(nil)
	assert.NotNil(t, libiptux.NewAboutDialog())
}

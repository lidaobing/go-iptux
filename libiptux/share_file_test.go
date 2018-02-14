package libiptux_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/lidaobing/go-iptux/libiptux"
	"github.com/mattn/go-gtk/gtk"
)

func TestNewShareFile(t *testing.T) {
	gtk.Init(nil)
	shareFile := libiptux.NewShareFile(nil)
	assert.NotNil(t, shareFile)
}
package libiptux_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gotk3/gotk3/gtk"
	"github.com/lidaobing/gtkmust"
	"github.com/gotk3/gotk3/glib"
	"github.com/lidaobing/go-iptux/libiptux"
)

func TestNewMainWindow(t *testing.T) {
	gtk.Init(nil)
	app := gtkmust.ApplicationNew("hello.world", glib.APPLICATION_FLAGS_NONE)
	assert.NotNil(t, libiptux.NewMainWindow(app))
}

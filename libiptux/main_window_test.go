package libiptux
import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/mattn/go-gtk/gtk"
)

func TestNewMainWindow(t *testing.T) {
	gtk.Init(nil)
	assert.NotNil(t, NewMainWindow())
}

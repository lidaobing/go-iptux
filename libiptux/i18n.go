package libiptux

import (
	"github.com/leonelquinteros/gotext"
)

func T(str string, vars ...interface{}) string {
	return gotext.Get(str, vars...)
}

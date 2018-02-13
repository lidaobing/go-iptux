package libiptux

import (
	"gopkg.in/leonelquinteros/gotext.v1"
)

func T(str string, vars ...interface{}) string {
	return gotext.Get(str, vars...)
}

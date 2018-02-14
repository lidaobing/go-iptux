package libiptux

import (
	"encoding/json"
	"os"
	"io/ioutil"

	"github.com/lidaobing/go-iptux/libiptux/errors"
)

type IptuxConfig struct {
	fileName string
	data map[string]interface{}
}

func NewIptuxConfig(fname string) (res *IptuxConfig, err error) {
	_, err = os.Stat(fname)
	if os.IsNotExist(err) {
		return &IptuxConfig{fname, nil}, nil
	}

	if err != nil {
		return nil, errors.WrapError(err)
	}

	res2 := &IptuxConfig{
		fileName: fname,
		data: nil,
	}

	file, err := os.Open(fname)
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, res2.data)
	if err != nil {
		return
	}

	res = res2
	return
}
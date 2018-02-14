package main

import (
	"github.com/mattn/go-gtk/gtk"
	"github.com/lidaobing/go-iptux/libiptux"
	"gopkg.in/leonelquinteros/gotext.v1"
)

func main() {
	gotext.SetDomain("go-iptux")

	gtk.Init(nil)
	shareFile := libiptux.NewMainWindow()
	shareFile.ShowAll()
	gtk.Main()
}

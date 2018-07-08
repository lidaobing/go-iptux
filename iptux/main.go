package main

import (
	"github.com/lidaobing/go-iptux/libiptux"
	"gopkg.in/leonelquinteros/gotext.v1"
	"github.com/lidaobing/gtkmust"
	"github.com/gotk3/gotk3/glib"
	"os"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gotext.SetDomain("go-iptux")

	const appName = "com.github.lidaobing.go-iptux"

	app := gtkmust.ApplicationNew(appName, glib.APPLICATION_FLAGS_NONE)
	app.Connect("activate", func() {
		win := libiptux.NewMainWindow(app)
		win.ShowAll()
		win.SetPosition(gtk.WIN_POS_MOUSE)
		win.SetDefaultSize(400, 300)
	})
	os.Exit(app.Run(os.Args))
	}

package main

import (
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/lidaobing/go-iptux/libiptux"
	"github.com/lidaobing/gtkmust"
	"gopkg.in/leonelquinteros/gotext.v1"
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

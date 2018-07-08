package libiptux

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/lidaobing/gtkmust"
)

const (
	Version = "0.0.1-git"
)

type AboutDialog struct {
	gtk.AboutDialog
}

func NewAboutDialog(parent gtk.IWindow) *AboutDialog {
	self := &AboutDialog{*gtkmust.AboutDialogNew()}
	self.SetTransientFor(parent)
	self.SetProgramName(T("iptux"))
	self.SetVersion(Version)
	self.SetCopyright("Copyright 2018 go-iptux Authors.")
	self.SetComments("A GTK+ based LAN Messenger.")
	self.SetWebsite("https://github.com/lidaobing/go-iptux")
	self.SetLicense("GPL 2+")
	self.SetLogoIconName("iptux")
	return self
}

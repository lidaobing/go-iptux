package libiptux

import "github.com/mattn/go-gtk/gtk"

type MainWindow struct {
	gtk.Window
}

func NewMainWindow() *MainWindow {
	res := &MainWindow{*gtk.NewWindow(gtk.WINDOW_TOPLEVEL)}
	res.SetTitle(T("iptux"))

	box := gtk.NewVBox(false, 0)
	res.Add(box)

	menubar := gtk.NewMenuBar()
	box.Add(menubar)

	menu := gtk.NewMenuItemWithMnemonic(T("_Tools"))
	menubar.Append(menu)

	submenu := gtk.NewMenu()
	menu.SetSubmenu(submenu)

	menuItem := gtk.NewImageMenuItemWithMnemonic(T("_Shared Management"))
	image := gtk.NewImageFromStock(gtk.STOCK_ABOUT, gtk.ICON_SIZE_MENU)
	menuItem.SetImage(&image.Widget)
	submenu.Append(menuItem)
	menuItem.Connect("activate", res.onSharedManagement)
	return res
}

func (self *MainWindow) onSharedManagement() {
	shareFile := NewShareFile(&self.Window)
	shareFile.Run()
}
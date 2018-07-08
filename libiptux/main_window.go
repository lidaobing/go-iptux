package libiptux

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/lidaobing/gtkmust"
)

type MainWindow struct {
	gtk.ApplicationWindow
}

func NewMainWindow(app *gtk.Application) *MainWindow {
	res := &MainWindow{*gtkmust.ApplicationWindowNew(app)}
	res.SetTitle(T("iptux"))

	box := gtkmust.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	res.Add(box)

	menubar := gtkmust.MenuBarNew()
	box.Add(menubar)
	menubar.Append(res.NewToolsMenu())
	menubar.Append(res.NewHelpMenu())

	label := gtkmust.LabelNew("Hello World")
	box.Add(label)

	return res
}

func (window *MainWindow) NewToolsMenu() gtk.IMenuItem {
	menu := gtkmust.MenuItemNewWithMnemonic(T("_Tools"))

	submenu := gtkmust.MenuNew()
	menu.SetSubmenu(submenu)

	menuItem := gtkmust.MenuItemNewWithMnemonic(T("_Shared Management"))

	//image := gtk.NewImageFromStock(gtk.STOCK_ABOUT, gtk.ICON_SIZE_MENU)
	//menuItem.SetImage(&image.Widget)
	submenu.Append(menuItem)
	menuItem.Connect("activate", window.onSharedManagement)
	return menu
}
func (window *MainWindow) NewHelpMenu() gtk.IMenuItem {
	menu := gtkmust.MenuItemNewWithMnemonic(T("_Help"))

	submenu := gtkmust.MenuNew()
	menu.SetSubmenu(submenu)

	menuItem := gtkmust.MenuItemNewWithMnemonic(T("_About"))
	//image := gtk.NewImageFromIconName(gtk.STOCK_ABOUT, gtk.ICON_SIZE_MENU)
	//image := gtk.NewImageFromIconName(gtk.STOCK_ABOUT)
	//menuItem.SetImage(&image.Widget)
	submenu.Append(menuItem)
	menuItem.Connect("activate", window.onAbout)
	return menu
}

func (self *MainWindow) onSharedManagement() {
	shareFile := NewShareFile(&self.Window)
	shareFile.Run()
	shareFile.Destroy()
}

func (self *MainWindow) onAbout() {
	dialog := NewAboutDialog(self)
	dialog.Run()
	dialog.Destroy()
}

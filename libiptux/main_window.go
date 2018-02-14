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
	menubar.Append(res.NewToolsMenu())
	menubar.Append(res.NewHelpMenu())

	return res
}

func (window *MainWindow) NewToolsMenu() gtk.IWidget {
	menu := gtk.NewMenuItemWithMnemonic(T("_Tools"))

	submenu := gtk.NewMenu()
	menu.SetSubmenu(submenu)

	menuItem := gtk.NewImageMenuItemWithMnemonic(T("_Shared Management"))
	image := gtk.NewImageFromIconName(gtk.STOCK_ABOUT, gtk.ICON_SIZE_MENU)
	//image := gtk.NewImageFromIconName(gtk.STOCK_ABOUT)
	menuItem.SetImage(&image.Widget)
	submenu.Append(menuItem)
	menuItem.Connect("activate", window.onSharedManagement)
	return menu
}
func (window *MainWindow) NewHelpMenu() gtk.IWidget {
	menu := gtk.NewMenuItemWithMnemonic(T("_Help"))

	submenu := gtk.NewMenu()
	menu.SetSubmenu(submenu)

	menuItem := gtk.NewImageMenuItemWithMnemonic(T("_About"))
	image := gtk.NewImageFromIconName(gtk.STOCK_ABOUT, gtk.ICON_SIZE_MENU)
	//image := gtk.NewImageFromIconName(gtk.STOCK_ABOUT)
	menuItem.SetImage(&image.Widget)
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
	dialog := NewAboutDialog()
	dialog.Run()
	dialog.Destroy()
}

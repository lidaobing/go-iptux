package libiptux

import (
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gdk"
	"fmt"
	"os"
)

type ShareFile struct {
	gtk.Dialog
	model gtk.ITreeModel
}

func NewShareFile(window *gtk.Window) *ShareFile {
	model := gtk.NewListStore(glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_UINT)


	res := &ShareFile{
		*gtk.NewDialog(),
		model,
	}

	sortable := gtk.NewTreeSortable(model)
	sortable.SetDefaultSortFunc(res.fileTreeCompareFunc)
	sortable.SetSortColumnId(gtk.TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, gtk.SORT_ASCENDING)

	res.SetTitle(T("Shared Files Management"))
	res.SetParent(window)
	res.AddButton(T("OK"), gtk.RESPONSE_OK)
	res.AddButton(T("Apply"), gtk.RESPONSE_APPLY)
	res.AddButton(T("Cancel"), gtk.RESPONSE_CANCEL)

	res.SetDefaultResponse(gtk.RESPONSE_OK)
	res.SetResizable(false)
	res.SetPosition(gtk.WIN_POS_CENTER);
	res.SetBorderWidth(5);
	res.SetSizeRequest(500, 350)

	entries := []gtk.TargetEntry{
		{"text/uri-list", 0, 0},
	}
	res.DragDestSet(gtk.DEST_DEFAULT_ALL, entries, gdk.ACTION_MOVE)
	// TODO: add DragDataReceived
	res.Connect("drag-data-received", nil)

	res.GetVBox().PackStart(res.createAllArea(), true, true, 0)
	return res
}

func (self *ShareFile) createAllArea() *gtk.HBox {
	res := gtk.NewHBox(false, 0)

	sw := gtk.NewScrolledWindow(nil, nil)
	sw.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	sw.SetShadowType(gtk.SHADOW_ETCHED_IN)
	sw.Add(self.createFileTree())
	res.Add(sw)

	vbox := gtk.NewVBox(false, 0);
	button := gtk.NewButtonWithLabel(T("Add Files"))
	vbox.PackStart(button, false, false, 0)
	button.Connect("clicked", self.addRegular)

	vbox.PackStart(gtk.NewButtonWithLabel(T("Add Folders")), false, false, 0)
	vbox.PackStart(gtk.NewButtonWithLabel(T("Delete Resources")), false, false, 0)
	vbox.PackEnd(gtk.NewButtonWithLabel(T("Clear Password")), false, false, 0)
	vbox.PackEnd(gtk.NewButtonWithLabel(T("Set Password")), false, false, 0)

	res.PackStart(vbox, false, false, 0)
	res.ShowAll()
	return res
}
func (self *ShareFile) createFileTree() *gtk.TreeView {
	res := gtk.NewTreeView()
	res.SetModel(self.model)
	res.SetHeadersVisible(true)
	res.SetRubberBanding(true)

	res.GetSelection().SetMode(gtk.SELECTION_MULTIPLE)

	column := gtk.NewTreeViewColumn()
	column.SetResizable(true)
	column.SetTitle(T("File"))

	var cell gtk.ICellRenderer
	cell = gtk.NewCellRendererPixbuf()
	column.PackStart(cell, false)
	column.AddAttribute(cell, "icon-name", 0)

	cell = gtk.NewCellRendererText()
	column.PackStart(cell, false)
	column.AddAttribute(cell, "text", 1)

	res.AppendColumn(column)

	column = gtk.NewTreeViewColumnWithAttributes(T("Size"), gtk.NewCellRendererText(), "text", 2)
	column.SetResizable(true)
	res.AppendColumn(column)

	column = gtk.NewTreeViewColumnWithAttributes(T("Type"), gtk.NewCellRendererText(), "text", 3)
	column.SetResizable(true)
	res.AppendColumn(column)
	return res
}

func (self *ShareFile) fileTreeCompareFunc(m *gtk.TreeModel, a *gtk.TreeIter, b *gtk.TreeIter) int {
	var aFilePath string
	var bFilePath string

	value := glib.ValueFromNative("")

	m.GetValue(a, 1, value)
	aFilePath = value.GetString()

	m.GetValue(b, 1, value)
	bFilePath = value.GetString()

	if aFilePath < bFilePath {
		return -1
	} else if aFilePath == bFilePath {
		return 0
	} else {
		return 1
	}
}

func (self *ShareFile) addRegular() {
	dialog := gtk.NewFileChooserDialog(
		T("Choose the files to share"),
		&self.Window,
		gtk.FILE_CHOOSER_ACTION_OPEN,
		gtk.STOCK_OPEN,
		gtk.RESPONSE_ACCEPT,
		gtk.STOCK_CANCEL,
		gtk.RESPONSE_CANCEL,
	)
	dialog.SetDefaultResponse(gtk.RESPONSE_ACCEPT)
	dialog.SetLocalOnly(false)
	dialog.SetSelectMultiple(true)
	dialog.SetCurrentFolder(glib.GetHomeDir())
	defer dialog.Destroy()

	var list *glib.SList

	switch dialog.Run() {
	case gtk.RESPONSE_ACCEPT:
		list = dialog.GetFilenames()
	default:
		return
	}

	for i := uint(0); i < list.Length(); i++ {
		fmt.Fprintln(os.Stdout,"%s", glib.GPtrToString(list.Nth(i).Data()))
	}
}
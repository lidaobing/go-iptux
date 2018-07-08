package libiptux

import (
	"C"
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/glib"
	"github.com/lidaobing/gtkmust"
	"unsafe"
)

type ShareFile struct {
	gtk.Dialog
	model gtk.ITreeModel
}

func NewShareFile(window *gtk.Window) *ShareFile {
	model := gtkmust.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_UINT)


	res := &ShareFile{
		*gtkmust.DialogNew(),
		model,
	}


	//sortable := gtk.NewTreeSortable(model)
	//sortable.SetDefaultSortFunc(res.fileTreeCompareFunc)
	//sortable.SetSortColumnId(gtk.TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, gtk.SORT_ASCENDING)

	res.SetTitle(T("Shared Files Management"))
	res.SetTransientFor(window)
	res.AddButton(T("OK"), gtk.RESPONSE_OK)
	res.AddButton(T("Apply"), gtk.RESPONSE_APPLY)
	res.AddButton(T("Cancel"), gtk.RESPONSE_CANCEL)

	res.SetDefaultResponse(gtk.RESPONSE_OK)
	res.SetResizable(false)
	res.SetPosition(gtk.WIN_POS_CENTER);
	res.SetBorderWidth(5);
	res.SetSizeRequest(500, 350)

	//entries := []gtk.TargetEntry{
	//	*gtkmust.TargetEntryNew("text/uri-list", 0, 0),
	//}
	//res.DragDestSet(gtk.DEST_DEFAULT_ALL, entries, gdk.ACTION_MOVE)
	// TODO: add DragDataReceived
	//res.Connect("drag-data-received", nil)
	gtkmust.MustBox(res.GetContentArea()).PackStart(res.createAllArea(), true, true, 0)
	return res
}

func (self *ShareFile) createAllArea() *gtk.Box {
	res := gtkmust.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0);
	sw := gtkmust.ScrolledWindowNew(nil, nil)

	sw.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	sw.SetShadowType(gtk.SHADOW_ETCHED_IN)
	sw.Add(self.createFileTree())
	res.Add(sw)

	vbox := gtkmust.BoxNew(gtk.ORIENTATION_VERTICAL, 0)

	button := gtkmust.ButtonNewWithLabel(T("Add Files"))
	vbox.PackStart(button, false, false, 0)
	button.Connect("clicked", self.addRegular)

	vbox.PackStart(gtkmust.ButtonNewWithLabel(T("Add Folders")), false, false, 0)
	vbox.PackStart(gtkmust.ButtonNewWithLabel(T("Delete Resources")), false, false, 0)
	vbox.PackEnd(gtkmust.ButtonNewWithLabel(T("Clear Password")), false, false, 0)
	vbox.PackEnd(gtkmust.ButtonNewWithLabel(T("Set Password")), false, false, 0)

	res.PackStart(vbox, false, false, 0)
	res.ShowAll()
	return res
}
func (self *ShareFile) createFileTree() *gtk.TreeView {
	res := gtkmust.TreeViewNew()
	res.SetModel(self.model)
	res.SetHeadersVisible(true)
	res.SetRubberBanding(true)

	gtkmust.MustTreeSelection(res.GetSelection()).SetMode(gtk.SELECTION_MULTIPLE)

	column := gtkmust.TreeViewColumnNew()
	column.SetResizable(true)
	column.SetTitle(T("File"))

	var cell gtk.ICellRenderer
	cell = gtkmust.CellRendererPixbufNew()
	column.PackStart(cell, false)
	column.AddAttribute(cell, "icon-name", 0)

	cell = gtkmust.CellRendererTextNew()
	column.PackStart(cell, false)
	column.AddAttribute(cell, "text", 1)

	res.AppendColumn(column)

	column = gtkmust.TreeViewColumnNewWithAttribute(T("Size"), gtkmust.CellRendererTextNew(), "text", 2)
	column.SetResizable(true)
	res.AppendColumn(column)

	column = gtkmust.TreeViewColumnNewWithAttribute(T("Type"), gtkmust.CellRendererTextNew(), "text", 3)
	column.SetResizable(true)
	res.AppendColumn(column)
	return res
}


func (self *ShareFile) addRegular() {
	dialog := gtkmust.FileChooserDialogNewWith2Buttons(
		T("Choose the files to share"),
		&self.Window,
		gtk.FILE_CHOOSER_ACTION_OPEN,
		T("Open"),
		gtk.RESPONSE_ACCEPT,
		T("Cancel"),
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
		list = gtkmust.MustSList(dialog.GetFilenames())
	default:
		return
	}

	for l := list; l != nil; l = l.Next() {
		l.Native()
		s := C.GoString(*((**C.char)(unsafe.Pointer(l.Native()))))
		fmt.Println(s)
	}
}

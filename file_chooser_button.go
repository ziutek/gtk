package gtk

/*
#include <gtk/gtkfilechooserbutton.h>
*/
import "C"

import (
	"github.com/ziutek/glib"
)

type FileChooserButton struct {
	HBox
}

func (f *FileChooserButton) g() *C.GtkFileChooserButton {
	return (*C.GtkFileChooserButton)(f.GetPtr())
}

func (f *FileChooserButton) AsFileChooserButton() *FileChooserButton {
	return f
}

func NewFileChooserButton(title string, action FileChooserAction) *FileChooserButton {
	f := new(FileChooserButton)
	f.SetPtr(glib.Pointer(C.gtk_file_chooser_button_new(
		(*C.gchar)(C.CString(title)),
		action.g(),
	)))
	return f
}

package gtk

/*
#include <stdlib.h>
#include <gtk/gtkfilechooser.h>

GtkFileChooser* _gtk_file_chooser_cast(GtkWidget* w) {
	return GTK_FILE_CHOOSER(w);
}

char* _gtk_file_chooser_get_filename(GtkFileChooser* f) {
	return gtk_file_chooser_get_filename(f);
}
*/
import "C"

import (
	"unsafe"
	"github.com/ziutek/glib"
)

type FileChooserAction C.GtkFileChooserAction

const (
	FILE_CHOOSER_ACTION_OPEN          = FileChooserAction(C.GTK_FILE_CHOOSER_ACTION_OPEN)
	FILE_CHOOSER_ACTION_SAVE          = FileChooserAction(C.GTK_FILE_CHOOSER_ACTION_SAVE)
	FILE_CHOOSER_ACTION_SELECT_FOLDER = FileChooserAction(C.GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER)
	FILE_CHOOSER_ACTION_CREATE_FOLDER = FileChooserAction(C.GTK_FILE_CHOOSER_ACTION_CREATE_FOLDER)
)

func (f FileChooserAction) g() C.GtkFileChooserAction {
	return C.GtkFileChooserAction(f)
}


type FileChooser C.GtkFileChooser

func (f *FileChooser) g() *C.GtkFileChooser {
	return (*C.GtkFileChooser)(f)
}

func (f *FileChooser) Type() glib.Type {
	return glib.TypeFromName("GtkFileChooser")
}

func (f *FileChooser) SetAction(action FileChooserAction) {
	C.gtk_file_chooser_set_action(f.g(), action.g())
}

func (f *FileChooser) GetAction() FileChooserAction {
	return FileChooserAction(C.gtk_file_chooser_get_action(f.g()))
}

func (f *FileChooser) GetFilename() string {
	n := C._gtk_file_chooser_get_filename(f.g())
	defer C.free(unsafe.Pointer(n))
	return C.GoString(n)
}

func FileChooserCast(w *Widget) *FileChooser {
	return (*FileChooser)(C._gtk_file_chooser_cast(w.g()))
}

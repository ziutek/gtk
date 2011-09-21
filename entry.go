package gtk

/*
#include <gtk/gtk.h>
*/
import "C"

import (
	"github.com/ziutek/glib"
)

type Entry struct {
	Widget
}

func (e *Entry) g() *C.GtkEntry {
	return (*C.GtkEntry)(e.GetPtr())
}

func (e *Entry) AsEntry() *Entry {
	return e
}

func (e *Entry) GetText() string {
	return C.GoString((*C.char)(C.gtk_entry_get_text(e.g())))
}

func (e *Entry) SetMaxLength(max int) {
	C.gtk_entry_set_max_length(e.g(), C.gint(max))
}

func NewEntry() *Entry {
	e := new(Entry)
	e.SetPtr(glib.Pointer(C.gtk_entry_new()))
	return e
}

package gtk

/*
#include <gtk/gtk.h>
*/
import "C"

import (
	"github.com/ziutek/glib"
)

type Settings struct {
	glib.Object
}

func (s *Settings) g() *C.GtkSettings {
	return (*C.GtkSettings)(s.GetPtr())
}

func (s *Settings) AsSettings() *Settings {
		return s
}

var settings_origin = (*C.gchar)(C.CString("Go code"))

func (s *Settings) Set(name string, val interface{}) {
	sval := C.GtkSettingsValue{
		origin: settings_origin,
		value: *v2g(glib.ValueOf(val)),
	}
	C.gtk_settings_set_property_value(s.g(), (*C.gchar)(C.CString(name)), &sval)
}

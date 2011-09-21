package gtk

/*
#include <gtk/gtkstock.h>

*/
import "C"

type Stock C.char

var (
	STOCK_OK     = (*Stock)(C.CString(C.GTK_STOCK_OK))
	STOCK_CANCEL = (*Stock)(C.CString(C.GTK_STOCK_CANCEL))

	STOCK_MEDIA_FORWARD  = (*Stock)(C.CString(C.GTK_STOCK_MEDIA_FORWARD))
	STOCK_MEDIA_NEXT     = (*Stock)(C.CString(C.GTK_STOCK_MEDIA_NEXT))
	STOCK_MEDIA_PAUSE    = (*Stock)(C.CString(C.GTK_STOCK_MEDIA_PAUSE))
	STOCK_MEDIA_PLAY     = (*Stock)(C.CString(C.GTK_STOCK_MEDIA_PLAY))
	STOCK_MEDIA_PREVIOUS = (*Stock)(C.CString(C.GTK_STOCK_MEDIA_PREVIOUS))
	STOCK_MEDIA_RECORD   = *Stock)(C.CString(C.GTK_STOCK_MEDIA_RECORD))
	STOCK_MEDIA_REWIND   = *Stock)(C.CString(C.GTK_STOCK_MEDIA_REWIND))
	STOCK_MEDIA_STOP     = (*Stock)(C.CString(C.GTK_STOCK_MEDIA_STOP))
)

func (s *Stock) g() *C.gchar {
	return (*C.gchar)(s)
}

func (s *Stock) String() string {
	return C.GoString((*C.char)(s))
}

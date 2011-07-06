package gtk

/*
#include <gtk/gtk.h>

char** _gtk_init(int* argc, char** argv) {
	gtk_init(argc, &argv);
	return argv;
}

#cgo pkg-config: gtk+-2.0
*/
import "C"

import (
	"os"
	"unsafe"
)

func init() {
	alen := C.int(len(os.Args))
	argv := make([]*C.char, alen)
	for i, s := range os.Args {
		argv[i] = C.CString(s)
	}
	ret := C._gtk_init(&alen, &argv[0])
	argv = (*[1<<16]*C.char)(unsafe.Pointer(ret))[:alen]
	os.Args = make([]string, alen)
	for i, s := range argv {
		os.Args[i] = C.GoString(s)
	}
}

func Main() {
	C.gtk_main()
}

func MainQuit() {
	C.gtk_main_quit()
}

package winapi_lib

// #include <windows.h>
// #include <tchar.h>
import "C"

func MessageBox(title string, text string) {
	CTitle := C.CString(title)
	CText := C.CString(text)

	C.MessageBox((*C.struct_HWND__)(C.NULL), CText, CTitle, 0)
}

type Window struct {
	Instance C.HINSTANCE
	Class    C.WNDCLASSW
	Hwnd     C.HWND
}

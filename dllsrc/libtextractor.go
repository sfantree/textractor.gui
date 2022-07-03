package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"github.com/gloomyzerg/textractor"
	"unsafe"
)

var result = &(textractor.Text{})
var cs = C.CString("")
var cg = string("")
var free = int(0)

//export Extractor
func Extractor(source *C.char) *C.char {
	//result, err := textractor.Extract(string(source))
	cg = C.GoString(source)
	result, _ = textractor.Extract(cg)
    if len(result.Content) > 0 {
		content := (result.Content)
		cs = C.CString(content)
		free = 1
		//C.print(cs)
	} else {
		cs = C.CString("")
		free = 1
	}
	return cs
}

//export CleanSource
func CleanSource(*C.char) {
	if (free == 1) {
		C.free(unsafe.Pointer(cs))
		//C.free(unsafe.Pointer(cg))
	}
}

func main()  {

}
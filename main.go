package main //import cgo-with-so

//头文件的位置，相对于源文件是当前目录，所以是 .，头文件在多个目录时写多个  #cgo CFLAGS:
// 从哪里加载动态库，位置与文件名，-ladd 加载 libadd.so 文件

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L ./lib -ladd -Wl,-rpath,lib
#include "lib/add.h"
*/
import "C"

import "fmt"

func main() {
	val := C.Add(C.CString("cgo"), 2022)
	fmt.Println("CGO: ", C.GoString(val))
}

package main

// #cgo CFLAGS: -g -I${SRCDIR}/c-lib
// #cgo LDFLAGS: -L${SRCDIR}/c-lib -lgreet
// #include <stdio.h>
// #include <string.h>
// #include <stdlib.h>
// #include "greet.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	C.hello()

	p := C.greet()
	s := C.GoString(p)

	name := C.CString("abhijit")
	// C.CString function allocates the memory for you, but won't release it, so it is your responsability to freed the memory with a call to C.free as following:
	defer C.free(unsafe.Pointer(name))

	pGreetWithName := C.greetWithName(name)
	sGreetWithName := C.GoString(pGreetWithName)

	fmt.Printf("Inside main.go: %v\n", s)
	fmt.Printf("greetWithName: %v\n", sGreetWithName)

	add := C.add(2, 3)

	fmt.Printf("add: %v\n", add)
}

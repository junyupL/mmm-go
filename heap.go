package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func sizeof[T any]() C.ulonglong {
	var a T
	return C.ulonglong(unsafe.Sizeof(a))
}

func new[T any]() *T {
	return (*T)(C.malloc(sizeof[T]()))
}

func delete[T any](ptr *T) {
	C.free(unsafe.Pointer(ptr))
}

func main() {
	var p = new[int]()
	*p = 4
	fmt.Print(*p)
	delete(p)
}

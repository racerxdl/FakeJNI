package main

import (
	"fmt"
	"github.com/racerxdl/FakeJNI/fakejni"
	"github.com/racerxdl/FakeJNI/gojni"
	"unsafe"
)

/*
#cgo LDFLAGS:  -L. -ldl
#include <stdlib.h>
#include <dlfcn.h>


typedef int (*JNI_OnLoad_t)(void *, void *);
typedef int (*Java_HelloJNI_sayHello_t)(void *, int);	// JString are actually ints

int callOnLoad(void *onload, void *vm, void *reserved) {
	if (onload != NULL) {
		return ((JNI_OnLoad_t)onload)(vm, reserved);
	}

	return -1;
}

void callJava_HelloJNI_sayHello(void *Java_HelloJNI_sayHello, void *vm, int arg) {
	if (Java_HelloJNI_sayHello != NULL) {
		((Java_HelloJNI_sayHello_t)Java_HelloJNI_sayHello)(vm, arg);
	}
}

*/
import "C"

func tryCallOnload(jvm, libhandle unsafe.Pointer) bool {
	// JNI_OnLoad might not exists, but it exists, we should call it
	funcName := C.CString("JNI_OnLoad")
	defer C.free(unsafe.Pointer(funcName))

	onload := C.dlsym(libhandle, funcName)
	if onload == nil {
		return false
	}
	C.callOnLoad(unsafe.Pointer(onload), jvm, nil)
	return true
}

func main() {
	jvm := fakejni.GetJVM()
	jnienv := fakejni.GetJNIEnv()
	jni := gojni.GetJNI()

	libName := C.CString("./java/libhello.so")
	defer C.free(unsafe.Pointer(libName))

	handle := C.dlopen(libName, C.RTLD_LAZY)
	if handle == nil {
		panic("error opening ./java/libhello.so")
	}

	defer func() {
		if r := C.dlclose(handle); r != 0 {
			panic("error closing ./java/libhello.so")
		}
	}()

	if !tryCallOnload(unsafe.Pointer(jvm), handle) {
		fmt.Printf("No OnLoad function, skipping")
	}

	funcName := C.CString("Java_HelloJNI_sayHello")
	defer C.free(unsafe.Pointer(funcName))

	Java_HelloJNI_sayHello := C.dlsym(unsafe.Pointer(handle), funcName)
	if Java_HelloJNI_sayHello == nil {
		panic("cannot find function Java_HelloJNI_sayHello")
	}

	str, _ := jni.NewStringUTF("Hello from Go!")
	C.callJava_HelloJNI_sayHello(Java_HelloJNI_sayHello, unsafe.Pointer(jnienv), C.int(str))
}

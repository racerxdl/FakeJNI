package fakejni

/*
#include "fakejni.h"
*/
import "C"

func GetJNIEnv() uintptr {
	return uintptr(C.GetJNIEnv())
}

func GetJVM() uintptr {
	return uintptr(C.GetJVM())
}

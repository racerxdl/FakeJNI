package fakejni

import "C"

type JavaVM uintptr

//export JVM_DestroyJavaVM
func JVM_DestroyJavaVM(jvm *JavaVM) JInt {
	log.Warnf("JVM_DestroyJavaVM(%d): NOT IMPLEMENTED!!", jvm)
	return 0
}

//export JVM_AttachCurrentThread
func JVM_AttachCurrentThread(jvm *JavaVM) JInt {
	log.Warnf("JVM_AttachCurrentThread(%d): NOT IMPLEMENTED!!", jvm)
	return 0
}

//export JVM_DetachCurrentThread
func JVM_DetachCurrentThread(jvm *JavaVM) JInt {
	log.Warnf("JVM_DetachCurrentThread(%d): NOT IMPLEMENTED!!", jvm)
	return 0
}

//export JVM_AttachCurrentThreadAsDaemon
func JVM_AttachCurrentThreadAsDaemon(jvm *JavaVM) JInt {
	log.Warnf("JVM_AttachCurrentThreadAsDaemon(%d): NOT IMPLEMENTED!!", jvm)
	return 0
}

//export JVM_LogDebug
func JVM_LogDebug(str *C.char) {
	goStr := C.GoString(str)
	log.Debug(goStr)
}

package fakejni

/*
#include "types.h"
#include <stdlib.h>
*/
import "C"
import (
	"github.com/racerxdl/FakeJNI/gojni"
	"github.com/sirupsen/logrus"
	"unsafe"
)

var log = logrus.New()

type CJNIEnv uintptr
type JClass C.int
type JMethodID C.int
type JObject C.int
type JFieldID C.int
type JString C.int
type JInt C.int32_t
type JLong C.int64_t
type JFloat C.float
type JDouble C.double
type JBool C.uint8_t
type JByte C.char
type JSize JInt

//export JNI_FindClass
func JNI_FindClass(env CJNIEnv, cClassName *C.char) JClass {
	className := C.GoString(cClassName)
	log.Debugf("JNI_FindClass(%s)", className)
	classId, err := gojni.GetJNI().FindClass(className)

	if err != nil {
		log.Errorf("Class %s not found! %s\n", className, err)
		return -1
	}

	return JClass(classId)
}

//export JNI_GetObjectMethodIDArgs
func JNI_GetObjectMethodIDArgs(env CJNIEnv, obj JObject, methodId JMethodID, args *C.struct_MethodArgType) C.int {
	log.Debugf("JNI_GetObjectMethodIDArgs(%d, %d, %d)", obj, methodId, args)
	class, _ := gojni.GetJNI().GetObjectClass(gojni.JObject(obj))
	typedArgs := gojni.GetJNI().GetMethodIDArgs(class, gojni.JMethodID(methodId))
	typedArgsLen := len(typedArgs)
	if args != nil {
		argsSlice := (*[1 << 20]C.struct_MethodArgType)(unsafe.Pointer(args))[:typedArgsLen:typedArgsLen]
		for i, v := range typedArgs {
			argsSlice[i]._type = C.int(v.ArgType)
		}
	}
	return C.int(typedArgsLen)
}

//export JNI_GetStaticMethodIDArgs
func JNI_GetStaticMethodIDArgs(env CJNIEnv, class JClass, methodId JMethodID, args *C.struct_MethodArgType) C.int {
	log.Debugf("JNI_GetStaticMethodIDArgs(%d, %d, %d)", class, methodId, args)
	typedArgs := gojni.GetJNI().GetStaticMethodIDArgs(gojni.JClass(class), gojni.JMethodID(methodId))
	typedArgsLen := len(typedArgs)
	if args != nil {
		argsSlice := (*[1 << 20]C.struct_MethodArgType)(unsafe.Pointer(args))[:typedArgsLen:typedArgsLen]
		for i, v := range typedArgs {
			argsSlice[i]._type = C.int(v.ArgType)
		}
	}
	return C.int(typedArgsLen)
}

//export JNI_GetStaticMethodID
func JNI_GetStaticMethodID(env CJNIEnv, c JClass, Cname, Carg *C.char) JMethodID {
	name := C.GoString(Cname)
	arg := ""
	if Carg != nil {
		arg = C.GoString(Carg)
	}
	className := gojni.GetJNI().GetClassName(gojni.JClass(c))
	log.Debugf("JNI_GetStaticMethodID(%s, %s, %s)", className, name, arg)

	method, err := gojni.GetJNI().GetStaticMethodID(gojni.JClass(c), name, arg)
	if err != nil {
		log.Errorf("Static Method %s(%s) in class %s not found! err: %s\n", name, arg, className, err)
	}

	return JMethodID(method)
}

//export JNI_CallStaticObjectMethod
func JNI_CallStaticObjectMethod(env CJNIEnv, c JClass, methodId JMethodID, args *C.struct_MethodArgType) JObject {
	className := gojni.GetJNI().GetClassName(gojni.JClass(c))
	argCount := int(JNI_GetStaticMethodIDArgs(env, c, methodId, nil))
	log.Debugf("JNI_CallStaticObjectMethod(%s, %d, %d)\n", className, methodId, argCount)
	goArgs := make([]gojni.TypedMethodArgs, argCount)

	if args != nil {
		argsSlice := (*[1 << 20]C.struct_MethodArgType)(unsafe.Pointer(args))[:argCount:argCount]
		for i := range goArgs {
			goArgs[i] = gojni.TypedMethodArgs{
				ArgType:     gojni.ArgType(argsSlice[i]._type),
				PtrValue:    uintptr(argsSlice[i].ptrValue),
				DoubleValue: float64(argsSlice[i].doubleValue),
				IntValue:    uint64(argsSlice[i].intValue),
			}
		}
	}

	return JObject(gojni.GetJNI().CallStaticObjectMethod(gojni.JClass(c), gojni.JMethodID(methodId), goArgs))
}

//export JNI_GetStaticFieldID
func JNI_GetStaticFieldID(env CJNIEnv, c JClass, Cname, Carg *C.char) JFieldID {
	name := C.GoString(Cname)
	arg := ""
	if Carg != nil {
		arg = C.GoString(Carg)
	}
	className := gojni.GetJNI().GetClassName(gojni.JClass(c))
	log.Debugf("GetStaticFieldID(%s, %s, %s)\n", className, name, arg)
	field, err := gojni.GetJNI().GetStaticFieldID(gojni.JClass(c), name, arg)
	if err != nil {
		log.Errorf("Static Field %s not found in class %s: %s\n", name, className, err)
	}

	return JFieldID(field)
}

//export JNI_GetMethodID
func JNI_GetMethodID(env CJNIEnv, c JClass, Cname, Carg *C.char) JMethodID {
	name := C.GoString(Cname)
	arg := ""
	if Carg != nil {
		arg = C.GoString(Carg)
	}
	className := gojni.GetJNI().GetClassName(gojni.JClass(c))
	log.Debugf("JNI_GetMethodID(%s, %s, %s)\n", className, name, arg)
	field, err := gojni.GetJNI().GetMethodID(gojni.JClass(c), name, arg)
	if err != nil {
		log.Errorf("Method %s not found in class %s: %s\n", name, className, err)
	}

	return JMethodID(field)
}

//export JNI_GetStaticObjectField
func JNI_GetStaticObjectField(env CJNIEnv, c JClass, f JFieldID) JObject {
	cName := gojni.GetJNI().GetClassName(gojni.JClass(c))
	fName := gojni.GetJNI().GetClassStaticFieldName(gojni.JClass(c), gojni.JFieldID(f))
	log.Debugf("JNI_GetStaticObjectField(%s, %s)\n", cName, fName)
	obj, err := gojni.GetJNI().GetStaticObjectField(gojni.JClass(c), gojni.JFieldID(f))
	if err != nil {
		log.Errorf("Field %s in class %s not found: %s\n", fName, cName, err)
	}
	return JObject(obj)
}

//export JNI_CallObjectMethod
func JNI_CallObjectMethod(env CJNIEnv, obj JObject, methodId JMethodID, args *C.struct_MethodArgType) JObject {
	argCount := int(JNI_GetObjectMethodIDArgs(env, obj, methodId, nil))
	goArgs := make([]gojni.TypedMethodArgs, argCount+1)

	goArgs[0] = gojni.TypedMethodArgs{ // Self
		ArgType:  gojni.ArgTypeJObject,
		IntValue: uint64(obj),
	}
	class, _ := gojni.GetJNI().GetObjectClass(gojni.JObject(obj))
	className := gojni.GetJNI().GetClassName(class)
	log.Debugf("JNI_CallObjectMethod(%s, %d, %d)\n", className, methodId, argCount)

	if args != nil {
		argsSlice := (*[1 << 20]C.struct_MethodArgType)(unsafe.Pointer(args))[:argCount:argCount]
		for i := range argsSlice {
			goArgs[i+1] = gojni.TypedMethodArgs{
				ArgType:     gojni.ArgType(argsSlice[i]._type),
				PtrValue:    uintptr(argsSlice[i].ptrValue),
				DoubleValue: float64(argsSlice[i].doubleValue),
				IntValue:    uint64(argsSlice[i].intValue),
			}
		}
	}

	return JObject(gojni.GetJNI().CallObjectMethod(gojni.JObject(obj), gojni.JMethodID(methodId), goArgs))
}

//export JNI_NewStringUTF
func JNI_NewStringUTF(env CJNIEnv, str *C.char) JString {
	data := C.GoString(str)
	log.Debugf("NewStringUTF(%s)\n", data)

	s, err := gojni.GetJNI().NewStringUTF(data)
	if err != nil {
		log.Errorf("Cannot create string %q: %s\n", data, err)
	}
	gojni.GetJNI().AddToLocalFrame(gojni.JObject(s))

	return JString(s)
}

//export JNI_GetStringUTFLength
func JNI_GetStringUTFLength(env CJNIEnv, str JString) JSize {
	log.Debugf("GetStringUTFLength(%d)\n", str)
	size, err := gojni.GetJNI().GetStringUTFLength(gojni.JString(str))
	if err != nil {
		log.Errorf("Cannot measure string %d length: %s\n", str, err)
	}

	return JSize(size)
}

//export JNI_GetStringUTFRegion
func JNI_GetStringUTFRegion(env CJNIEnv, str JString, start, l JSize, buf *C.char) {
	log.Debugf("GetStringUTFRegion(%d, %d, %d, %p)\n", str, start, l, buf)
	region, err := gojni.GetJNI().GetStringUTFRegion(gojni.JString(str), int(start), int(l))
	if err != nil {
		log.Errorf("error getting string %d: %s\n", str, err)
	}

	castedBuff := (*[1 << 20]byte)(unsafe.Pointer(buf))[:l:l]
	copy(castedBuff, region)
}

//export JNI_DeleteLocalRef
func JNI_DeleteLocalRef(env CJNIEnv, obj JObject) {
	log.Debugf("DeleteLocalRef(%d)\n", obj)
	gojni.GetJNI().DelFromLocalFrame(gojni.JObject(obj))
	gojni.GetJNI().DeallocObject(gojni.JObject(obj))
}

//export JNI_PushLocalFrame
func JNI_PushLocalFrame(env CJNIEnv, capacity JInt) JInt {
	log.Debugf("JNI_PushLocalFrame(%d)\n", capacity)
	gojni.GetJNI().PushLocalFrame(int(capacity))
	return 0
}

//export JNI_PopLocalFrame
func JNI_PopLocalFrame(env CJNIEnv, result JObject) JObject {
	log.Debugf("JNI_PopLocalFrame(%d)\n", result)
	gojni.GetJNI().PopLocalFrame(gojni.JObject(result))
	return result
}

//export JNI_GetObjectClass
func JNI_GetObjectClass(env CJNIEnv, obj JObject) JClass {
	log.Debugf("JNI_GetObjectClass(%d)\n", obj)
	class, err := gojni.GetJNI().GetObjectClass(gojni.JObject(obj))
	if err != nil {
		log.Errorf("cannot get object class for %d: %s\n", obj, err)
	}

	return JClass(class)
}

//export JNI_ExceptionOccurred
func JNI_ExceptionOccurred(env CJNIEnv) {
	log.Debug("JNI_ExceptionOccurred()\n")
	gojni.GetJNI().SetExceptionOccured()
}

//export JNI_ExceptionClear
func JNI_ExceptionClear() {
	log.Debug("JNI_ExceptionClear()\n")
	gojni.GetJNI().UnsetExceptionOccured()
}

//export JNI_NewGlobalRef
func JNI_NewGlobalRef(env CJNIEnv, obj JObject) JObject {
	log.Debugf("JNI_NewGlobalRef(%d)\n", obj)
	return JObject(gojni.GetJNI().NewGlobalReference(gojni.JObject(obj)))
}

//export JNI_DeleteGlobalRef
func JNI_DeleteGlobalRef(env CJNIEnv, obj JObject) {
	log.Debugf("JNI_DeleteGlobalRef(%d)\n", obj)
	gojni.GetJNI().DelGlobalReference(gojni.JObject(obj))
}

//export JNI_GetStringUTFChars
func JNI_GetStringUTFChars(env CJNIEnv, str JString, isCopy *JBool) *C.char {
	log.Debugf("JNI_GetStringUTFChars(%d, %p)\n", str, isCopy)
	val, err := gojni.GetJNI().GetString(gojni.JString(str))
	if err != nil {
		return nil
	}
	return C.CString(val)
}

//export JNI_ReleaseStringUTFChars
func JNI_ReleaseStringUTFChars(env CJNIEnv, str JString, utf *C.char) {
	C.free(unsafe.Pointer(utf))
}

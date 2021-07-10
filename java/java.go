package java

import (
	"fmt"
	"github.com/racerxdl/FakeJNI/gojni"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func class_getClassLoader(jni *gojni.JNIState, args gojni.JNIMethodArgs) gojni.JObject {
	log.Infof("Class.getClassLoader(%d)\n", len(args))
	class, _ := jni.FindClass("java/lang/ClassLoader")
	return jni.AllocObject(gojni.ArgTypeJObject, class, "CLASS LOADER")
}

func classLoader_findClass(jni *gojni.JNIState, args gojni.JNIMethodArgs) gojni.JObject {
	arg0 := args[1]
	if arg0.ArgType != gojni.ArgTypeJString {
		panic(fmt.Errorf("expected string on classLoader_findClass, got %d", arg0.ArgType))
	}
	_, arg0valI, _ := jni.GetObject(gojni.JObject(arg0.IntValue))
	arg0Val := arg0valI.(string)

	class, err := jni.FindClass(arg0Val)
	if err != nil {
		panic(fmt.Errorf("class %s not found", arg0Val))
	}

	return gojni.JObject(class)
}

// InitBasicJava intializes some basic java types in the specified JNI State
func InitBasicJava(jni *gojni.JNIState) {
	// Initialize default java
	jni.AddClass("fullyNullClass")
	jni.AddClass("java/lang/Class").
		AddMethod("getClassLoader", "()Ljava/lang/ClassLoader;", class_getClassLoader)
	jni.AddClass("java/lang/ClassLoader").
		AddMethod("findClass", "(Ljava/lang/String;)Ljava/lang/Class;", classLoader_findClass)
}

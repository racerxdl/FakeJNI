#pragma once
#include <stdlib.h>
#include "types.h"
#include "jni_map.h"
#include "jvm_map.h"

extern int  JNI_GetObjectMethodIDArgs(JNIEnv* env, jclass classId, jmethodID methodId, MethodArgType *args);
extern int  JNI_GetStaticMethodIDArgs(JNIEnv* env, jclass classId, jmethodID methodId, MethodArgType *args);
extern jobject JNI_CallStaticObjectMethod(JNIEnv* env, jclass classId, jmethodID methodId, MethodArgType *args);
extern jobject JNI_CallObjectMethod(JNIEnv* env, jclass classId, jmethodID methodId, MethodArgType *args);

void convertVaArgs(va_list args, int numArgs, MethodArgType *typedArgs) {
    // Extract variadic arguments
    // We don't actually need exact types since golang can cast it
    // It just need to be the correct size.

    for (int i = 0; i < numArgs; i++) {
        switch (typedArgs[i].type) {
            case ARG_TYPE_UINT32:   // uint32
                typedArgs[i].intValue = va_arg(args, uint32_t);
                break;
            case ARG_TYPE_UINT64:   // uint64
                typedArgs[i].intValue = va_arg(args, uint64_t);
                break;
            case ARG_TYPE_DOUBLE:   // float64
                typedArgs[i].doubleValue = va_arg(args, double);
                break;
            case ARG_TYPE_PTR:      // pointer to void
                typedArgs[i].ptrValue = va_arg(args, void *);
                break;
            case ARG_TYPE_JCLASS:   // JClass
            case ARG_TYPE_JOBJECT:  // JObject
            case ARG_TYPE_JSTRING:  // JString
                // All of these objects are just IDs. So we just read as int
                typedArgs[i].intValue = va_arg(args, int);
                break;
        }
    }
}

jobject JMP_CallStaticObjectMethod(JNIEnv* env, jclass classId, jmethodID methodId, ...) {
    MethodArgType *typedArgs = NULL;
    // Fetch number of arguments for the specified method
    int numArgs = JNI_GetStaticMethodIDArgs(env, classId, methodId, typedArgs);
    // if there is any args, fetch types
    if (numArgs > 0) {
        typedArgs = malloc(numArgs * sizeof(MethodArgType));
        JNI_GetStaticMethodIDArgs(env, classId, methodId, typedArgs);
    }

    va_list args;
    va_start(args, methodId);
    convertVaArgs(args, numArgs, typedArgs);
    va_end(args);

    jobject obj = JNI_CallStaticObjectMethod(env, classId, methodId, typedArgs);
    if (typedArgs != NULL) {
        free(typedArgs);
    }
    return obj;
}

jobject JMP_CallObjectMethodV(JNIEnv* env, jobject objId, jmethodID methodId, va_list args) {
    MethodArgType *typedArgs = NULL;
    // Fetch number of arguments for the specified method
    int numArgs = JNI_GetObjectMethodIDArgs(env, objId, methodId, typedArgs);
    // if there is any args, fetch types
    if (numArgs > 0) {
        typedArgs = malloc(numArgs * sizeof(MethodArgType));
        JNI_GetObjectMethodIDArgs(env, objId, methodId, typedArgs);
    }

    convertVaArgs(args, numArgs, typedArgs);

    jobject obj = JNI_CallObjectMethod(env, objId, methodId, typedArgs);
    if (typedArgs != NULL) {
        free(typedArgs);
    }
    return obj;
}

jobject JMP_CallObjectMethod(JNIEnv* env, jobject objId, jmethodID methodId, ...) {
    va_list args;
    va_start(args, methodId);
    jobject result = JMP_CallObjectMethodV(env, objId, methodId, args);
    va_end(args);
    return result;
}

const struct JNIInvokeInterface gInvokeInterface = {
    NULL,
    NULL,
    NULL,

    JVM_DestroyJavaVM,
    JVM_AttachCurrentThread,
    JVM_DetachCurrentThread,
    JVM_GetEnv,
    JVM_AttachCurrentThreadAsDaemon,
};

const JavaVM jvm = &gInvokeInterface;

const struct JNINativeInterface gNativeInterface = {
    NULL,
    NULL,
    NULL,
    NULL,
    JNI_GetVersion,
    JNI_DefineClass,
    JNI_FindClass,
    JNI_FromReflectedMethod,
    JNI_FromReflectedField,
    JNI_ToReflectedMethod,
    JNI_GetSuperclass,
    JNI_IsAssignableFrom,
    JNI_ToReflectedField,
    JNI_Throw,
    JNI_ThrowNew,
    JNI_ExceptionOccurred,
    JNI_ExceptionDescribe,
    JNI_ExceptionClear,
    JNI_FatalError,
    JNI_PushLocalFrame,
    JNI_PopLocalFrame,
    JNI_NewGlobalRef,
    JNI_DeleteGlobalRef,
    JNI_DeleteLocalRef,
    JNI_IsSameObject,
    JNI_NewLocalRef,
    JNI_EnsureLocalCapacity,
    JNI_AllocObject,
    JNI_NewObject,
    JNI_NewObjectV,
    JNI_NewObjectA,
    JNI_GetObjectClass,
    JNI_IsInstanceOf,
    JNI_GetMethodID,
    JMP_CallObjectMethod,
    JMP_CallObjectMethodV,
    JNI_CallObjectMethodA,
    JNI_CallBooleanMethod,
    JNI_CallBooleanMethodV,
    JNI_CallBooleanMethodA,
    JNI_CallByteMethod,
    JNI_CallByteMethodV,
    JNI_CallByteMethodA,
    JNI_CallCharMethod,
    JNI_CallCharMethodV,
    JNI_CallCharMethodA,
    JNI_CallShortMethod,
    JNI_CallShortMethodV,
    JNI_CallShortMethodA,
    JNI_CallIntMethod,
    JNI_CallIntMethodV,
    JNI_CallIntMethodA,
    JNI_CallLongMethod,
    JNI_CallLongMethodV,
    JNI_CallLongMethodA,
    JNI_CallFloatMethod,
    JNI_CallFloatMethodV,
    JNI_CallFloatMethodA,
    JNI_CallDoubleMethod,
    JNI_CallDoubleMethodV,
    JNI_CallDoubleMethodA,
    JNI_CallVoidMethod,
    JNI_CallVoidMethodV,
    JNI_CallVoidMethodA,
    JNI_CallNonvirtualObjectMethod,
    JNI_CallNonvirtualObjectMethodV,
    JNI_CallNonvirtualObjectMethodA,
    JNI_CallNonvirtualBooleanMethod,
    JNI_CallNonvirtualBooleanMethodV,
    JNI_CallNonvirtualBooleanMethodA,
    JNI_CallNonvirtualByteMethod,
    JNI_CallNonvirtualByteMethodV,
    JNI_CallNonvirtualByteMethodA,
    JNI_CallNonvirtualCharMethod,
    JNI_CallNonvirtualCharMethodV,
    JNI_CallNonvirtualCharMethodA,
    JNI_CallNonvirtualShortMethod,
    JNI_CallNonvirtualShortMethodV,
    JNI_CallNonvirtualShortMethodA,
    JNI_CallNonvirtualIntMethod,
    JNI_CallNonvirtualIntMethodV,
    JNI_CallNonvirtualIntMethodA,
    JNI_CallNonvirtualLongMethod,
    JNI_CallNonvirtualLongMethodV,
    JNI_CallNonvirtualLongMethodA,
    JNI_CallNonvirtualFloatMethod,
    JNI_CallNonvirtualFloatMethodV,
    JNI_CallNonvirtualFloatMethodA,
    JNI_CallNonvirtualDoubleMethod,
    JNI_CallNonvirtualDoubleMethodV,
    JNI_CallNonvirtualDoubleMethodA,
    JNI_CallNonvirtualVoidMethod,
    JNI_CallNonvirtualVoidMethodV,
    JNI_CallNonvirtualVoidMethodA,
    JNI_GetFieldID,
    JNI_GetObjectField,
    JNI_GetBooleanField,
    JNI_GetByteField,
    JNI_GetCharField,
    JNI_GetShortField,
    JNI_GetIntField,
    JNI_GetLongField,
    JNI_GetFloatField,
    JNI_GetDoubleField,
    JNI_SetObjectField,
    JNI_SetBooleanField,
    JNI_SetByteField,
    JNI_SetCharField,
    JNI_SetShortField,
    JNI_SetIntField,
    JNI_SetLongField,
    JNI_SetFloatField,
    JNI_SetDoubleField,
    JNI_GetStaticMethodID,
    JMP_CallStaticObjectMethod,
    JNI_CallStaticObjectMethodV,
    JNI_CallStaticObjectMethodA,
    JNI_CallStaticBooleanMethod,
    JNI_CallStaticBooleanMethodV,
    JNI_CallStaticBooleanMethodA,
    JNI_CallStaticByteMethod,
    JNI_CallStaticByteMethodV,
    JNI_CallStaticByteMethodA,
    JNI_CallStaticCharMethod,
    JNI_CallStaticCharMethodV,
    JNI_CallStaticCharMethodA,
    JNI_CallStaticShortMethod,
    JNI_CallStaticShortMethodV,
    JNI_CallStaticShortMethodA,
    JNI_CallStaticIntMethod,
    JNI_CallStaticIntMethodV,
    JNI_CallStaticIntMethodA,
    JNI_CallStaticLongMethod,
    JNI_CallStaticLongMethodV,
    JNI_CallStaticLongMethodA,
    JNI_CallStaticFloatMethod,
    JNI_CallStaticFloatMethodV,
    JNI_CallStaticFloatMethodA,
    JNI_CallStaticDoubleMethod,
    JNI_CallStaticDoubleMethodV,
    JNI_CallStaticDoubleMethodA,
    JNI_CallStaticVoidMethod,
    JNI_CallStaticVoidMethodV,
    JNI_CallStaticVoidMethodA,
    JNI_GetStaticFieldID,
    JNI_GetStaticObjectField,
    JNI_GetStaticBooleanField,
    JNI_GetStaticByteField,
    JNI_GetStaticCharField,
    JNI_GetStaticShortField,
    JNI_GetStaticIntField,
    JNI_GetStaticLongField,
    JNI_GetStaticFloatField,
    JNI_GetStaticDoubleField,
    JNI_SetStaticObjectField,
    JNI_SetStaticBooleanField,
    JNI_SetStaticByteField,
    JNI_SetStaticCharField,
    JNI_SetStaticShortField,
    JNI_SetStaticIntField,
    JNI_SetStaticLongField,
    JNI_SetStaticFloatField,
    JNI_SetStaticDoubleField,
    JNI_NewString,
    JNI_GetStringLength,
    JNI_GetStringChars,
    JNI_ReleaseStringChars,
    JNI_NewStringUTF,
    JNI_GetStringUTFLength,
    JNI_GetStringUTFChars,
    JNI_ReleaseStringUTFChars,
    JNI_GetArrayLength,
    JNI_NewObjectArray,
    JNI_GetObjectArrayElement,
    JNI_SetObjectArrayElement,
    JNI_NewBooleanArray,
    JNI_NewByteArray,
    JNI_NewCharArray,
    JNI_NewShortArray,
    JNI_NewIntArray,
    JNI_NewLongArray,
    JNI_NewFloatArray,
    JNI_NewDoubleArray,
    JNI_GetBooleanArrayElements,
    JNI_GetByteArrayElements,
    JNI_GetCharArrayElements,
    JNI_GetShortArrayElements,
    JNI_GetIntArrayElements,
    JNI_GetLongArrayElements,
    JNI_GetFloatArrayElements,
    JNI_GetDoubleArrayElements,
    JNI_ReleaseBooleanArrayElements,
    JNI_ReleaseByteArrayElements,
    JNI_ReleaseCharArrayElements,
    JNI_ReleaseShortArrayElements,
    JNI_ReleaseIntArrayElements,
    JNI_ReleaseLongArrayElements,
    JNI_ReleaseFloatArrayElements,
    JNI_ReleaseDoubleArrayElements,
    JNI_GetBooleanArrayRegion,
    JNI_GetByteArrayRegion,
    JNI_GetCharArrayRegion,
    JNI_GetShortArrayRegion,
    JNI_GetIntArrayRegion,
    JNI_GetLongArrayRegion,
    JNI_GetFloatArrayRegion,
    JNI_GetDoubleArrayRegion,
    JNI_SetBooleanArrayRegion,
    JNI_SetByteArrayRegion,
    JNI_SetCharArrayRegion,
    JNI_SetShortArrayRegion,
    JNI_SetIntArrayRegion,
    JNI_SetLongArrayRegion,
    JNI_SetFloatArrayRegion,
    JNI_SetDoubleArrayRegion,
    JNI_RegisterNatives,
    JNI_UnregisterNatives,
    JNI_MonitorEnter,
    JNI_MonitorExit,
    JNI_GetJavaVM,
    JNI_GetStringRegion,
    JNI_GetStringUTFRegion,
    JNI_GetPrimitiveArrayCritical,
    JNI_ReleasePrimitiveArrayCritical,
    JNI_GetStringCritical,
    JNI_ReleaseStringCritical,
    JNI_NewWeakGlobalRef,
    JNI_DeleteWeakGlobalRef,
    JNI_ExceptionCheck,
    JNI_NewDirectByteBuffer,
    JNI_GetDirectBufferAddress,
    JNI_GetDirectBufferCapacity,
    JNI_GetObjectRefType,
};

const struct JNINativeInterface *nativeInterfacePtr = &gNativeInterface;
JNIEnv *env = (JNIEnv *)&nativeInterfacePtr;


jint JVM_GetEnv(JavaVM* vm, void** argEnv, jint version) {
    JVM_LogDebug("JVM_GetEnv()\n");
    *argEnv = env;
    return 0;
}

void *GetJNIEnv() {
    return (void *)env;
}

void *GetJVM() {
    return (void *) &jvm;
}
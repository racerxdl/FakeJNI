package android

import (
	"fmt"
	"github.com/racerxdl/FakeJNI/gojni"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func identificatorGetContext(jni *gojni.JNIState, args gojni.JNIMethodArgs) gojni.JObject {
	log.Debugf("IdentificatorImpl::getContext()")
	class, _ := jni.FindClass("android/content/Context")
	return jni.AllocObject(gojni.ArgTypeJObject, class, "CONTEXT")
}

func androidContentContextGetTELEPHONY_SERVICE(jni *gojni.JNIState, args gojni.JNIMethodArgs) gojni.JObject {
	log.Debugf("android/content/Context::TELEPHONY_SERVICE")
	return jni.AllocObject(gojni.ArgTypeJObject, -100, "TELEPHONY_SERVICE")
}
func androidContentContextGetWIFI_SERVICE(jni *gojni.JNIState, args gojni.JNIMethodArgs) gojni.JObject {
	log.Debugf("android/content/Context::WIFI_SERVICE")
	return jni.AllocObject(gojni.ArgTypeJObject, -100, "WIFI_SERVICE")
}

func androidContentContextGetSystemService(jni *gojni.JNIState, args gojni.JNIMethodArgs) gojni.JObject {
	arg0 := args[1]
	if arg0.ArgType != gojni.ArgTypeJString {
		panic(fmt.Errorf("expected string on getSystemService, got %d", arg0.ArgType))
	}
	_, arg0valI, _ := jni.GetObject(gojni.JObject(arg0.IntValue))
	arg0Val := arg0valI.(string)

	log.Debugf("android/content/Context.getSystemService(%q)\n", arg0Val)
	switch arg0Val {
	case "TELEPHONY_SERVICE":
		class, _ := jni.FindClass("android/telephony/TelephonyManager")
		return jni.AllocObject(gojni.ArgTypeJObject, class, "TELEPHONY_SERVICE")
	case "WIFI_SERVICE":
		class, _ := jni.FindClass("android/net/wifi/WifiManager")
		return jni.AllocObject(gojni.ArgTypeJObject, class, "WIFI_MANAGER")
	default:
		log.Errorf("Unexpected service: %s\n", arg0Val)
		return gojni.InvalidObject
	}
}
func androidContentContextGetContentResolver(jni *gojni.JNIState, args gojni.JNIMethodArgs) gojni.JObject {
	log.Debugf("android/content/Context.GetContentResolver()")
	class, _ := jni.FindClass("android/content/ContentResolver")
	return jni.AllocObject(gojni.ArgTypeJObject, class, "CONTENT_RESOLVER")
}

func androidtelephonyTelephonyManagerGetSimSerialNumber(jni *gojni.JNIState, args gojni.JNIMethodArgs) gojni.JObject {
	log.Debugf("android/telephony/TelephonyManager.getSimSerialNumber()")
	str, _ := jni.NewStringUTF(androidData.simSerialNumber)
	return gojni.JObject(str)
}

func androidproviderSettingsSecureGetString(jni *gojni.JNIState, args gojni.JNIMethodArgs) gojni.JObject {
	// Expected Landroid/content/ContentResolver;Ljava/lang/String; and return Ljava/lang/String;
	arg0 := args[1]
	if arg0.ArgType != gojni.ArgTypeJString {
		panic(fmt.Errorf("expected string on SettingsSecureGetString, got %d", arg0.ArgType))
	}
	_, arg0valI, _ := jni.GetObject(gojni.JObject(arg0.IntValue))
	arg0Val := arg0valI.(string)
	log.Debugf("android/provider/Settings$Secure::getString(%d, %q)\n", args[0].IntValue, arg0Val)
	val := androidData.secureStorage[arg0Val]
	str, _ := jni.NewStringUTF(val)
	return gojni.JObject(str)
}

func getMacAddress(jni *gojni.JNIState, args gojni.JNIMethodArgs) gojni.JObject {
	log.Debugf("android/net/wifi/WiFiInfo.getMacAddress()")
	str, _ := jni.NewStringUTF(androidData.macaddress)
	return gojni.JObject(str)
}

func getConnectionInfo(jni *gojni.JNIState, args gojni.JNIMethodArgs) gojni.JObject {
	log.Debugf("android/net/wifi/WifiManager.getConnectionInfo()")
	class, _ := jni.FindClass("android/net/wifi/WifiInfo")
	return jni.AllocObject(gojni.ArgTypeJObject, class, "WIFI_INFO")
}

func getDeviceId(jni *gojni.JNIState, args gojni.JNIMethodArgs) gojni.JObject {
	log.Debugf("android/telephony/TelephonyManager.getDeviceId()")
	str, _ := jni.NewStringUTF(androidData.deviceId)
	return gojni.JObject(str)
}

func getBuildHost(jni *gojni.JNIState, args gojni.JNIMethodArgs) gojni.JObject {
	log.Debugf("android/os/Build::HOST")
	str, _ := jni.NewStringUTF(androidData.hostname)
	return gojni.JObject(str)
}

// InitAndroid initializes basic android functionality for JNI
func InitAndroid(jni *gojni.JNIState) {
	jni.AddClass("android/content/ContentResolver")
	jni.AddClass("br/com/dnofd/heartbeat/wrapper/IdentificatorImpl").
		AddStaticMethod("getContext", "()Landroid/content/Context;", identificatorGetContext)

	jni.AddClass("android/content/Context").
		AddStaticFieldGetter("TELEPHONY_SERVICE", androidContentContextGetTELEPHONY_SERVICE).
		AddStaticFieldGetter("WIFI_SERVICE", androidContentContextGetWIFI_SERVICE).
		AddMethod("getSystemService", "(Ljava/lang/String;)Ljava/lang/Object;", androidContentContextGetSystemService).
		AddMethod("getContentResolver", "()Landroid/content/ContentResolver;", androidContentContextGetContentResolver)

	jni.AddClass("android/telephony/TelephonyManager").
		AddMethod("getSimSerialNumber", "()Ljava/lang/String;", androidtelephonyTelephonyManagerGetSimSerialNumber).
		AddMethod("getDeviceId", "()Ljava/lang/String;", getDeviceId)

	jni.AddClass("android/provider/Settings$Secure").
		AddStaticMethod("getString", "(Landroid/content/ContentResolver;Ljava/lang/String;)Ljava/lang/String;", androidproviderSettingsSecureGetString)

	jni.AddClass("android/net/wifi/WifiInfo").
		AddMethod("getMacAddress", "()Ljava/lang/String;", getMacAddress)

	jni.AddClass("android/net/wifi/WifiManager").
		AddMethod("getConnectionInfo", "()Landroid/net/wifi/WifiInfo;", getConnectionInfo)

	jni.AddClass("android/os/Build").
		AddStaticFieldGetter("HOST", getBuildHost)
}

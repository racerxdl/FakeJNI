package gojni

import "regexp"

var methodArgsRegex = regexp.MustCompile(`\((.*?)\)(.*)`)
var javaTypeRegex = regexp.MustCompile(`L(.*?);`)

// JNIMethodArgs is the argument values for a JNI Call
type JNIMethodArgs []TypedMethodArgs

// JNICallHandle is a handler for a JNI Call
type JNICallHandle func(jni *JNIState, args JNIMethodArgs) JObject

// TypedMethodArgs represent a single argument from a called JNI Method / Function
type TypedMethodArgs struct {
	javaArgName string
	// ArgType is the type of the current argument
	ArgType ArgType
	// PtrValue is the content of the arg when ArgType is ArgTypePtr
	PtrValue uintptr
	// IntValue is the content of the arg when ArgType is ArgTypeU32, ArgTypeU64, ArgTypeJClass, ArgTypeJObject, ArgTypeJString
	IntValue uint64
	// DoubleValue is the content of the arg when ArgType is ArgTypeDouble
	DoubleValue float64
}

func createTypedMethodArgsFromString(argsString string) (args JNIMethodArgs) {
	parsedArgs := methodArgsRegex.FindAllStringSubmatch(argsString, -1)
	if len(parsedArgs) == 1 && len(parsedArgs[0]) == 3 {
		serializedArgs := parsedArgs[0][1]
		javaArgs := javaTypeRegex.FindAllString(serializedArgs, -1)
		for _, pa := range javaArgs {
			mappedType, ok := javaTypeMap[pa]
			if !ok {
				log.Warnf("unmapped type %s. default to java object\n", pa)
				mappedType = ArgTypeJObject
			}
			args = append(args, TypedMethodArgs{
				javaArgName: pa,
				ArgType:     mappedType,
			})
		}
	}

	return args
}

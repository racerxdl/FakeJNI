package gojni

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

const (
	ArgTypeU32     ArgType = 0
	ArgTypeU64     ArgType = 1
	ArgTypeDouble  ArgType = 2
	ArgTypePtr     ArgType = 3
	ArgTypeJClass  ArgType = 4
	ArgTypeJObject ArgType = 5
	ArgTypeJString ArgType = 6
)

const (
	InvalidClass  = 0
	InvalidObject = 0
	InvalidField  = 0
	InvalidMethod = 0
)

// JNI Class Objects needs to have a corresponding jclass id. So we offset all alocated objects
const nonClassObjectOffset = 1000
const globalObjectOffset = 4000

// javaTypeMap holds primitive type maps
var javaTypeMap = map[string]ArgType{
	"Ljava/lang/String;": ArgTypeJString,
	"Ljava/lang/Object;": ArgTypeJObject,
}

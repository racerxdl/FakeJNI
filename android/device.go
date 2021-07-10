package android

type androidHolder struct {
	deviceId        string
	hostname        string
	macaddress      string
	simSerialNumber string
	secureStorage   map[string]string
}

var androidData = &androidHolder{
	deviceId:        "ABCD",
	hostname:        "GoFakeJNI",
	macaddress:      "12:34:56:78:90:AB",
	simSerialNumber: "111111111111111111111",
	secureStorage: map[string]string{
		"android_id": "fakeid",
	},
}

// SetDeviceID sets the Fake Android Device ID
func SetDeviceID(id string) {
	androidData.deviceId = id
}

// SetHostname sets the Fake Android Hostname
func SetHostname(hostname string) {
	androidData.hostname = hostname
}

// SetMacAddress sets the Fake Android Device Mac Address in the format 12:34:56:78:90:AB
func SetMacAddress(macaddress string) {
	androidData.macaddress = macaddress
}

// SetSimSerialNumber sets the Fake Android Simcard Serial
func SetSimSerialNumber(serial string) {
	androidData.simSerialNumber = serial
}

// SetSecureStorageValue sets the a value inside Fake Android Secure Storage
func SetSecureStorageValue(key, value string) {
	androidData.secureStorage[key] = value
}

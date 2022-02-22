package leafMandatory

type Mandatory struct {
	traceID         string
	ipAddress       []string
	valid           bool
	bringDeviceType bool
	language        string

	authorization Authorization
	deviceType    DeviceType
	device        Device
	userAgent     UserAgent
	os            OS
	user          User
}

func (m Mandatory) TraceID() string {
	return m.traceID
}

func (m Mandatory) IpAddresses() []string {
	return m.ipAddress
}

func (m Mandatory) Language() string {
	if m.language == "" {
		return "en"
	}
	return m.language
}

func (m Mandatory) Authorization() Authorization {
	return m.authorization
}

func (m Mandatory) DeviceType() DeviceType {
	return m.deviceType
}

func (m Mandatory) Device() Device {
	return m.device
}

func (m Mandatory) UserAgent() UserAgent {
	return m.userAgent
}

func (m Mandatory) OS() OS {
	return m.os
}

func (m Mandatory) User() User {
	return m.user
}

/*
===========================
	Utilities Function
===========================
*/
func (m Mandatory) JSON() map[string]interface{} {
	return map[string]interface{}{
		"trace_id":         m.TraceID(),
		"ip_addresses":     m.IpAddresses(),
		"authorization":    m.Authorization().JSON(),
		"device_type":      m.DeviceType().JSON(),
		"device":           m.Device().JSON(),
		"user_agent":       m.UserAgent().JSON(),
		"operating_system": m.OS().JSON(),
		"user":             m.User().JSON(),
	}
}

func (m Mandatory) Valid() bool {
	return m.valid
}

func (m Mandatory) IsUserLogin() bool {
	return m.user.IsLogin()
}

func (m Mandatory) IsMobileApp() bool {
	return Android == m.deviceType || Ios == m.deviceType
}

func (m Mandatory) IsWebApp() bool {
	return Web == m.deviceType || MobileWeb == m.deviceType
}

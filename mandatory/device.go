package leafMandatory

type Device struct {
	appVersion string
	deviceID   string
	family     string
	brand      string
	model      string
}

func (d Device) AppVersion() string {
	return d.appVersion
}

func (d Device) DeviceID() string {
	return d.deviceID
}

func (d Device) Family() string {
	return d.family
}

func (d Device) Brand() string {
	return d.brand
}

func (d Device) Model() string {
	return d.model
}

func (d Device) JSON() map[string]interface{} {
	return map[string]interface{}{
		"app_version": d.AppVersion(),
		"device_id":   d.DeviceID(),
		"family":      d.Family(),
		"brand":       d.Brand(),
		"model":       d.Model(),
	}
}

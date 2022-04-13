package leafMandatory

import (
	"bytes"
	"github.com/ua-parser/uap-go/uaparser"
	"strings"
)

type MandatoryBuilder struct {
	Mandatory
	uaParser *uaparser.Parser
}

func NewMandatoryBuilder() (MandatoryBuilder, error) {
	parser, err := uaparser.NewFromBytes(bytes.NewBufferString(regexes).Bytes())
	if err != nil {
		return MandatoryBuilder{}, err
	}

	return MandatoryBuilder{
		Mandatory: Mandatory{},
		uaParser:  parser,
	}, nil
}

func (m MandatoryBuilder) Build() Mandatory {
	m.valid = true
	return m.Mandatory
}

func (m MandatoryBuilder) WithTraceID(traceID string) MandatoryBuilder {
	m.traceID = traceID
	return m
}

func (m MandatoryBuilder) WithIpAddresses(ipAddress []string) MandatoryBuilder {
	m.ipAddress = ipAddress
	return m
}

func (m MandatoryBuilder) WithAuthorization(authorization string) MandatoryBuilder {
	m.authorization.authorization = authorization
	m.authorization.token = strings.ReplaceAll(authorization, "Bearer ", "")
	return m
}

func (m MandatoryBuilder) WithApiKey(apiKey string) MandatoryBuilder {
	m.authorization.apiKey = apiKey
	return m
}

func (m MandatoryBuilder) WithServiceSecret(ID, secret string) MandatoryBuilder {
	m.authorization.serviceID = ID
	m.authorization.serviceSecret = secret
	return m
}

func (m MandatoryBuilder) WithUserAgent(userAgent string) MandatoryBuilder {
	client := m.uaParser.Parse(userAgent)
	m.userAgent.value = userAgent
	m.userAgent.family = client.UserAgent.Family
	m.userAgent.major = client.UserAgent.Major
	m.userAgent.minor = client.UserAgent.Minor
	m.userAgent.patch = client.UserAgent.Patch
	m.os.family = client.Os.Family
	m.os.major = client.Os.Major
	m.os.minor = client.Os.Minor
	m.os.patch = client.Os.Patch
	m.os.patchMinor = client.Os.PatchMinor
	m.device.family = client.Device.Family
	m.device.brand = client.Device.Brand
	m.device.model = client.Device.Model

	if m.bringDeviceType {
		return m
	}

	switch strings.ToLower(m.os.family) {
	case "android":
		m.deviceType = Android
		break
	case "ios":
		m.deviceType = Ios
		break
	}

	if m.device.deviceID == "" {
		if strings.Contains(strings.ToLower(m.userAgent.family), "mobile") {
			m.deviceType = MobileWeb
		} else {
			m.deviceType = Web
		}
	}
	return m
}

func (m MandatoryBuilder) WithApplication(deviceID, appsVersion string) MandatoryBuilder {
	m.device.deviceID = deviceID
	m.device.appVersion = appsVersion
	if m.userAgent.value != "" {
		return m.WithUserAgent(m.userAgent.value)
	}
	return m
}

func (m MandatoryBuilder) WithDeviceType(deviceType string) MandatoryBuilder {
	if "" == deviceType {
		return m
	}

	var err error
	m.deviceType, err = DeviceFromStringCode(deviceType)
	if err != nil {
		return m
	}

	m.bringDeviceType = true
	return m
}

func (m MandatoryBuilder) WithUser(ID uint64, email string) MandatoryBuilder {
	m.user.login = true
	m.user.id = ID
	m.user.email = email
	return m
}

func (m MandatoryBuilder) WithUserPhone(ID uint64, email string, phone string) MandatoryBuilder {
	m.user.login = true
	m.user.id = ID
	m.user.email = email
	m.user.phone = phone
	return m
}

func (m MandatoryBuilder) WithPhone(ID uint64, phone string) MandatoryBuilder {
	m.user.login = true
	m.user.id = ID
	m.user.phone = phone
	return m
}

func (m MandatoryBuilder) WithLanguage(language string) MandatoryBuilder {
	m.language = language
	return m
}

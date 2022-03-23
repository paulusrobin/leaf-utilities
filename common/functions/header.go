package leafFunctions

import (
	"context"
	leafHeader "github.com/paulusrobin/leaf-utilities/common/header"
	leafMandatory "github.com/paulusrobin/leaf-utilities/mandatory"
	"net/http"
	"strings"
)

func AppendMandatoryHeader(ctx context.Context, headers http.Header) http.Header {
	mandatory := leafMandatory.FromContext(ctx)
	if !mandatory.Valid() {
		return headers
	}

	headers = AppendHeaderIfNotExist(headers, leafHeader.TraceID, mandatory.TraceID())
	headers = AppendHeaderIfNotExist(headers, leafHeader.Language, mandatory.TraceID())
	headers = AppendHeaderIfNotExist(headers, leafHeader.AppVersion, mandatory.Device().AppVersion())
	headers = AppendHeaderIfNotExist(headers, leafHeader.ServiceID, mandatory.Authorization().ServiceID())
	headers = AppendHeaderIfNotExist(headers, leafHeader.ServiceSecret, mandatory.Authorization().ServiceSecret())
	headers = AppendHeaderIfNotExist(headers, leafHeader.Authorization, mandatory.Authorization().Authorization())
	headers = AppendHeaderIfNotExist(headers, leafHeader.ApiKey, mandatory.Authorization().ApiKey())
	headers = AppendHeaderIfNotExist(headers, leafHeader.DeviceID, mandatory.Device().DeviceID())
	headers = AppendHeaderIfNotExist(headers, leafHeader.DeviceType, mandatory.DeviceType().Info().Code())
	headers = AppendHeaderIfNotExist(headers, leafHeader.UserAgent, mandatory.UserAgent().Value())
	headers = AppendHeaderIfNotExist(headers, leafHeader.IpAddress, strings.Join(mandatory.IpAddresses(), ","))

	headers = AppendHeaderIfNotExist(headers, leafHeader.XTraceID, mandatory.TraceID())
	headers = AppendHeaderIfNotExist(headers, leafHeader.XAppVersion, mandatory.Device().AppVersion())
	headers = AppendHeaderIfNotExist(headers, leafHeader.XServiceID, mandatory.Authorization().ServiceID())
	headers = AppendHeaderIfNotExist(headers, leafHeader.XServiceSecret, mandatory.Authorization().ServiceSecret())
	headers = AppendHeaderIfNotExist(headers, leafHeader.XApiKey, mandatory.Authorization().ApiKey())

	return headers
}

func AppendHeaderIfNotExist(headers http.Header, headerName, value string) http.Header {
	existing := headers.Get(headerName)
	if "" != existing {
		return headers
	}
	headers.Add(headerName, value)
	return headers
}

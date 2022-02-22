package leafMandatory

type UserAgent struct {
	value  string
	family string
	major  string
	minor  string
	patch  string
}

func (u UserAgent) Value() string {
	return u.value
}

func (u UserAgent) Family() string {
	return u.family
}

func (u UserAgent) Major() string {
	return u.major
}

func (u UserAgent) Minor() string {
	return u.minor
}

func (u UserAgent) Patch() string {
	return u.patch
}

func (u UserAgent) JSON() map[string]interface{} {
	return map[string]interface{}{
		"value":  u.Value(),
		"family": u.Family(),
		"major":  u.Major(),
		"minor":  u.Minor(),
		"patch":  u.Patch(),
	}
}

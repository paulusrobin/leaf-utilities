package leafMandatory

type OS struct {
	name    string
	version string

	family     string
	major      string
	minor      string
	patch      string
	patchMinor string
}

func (o OS) Name() string {
	return o.name
}

func (o OS) Version() string {
	return o.version
}

func (o OS) Family() string {
	return o.family
}

func (o OS) Major() string {
	return o.major
}

func (o OS) Minor() string {
	return o.minor
}

func (o OS) Patch() string {
	return o.patch
}

func (o OS) PatchMinor() string {
	return o.patchMinor
}

func (o OS) JSON() map[string]interface{} {
	return map[string]interface{}{
		"name":        o.Name(),
		"version":     o.Version(),
		"family":      o.Family(),
		"major":       o.Major(),
		"minor":       o.Minor(),
		"patch":       o.Patch(),
		"patch_minor": o.PatchMinor(),
	}
}

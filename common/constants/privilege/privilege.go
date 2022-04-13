package leafPrivilege

const (
	Granted = "GRANTED"
	Trusted = "TRUSTED"
	Public  = "PUBLIC"
)

func Exist(privilege string) bool {
	switch privilege {
	case Public, Trusted, Granted:
		return true
	default:
		return false
	}
}

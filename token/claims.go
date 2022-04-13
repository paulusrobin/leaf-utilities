package leafToken

type (
	UserLogin struct {
		ID    uint64 `json:"id"`
		Email string `json:"email"`
		Phone string `json:"phone"`
	}
	Claims interface {
		User() UserLogin
		Valid() error
	}
)

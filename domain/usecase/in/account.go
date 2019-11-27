package in

// NewAccount -
type NewAccount struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

// EditAccount -
type EditAccount struct {
	ID       string `json:"id" validate:"required,numeric"`
	Password string `validate:"required"`
}

// FetchAccount -
type FetchAccount struct {
	ID string `validate:"required,numeric"`
}

// FetchAccountByEmail -
type FetchAccountByEmail struct {
	Email string `validate:"required,email"`
}

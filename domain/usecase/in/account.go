package in

// NewAccount -
type NewAccount struct {
	Email    string `validate:"required"`
	Name     string
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

// FetchAccountByName -
type FetchAccountByName struct {
	Name string `validate:"required,alphanum"`
}

// LoginAccountByEmail -
type LoginAccountByEmail struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

// LoginAccountByName-
type LoginAccountByName struct {
	Name     string `validate:"required,alphanum"`
	Password string `validate:"required"`
}

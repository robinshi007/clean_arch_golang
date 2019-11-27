package in

// NewUser -
type NewUser struct {
	Name string `json:"name" validate:"required"`
}

// EditUser -
type EditUser struct {
	ID   string `json:"id" validate:"required,numeric"`
	Name string `json:"name" validate:"required"`
}

// FetchUser -
type FetchUser struct {
	ID string `json:"id" validate:"required,numeric"`
}

// FetchUserByName -
type FetchUserByName struct {
	Name string `json:"name" validate:"required"`
}

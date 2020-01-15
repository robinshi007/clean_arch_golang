package in

// NewRedirect -
type NewRedirect struct {
	URL string `json:"url" validate:"required"`
}

// FetchRedirects -
type FetchRedirects struct {
	Offset string `validate:"required,numeric"`
	Limit  string `validate:"required,numeric"`
}

// FetchRedirect -
type FetchRedirect struct {
	ID string `json:"id" validate:"required"`
}

// FetchRedirectByCode -
type FetchRedirectByCode struct {
	Code string `json:"code" validate:"required"`
}

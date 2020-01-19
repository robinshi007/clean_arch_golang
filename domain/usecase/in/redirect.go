package in

// NewRedirect -
type NewRedirect struct {
	URL string `json:"url" validate:"required,url"`
}

// FetchRedirect -
type FetchRedirect struct {
	ID string `json:"id" validate:"required"`
}

// FetchRedirectByCode -
type FetchRedirectByCode struct {
	Code string `json:"code" validate:"required"`
}

// FetchRedirectByURL -
type FetchRedirectByURL struct {
	URL string `json:"url" validate:"required"`
}

// FetchOrCreateRedirect -
type FetchOrCreateRedirect struct {
	URL string `json:"url" validate:"required"`
}

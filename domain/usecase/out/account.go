package out

// Account -
type Account struct {
	ID    int64  `json:"id" xml:"id"`
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

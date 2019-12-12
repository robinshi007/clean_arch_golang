package out

// Account -
type Account struct {
	ID    int64  `json:"id" msgpack:"id"`
	Email string `json:"email" msgpack:"email"`
}

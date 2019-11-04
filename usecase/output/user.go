package output

// UserID -
type UserID int64

// User -
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

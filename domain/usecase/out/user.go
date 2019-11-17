package out

// UserID -
type UserID int64

// User -
type User struct {
	ID   int64  `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

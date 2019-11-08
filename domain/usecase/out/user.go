package out

import (
	"encoding/xml"
)

// UserID -
type UserID int64

// User -
type User struct {
	XMLName xml.Name `json:"-" xml:"user"`
	ID      int64    `json:"id" xml:"id"`
	Name    string   `json:"name" xml:"name"`
}

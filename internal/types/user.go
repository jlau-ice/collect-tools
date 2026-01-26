package types

import "time"

type User struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Sex       int       `json:"sex"`
	Birth     time.Time `json:"birth"`
	Account   string    `json:"account"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

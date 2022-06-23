package user

import "time"

type User struct {
	ID             int
	Fullname       string
	Occupation     string
	Email          string
	Password       string
	ProfilePicture string
	Role           string
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
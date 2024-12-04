package modals

import "gorm.io/gorm"

type User struct {
	gorm.Model
	CommonModal
	UserDto
}

type UserDto struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       uint16 `json:"age"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

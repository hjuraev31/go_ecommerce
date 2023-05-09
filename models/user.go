package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json: "id"`
	FirstName string `json: "first_name"`
	LastName  string `json: "last_name"`
	CartId    string `json: "cart_id"`
	Email     string `json: "email"`
	Password  []byte `json: "-"`
	Cart      Cart   `json: "cart";gorm:"foreignKey:CartId`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

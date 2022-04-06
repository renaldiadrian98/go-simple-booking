package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	RoleId    int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserWithRelation struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	RoleId    int       `json:"role_id"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CheckExistingEmail(email string) (User, bool) {
	var user User
	err := DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, false
	}
	return user, true
}

func GetUser(userId string) User {
	var user User
	DB.Where("id = ?", userId).Find(&user)
	return user
}

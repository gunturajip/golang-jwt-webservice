package models

import (
	"golang-jwt-auth/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your full name is required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string    `gorm:"not null" json:"-" form:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

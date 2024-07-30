package repositoryimpl

import (
	"gorm.io/gorm"

	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/user/domain"
)

var tableNameUser string

type UserDO struct {
	gorm.Model

	Username string `gorm:"column:username;uniqueIndex:username_index;size:255"`
	Password string `gorm:"column:password;index:password_index;size:255"`
	Email    string `gorm:"column:email;size:255"`
}

func (do *UserDO) toUser() (user domain.User, err error) {

	if user.Username, err = cmprimitive.NewUsername(do.Username); err != nil {
		return
	}

	if user.Email, err = domain.NewEmail(do.Email); err != nil {
		return
	}

	return
}

func (do *UserDO) TableName() string {
	return tableNameUser
}

package repositoryimpl

import (
	"gorm.io/gorm"

	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/user-server/domain"
)

const tableNameUser = "user"

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

	if user.Email, err = cmprimitive.NewEmail(do.Email); err != nil {
		return
	}

	return
}

func (do *UserDO) TableName() string {
	return tableNameUser
}

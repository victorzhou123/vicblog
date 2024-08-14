package controller

import (
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/user/app"
	"github.com/victorzhou123/vicblog/user/domain"
)

type reqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *reqLogin) toCmd() (cmd app.AccountCmd, err error) {

	if cmd.Username, err = cmprimitive.NewUsername(req.Username); err != nil {
		return
	}

	if cmd.Password, err = domain.NewPassword(req.Password); err != nil {
		return
	}

	return
}

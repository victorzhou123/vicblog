package controller

import (
	"context"

	"github.com/victorzhou123/vicblog/common/controller/rpc"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/user-server/app"
	"github.com/victorzhou123/vicblog/user-server/domain"
)

var successResp = &rpc.RespLogin{
	Info: &rpc.ResponseInfo{
		Code: "",
		Msg:  "success",
	},
}
var failedResp = &rpc.RespLogin{
	Info: &rpc.ResponseInfo{
		Code: "error",
		Msg:  "error",
	},
}

type server struct {
	rpc.UnimplementedLoginServiceServer

	login app.LoginService
}

func NewLoginRpcServer(login app.LoginService) rpc.LoginServiceServer {
	return &server{
		login: login,
	}
}

func (s *server) Login(ctx context.Context, req *rpc.ReqLogin) (*rpc.RespLogin, error) {

	username, err := cmprimitive.NewUsername(req.GetUsername())
	if err != nil {
		return failedResp, err
	}

	password, err := domain.NewPassword(req.GetPassword())
	if err != nil {
		return failedResp, err
	}

	dto, err := s.login.Login(&app.AccountCmd{
		Username: username,
		Password: password,
	})
	if err != nil {
		return failedResp, err
	}

	successResp.Data = &rpc.UserAndToken{
		Username: dto.Username,
		Email:    dto.Email,
		Token:    dto.Token,
	}

	return successResp, nil
}

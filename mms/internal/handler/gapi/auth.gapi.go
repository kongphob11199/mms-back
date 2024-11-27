package gapi

import (
	"context"

	"mms/internal/dto"
	pb "mms/internal/pb"
	"mms/internal/service"
	"mms/internal/utils"
)

type AuthHandlerGrpc struct {
	pb.UnimplementedAuthServiceServer
	auth service.AuthService
}

func NewAuthHandlerGrpcHandler(auth service.AuthService) *AuthHandlerGrpc {
	authServer := AuthHandlerGrpc{
		auth: auth,
	}
	return &authServer
}

func (a *AuthHandlerGrpc) Login(ctx context.Context,req *pb.LoginRequest) (*pb.LoginResponse, error) {
	newReq := &dto.AuthLoginReq{
		Username: req.Username,
		Password: req.Password,
	}
	res, err := a.auth.Login(newReq)

	if err != nil {
		return nil, err
	}

	newRes := &pb.LoginResponse{
		Response: utils.ConvertToPbResponse(&dto.StatusResp{Response: res.Response}),
		Token:    res.Token,
	}

	return newRes, nil
}

func (a *AuthHandlerGrpc) CheckAuth(ctx context.Context, req *pb.EmptyAuth) (*pb.StatusResponse, error) {
	res, err := a.auth.CheckAuth()
	if err != nil {
		return nil, err
	}

	newRes := &pb.StatusResponse{
		Response: utils.ConvertToPbResponse(res),
	}

	return newRes, nil
}

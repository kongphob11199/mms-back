package gapi

import (
	"context"
	"log"

	"mms/internal/dto"
	"mms/internal/middleware"
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

func (a *AuthHandlerGrpc) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
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

func (a *AuthHandlerGrpc) CheckAuth(ctx context.Context, req *pb.EmptyAuth) (*pb.AuthResponse, error) {
	claimsFromCtx, ok := ctx.Value("claims").(*middleware.ClaimsContextKey)
	if !ok {
		log.Panic("claims not found or invalid type")
	}
	newReq := &dto.AuthUpdateTokenReq{
		Token:  claimsFromCtx.Token,
		UserId: claimsFromCtx.UserId,
	}
	userRes, err := a.auth.CheckAuth(newReq)
	if err != nil {
		return nil, err
	}
	log.Println("userRes : ", userRes)
	pbUser := &pb.User{
		UserId:     int32(userRes.UserId),
		Firstname:  userRes.Firstname,
		Lastname:   userRes.Lastname,
		Birthday:   utils.TimeToTimestamp(userRes.Birthday),
		Gender:     utils.ConvertToPbGender(userRes.Gender),
		Role:       utils.ConvertToPbRole(userRes.Role),
		CreateAt:   utils.TimeToTimestamp(userRes.CreateAt),
		CreateBy:   userRes.CreateBy,
		UpdateAt:   utils.TimeToTimestamp(userRes.UpdateAt),
		UpdateBy:   userRes.UpdateBy,
		StatusUser: utils.ConvertToPbStatusUser(userRes.StatusUser),
		Username:   userRes.Username,
	}

	newRes := &pb.AuthResponse{User: pbUser}

	return newRes, nil
}

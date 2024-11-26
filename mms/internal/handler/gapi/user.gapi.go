package gapi

import (
	"context"
	"fmt"
	"log"
	"mms/internal/dto"
	"mms/internal/models"
	pb "mms/internal/pb"
	"mms/internal/service"
	"mms/internal/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type UserHandlerGrpc struct {
	pb.UnimplementedUserServiceServer
	user service.UserService
}

func NewUserHandlerGrpcHandler(user service.UserService) *UserHandlerGrpc {
	userServer := UserHandlerGrpc{
		user: user,
	}

	return &userServer
}

func (u *UserHandlerGrpc) FindAll(ctx context.Context, req *pb.Empty) (*pb.UsersResponse, error) {
	log.Println("user gapi : FindAll()")
	userRes, total, err := u.user.FindAll()

	if err != nil {
		return nil, err
	}

	var pbUsers []*pb.User

	for _, user := range *userRes {
		pbUser := &pb.User{
			UserId:     int32(user.UserId),
			Firstname:  user.Firstname,
			Lastname:   user.Lastname,
			Birthday:   utils.TimeToTimestamp(user.Birthday),
			Gender:     utils.ConvertToPbGender(user.Gender),
			Role:       utils.ConvertToPbRole(user.Role),
			CreateAt:   utils.TimeToTimestamp(user.CreateAt),
			CreateBy:   user.CreateBy,
			UpdateAt:   utils.TimeToTimestamp(user.UpdateAt),
			UpdateBy:   user.UpdateBy,
			StatusUser: utils.ConvertToPbStatusUser(user.StatusUser),
			Username:   user.Username,
		}
		pbUsers = append(pbUsers, pbUser)
	}

	res := &pb.UsersResponse{Users: pbUsers}

	md := metadata.Pairs("x-total", fmt.Sprintf("%d", total))
	grpc.SendHeader(ctx, md)

	return res, nil
}

func (u *UserHandlerGrpc) FindPagination(ctx context.Context, req *pb.UserPaginationRequest) (*pb.UsersResponse, error) {
	newReq := &dto.UserPaginationReq{
		Page:       req.Page,
		PageLimit:  req.PageLimit,
		Search:     req.Search,
		Role:       utils.ConvertPbRolesToModelsRoles(req.Role),
		StatusUser: utils.ConvertPbStatusUsersToModelsStatusUsers(req.StatusUser),
	}
	userRes, total, err := u.user.FindPagination(newReq)

	if err != nil {
		return nil, err
	}

	var pbUsers []*pb.User

	for _, user := range *userRes {
		pbUser := &pb.User{
			UserId:     int32(user.UserId),
			Firstname:  user.Firstname,
			Lastname:   user.Lastname,
			Birthday:   utils.TimeToTimestamp(user.Birthday),
			Gender:     utils.ConvertToPbGender(user.Gender),
			Role:       utils.ConvertToPbRole(user.Role),
			CreateAt:   utils.TimeToTimestamp(user.CreateAt),
			CreateBy:   user.CreateBy,
			UpdateAt:   utils.TimeToTimestamp(user.UpdateAt),
			UpdateBy:   user.UpdateBy,
			StatusUser: utils.ConvertToPbStatusUser(user.StatusUser),
			Username:   user.Username,
		}
		pbUsers = append(pbUsers, pbUser)
	}

	res := &pb.UsersResponse{Users: pbUsers}

	md := metadata.Pairs("x-total", fmt.Sprintf("%d", total))
	grpc.SendHeader(ctx, md)

	return res, nil
}

func (u *UserHandlerGrpc) FindById(ctx context.Context, req *pb.UserFindIdRequest) (*pb.UserResponse, error) {
	newUserIdReq := req.UserId
	userRes, err := u.user.FindById(newUserIdReq)

	if err != nil {
		return nil, err
	}

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

	res := &pb.UserResponse{User: pbUser}

	return res, nil
}

func (u *UserHandlerGrpc) CreateCustomer(ctx context.Context, req *pb.CreateUserCustomerRequest) (*pb.StatusResponse, error) {
	log.Println("user gapi : CreateCustomer()", req)
	newReq := &dto.CreateUserCustomerReq{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Gender:    models.Gender(req.GetGender().String()),
		Birthday:  utils.TimestampToTime(req.Birthday),
		Username:  req.Username,
		Password:  req.Password,
	}

	status, err := u.user.CreateCustomer(newReq)

	if err != nil {
		return nil, err
	}

	res := &pb.StatusResponse{
		Response: utils.ConvertToPbResponse(status),
	}

	return res, nil
}

func (u *UserHandlerGrpc) UpdateCustomer(ctx context.Context, req *pb.UpdateUserCustomerRequest) (*pb.StatusResponse, error) {
	newUserIdReq := req.UserId
	newReq := &dto.UpdateUserCustomerReq{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Gender:    models.Gender(req.GetGender().String()),
		Birthday:  utils.TimestampToTime(req.Birthday),
		Username:  req.Username,
	}

	status, err := u.user.UpdateCustomer(newUserIdReq, newReq)

	if err != nil {
		return nil, err
	}

	res := &pb.StatusResponse{
		Response: utils.ConvertToPbResponse(status),
	}

	return res, nil
}

func (u *UserHandlerGrpc) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.StatusResponse, error) {

	newReq := &dto.CreateUserReq{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Gender:    models.Gender(req.GetGender().String()),
		Role:      models.Role(req.GetRole().String()),
		Birthday:  utils.TimestampToTime(req.Birthday),
		Username:  req.Username,
		Password:  req.Password,
	}

	status, err := u.user.Create(newReq)

	if err != nil {
		return nil, err
	}

	res := &pb.StatusResponse{
		Response: utils.ConvertToPbResponse(status),
	}

	return res, nil
}

func (u *UserHandlerGrpc) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.StatusResponse, error) {
	newUserIdReq := req.UserId
	newReq := &dto.UpdateUserReq{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Gender:    models.Gender(req.GetGender().String()),
		Role:      models.Role(req.GetRole().String()),
		Birthday:  utils.TimestampToTime(req.Birthday),
		Username:  req.Username,
	}

	status, err := u.user.Update(newUserIdReq, newReq)

	if err != nil {
		return nil, err
	}

	res := &pb.StatusResponse{
		Response: utils.ConvertToPbResponse(status),
	}

	return res, nil
}

func (u *UserHandlerGrpc) Delete(ctx context.Context, req *pb.UserFindIdRequest) (*pb.StatusResponse, error) {
	newUserIdReq := req.UserId

	status, err := u.user.Delete(newUserIdReq)

	if err != nil {
		return nil, err
	}

	res := &pb.StatusResponse{
		Response: utils.ConvertToPbResponse(status),
	}

	return res, nil
}

func (u *UserHandlerGrpc) UpdateStatus(ctx context.Context, req *pb.UserStatusRequest) (*pb.StatusResponse, error) {
	newReq := &dto.UserStatusReq{
		UserId:     req.UserId,
		StatusUser: models.StatusUser(req.GetStatusUser().String()),
	}

	status, err := u.user.UpdateStatus(newReq)

	if err != nil {
		return nil, err
	}

	res := &pb.StatusResponse{
		Response: utils.ConvertToPbResponse(status),
	}

	return res, nil
}

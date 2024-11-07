package utils

import (
	"mms/internal/dto"
	"mms/internal/models"
	pb "mms/internal/pb"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TimeToTimestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

func TimestampToTime(ts *timestamppb.Timestamp) time.Time {
	if ts != nil {
		return ts.AsTime()
	}
	return time.Time{}
}

func ConvertToPbGender(gender models.Gender) pb.Gender {
	switch gender {
	case models.MALE:
		return pb.Gender_MALE
	case models.FEMALE:
		return pb.Gender_FEMALE
	default:
		return pb.Gender_UNKNOWN
	}
}

func ConvertToPbRole(role models.Role) pb.Role {
	switch role {
	case models.CUSTOMER:
		return pb.Role_CUSTOMER
	case models.EMPLOYEE:
		return pb.Role_EMPLOYEE
	case models.ADMIN:
		return pb.Role_ADMIN
	case models.SUPERADMIN:
		return pb.Role_SUPERADMIN
	default:
		return pb.Role_CUSTOMER
	}
}

func ConvertToPbStatusUser(statusUser models.StatusUser) pb.StatusUser {
	switch statusUser {
	case models.ACTIVE:
		return pb.StatusUser_ACTIVE
	case models.INACTIVE:
		return pb.StatusUser_INACTIVE
	case models.DELETE:
		return pb.StatusUser_DELETE
	default:
		return pb.StatusUser_ACTIVE
	}
}

func ConvertToPbResponse(status *dto.StatusResp) pb.Response {
	switch status.Response {
	case dto.OK:
		return pb.Response_OK
	case dto.ERROR:
		return pb.Response_ERROR
	default:
		return pb.Response_ERROR
	}
}

func ConvertPbRolesToModelsRoles(pbRoles []pb.Role) []models.Role {
	var modelsRoles []models.Role
	for _, role := range pbRoles {
		modelsRoles = append(modelsRoles, models.Role(role))
	}
	return modelsRoles
}

func ConvertPbStatusUsersToModelsStatusUsers(pbStatusUser []pb.StatusUser) []models.StatusUser {
	var modelsStatusUsers []models.StatusUser
	for _, StatusUser := range pbStatusUser {
		modelsStatusUsers = append(modelsStatusUsers, models.StatusUser(StatusUser))
	}
	return modelsStatusUsers
}

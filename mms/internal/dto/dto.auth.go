package dto

type AuthLoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthLoginResp struct {
	Token    string   `json:"token"`
	Response Response `json:"response"`
}

type AuthUpdateTokenReq struct {
	UserId uint32 `json:"user_id" validate:"required"`
	Token  string `json:"token" validate:"required"`
}

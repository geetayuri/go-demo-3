package mainEntities

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type AuthRepository interface {
	SignUsersAccessToken(req *UsersPassport, xxhashUA *string, xxhashOR *string, xxhashRF *string, rType *RefreshTokenRequestType) (string, error)
	CheckRefreshToken(tokensum *string) error
	InsertRefreshToken(tokensum *string,refreshtoken *string, username *string) error
	SignUsersRefreshToken(userid *int, username *string, xxhashUA *string, xxhashOR *string, xxhashRF *string) (string, error)
}

type AuthUsecase interface {
	Login(req *UsersCredentials, xxhashUA *string, xxhashOR *string, xxhashRF *string) (*UsersLoginRes, error)
	RefreshToken(userid *int, username *string, xxhashUA *string, xxhashOR *string, xxhashRF *string, tokenSum *string, rType *RefreshTokenRequestType) (*UsersLoginRes, error)
}

type RefreshTokenRequestType struct {
	Type string `json:"type" db:"type" form:"type"`
}

type UsersCredentials struct {
	Username string `json:"username" db:"username" form:"username"`
	Password string `json:"password" db:"password" form:"password"`
}

type UsersPassport struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UsersClaims struct {
	Id       int    `json:"user_id"`
	Username string `json:"username"`
	Uh       string `json:"uh"`
	Oh       string `json:"oh"`
	Rh       string `json:"rh"`
	jwt.RegisteredClaims
}

type UsersLoginRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UsersRepository interface {
	FindOneUser(username string) (*UsersPassport, error)
	Register(req *UsersRegisterReq) (*UsersRegisterRes, error)
}

// import "context"

type UsersUsecase interface {
	Register(ctx context.Context, req *UsersRegisterReq) (*UsersRegisterRes, error)
}

// type UsersRepository interface {
// 	Register(ctx context.Context, req *UsersRegisterReq) (*UsersRegisterRes, error)
// }

type UsersRegisterReq struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UsersRegisterRes struct {
	Id       uint64 `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}

package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const SERECTKEY = "N@Tura1Bui1de12"

type UserClaims struct {
	Email                string    `json:"email"`
	FirstName            string    `json:"firstname"`
	LastName             string    `json:"lastname"`
	Gender               string    `json:"gender"`
	Image                string    `json:"image"`
	Uuid                 uuid.UUID `json:"uuid"`
	Role                 *Role     `json:"role"`
	DeviceID             string    `json:"device_id"`
	jwt.RegisteredClaims `json:"-"`
}

type Role struct {
	Uuid uuid.UUID
	Name string
}

type AuthenticateClaims struct {
	UUID                 uuid.UUID `json:"uuid"`
	Code                 string    `json:"code"`
	jwt.RegisteredClaims `json:"-"`
}

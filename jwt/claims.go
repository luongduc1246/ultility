package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type UserClaims struct {
	Email                string        `json:"email"`
	Phone                string        `json:"phone"`
	CountryCode          string        `json:"country_code"`
	FirstName            string        `json:"firstname"`
	LastName             string        `json:"lastname"`
	Gender               string        `json:"gender"`
	Image                string        `json:"image"`
	Uuid                 uuid.UUID     `json:"uuid"`
	Role                 *Role         `json:"role"`
	Organization         *Organization `json:"organization"`
	DeviceID             string        `json:"device_id"`
	jwt.RegisteredClaims `json:"-"`
}

type Role struct {
	Uuid uuid.UUID
	Name string
}

type Organization struct {
	ID         int       `json:"organization_id"`
	Uuid       uuid.UUID `json:"organization_uuid"`
	Name       string    `json:"organization_name"`
	Substitute string    `json:"organization_substitute"`
	Role       *Role     `json:"organization_role"`
}

type AuthenticateClaims struct {
	UUID                 uuid.UUID `json:"uuid"`
	Code                 string    `json:"code"`
	jwt.RegisteredClaims `json:"-"`
}

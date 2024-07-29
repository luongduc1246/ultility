package jwt

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func TestCreateUserToken(t *testing.T) {
	uc := UserClaims{
		FirstName: "Duc",
		LastName:  "Luong",
		Uuid:      uuid.New(),
		Role: &Role{
			Uuid: uuid.UUID{},
			Name: "akdjfl",
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(365 * 24 * time.Hour)),
		},
	}
	te, err := CreateUserToken(&uc)
	t.Log(te, err)
}

func BenchmarkU(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uc := UserClaims{
			FirstName: "Duc",
			LastName:  "Luong",
			Uuid:      uuid.New(),
			Role: &Role{
				Uuid: uuid.UUID{},
				Name: "akdjfl",
			},
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(365 * 24 * time.Hour)),
			},
		}
		CreateUserToken(&uc)
	}
}

type UserClaim struct {
	Email                string                            `json:"email"`
	FirstName            string                            `json:"firstname"`
	LastName             string                            `json:"lastname"`
	UUID                 uuid.UUID                         `json:"uuid"`
	Role                 map[string]map[string]interface{} `json:"roles"`
	jwt.RegisteredClaims `json:"-"`
}

func CreateUserT(userclaim *UserClaim) (string, error) {
	userclaim.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(365 * 24 * time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userclaim)
	serect, err := token.SignedString([]byte(SERECTKEY))
	if err != nil {
		return "", err
	}
	return serect, nil
}

func BenchmarkT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uc := UserClaim{
			FirstName: "Duc",
			LastName:  "Luong",
			UUID:      uuid.New(),
			Role: map[string]map[string]interface{}{"admin": {
				"create.name":       1,
				"create.adf":        1,
				"create.adfad":      1,
				"create.sf":         1,
				"create.addfad":     1,
				"create.adfdad":     1,
				"create.sfsd":       1,
				"create.adfddad":    1,
				"create.adfdafadfd": 1,
				"create.sdf":        1,
				"create.adfaa":      1,
			}},
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(365 * 24 * time.Hour)),
			},
		}
		CreateUserT(&uc)
	}
}

func TestParseUserToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IiIsImZpcnN0bmFtZSI6IkR1YyIsImxhc3RuYW1lIjoiTHVvbmciLCJ1dWlkIjoiMjIxN2YzMTQtZWQ5ZC00YjdkLTg5NDItOTI2YTQ3ZDRmYjVjIiwicm9sZXMiOnsiTmFtZSI6ImFkbWluIiwicGVybWlzc2lvbnMiOlsiY3JlYXRlLm5hbWUiLCJkZWxldGUuYmFiZSIsImxhZHNrZmpsIiwibGFrZGpmbCIsImxhc2tkZmpsIiwibGFka2pmbGFqZiIsImxhZGprZmwiLCJsYWRqZmtsIiwibGFrZGZqIiwiYWxkamZsYWwiXX19.VCdwbi1JeuGuRUz9Sm2Uj8Xa_1dRWean9LYzkmUh_oU"
	user, err := ParseUserToken(token)
	fmt.Println(user, err)
}

func TestCreateAuthToken(t *testing.T) {
	au := AuthenticateClaims{
		UUID: uuid.New(),
		Code: "aldsfjl",
	}
	tk, err := CreateAuthenticateToken(&au)
	fmt.Println(tk, err)
}
func TestPearseAuthToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1dWlkIjoiYzhhOWQ4MmEtMjgyMi00ZTczLWFjZTYtOGU4ZGU4MGI2YmE0IiwiY29kZSI6ImFsZHNmamwifQ.8lelgYYOMD1QGCWJG0Rk4Q4rao1t2MUuxcCAZ09SfLY"
	au, err := ParseAuthenticateToken(token)
	fmt.Println(au, err)
}

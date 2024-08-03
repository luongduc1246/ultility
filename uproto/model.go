package proto

import (
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Status int

type Role struct {
	gorm.Model
	Uuid                  uuid.UUID
	Name                  string
	Status                *Status
	Description           *string
	Test                  string
	User                  UserRole
	Mame                  *string
	Address               []*string
	Users                 []*UserRole
	Meme                  []string
	SliceInt              []*int32
	MyTime                time.Time
	MyDataJson            *datatypes.JSON
	MyDataJsonToString    datatypes.JSON
	MyDataJsonMap         *datatypes.JSONMap
	MyDataJsonMapToString *datatypes.JSONMap
	MyDataJsonMapP        datatypes.JSON
	MyDataDate            datatypes.Date
	MyDataTime            datatypes.Time
	MyByte                datatypes.Date
}

type Test struct {
	Status  *int32
	Test    string
	Users   []TestS
	Mame    *wrapperspb.StringValue
	Meme    []*wrapperspb.StringValue
	Address []string
}

type TestS struct {
	Name    string
	Address []*wrapperspb.StringValue
}

type TestRequest struct {
	Uuid                  string
	Name                  *wrapperspb.StringValue
	Status                *wrapperspb.Int32Value
	SliceInt              []*wrapperspb.Int32Value
	Description           string
	Test                  *string
	Mame                  *wrapperspb.StringValue
	User                  *User
	Users                 []*User
	Address               []string
	Meme                  []*wrapperspb.StringValue
	MyTime                int64
	MyDataJsonToString    string
	MyDataJsonMapToString string
	MyDataJsonMapP        string
	MyDataJson            *structpb.Struct
	MyDataJsonMap         *structpb.Struct
	MyDataDate            *timestamppb.Timestamp
	MyDataTime            *durationpb.Duration
	MyByte                []byte
}

func (r Role) GetName() string {
	return "name"
}

type UserRole struct {
	Uuid    uuid.UUID
	Name    string
	Address []string
}

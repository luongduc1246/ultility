package proto

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func TestAssignProtoToGormModel(t *testing.T) {
	m, _ := structpb.NewStruct(map[string]interface{}{
		"firstName": "John",
		"lastName":  "Smith",
		"isAlive":   true,
		"age":       27,
		"phoneNumbers": []interface{}{
			map[string]interface{}{
				"type":   "home",
				"number": "212 555-1234",
			},
			map[string]interface{}{
				"type":   "office",
				"number": "646 555-4567",
			},
		},
		"children": []interface{}{},
		"spouse":   nil,
	})
	// json, _ := json.Marshal(map[string]interface{}{
	// 	"name": "DUc",
	// })
	// date, _ := time.Now().MarshalJSON()

	rq := TestRequest{
		Uuid: uuid.New().String(),
		Name: &wrapperspb.StringValue{
			Value: "aldkfjl",
		},
		SliceInt: []*wrapperspb.Int32Value{
			{
				Value: 123,
			},
			{
				Value: 124,
			},
		},
		Status: &wrapperspb.Int32Value{
			Value: 245,
		},

		User: &User{
			Uuid: uuid.New().String(),
			Name: "aljdflak",
		},
		Description: "ladskjflajdfl",
		Address:     []string{"aldskfjkl", "laksdfjl"},
		Meme: []*wrapperspb.StringValue{
			{
				Value: "aldfjadlkjfl",
			},
			{
				Value: "namu",
			},
		},
		Users:          []*User{{Uuid: uuid.New().String(), Name: "Test"}},
		MyDataJson:     m,
		MyDataJsonMap:  m,
		MyDataJsonMapP: "[\"manager\"]",
		MyDataDate:     timestamppb.Now(),
		MyTime:         0,
		MyDataTime:     durationpb.New(3600000000000),
		// MyByte:        date,
	}
	role := Role{}
	AssignMessageToGormModel(&rq, &role)

	fmt.Printf("%+v \n", role.MyDataJsonMapP)
}
func TestMessageToGorm(t *testing.T) {
	m, _ := structpb.NewStruct(map[string]interface{}{
		"firstName": "John",
		"lastName":  "Smith",
		"isAlive":   true,
		"age":       27,
		"phoneNumbers": []interface{}{
			map[string]interface{}{
				"type":   "home",
				"number": "212 555-1234",
			},
			map[string]interface{}{
				"type":   "office",
				"number": "646 555-4567",
			},
		},
		"children": []interface{}{},
		"spouse":   nil,
	})
	rq := TestRequest{
		Uuid: uuid.New().String(),
		Name: &wrapperspb.StringValue{
			Value: "aldkfjl",
		},
		SliceInt: []*wrapperspb.Int32Value{
			{
				Value: 123,
			},
			{
				Value: 124,
			},
		},
		Status: &wrapperspb.Int32Value{
			Value: 245,
		},

		User: &User{
			Uuid: uuid.New().String(),
			Name: "aljdflak",
		},
		Description: "ladskjflajdfl",
		Address:     []string{"aldskfjkl", "laksdfjl"},
		Meme: []*wrapperspb.StringValue{
			{
				Value: "aldfjadlkjfl",
			},
			{
				Value: "namu",
			},
		},
		Users:         []*User{{Uuid: uuid.New().String(), Name: "Test"}},
		MyDataJson:    m,
		MyDataJsonMap: m,
		MyDataDate:    timestamppb.Now(),
		MyTime:        0,
		MyDataTime:    durationpb.New(3600000000000),
		// MyByte:        date,
	}
	role := Role{}
	mtg := NewMessageToGorm()
	mtg.Parse(&rq, &role)

	fmt.Printf("%+v \n", mtg.Associations["Users"].Model)
}

func TestGormToMessage(t *testing.T) {
	user1Attrs := `{"age":18,"name":"json-1","orgs":{"orga":"orga"},"tags":["tag1","tag2"],"admin":true}`
	dataJson := datatypes.JSON([]byte(user1Attrs))
	dataJsonMap := make(datatypes.JSONMap)
	dataJsonMap.UnmarshalJSON([]byte(`{"id": "369", "comp": "Fragment", "children": [{"id": "343", "comp": "Container", "props": {"style": {"sm": {"display": "flex", "background": "bg-[#50d71e]", "flexDirection": "row"}}}}]}`))
	m := make(map[string]interface{})
	m["a"] = "lasdkjfl"
	rq := TestRequest{}
	dt := datatypes.Date(time.Now())
	des := "laksdjfllkdjf"
	s := "toi la meme"
	b := Status(3)
	role := Role{
		Model:       gorm.Model{},
		Uuid:        uuid.New(),
		Name:        "Hai di",
		Status:      &b,
		Description: &des,
		Test:        des,
		User:        UserRole{Name: "aldjfl"},
		Address: []*string{
			&des, &s,
		},
		Users: []*UserRole{
			{
				Name:    "Test Slice user1",
				Address: []string{"aldskfjl", "lasdfjlsad"},
			},
			{
				Name: "Test user 2",
			},
		},
		Meme: []string{
			des, s,
		},
		Mame:                  &s,
		MyTime:                time.Now(),
		MyDataJson:            &dataJson,
		MyDataJsonToString:    dataJson,
		MyDataJsonMapToString: &dataJsonMap,
		// MyDataJsonMap: datatypes.JSONMap(m),
		MyDataDate: dt,
		MyDataTime: 3600,
		MyByte:     datatypes.Date{},
	}
	// te := Test{}
	AssignGormModelToMessage(&role, &rq)
	fmt.Printf("%+v \n", rq.MyDataJsonMapToString)

}

func Abc() {
	a := "asdlfjlj"
	var b *string
	b = &a
	fmt.Println(b)
	rv := reflect.ValueOf(b)

	fmt.Println(rv)
}

func TestXxx(t *testing.T) {
	var a interface{}
	b := map[string]interface{}{
		"lasdkjfl": "lsdkjfla",
	}
	a = b
	fmt.Println(a)
	dsm, _ := a.(map[string]interface{})
	s := datatypes.JSONMap(dsm)
	fmt.Println(s)
}

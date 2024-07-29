package gormdb

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gorm.io/datatypes"
)

type AssociationModel struct {
	Model        interface{}                  // lưu con trỏ của model
	Associations map[string]*AssociationModel // lưu con trỏ của Associations
}

func NewAssociationModel() *AssociationModel {
	return &AssociationModel{
		Model:        nil,
		Associations: make(map[string]*AssociationModel),
	}
}

// phân tích model từ message protobuf
func (asm *AssociationModel) Parse(message, object interface{}) {
	if asm.Model == nil {
		asm.Model = object
	}
	valOb := reflect.Indirect(reflect.ValueOf(object))
	valMessage := reflect.Indirect(reflect.ValueOf(message))
	typeOb := valOb.Type()
	typeMess := valMessage.Type()
	if !valOb.CanSet() {
		return
	}
	if typeMess.Kind() == reflect.Struct && typeOb.Kind() == reflect.Struct {
		for i := 0; i < typeMess.NumField(); i++ {
			fieldMessType := typeMess.Field(i)
			_, ok := typeOb.FieldByName(fieldMessType.Name)
			if ok {
				valFieldMess := valMessage.FieldByName(fieldMessType.Name)
				valFieldOb := valOb.FieldByName(fieldMessType.Name)
				if !valFieldMess.IsZero() {
					switch valFieldOb.Interface().(type) {
					case uuid.UUID:
						assignToUuid(valFieldMess, valFieldOb)
					case datatypes.JSONMap, *datatypes.JSONMap:
						assignToDataJsonMap(valFieldMess, valFieldOb)
					case datatypes.JSON, *datatypes.JSON:
						assignToDataJson(valFieldMess, valFieldOb)
					case datatypes.Date, *datatypes.Date:
						assignToDataDate(valFieldMess, valFieldOb)
					case datatypes.Time, *datatypes.Time:
						assignToDataTime(valFieldMess, valFieldOb)
					case time.Time, *time.Time:
						assignToTime(valFieldMess, valFieldOb)
					default:
						switch valFieldMess.Kind() {
						case reflect.Pointer:
							if !valFieldMess.IsNil() {
								switch valFieldMess.Interface().(type) {
								case *wrapperspb.BoolValue, *wrapperspb.BytesValue, *wrapperspb.FloatValue, *wrapperspb.DoubleValue,
									*wrapperspb.Int32Value, *wrapperspb.Int64Value, *wrapperspb.StringValue,
									*wrapperspb.UInt32Value, *wrapperspb.UInt64Value:
									val := valFieldMess.Elem().FieldByName("Value")
									if val.IsValid() {
										if valFieldOb.Kind() == reflect.Pointer {
											newVal := reflect.New(valFieldOb.Type().Elem())
											if val.CanConvert(valFieldOb.Type().Elem()) {
												newVal.Elem().Set(val.Convert(valFieldOb.Type().Elem()))
											}
											valFieldOb.Set(newVal)
										} else {
											if val.CanConvert(valFieldOb.Type()) {
												valFieldOb.Set(val.Convert(valFieldOb.Type()))
											}
										}
									}
								default:
									if valFieldMess.Elem().Kind() == reflect.Struct {
										newMtg := NewAssociationModel()
										asm.Associations[fieldMessType.Name] = newMtg
										if valFieldOb.Kind() == reflect.Pointer {
											newElemObVal := reflect.New(valFieldOb.Type().Elem())
											newMtg.Parse(valFieldMess.Interface(), newElemObVal.Interface())
											valFieldOb.Set(newElemObVal)
										} else {
											newElemObVal := reflect.New(valFieldOb.Type())
											newMtg.Parse(valFieldMess.Interface(), newElemObVal.Interface())
											valFieldOb.Set(newElemObVal.Elem())

										}
									} else {
										if valFieldOb.Kind() == reflect.Pointer {
											if valFieldMess.Elem().CanConvert(valFieldOb.Type().Elem()) {
												newVal := reflect.New(valFieldOb.Type().Elem())
												newVal.Elem().Set(valFieldMess.Elem().Convert(valFieldOb.Type().Elem()))
												valFieldOb.Set(newVal)
											}
										} else {
											if valFieldMess.Elem().CanConvert(valFieldOb.Type()) {
												valFieldOb.Set(valFieldMess.Elem().Convert(valFieldOb.Type()))
											}
										}
									}
								}
							}
						case reflect.Array, reflect.Slice:
							if sliceLength := valFieldMess.Len(); sliceLength > 0 {
								//type one elem of slice
								elemObType := valFieldOb.Type().Elem()
								// elemMessType := valFieldMess.Type().Elem()
								for j := 0; j < sliceLength; j++ {
									valIndex := valFieldMess.Index(j)
									if valIndex.IsValid() {
										switch valIndex.Interface().(type) {
										case *wrapperspb.BoolValue, *wrapperspb.BytesValue, *wrapperspb.FloatValue, *wrapperspb.DoubleValue,
											*wrapperspb.Int32Value, *wrapperspb.Int64Value, *wrapperspb.StringValue,
											*wrapperspb.UInt32Value, *wrapperspb.UInt64Value:
											val := valIndex.Elem().FieldByName("Value")
											if val.IsValid() {
												if elemObType.Kind() == reflect.Pointer {
													newVal := reflect.New(elemObType.Elem())
													if val.CanConvert(elemObType.Elem()) {
														newVal.Elem().Set(val.Convert(elemObType.Elem()))
													}
													valFieldOb.Set(reflect.Append(valFieldOb, newVal))
												} else {
													newVal := reflect.New(elemObType)
													if val.CanConvert(elemObType) {
														newVal.Elem().Set(val.Convert(elemObType))
													}
													valFieldOb.Set(reflect.Append(valFieldOb, newVal.Elem()))
												}
											}
										default:
											if valIndex.Kind() == reflect.Pointer {
												if valIndex.Elem().Kind() == reflect.Struct {
													newMtg := NewAssociationModel()
													asm.Associations[fieldMessType.Name] = newMtg
													newMtg.Model = valFieldOb
													if elemObType.Kind() == reflect.Pointer {
														newElemObVal := reflect.New(elemObType.Elem())
														newMtg.Parse(valIndex.Interface(), newElemObVal.Interface())
														valFieldOb.Set(reflect.Append(valFieldOb, newElemObVal))
													} else {
														newElemObVal := reflect.New(elemObType)
														newMtg.Parse(valIndex.Interface(), newElemObVal.Interface())
														valFieldOb.Set(reflect.Append(valFieldOb, newElemObVal.Elem()))
													}
													newMtg.Model = valFieldOb.Interface()
												} else {
													if elemObType.Kind() == reflect.Pointer {
														newElemObVal := reflect.New(elemObType.Elem())
														if valIndex.Elem().CanConvert(elemObType.Elem()) {
															newElemObVal.Elem().Set(valIndex.Elem().Convert(elemObType.Elem()))
														}
														valFieldOb.Set(reflect.Append(valFieldOb, newElemObVal))
													} else {
														newElemObVal := reflect.New(elemObType)
														if valIndex.Elem().CanConvert(elemObType) {
															newElemObVal.Elem().Set(valIndex.Elem().Convert(elemObType))
														}
														valFieldOb.Set(reflect.Append(valFieldOb, newElemObVal.Elem()))
													}
												}
											} else {
												if valIndex.Kind() == reflect.Struct {
													newMtg := NewAssociationModel()
													asm.Associations[fieldMessType.Name] = newMtg
													newMtg.Model = valFieldOb
													if elemObType.Kind() == reflect.Pointer {
														newElemObVal := reflect.New(elemObType)
														newMtg.Parse(valIndex.Interface(), newElemObVal.Interface())
														valFieldOb.Set(reflect.Append(valFieldOb, newElemObVal))
													} else {
														newElemObVal := reflect.New(elemObType)
														newMtg.Parse(valIndex.Interface(), newElemObVal.Interface())
														valFieldOb.Set(reflect.Append(valFieldOb, newElemObVal.Elem()))
													}
													newMtg.Model = valFieldOb.Interface()
												} else {
													if elemObType.Kind() == reflect.Pointer {
														newElemObVal := reflect.New(elemObType.Elem())
														if valIndex.CanConvert(elemObType.Elem()) {
															newElemObVal.Elem().Set(valIndex.Convert(elemObType.Elem()))
														}
														valFieldOb.Set(reflect.Append(valFieldOb, newElemObVal))
													} else {
														newElemObVal := reflect.New(elemObType)
														if valIndex.CanConvert(elemObType) {
															newElemObVal.Elem().Set(valIndex.Convert(elemObType))
														}
														valFieldOb.Set(reflect.Append(valFieldOb, newElemObVal.Elem()))
													}
												}
											}
										}
									}
								}
							}
						default:
							if valFieldOb.Kind() == reflect.Pointer {
								if valFieldMess.CanConvert(valFieldOb.Type().Elem()) {
									newVal := reflect.New(valFieldOb.Type().Elem())
									newVal.Elem().Set(valFieldMess.Convert(valFieldOb.Type().Elem()))
									valFieldOb.Set(newVal)
								}
							} else {
								if valFieldMess.CanConvert(valFieldOb.Type()) {
									valFieldOb.Set(valFieldMess.Convert(valFieldOb.Type()))
								}
							}
						}
					}
				}
			}
		}
	}
}

func assignToUuid(message, gorm reflect.Value) {
	switch message.Interface().(type) {
	case string:
		uid, err := uuid.Parse(message.String())
		if err != nil {
			return
		}
		valNewUuid := reflect.ValueOf(uid)
		gorm.Set(valNewUuid)
	case []byte:
		uid := uuid.UUID{}
		uid.UnmarshalBinary(message.Bytes())
		gorm.Set(reflect.ValueOf(uid))
	}
}
func assignToDataJsonMap(message, gorm reflect.Value) {
	switch message.Interface().(type) {
	case *structpb.Struct:
		if message.IsNil() {
			return
		}
		valFunc := message.MethodByName("AsMap")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 {
				valMap := callFunc[0]
				mVal, ok := valMap.Interface().(map[string]interface{})
				if !ok {
					return
				}
				dataJsm := datatypes.JSONMap(mVal)
				if gorm.Kind() == reflect.Pointer {
					gorm.Set(reflect.ValueOf(&dataJsm))
				} else {
					gorm.Set(reflect.ValueOf(dataJsm))
				}

				// // cách khác để tạo datajson Map
				// gorm.Set(reflect.MakeMap(reflect.TypeOf(datatypes.JSONMap{})))
				// iter := valMap.MapRange()
				// for iter.Next() {
				// 	gorm.SetMapIndex(iter.Key(), iter.Value())
				// }
			}
		}
	case []byte:
		jsonMap := datatypes.JSONMap{}
		err := jsonMap.UnmarshalJSON(message.Bytes())
		if err != nil {
			return
		}
		if gorm.Kind() == reflect.Pointer {
			gorm.Set(reflect.ValueOf(&jsonMap))
		} else {
			gorm.Set(reflect.ValueOf(jsonMap))
		}
	case string:
		m := make(map[string]interface{})
		e := json.Unmarshal([]byte(message.String()), &m)
		if e != nil {
			return
		}
		if len(m) > 0 {
			dataJsm := datatypes.JSONMap(m)
			if gorm.Kind() == reflect.Pointer {
				gorm.Set(reflect.ValueOf(&dataJsm))
			} else {
				gorm.Set(reflect.ValueOf(dataJsm))
			}
		}
	}
}
func assignToDataJson(message, gorm reflect.Value) {
	switch message.Interface().(type) {
	case *structpb.Struct:
		if message.IsNil() {
			return
		}
		valFunc := message.MethodByName("MarshalJSON")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 && callFunc[1].IsNil() {
				val := callFunc[0]
				dataJson := datatypes.JSON(val.Bytes())
				if gorm.Kind() == reflect.Pointer {
					gorm.Set(reflect.ValueOf(&dataJson))
				} else {
					gorm.Set(reflect.ValueOf(dataJson))
				}
			}
		}
	case []byte:
		js := datatypes.JSON{}
		err := js.UnmarshalJSON(message.Bytes())
		if err != nil {
			return
		}
		if gorm.Kind() == reflect.Pointer {
			gorm.Set(reflect.ValueOf(&js))
		} else {
			gorm.Set(reflect.ValueOf(js))
		}
	case string:
		jr := json.RawMessage(message.String())
		js, err := jr.MarshalJSON()
		if err != nil {
			return
		}
		jsData := datatypes.JSON{}
		jsData.UnmarshalJSON(js)
		if gorm.Kind() == reflect.Pointer {
			gorm.Set(reflect.ValueOf(&js))
		} else {
			gorm.Set(reflect.ValueOf(js))
		}
	}
}
func assignToDataDate(message, gorm reflect.Value) {
	switch message.Interface().(type) {
	case *timestamppb.Timestamp:
		if message.IsNil() {
			return
		}
		valFunc := message.MethodByName("AsTime")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 {
				val := callFunc[0]
				typeDate := reflect.TypeOf(datatypes.Date{})
				if val.CanConvert(typeDate) {
					valCon := val.Convert(typeDate)
					if gorm.Kind() == reflect.Pointer {
						if valCon.CanAddr() {
							gorm.Set(valCon.Addr())
						}
					} else {
						gorm.Set(valCon)
					}
				}
			}
		}
	case []byte:
		t := time.Time{}
		err := t.UnmarshalBinary(message.Bytes())
		if err != nil {
			return
		}
		dataDate := datatypes.Date(t)
		if gorm.Kind() == reflect.Pointer {
			gorm.Set(reflect.ValueOf(&dataDate))
		} else {
			gorm.Set(reflect.ValueOf(dataDate))
		}
	case int64, int32:
		t := time.Unix(message.Int(), 0)
		dataDate := datatypes.Date(t)
		if gorm.Kind() == reflect.Pointer {
			gorm.Set(reflect.ValueOf(&dataDate))
		} else {
			gorm.Set(reflect.ValueOf(dataDate))
		}
	case string:
		t, err := time.Parse(time.RFC1123, message.String())
		if err != nil {
			return
		}
		dataDate := datatypes.Date(t)
		if gorm.Kind() == reflect.Pointer {
			gorm.Set(reflect.ValueOf(&dataDate))
		} else {
			gorm.Set(reflect.ValueOf(dataDate))
		}
	}
}
func assignToDataTime(message, gorm reflect.Value) {
	switch message.Interface().(type) {
	case *durationpb.Duration:
		if message.IsNil() {
			return
		}
		valFunc := message.MethodByName("AsDuration")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 {
				val := callFunc[0]
				td := datatypes.Time(val.Int())
				if gorm.Kind() == reflect.Pointer {

					gorm.Set(reflect.ValueOf(&td))

				} else {
					gorm.Set(reflect.ValueOf(td))
				}
			}
		}
	case string:
		td, err := time.ParseDuration(message.String())
		if err != nil {
			return
		}
		typeTime := datatypes.Time(td)
		if gorm.Kind() == reflect.Pointer {
			gorm.Set(reflect.ValueOf(&typeTime))
		} else {
			gorm.Set(reflect.ValueOf(typeTime))
		}
	}
}
func assignToTime(message, gorm reflect.Value) {
	switch message.Interface().(type) {
	case *timestamppb.Timestamp:
		if message.IsNil() {
			return
		}
		valFunc := message.MethodByName("AsTime")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 {
				val := callFunc[0]
				if gorm.Kind() == reflect.Pointer {
					if val.CanAddr() {
						gorm.Set(val.Addr())
					}
				} else {
					gorm.Set(val)
				}
			}
		}
	case []byte:
		t := time.Time{}
		err := t.UnmarshalBinary(message.Bytes())
		if err != nil {
			return
		}
		if gorm.Kind() == reflect.Pointer {
			gorm.Set(reflect.ValueOf(&t))
		} else {
			gorm.Set(reflect.ValueOf(t))
		}
	case int64, int32:
		t := time.Unix(message.Int(), 0)
		if gorm.Kind() == reflect.Pointer {
			gorm.Set(reflect.ValueOf(&t))
		} else {
			gorm.Set(reflect.ValueOf(t))
		}
	case string:
		t, err := time.Parse(time.RFC1123, message.String())
		if err != nil {
			return
		}
		if gorm.Kind() == reflect.Pointer {
			gorm.Set(reflect.ValueOf(&t))
		} else {
			gorm.Set(reflect.ValueOf(t))
		}
	}
}

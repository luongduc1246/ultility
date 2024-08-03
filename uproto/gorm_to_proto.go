package proto

import (
	"reflect"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func AssignGormModelToMessage(object, message interface{}) {
	valOb := reflect.Indirect(reflect.ValueOf(object))
	valMessage := reflect.Indirect(reflect.ValueOf(message))
	typeOb := valOb.Type()
	typeMess := valMessage.Type()
	if !valMessage.CanSet() {
		return
	}
	if typeMess.Kind() == reflect.Struct && typeOb.Kind() == reflect.Struct {
		for i := 0; i < typeMess.NumField(); i++ {
			fieldMessType := typeMess.Field(i)
			_, ok := typeOb.FieldByName(fieldMessType.Name)
			if ok {
				valFieldMess := valMessage.FieldByName(fieldMessType.Name)
				valFieldOb := valOb.FieldByName(fieldMessType.Name)
				if valFieldOb.Kind() != reflect.Interface {
					if !valFieldOb.IsZero() {
						switch valFieldOb.Interface().(type) {
						case uuid.UUID:
							parseFromUuid(valFieldOb, valFieldMess)
						case datatypes.JSON, *datatypes.JSON:
							parseFromDataJson(valFieldOb, valFieldMess)
						case datatypes.JSONMap, *datatypes.JSONMap:
							parseFromDataJsonMap(valFieldOb, valFieldMess)
						case datatypes.Date, *datatypes.Date:
							parseFromDataDate(valFieldOb, valFieldMess)
						case datatypes.Time, *datatypes.Time:
							parseFromDataTime(valFieldOb, valFieldMess)
						case time.Time, *time.Time:
							parseFromTime(valFieldOb, valFieldMess)
						case gorm.DeletedAt:
							parseFromDeleteAt(valFieldOb, valFieldMess)
						default:
							switch valFieldMess.Kind() {
							case reflect.Pointer:
								switch valFieldMess.Interface().(type) {
								case *wrapperspb.BoolValue, *wrapperspb.BytesValue, *wrapperspb.FloatValue, *wrapperspb.DoubleValue,
									*wrapperspb.Int32Value, *wrapperspb.Int64Value, *wrapperspb.StringValue,
									*wrapperspb.UInt32Value, *wrapperspb.UInt64Value:
									newVal := reflect.New(valFieldMess.Type().Elem())
									val := newVal.Elem().FieldByName("Value")
									if valFieldOb.Kind() == reflect.Pointer {
										valFieldOb = valFieldOb.Elem()
									}
									if valFieldOb.CanConvert(val.Type()) {
										val.Set(valFieldOb.Convert(val.Type()))
									}
									valFieldMess.Set(newVal)
								default:
									if valFieldMess.Type().Elem().Kind() == reflect.Struct {
										newVal := reflect.New(valFieldMess.Type().Elem())
										AssignGormModelToMessage(valFieldOb.Interface(), newVal.Interface())
										valFieldMess.Set(newVal)
									} else {
										if valFieldOb.Kind() == reflect.Pointer {
											newValFielMess := reflect.New(valFieldMess.Type().Elem())
											if valFieldOb.Elem().CanConvert(newValFielMess.Type().Elem()) {
												val := valFieldOb.Elem().Convert(newValFielMess.Type().Elem())
												newValFielMess.Elem().Set(val)
												valFieldMess.Set(newValFielMess)
											}
										} else {
											// tạo mới để lấy địa chỉ con trỏ
											newValFielMess := reflect.New(valFieldMess.Type().Elem())
											// ép kiểu để gán giá trị
											if valFieldOb.CanConvert(newValFielMess.Type().Elem()) {
												val := valFieldOb.Convert(newValFielMess.Type().Elem())
												newValFielMess.Elem().Set(val)
												valFieldMess.Set(newValFielMess)
											}

										}
									}
								}
							case reflect.Slice, reflect.Array:
								if sliceLength := valFieldOb.Len(); sliceLength > 0 {
									//type one elem of slice
									elemObType := valFieldOb.Type().Elem()
									elemMessType := valFieldMess.Type().Elem()
									valElemMess := reflect.New(elemMessType).Elem()
									switch valElemMess.Interface().(type) {
									case *wrapperspb.BoolValue, *wrapperspb.BytesValue, *wrapperspb.FloatValue, *wrapperspb.DoubleValue,
										*wrapperspb.Int32Value, *wrapperspb.Int64Value, *wrapperspb.StringValue,
										*wrapperspb.UInt32Value, *wrapperspb.UInt64Value:
										for j := 0; j < sliceLength; j++ {
											valIndex := valFieldOb.Index(j)
											if !valIndex.IsZero() {
												newElemMess := reflect.New(elemMessType.Elem())
												val := newElemMess.Elem().FieldByName("Value")
												if valIndex.Kind() == reflect.Pointer {
													if valIndex.Elem().CanConvert(val.Type()) {
														val.Set(valIndex.Elem().Convert(val.Type()))
													}
												} else {
													if valIndex.CanConvert(val.Type()) {
														val.Set(valIndex.Convert(val.Type()))
													}
												}
												valFieldMess.Set(reflect.Append(valFieldMess, newElemMess))
											}
										}
									default:
										if elemMessType.Kind() == reflect.Pointer {
											// làm việc với struct
											if elemMessType.Elem().Kind() == reflect.Struct {
												for j := 0; j < sliceLength; j++ {
													valIndex := valFieldOb.Index(j)
													newElemMessVal := reflect.New(elemMessType.Elem())
													AssignGormModelToMessage(valIndex.Interface(), newElemMessVal.Interface())
													valFieldMess.Set(reflect.Append(valFieldMess, newElemMessVal))
												}
											} else {
												for j := 0; j < sliceLength; j++ {
													newElemMessVal := reflect.New(elemMessType.Elem())
													valIndex := valFieldOb.Index(j)
													if elemObType.Kind() == reflect.Pointer {
														if valIndex.Elem().CanConvert(elemMessType.Elem()) {
															newElemMessVal.Elem().Set(valIndex.Elem().Convert(elemMessType.Elem()))
														}
													} else {
														if valIndex.CanConvert(elemMessType.Elem()) {
															newElemMessVal.Elem().Set(valIndex.Convert(elemMessType.Elem()))
														}
													}
													valFieldMess.Set(reflect.Append(valFieldMess, newElemMessVal))
												}
											}
										} else {
											// làm việc với struct
											if elemMessType.Kind() == reflect.Struct {
												for j := 0; j < sliceLength; j++ {
													valIndex := valFieldOb.Index(j)
													newElemMessVal := reflect.New(elemMessType)
													AssignGormModelToMessage(valIndex.Interface(), newElemMessVal.Interface())
													valFieldMess.Set(reflect.Append(valFieldMess, newElemMessVal.Elem()))
												}
											} else {
												for j := 0; j < sliceLength; j++ {
													newElemMessVal := reflect.New(elemMessType).Elem()
													valIndex := valFieldOb.Index(j)
													if elemObType.Kind() == reflect.Pointer {
														if valIndex.Elem().CanConvert(elemMessType) {
															newElemMessVal.Set(valIndex.Elem().Convert(elemMessType))
														}
													} else {
														if valIndex.CanConvert(elemMessType) {
															newElemMessVal.Set(valIndex.Convert(elemMessType))
														}
													}
													valFieldMess.Set(reflect.Append(valFieldMess, newElemMessVal))
												}
											}
										}
									}
								}
							default:
								if valFieldOb.Kind() == reflect.Pointer {
									valFieldOb = valFieldOb.Elem()
									if valFieldOb.CanConvert(valFieldMess.Type()) {
										valFieldMess.Set(valFieldOb.Convert(valFieldMess.Type()))
									}
								} else {
									if valFieldOb.CanConvert(valFieldMess.Type()) {
										valFieldMess.Set(valFieldOb.Convert(valFieldMess.Type()))
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func parseFromUuid(object, message reflect.Value) {
	switch message.Interface().(type) {
	case string:
		valFunc := object.MethodByName("String")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 {
				val := callFunc[0]
				if val.CanConvert(message.Type()) {
					message.Set(val.Convert(message.Type()))
				}
			}
		}
	case *string:
		valFunc := object.MethodByName("String")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 {
				val := callFunc[0]
				newVal := reflect.New(message.Type().Elem())
				newVal.Elem().Set(val)
				if newVal.CanConvert(message.Type()) {
					message.Set(newVal.Convert(message.Type()))
				}
			}
		}
	case []byte:
		valFunc := object.MethodByName("MarshalBinary")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 {
				val := callFunc[0]
				if val.CanConvert(message.Type()) {
					message.Set(val.Convert(message.Type()))
				}
			}
		}
	}
}

func parseFromDataJson(object, message reflect.Value) {
	switch message.Interface().(type) {
	case []byte:
		valFunc := object.MethodByName("MarshalJSON")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 && callFunc[1].IsNil() {
				valMap := callFunc[0]
				message.Set(valMap)
			}
		}
	case *structpb.Struct:
		valFunc := object.MethodByName("MarshalJSON")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 && callFunc[1].IsNil() {
				valMap := callFunc[0]
				data := structpb.Struct{}
				data.UnmarshalJSON(valMap.Bytes())
				message.Set(reflect.ValueOf(&data))
			}
		}
	case string:
		valFunc := object.MethodByName("String")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 {
				valMap := callFunc[0]
				message.Set(valMap)
			}
		}
	case *string:
		valFunc := object.MethodByName("String")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 {
				val := callFunc[0]
				newVal := reflect.New(message.Type().Elem())
				newVal.Elem().Set(val)
				if newVal.CanConvert(message.Type()) {
					message.Set(newVal.Convert(message.Type()))
				}
			}
		}
	}
}
func parseFromDataJsonMap(object, message reflect.Value) {
	switch message.Interface().(type) {
	case []byte:
		valFunc := object.MethodByName("MarshalJSON")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 && callFunc[1].IsNil() {
				val := callFunc[0]
				message.Set(val)
			}
		}
	case *structpb.Struct:
		valFunc := object.MethodByName("MarshalJSON")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 && callFunc[1].IsNil() {
				valMap := callFunc[0]
				data := structpb.Struct{}
				data.UnmarshalJSON(valMap.Bytes())
				message.Set(reflect.ValueOf(&data))
			}
		}
	case string:
		valFunc := object.MethodByName("MarshalJSON")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 && callFunc[1].IsNil() {
				val := callFunc[0]
				s := string(val.Bytes())
				message.Set(reflect.ValueOf(s))
			}
		}
	case *string:
		valFunc := object.MethodByName("MarshalJSON")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 && callFunc[1].IsNil() {
				val := callFunc[0]
				s := string(val.Bytes())
				newVal := reflect.New(message.Type().Elem())
				newVal.Elem().Set(reflect.ValueOf(s))
				if newVal.CanConvert(message.Type()) {
					message.Set(newVal.Convert(message.Type()))
				}
			}
		}
	}
}
func parseFromDataDate(object, message reflect.Value) {
	if object.Kind() == reflect.Pointer {
		object = object.Elem()
	}
	if !object.CanConvert(reflect.TypeOf(time.Time{})) {
		return
	} else {
		object = object.Convert(reflect.TypeOf(time.Time{}))
	}

	switch message.Interface().(type) {
	case []byte:
		valFunc := object.MethodByName("MarshalBinary")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 && callFunc[1].IsNil() {
				val := callFunc[0]
				message.Set(val)
			}
		}
	case *timestamppb.Timestamp:
		t, ok := object.Interface().(time.Time)
		if !ok {
			return
		}
		tstam := timestamppb.New(t)
		message.Set(reflect.ValueOf(tstam))
	case string:
		valFunc := object.MethodByName("String")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 {
				val := callFunc[0]
				message.Set(val)
			}
		}
	case int, int32, int64:
		valFunc := object.MethodByName("Unix")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 {
				val := callFunc[0]
				if val.CanConvert(message.Type()) {
					message.Set(val.Convert(message.Type()))
				}
			}
		}
	}
}

func parseFromDataTime(object, message reflect.Value) {
	if object.Kind() == reflect.Pointer {
		object = object.Elem()
	}
	var a time.Duration
	if !object.CanConvert(reflect.TypeOf(a)) {
		return
	} else {
		object = object.Convert(reflect.TypeOf(a))
	}

	switch message.Interface().(type) {
	case *durationpb.Duration:
		t, ok := object.Interface().(time.Duration)
		if !ok {
			return
		}
		duration := durationpb.New(t)
		message.Set(reflect.ValueOf(duration))
	case string:
		valFunc := object.MethodByName("String")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 {
				val := callFunc[0]
				message.Set(val)
			}
		}
	}
}
func parseFromTime(object, message reflect.Value) {
	if object.Kind() == reflect.Pointer {
		object = object.Elem()
	}
	switch message.Interface().(type) {
	case []byte:
		valFunc := object.MethodByName("MarshalBinary")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 && callFunc[1].IsNil() {
				val := callFunc[0]
				message.Set(val)
			}
		}
	case *timestamppb.Timestamp:
		t, ok := object.Interface().(time.Time)
		if !ok {
			return
		}
		tstam := timestamppb.New(t)
		message.Set(reflect.ValueOf(tstam))
	case string:
		valFunc := object.MethodByName("String")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 {
				val := callFunc[0]
				message.Set(val)
			}
		}
	case int32, int64, uint32, uint64:
		valFunc := object.MethodByName("Unix")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 {
				val := callFunc[0]
				if val.CanConvert(message.Type()) {
					message.Set(val.Convert(message.Type()))
				}
			}
		}
	case *int32, *int64, *uint32, *uint64:
		valFunc := object.MethodByName("Unix")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 {
				val := callFunc[0]
				newVal := reflect.New(message.Type().Elem())
				newVal.Elem().Set(val)
				if newVal.CanConvert(message.Type()) {
					message.Set(newVal.Convert(message.Type()))
				}
			}
		}
	}
}
func parseFromDeleteAt(object, message reflect.Value) {
	if object.Kind() == reflect.Pointer {
		object = object.Elem()
	}
	object = object.FieldByName("Time")
	switch message.Interface().(type) {
	case []byte:
		valFunc := object.MethodByName("MarshalBinary")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 && callFunc[1].IsNil() {
				val := callFunc[0]
				message.Set(val)
			}
		}
	case *timestamppb.Timestamp:
		t, ok := object.Interface().(time.Time)
		if !ok {
			return
		}
		tstam := timestamppb.New(t)
		message.Set(reflect.ValueOf(tstam))
	case string:
		valFunc := object.MethodByName("String")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 {
				val := callFunc[0]
				message.Set(val)
			}
		}
	case int32, int64, uint32, uint64:
		valFunc := object.MethodByName("Unix")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 {
				val := callFunc[0]
				if val.CanConvert(message.Type()) {
					message.Set(val.Convert(message.Type()))
				}
			}
		}
	case *int32, *int64, *uint32, *uint64:
		valFunc := object.MethodByName("Unix")
		if valFunc.IsValid() {
			callFunc := valFunc.Call([]reflect.Value{})
			if len(callFunc) > 0 {
				val := callFunc[0]
				newVal := reflect.New(message.Type().Elem())
				newVal.Elem().Set(val)
				if newVal.CanConvert(message.Type()) {
					message.Set(newVal.Convert(message.Type()))
				}
			}
		}
	}
}

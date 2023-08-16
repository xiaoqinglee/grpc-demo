package pkg

import (
	"encoding/json"
	"github.com/k0kubun/pp/v3"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/structpb"
)

func ProtoMessageToMap(m proto.Message) map[string]interface{} {
	jsonBytes, _ := protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(m)
	var data map[string]interface{}
	_ = json.Unmarshal(jsonBytes, &data)
	return data
}
func ProtoMessageToMap2(m proto.Message) map[string]interface{} {
	jsonBytes, _ := protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(m)
	newStruct := new(structpb.Struct)
	_ = protojson.Unmarshal(jsonBytes, newStruct)
	return newStruct.AsMap()
}

func MapToProtoMessage(mp map[string]interface{}, m proto.Message) {
	jsonBytes, _ := json.Marshal(mp)
	_ = protojson.Unmarshal(jsonBytes, m)
}
func MapToProtoMessage2(mp map[string]interface{}, m proto.Message) {
	newStruct, _ := structpb.NewStruct(mp)
	jsonBytes, _ := protojson.Marshal(newStruct)
	_ = protojson.Unmarshal(jsonBytes, m)
}

func ProtoMessageFieldDistinguishNull(m proto.Message) map[string]bool {
	fieldDistinguishesNull := make(map[string]bool)
	//迭代消息类型字段
	fields := m.ProtoReflect().Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		fieldDistinguishesNull[fields.Get(i).TextName()] = fields.Get(i).HasPresence()
	}
	return fieldDistinguishesNull
}
func ProtoMessageNonZeroValueFields(m proto.Message) map[string]interface{} {
	instancePresentFieldToValue := make(map[string]interface{})
	//迭代消息实例字段
	m.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) (continueIteration bool) {
		if v.IsValid() { //当底层字段不是零值的时候返回true
			instancePresentFieldToValue[fd.TextName()] = v.Interface()
		}
		return true
	})
	return instancePresentFieldToValue
}

func PrintProtoMessage(m proto.Message) {
	pp.Println(ProtoMessageToMap(m))
}

// proto message and json encoding:
//
// Package protojson marshals and unmarshals protocol buffer messages as JSON format.
// It follows the guide at https://protobuf.dev/programming-guides/proto3#json.
//
// This package produces a different output than the standard "encoding/json" package,
// which does not operate correctly on protocol buffer messages.
//
// https://github.com/golang/protobuf/issues/1388#issuecomment-988167110
// https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson
// https://protobuf.dev/programming-guides/proto3/#json

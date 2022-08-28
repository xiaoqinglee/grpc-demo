package main

import (
	"github.com/k0kubun/pp/v3"
	"grpc-demo/pkg"
	"grpc-demo/tell"
)

func main() {
	nonOptionalMsg := &tell.NonOptionalFieldMessage{
		Id:    0,
		Id2:   42,
		Ping:  "",
		Ping2: "42",
	}

	//结果
	//map[string]interface {}{
	//	"id":    "0",
	//	"id2":   "42",
	//	"ping":  "",
	//	"ping2": "42",
	//}
	pp.Println(pkg.ProtoMessageToMap(nonOptionalMsg))
	pp.Println(pkg.ProtoMessageToMap2(nonOptionalMsg))
	pp.Println("=================================")

	//和使用这个一样
	//mp := map[string]interface{}{
	//	"id":    0,
	//	"id2":   42,
	//	"ping":  "",
	//	"ping2": "42",
	//}
	mp := map[string]interface{}{
		"id":    "0",
		"id2":   "42",
		"ping":  "",
		"ping2": "42",
	}
	pkg.MapToProtoMessage(mp, nonOptionalMsg)
	pp.Println(nonOptionalMsg.Id)
	pp.Println(nonOptionalMsg.Id2)
	pp.Println(nonOptionalMsg.Ping)
	pp.Println(nonOptionalMsg.Ping2)
	pkg.MapToProtoMessage2(mp, nonOptionalMsg)
	pp.Println(nonOptionalMsg.Id)
	pp.Println(nonOptionalMsg.Id2)
	pp.Println(nonOptionalMsg.Ping)
	pp.Println(nonOptionalMsg.Ping2)
	pp.Println("=================================")

	pp.Println(pkg.ProtoMessageFieldDistinguishNull(nonOptionalMsg))
	pp.Println("=================================")

	pp.Println(pkg.ProtoMessageNonZeroValueFields(nonOptionalMsg))
	pp.Println("===================================================================================================")

	var num int64 = 100
	str := "kim"
	optionalMsg := &tell.OptionalFieldMessage{
		Id:    nil,
		Id2:   &num,
		Ping:  nil,
		Ping2: &str,
	}

	//结果:
	//map[string]interface {}{
	//	"id2":   "100",
	//	"ping2": "kim",
	//}
	pp.Println(pkg.ProtoMessageToMap(optionalMsg))
	pp.Println(pkg.ProtoMessageToMap2(optionalMsg))
	pp.Println("=================================")

	//和使用这个结果一样
	//mp = map[string]interface{}{
	//	"id2":   100,
	//	"ping2": "kim",
	//}
	mp = map[string]interface{}{
		"id2":   "100",
		"ping2": "kim",
	}
	pkg.MapToProtoMessage(mp, optionalMsg)
	pp.Println(optionalMsg.Id)
	pp.Println(optionalMsg.Id2)
	pp.Println(optionalMsg.Ping)
	pp.Println(optionalMsg.Ping2)
	pkg.MapToProtoMessage2(mp, optionalMsg)
	pp.Println(optionalMsg.Id)
	pp.Println(optionalMsg.Id2)
	pp.Println(optionalMsg.Ping)
	pp.Println(optionalMsg.Ping2)
	pp.Println("=================================")

	pp.Println(pkg.ProtoMessageFieldDistinguishNull(optionalMsg))
	pp.Println("=================================")

	pp.Println(pkg.ProtoMessageNonZeroValueFields(optionalMsg))
	pp.Println("=================================")
}

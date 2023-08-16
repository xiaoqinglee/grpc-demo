package main

import (
	"context"
	"github.com/k0kubun/pp/v3"
	"google.golang.org/grpc"
	pb "grpc-demo/tell"
	"io"
	"log"
	"net"
)

type TellEvenServer struct {
	pb.UnimplementedTellEvenNumberServiceServer
}

func (server *TellEvenServer) IsEven(ctx context.Context, req *pb.Number) (resp *pb.NumberIsEven, err error) {
	resp = &pb.NumberIsEven{
		Num:    req.Num,
		IsEven: numberIsEvenNumber(req.Num),
	}
	return
}

func (server *TellEvenServer) IsEvenUsingList(ctx context.Context, req *pb.Numbers) (resp *pb.NumberIsEvenList, err error) {
	resp = &pb.NumberIsEvenList{NumIsEvenList: nil}
	for _, pNumber := range req.Nums {
		num := pNumber.Num
		resp.NumIsEvenList = append(resp.NumIsEvenList, &pb.NumberIsEven{
			Num:    num,
			IsEven: numberIsEvenNumber(num),
		})
	}
	return
}
func (server *TellEvenServer) IsEvenServerStreaming(req *pb.Numbers, stream pb.TellEvenNumberService_IsEvenServerStreamingServer) error {
	for _, pNumber := range req.Nums {
		num := pNumber.Num
		respElem := &pb.NumberIsEven{
			Num:    num,
			IsEven: numberIsEvenNumber(num),
		}

		if err := stream.Send(respElem); err != nil {
			return err
		}
		//通过 stream.Context() 感知 context 是否被取消了
	}
	return nil
}
func (server *TellEvenServer) IsEvenClientStreaming(stream pb.TellEvenNumberService_IsEvenClientStreamingServer) error {
	resp := &pb.NumberIsEvenList{NumIsEvenList: nil}
	for {
		reqElem, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(resp)
		}
		if err != nil {
			return err
		}

		num := reqElem.Num
		resp.NumIsEvenList = append(resp.NumIsEvenList, &pb.NumberIsEven{
			Num:    num,
			IsEven: numberIsEvenNumber(num),
		})
	}
}
func (server *TellEvenServer) IsEvenBidiStreaming(stream pb.TellEvenNumberService_IsEvenBidiStreamingServer) error {
	for {
		reqElem, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		respElem := &pb.NumberIsEven{
			Num:    reqElem.Num,
			IsEven: numberIsEvenNumber(reqElem.Num),
		}

		if err := stream.Send(respElem); err != nil {
			return err
		}
		//通过 stream.Context() 感知 context 是否被取消了
	}
}

func (server *TellEvenServer) HeartBeat(ctx context.Context, req *pb.HeartBeatPing) (resp *pb.HeartBeatPong, err error) {
	//Ping     string  `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
	//PingNote *string `protobuf:"bytes,2,opt,name=ping_note,json=pingNote,proto3,oneof" json:"ping_note,omitempty"`
	pp.Println("test ping and ping_note:")
	pp.Println(req.Ping)
	pp.Println(req.PingNote)
	if req.PingNote != nil {
		pp.Println("ping_note field is present")
	} else {
		pp.Println("ping_note field is not present")
	}
	resp = &pb.HeartBeatPong{
		Pong:     "pong",
		PongNote: req.PingNote,
	}
	return
}

func numberIsEvenNumber(num int64) bool {
	return num%2 == 0
}

func newBizServer() *TellEvenServer {
	return &TellEvenServer{}
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:5001")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTellEvenNumberServiceServer(grpcServer, newBizServer())
	log.Fatalln(grpcServer.Serve(listener))
}

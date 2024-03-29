package main

import (
	"context"
	"github.com/k0kubun/pp/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-demo/pkg"
	pb "grpc-demo/tell"
	"io"
	"log"
	"time"
)

func runFirst(client pb.TellEvenNumberServiceClient) {
	pNumberIsEven, err := client.IsEven(context.Background(), &pb.Number{
		Num: 42,
	})
	if err != nil {
		log.Fatalln(err)
	}
	pkg.PrintProtoMessage(pNumberIsEven)

	req := &pb.Numbers{
		Nums: nil,
	}
	for i := 0; i < 10; i++ {
		req.Nums = append(req.Nums, &pb.Number{
			Num: int64(i),
		})
	}
	pNumberIsEvenList, err := client.IsEvenUsingList(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	pkg.PrintProtoMessage(pNumberIsEvenList)
}

func runSecond(client pb.TellEvenNumberServiceClient) {
	req := &pb.Numbers{
		Nums: nil,
	}
	for i := 0; i < 10; i++ {
		req.Nums = append(req.Nums, &pb.Number{
			Num: int64(i),
		})
	}
	serverStream, err := client.IsEvenServerStreaming(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		pNumberIsEven, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		pkg.PrintProtoMessage(pNumberIsEven)
	}
}

func runThird(client pb.TellEvenNumberServiceClient) {
	clientStream, err := client.IsEvenClientStreaming(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < 10; i++ {
		if err := clientStream.Send(&pb.Number{Num: int64(i)}); err != nil {
			log.Fatalln(err)
		}
	}
	pNumberIsEvenList, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	pkg.PrintProtoMessage(pNumberIsEvenList)
}

func runFourth(client pb.TellEvenNumberServiceClient) {
	bidiStream, err := client.IsEvenBidiStreaming(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	waitChan := make(chan struct{})
	go func() {
		for {
			pNumIsEven, err := bidiStream.Recv()
			if err == io.EOF {
				// read done.
				close(waitChan)
				return
			}
			if err != nil {
				log.Fatalln(err)
			}
			pkg.PrintProtoMessage(pNumIsEven)
		}
	}()
	for i := 0; i < 10; i++ {
		if err := bidiStream.Send(&pb.Number{Num: int64(i)}); err != nil {
			log.Fatalln(err)
		}
	}
	bidiStream.CloseSend()
	<-waitChan
}

func runFifth(client pb.TellEvenNumberServiceClient) {
	pHeartBeatPong, err := client.HeartBeat(context.Background(), &pb.HeartBeatPing{
		Ping:     "ping",
		PingNote: nil,
	})
	if err != nil {
		log.Fatalln(err)
	}
	pp.Println("test pong and ping_note:")
	pp.Println(pHeartBeatPong.Pong)
	pp.Println(pHeartBeatPong.PongNote)
	if pHeartBeatPong.PongNote != nil {
		pp.Println("pong_note field is present")
	} else {
		pp.Println("pong_note field is not present")
	}
	pp.Println("again:")
	betterCallSaul := "kim&wexler"
	pHeartBeatPong, err = client.HeartBeat(context.Background(), &pb.HeartBeatPing{
		Ping:     "ping",
		PingNote: &betterCallSaul,
	})
	if err != nil {
		log.Fatalln(err)
	}
	pp.Println("test pong and ping_note:")
	pp.Println(pHeartBeatPong.Pong)
	pp.Println(pHeartBeatPong.PongNote)
	if pHeartBeatPong.PongNote != nil {
		pp.Println("pong_note field is present")
	} else {
		pp.Println("pong_note field is not present")
	}

}

func runSixth(client pb.TellEvenNumberServiceClient) {
	pNumberIsEven, err := client.IsEven(context.Background(), &pb.Number{
		Num: 42,
	})
	if err != nil {
		pp.Println(err)
	} else {
		pkg.PrintProtoMessage(pNumberIsEven)
	}
}

func TestOffLineAndBackOnLineAgain() {
	//"1:"
	//map[string]interface {}{
	//  "isEven": true,
	//  "num":    "42",
	//}
	//"2:"
	//&status.Error{
	//  s: &status.Status{
	//    s: &status.Status{
	//      state: impl.MessageState{
	//        NoUnkeyedLiterals: pragma.NoUnkeyedLiterals{},
	//        DoNotCompare:      pragma.DoNotCompare{},
	//        DoNotCopy:         pragma.DoNotCopy{},
	//        atomicMessageInfo: (*impl.MessageInfo)(nil),
	//      },
	//      sizeCache:     0,
	//      unknownFields: []uint8(nil),
	//      Code:          14,
	//      Message:       "connection error: desc = \"transport: Error while dialing dial tcp [::1]:5001: connect: connection refused\"",
	//      Details:       []*anypb.Any(nil),
	//    },
	//  },
	//}
	//"3:"
	//map[string]interface {}{
	//  "isEven": true,
	//  "num":    "42",
	//}

	conn, err := grpc.Dial("localhost:5001",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewTellEvenNumberServiceClient(conn)

	pp.Println("1:")
	runSixth(client)
	time.Sleep(time.Second * 15)

	pp.Println("2:")
	runSixth(client)
	time.Sleep(time.Second * 15)

	pp.Println("3:")
	runSixth(client)
}

func main() {
	//使用localhost没有问题
	conn, err := grpc.Dial("localhost:5001",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewTellEvenNumberServiceClient(conn)
	//runFirst(client)
	//runSecond(client)
	//runThird(client)
	//runFourth(client)
	runFifth(client)
}

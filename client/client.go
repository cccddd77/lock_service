package main

import (
	"context"
	"flag"
	"log"
	"time"
	"math/rand"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/cccddd77/lock_service/lockservice"
)

var (
	serverAddr         = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

type request struct {
	CliID int64
	Operator string
}

func sendReq(client pb.LockServiceClient) {
	log.Print("Sending requests.")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	lockedClient := int64(0)
	for i := 0; i < 50; i++ {
		req := request{CliID: rand.Int63n(7) + 1}
		if lockedClient != req.CliID {
			req.Operator = "acquire"
		} else {
			switch rand.Intn(2) {
			case 0:
				req.Operator = "acquire"
			case 1:
				req.Operator = "release"
			}
		}

		log.Print("Sends: CliID ", req.CliID, " to ", req.Operator)
		switch req.Operator {
		case "acquire":
			rsp, err := client.DoLock(ctx, &pb.Req{CliID: req.CliID, Operator: req.Operator})
			if err == nil {
				log.Print("Receive: ", rsp.CliID, " ", rsp.Operator, " ", rsp.Msg)
				if rsp.Msg == "success" {
					lockedClient = rsp.CliID
				}
			}
		case "release":
			rsp, err := client.UnLock(ctx, &pb.Req{CliID: req.CliID, Operator: req.Operator})
			if err == nil {
				log.Print("Receive: ", rsp.CliID, " ", rsp.Operator, " ", rsp.Msg)
				if rsp.Msg == "success" {
					lockedClient = 0
				}
			}
		}
	}
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()
	client := pb.NewLockServiceClient(conn)

	sendReq(client)
}

package main

import (
	"context"
	"flag"
	"log"
	"time"

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

var testDate = [7]request{
	{1, "acquire"},
	{2, "release"},
	{3, "acquire"},
	{4, "acquire"},
	{5, "release"},
	{6, "release"},
	{7, "acquire"},
}

func sendReq(client pb.LockServiceClient) {
	log.Print("Sending requests.")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for _, i := range testDate {
		log.Print("Send: ", i.CliID, " ", i.Operator)
		switch i.Operator {
		case "acquire":
			rsp, err := client.DoLock(ctx, &pb.Req{CliID: i.CliID, Operator: i.Operator})
			if err == nil {
				log.Print("Receive: ", rsp.CliID, " ", rsp.Operator, " ", rsp.Msg)
			}
		case "release":
			rsp, err := client.UnLock(ctx, &pb.Req{CliID: i.CliID, Operator: i.Operator})
			if err == nil {
				log.Print("Receive: ", rsp.CliID, " ", rsp.Operator, " ", rsp.Msg)
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

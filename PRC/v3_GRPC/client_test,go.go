package main

import (
	message2 "Devops/PRC/v3_GRPC/message"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
)

func TestClient(t *testing.T) {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	encryptionServiceClient := message2.NewEncryptionServiceClient(conn)
	orderRequest := &message2.EncryptionRequest{Str: "mclik"}
	result, err := encryptionServiceClient.Encryption(context.Background(), orderRequest)
	if result != nil {
		fmt.Println(result.Result)
	}
}

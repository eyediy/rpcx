package main

import (
	"fmt"
	"time"

	"github.com/eyediy/rpcx"
)

type Args struct {
	A int `msg:"a"`
	B int `msg:"b"`
}

type Reply struct {
	C int `msg:"c"`
}

func main() {
	s := &rpcx.DirectClientSelector{Network: "tcp", Address: "127.0.0.1:8972", DialTimeout: 10 * time.Second}
	client := rpcx.NewClient(s)
	client.Timeout = 1 * time.Nanosecond
	client.ReadTimeout = 1 * time.Nanosecond
	client.WriteTimeout = 1 * time.Nanosecond

	args := &Args{7, 8}
	var reply Reply
	divCall := client.Go("Arith.Mul", args, &reply, nil)
	replyCall := <-divCall.Done // will be equal to divCall
	if replyCall.Error != nil {
		fmt.Printf("error for Arith: %d*%d, %v \n", args.A, args.B, replyCall.Error)
	} else {
		fmt.Printf("Arith: %d*%d=%d \n", args.A, args.B, reply.C)
	}

	client.Close()
}

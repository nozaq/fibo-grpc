package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nozaq/fibo-grpc/fiborpc"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

var (
	addr      = flag.String("addr", "localhost:8888", "the server address to connect")
	fiboIndex = flag.Int("index", -1, "the index to be provided to the fibonacci calculation")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to the server: %s", err.Error())
	}
	defer conn.Close()

	client := fiborpc.NewLeonardoClient(conn)
	res, err := client.Fibo(context.Background(), &fiborpc.FiboRequest{
		Index: int32(*fiboIndex),
	})
	if err != nil {
		log.Fatalf("an error occurred during the rpc call: %s", err.Error())
	}

	fmt.Printf("fibo(%d) = %d\n", *fiboIndex, res.Value)
}

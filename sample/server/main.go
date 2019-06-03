package main

import (
	"context"
	"github.com/gorilla/handlers"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"server/gen/go/sample"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"strconv"
)

func main() {
	s := server{}
	grpcServer := grpc.NewServer()
	sample.RegisterTestServiceServer(grpcServer, s)

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = sample.RegisterTestServiceHandlerFromEndpoint(ctx, mux, "localhost:9000", opts)

	// Graceful Shutdown for pure grpc
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		if err := http.ListenAndServe(":9001", handlers.LoggingHandler(os.Stdout, mux)); err != nil {
			log.Fatal(err)
		}
	}()
	<-stop
	grpcServer.GracefulStop()

}

type server struct {
}

func (s server) TestEndpointGet(ctx context.Context, in *sample.TestEndpointRequest) (*sample.TestEndpointResponse, error) {

	num, _ := strconv.Atoi(in.Id)
	return &sample.TestEndpointResponse{Message: &sample.TestMessage{
		Id:         "hello",
		Boolean:    true,
		Int32Type:  int32(num),
		Int64Type:  1,
		Uint32Type: 1,
		Uint64Type: 1,
		Enum:       1,
	}}, nil
}

func (s server) TestEndpointPost(context.Context, *sample.TestEndpointRequest) (*sample.TestEndpointResponse, error) {
	return &sample.TestEndpointResponse{
		Message: &sample.TestMessage{
			Id:         "hello",
			Boolean:    true,
			Int32Type:  1,
			Int64Type:  1,
			Uint32Type: 1,
			Uint64Type: 1,
			Enum:       1,
		},
	}, nil
}

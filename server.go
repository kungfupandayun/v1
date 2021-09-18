package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/bigbluedisco/tech-challenge/backend/v1/order"
	orderpb "github.com/bigbluedisco/tech-challenge/backend/v1/order/rpc"
	"github.com/bigbluedisco/tech-challenge/backend/v1/product"
	productpb "github.com/bigbluedisco/tech-challenge/backend/v1/product/rpc"
	"github.com/bigbluedisco/tech-challenge/backend/v1/store"
	"google.golang.org/grpc"
)

func main() {
	log := log.Default()
	ps := store.NewProductStore()
	ords := store.NewOrderStore()

	lis, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	opts := []grpc.ServerOption{}
	srv := grpc.NewServer(opts...)
	defer srv.Stop()

	prd := product.NewService(ps)
	productpb.RegisterServiceServer(srv, prd)

	ord := order.NewService(ords)
	orderpb.RegisterServiceServer(srv, ord)

	go func() {
		log.Println("server started on port 8000...")
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// wait for control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Println("stopping the server")
}

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	petv1 "github.com/bufbuild/buf-tour/petstore/gen/proto/go/pet/v1"
	petv1connect "github.com/bufbuild/buf-tour/petstore/gen/proto/go/pet/v1/petv1connect"
	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

func main() {
	if false {

		if err := run(); err != nil {
			log.Fatal(err)
		}
	} else {
		petter := &PetServer{}
		mux := http.NewServeMux()
		path, handler := petv1connect.NewPetStoreServiceHandler(petter)
		mux.Handle(path, handler)
		fmt.Println("starting server on", path)
		http.ListenAndServe(
			"localhost:8080",
			h2c.NewHandler(mux, &http2.Server{}),
		)
	}
}

func run() error {
	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	petv1.RegisterPetStoreServiceServer(server, &petStoreServiceServer{})
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

// petStoreServiceServer implements the PetStoreService API.
type petStoreServiceServer struct {
	petv1.UnimplementedPetStoreServiceServer
}

type PetServer struct{}

func (s *PetServer) PutPet(
	ctx context.Context,
	req *connect.Request[petv1.PutPetRequest],
) (*connect.Response[petv1.PutPetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&petv1.PutPetResponse{})
	res.Header().Set("PetV", "v1")
	return res, nil
}
func (s *PetServer) GetPet(
	ctx context.Context,
	req *connect.Request[petv1.GetPetRequest],
) (*connect.Response[petv1.GetPetResponse], error) {
	log.Println("Request headers: ", req.Header())
	log.Println("GET, ", req.Msg)
	res := connect.NewResponse(&petv1.GetPetResponse{})
	res.Header().Set("PetV", "v1")
	return res, nil
}
func (s *PetServer) DeletePet(
	ctx context.Context,
	req *connect.Request[petv1.DeletePetRequest],
) (*connect.Response[petv1.DeletePetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&petv1.DeletePetResponse{})
	res.Header().Set("PetV", "v1")
	return res, nil
}

// PutPet adds the pet associated with the given request into the PetStore.
func (s *petStoreServiceServer) PutPet(ctx context.Context, req *petv1.PutPetRequest) (*petv1.PutPetResponse, error) {
	name := req.GetName()
	petType := req.GetPetType()
	log.Println("Got a request to create a", petType, "named", name)

	return &petv1.PutPetResponse{}, nil
}

package main

import (
	"context"
	"log"
	"net/http"
	"sync/atomic"

	pb "comp-project/gen/go/proto/counter"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var counter int32

type server struct {
	pb.UnimplementedCounterServiceServer
}

func (s *server) IncrementCounter(ctx context.Context, req *pb.IncrementRequest) (*pb.CounterResponse, error) {
	newValue := atomic.AddInt32(&counter, 1)
	return &pb.CounterResponse{Value: newValue}, nil
}

func main() {
	// Create gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterCounterServiceServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	// Create gRPC-Web wrapper
	wrappedGrpc := grpcweb.WrapServer(grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
		grpcweb.WithWebsockets(true),
		grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool { return true }),
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
	)

	// Create HTTP server
	httpServer := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set CORS headers for all responses
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Max-Age", "86400")

			// Handle preflight requests
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			// Handle gRPC-Web requests (including WebSocket upgrades)
			if wrappedGrpc.IsGrpcWebRequest(r) || wrappedGrpc.IsAcceptableGrpcCorsRequest(r) || wrappedGrpc.IsGrpcWebSocketRequest(r) {
				wrappedGrpc.ServeHTTP(w, r)
				return
			}

			// Serve static files for the Svelte app
			http.FileServer(http.Dir("svelte-client/public")).ServeHTTP(w, r)
		}),
	}

	// Start HTTP server
	log.Printf("Starting server on :8080")
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

package server

import (
	"context"
	"log"
	"time"

	"github.com/gopherland/labs2/grpc/internal/generated"
	"google.golang.org/grpc"
)

// Logger intercepts and logs a request.
func Logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	<<!!YOUR_CODE!!>> -- Define a logger interceptor to print the book and word from the request
}

// Measure intercepts and times a request.
func Measure(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	<<!!YOUR_CODE!!>> -- Define your measure interceptor to compute how long the call duration
}

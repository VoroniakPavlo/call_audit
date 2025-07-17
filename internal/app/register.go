package app

import (
	"google.golang.org/grpc"
)

// serviceRegistration holds information for initializing and registering a gRPC service.
type serviceRegistration struct {
	init     func(*App) (interface{}, error)                    // Initialization function for *App
	register func(grpcServer *grpc.Server, service interface{}) // Registration function for gRPC server
	name     string                                             // Service name for logging
}

// RegisterServices initializes and registers all necessary gRPC services.
func RegisterServices(grpcServer *grpc.Server, appInstance *App) {
	//TODO: Add more services as needed
	return
}

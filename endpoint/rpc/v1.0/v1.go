package v1

import (
	"google.golang.org/grpc"

	"clean_arch/domain/usecase"

	"clean_arch/endpoint/rpc/v1.0/protocol"
)

// Apply -
func Apply(server *grpc.Server, service usecase.UserUsecase) {
	protocol.RegisterUserServiceServer(server, NewUserService(service))
}

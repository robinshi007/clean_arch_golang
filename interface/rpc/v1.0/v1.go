package v1

import (
	"google.golang.org/grpc"

	"clean_arch/usecase"

	"clean_arch/interface/rpc/v1.0/protocol"
)

// Apply -
func Apply(server *grpc.Server, service usecase.UserUsecase) {
	protocol.RegisterUserServiceServer(server, NewUserService(service))
}

package rpc

import (
	"google.golang.org/grpc"

	"clean_arch/domain/usecase"

	"clean_arch/interface/rpc/v1.0"
)

// Apply -
func Apply(server *grpc.Server, service usecase.UserUsecase) {
	v1.Apply(server, service)
}

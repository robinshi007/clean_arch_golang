package v1

import (
	"context"
	"strconv"

	"clean_arch/usecase"
	"clean_arch/usecase/output"

	"clean_arch/interface/rpc/v1.0/protocol"
)

type userService struct {
	userUsecase usecase.UserUsecase
}

// NewUserService -
func NewUserService(userUsecase usecase.UserUsecase) *userService {
	return &userService{
		userUsecase: userUsecase,
	}
}

func (s *userService) ListUser(ctx context.Context, in *protocol.ListUserRequestType) (*protocol.ListUserResponseType, error) {
	users, err := s.userUsecase.GetAll(ctx, 10)
	if err != nil {
		return nil, err
	}
	res := &protocol.ListUserResponseType{
		Users: toUser(users),
	}
	return res, nil
}

func toUser(users []*output.User) []*protocol.User {
	res := make([]*protocol.User, len(users))
	for i, user := range users {
		res[i] = &protocol.User{
			Id:   strconv.FormatInt(user.ID, 10),
			Name: user.Name,
		}
	}
	return res
}

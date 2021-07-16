package implementation

import (
	"context"
	pb "user-service/infrastructure/grpc/service/user_service"
	"user-service/usecase"
)

type UserServiceImpl struct{
	pb.UnimplementedUserDetailsServer
	ProfileInfoUsecase usecase.ProfileInfoUseCase
}

func NewUserServiceImpl(profileInfoUsecase usecase.ProfileInfoUseCase) *UserServiceImpl {
	return &UserServiceImpl{ProfileInfoUsecase: profileInfoUsecase}
}

func (u *UserServiceImpl) GetUsername(ctx context.Context, in *pb.UserId) (*pb.Username, error) {

	user, err := u.ProfileInfoUsecase.GetById(in.UserId, ctx)

	if err != nil {
		return nil, err
	}

	return &pb.Username{Username: user.Profile.Username}, nil
}

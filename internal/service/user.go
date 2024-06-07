package service

import (
	"context"

	pb "kratos-test/api/realworld/v1"
)

func (s *RealworldService) Login(ctx context.Context, req *pb.LoginReq) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *RealworldService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	u, err := s.uc.Registry(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResp{
		User: &pb.RegisterResp_User{
			Email:    u.Email,
			Username: u.Username,
			Token:    u.Token,
		},
	}, nil
}
func (s *RealworldService) CurrentUser(ctx context.Context, req *pb.Empty) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *RealworldService) UpdateUser(ctx context.Context, req *pb.UserReq) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}

func (s *RealworldService) GetProfile(ctx context.Context, req *pb.GetProfileReq) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
func (s *RealworldService) FollowProfile(ctx context.Context, req *pb.GetProfileReq) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
func (s *RealworldService) UnfollowProfile(ctx context.Context, req *pb.GetProfileReq) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}

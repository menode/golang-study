package service

import (
	"context"

	pb "kratos-test/api/realworld/v1"
	"kratos-test/internal/biz"
)

func (s *RealworldService) Login(ctx context.Context, req *pb.LoginReq) (*pb.UserReply, error) {
	u, err := s.uc.Login(ctx, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &pb.UserReply{
		User: &pb.UserReply_User{
			Email:    u.Email,
			Username: u.Username,
			Token:    u.Token,
			Image:    u.Image,
			Bio:      u.Bio,
		},
	}, nil
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
	u, err := s.uc.CurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.UserReply{
		User: &pb.UserReply_User{
			Email:    u.Email,
			Username: u.Username,
			Image:    u.Image,
			Bio:      u.Bio,
		},
	}, nil
}
func (s *RealworldService) UpdateUser(ctx context.Context, req *pb.UserReq) (*pb.UserReply, error) {
	u, err := s.uc.UpdateUser(ctx, &biz.UserUpdate{
		Email:    req.User.Email,
		Username: req.User.Username,
		Password: req.User.Password,
		Image:    req.User.Image,
		Bio:      req.User.Bio,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UserReply{
		User: &pb.UserReply_User{
			Email:    u.Email,
			Username: u.Username,
			Image:    u.Image,
			Bio:      u.Bio,
		},
	}, nil
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

package service

import (
	"context"

	v1 "kratos-test/api/realworld/v1"
	"kratos-test/internal/biz"
)

// GreeterService is a greeter service.
type RealWorldService struct {
	v1.UnimplementedRealworldServer

	uc *biz.GreeterUsecase
}

// NewRealworldService new a greeter service.
func NewRealworldService(uc *biz.GreeterUsecase) *RealWorldService {
	return &RealWorldService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *RealWorldService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

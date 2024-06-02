package service

import (
	pb "kratos-test/api/realworld/v1"
	"kratos-test/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealworldService)

type RealworldService struct {
	pb.UnimplementedRealworldServer
	uc *biz.UserUsecase

	log *log.Helper
}

func NewRealworldService(uc *biz.UserUsecase) *RealworldService {
	return &RealworldService{uc: uc, log: log.NewHelper(log.DefaultLogger)}
}

package data

import (
	"context"

	"kratos-test/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	gorm.Model
	Email     string `gorm:"size:500"`
	Username  string `gorm:"size:500"`
	Bio       string `gorm:"size:500"`
	Image     string `gorm:"size:500"`
	Password  string `gorm:"size:500"`
	Following uint32
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	user := User{
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
		Bio:      u.Bio,
		Image:    u.Image,
	}
	rv := r.data.db.Where(User{Email: u.Email}).Or(User{Username: u.Username}).FirstOrCreate(&user)
	return rv.Error
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.UserLogin, error) {
	var user User
	rv := r.data.db.Where("email = ?", email).First(&user)

	if errors.Is(rv.Error, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("user", email)
	}

	if rv.Error != nil {
		return nil, errors.New(400, "not found", rv.Error.Error())
	}
	return &biz.UserLogin{
		Email:        user.Email,
		Username:     user.Username,
		PasswordHash: user.Password,
		Image:        user.Image,
		Bio:          user.Bio,
	}, nil
}

type profileRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

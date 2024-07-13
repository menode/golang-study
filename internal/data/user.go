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

type profileRepo struct {
	data *Data
	log  *log.Helper
}

type FollowUser struct {
	gorm.Model
	UserID      uint32
	FollowingID uint32
}

type User struct {
	gorm.Model
	Email        string `gorm:"size:500"`
	Username     string `gorm:"size:500"`
	Bio          string `gorm:"size:500"`
	Image        string `gorm:"size:500"`
	PasswordHash string `gorm:"size:500"`
	Following    uint32
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// NewGreeterRepo .
func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	user := User{
		Email:        u.Email,
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
		Bio:          u.Bio,
		Image:        u.Image,
	}
	rv := r.data.db.Where(User{Email: u.Email}).Or(User{Username: u.Username}).FirstOrCreate(&user)
	return rv.Error
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	var user User
	rv := r.data.db.Where("email = ?", email).First(&user)

	if errors.Is(rv.Error, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("user", email)
	}

	if rv.Error != nil {
		return nil, errors.New(400, "not found", rv.Error.Error())
	}
	return &biz.User{
		ID:           user.ID,
		Email:        user.Email,
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
		Image:        user.Image,
		Bio:          user.Bio,
	}, nil
}

func (r *userRepo) GetUserByID(ctx context.Context, id uint) (*biz.User, error) {
	var user User
	rv := r.data.db.Where("id = ?", id).First(&user)

	if errors.Is(rv.Error, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("user", string(id))
	}

	if rv.Error != nil {
		return nil, errors.New(400, "not found", rv.Error.Error())
	}
	return &biz.User{
		Email:        user.Email,
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
		Image:        user.Image,
		Bio:          user.Bio,
	}, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	user := User{
		Model: gorm.Model{
			ID: u.ID,
		},
		Email:        u.Email,
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
		Bio:          u.Bio,
		Image:        u.Image,
	}
	rv := r.data.db.Save(&user)
	return &biz.User{
		Email:        user.Email,
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
		Image:        user.Image,
		Bio:          user.Bio,
		ID:           user.ID,
	}, rv.Error
}

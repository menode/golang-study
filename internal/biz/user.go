package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}
type UserLogin struct {
	Email    string
	Username string
	Token    string
	Image    string
	Bio      string
}

type UserUpdate struct {
	Email    string
	Username string
	Password string
	Bio      string
	Image    string
}

// password 生成hash
func hashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// 哈希和密码对比
func verifyPassword(hashed, input string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input)); err != nil {
		return false
	}
	return true
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	// GetUserByEmail(ctx context.Context, email string) (*User, error)
	// GetUserByUsername(ctx context.Context, username string) (*User, error)
	// GetUserByID(ctx context.Context, id uint) (*User, error)
	// UpdateUser(ctx context.Context, user *User) error
}

type ProfileRepo interface {
}

type UserUsecase struct {
	ur UserRepo
	pr ProfileRepo

	log *log.Helper
}

func NewUserUsecase(ur UserRepo, pr ProfileRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, pr: pr, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) generateToken(userID uint) string {
	// return auth.generateToken(userID)
	return string(userID)
}
func (uc *UserUsecase) CreateUser(ctx context.Context, u *User) error {
	if err := uc.ur.CreateUser(ctx, u); err != nil {
		return err
	}
	return nil
}
func (uc *UserUsecase) Registry(ctx context.Context, username, email, password string) (*UserLogin, error) {
	u := &User{
		Email:        email,
		Username:     username,
		PasswordHash: hashPassword(password),
	}
	if err := uc.ur.CreateUser(ctx, u); err != nil {
		return nil, err
	}
	return &UserLogin{
		Email:    email,
		Username: username,
		Token:    uc.generateToken(u.ID),
	}, nil
}

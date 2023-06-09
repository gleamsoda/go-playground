package usecase

import (
	"context"
	"fmt"
	"time"

	"playground/config"
	"playground/domain"
	"playground/internal"
	"playground/internal/token"
)

type UserUsecase struct {
	userRepo    domain.UserRepository
	sessionRepo domain.SessionRepository
	tokenMaker  token.Maker
	cfg         config.Config
}

var _ domain.UserUsecase = (*UserUsecase)(nil)

func NewUserUsecase(userRepo domain.UserRepository, sessionRepo domain.SessionRepository, tokenMaker token.Maker, cfg config.Config) domain.UserUsecase {
	return &UserUsecase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		tokenMaker:  tokenMaker,
		cfg:         cfg,
	}
}

func (u *UserUsecase) Create(ctx context.Context, arg domain.CreateUserInputParams) (*domain.User, error) {
	hashedPassword, err := internal.HashPassword(arg.Password)
	if err != nil {
		return nil, err
	}
	usr := domain.NewUser(
		arg.Username,
		arg.FullName,
		arg.Email,
		hashedPassword,
	)
	return u.userRepo.Create(ctx, usr)
}

func (u *UserUsecase) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	return u.userRepo.GetByUsername(ctx, username)
}

func (u *UserUsecase) Login(ctx context.Context, arg domain.LoginUserInputParams) (*domain.LoginUserOutputParams, error) {
	usr, err := u.userRepo.GetByUsername(ctx, arg.Username)
	if err != nil {
		return nil, err
	}
	if err := internal.CheckPassword(arg.Password, usr.HashedPassword); err != nil {
		return nil, err
	}

	aToken, aPayload, err := u.tokenMaker.CreateToken(
		usr.ID,
		u.cfg.AccessTokenDuration,
	)
	if err != nil {
		return nil, err
	}

	rToken, rPayload, err := u.tokenMaker.CreateToken(
		usr.ID,
		u.cfg.RefreshTokenDuration,
	)
	if err != nil {
		return nil, err
	}

	if err := u.sessionRepo.Create(ctx, domain.NewSession(
		rPayload.ID,
		usr.ID,
		rToken,
		arg.UserAgent,
		arg.ClientIP,
		false,
		rPayload.ExpiredAt,
	)); err != nil {
		return nil, err
	}

	return &domain.LoginUserOutputParams{
		SessionID:             rPayload.ID,
		AccessToken:           aToken,
		AccessTokenExpiresAt:  aPayload.ExpiredAt,
		RefreshToken:          rToken,
		RefreshTokenExpiresAt: rPayload.ExpiredAt,
		User:                  usr,
	}, nil
}

func (u *UserUsecase) RenewAccessToken(ctx context.Context, refreshToken string) (*domain.RenewAccessTokenOutputParams, error) {
	rPayload, err := u.tokenMaker.VerifyToken(refreshToken)
	if err != nil {
		return nil, err
	}

	sess, err := u.sessionRepo.Get(ctx, rPayload.ID)
	if err != nil {
		return nil, err
	}
	if sess.IsBlocked {
		err := fmt.Errorf("blocked session")
		return nil, err
	}
	if sess.UserID != rPayload.UserID {
		err := fmt.Errorf("incorrect session user")
		return nil, err
	}
	if sess.RefreshToken != refreshToken {
		err := fmt.Errorf("mismatched session token")
		return nil, err
	}
	if time.Now().After(sess.ExpiresAt) {
		err := fmt.Errorf("expired session")
		return nil, err
	}

	aToken, aPayload, err := u.tokenMaker.CreateToken(
		rPayload.UserID,
		u.cfg.AccessTokenDuration,
	)
	if err != nil {
		return nil, err
	}

	resp := &domain.RenewAccessTokenOutputParams{
		AccessToken:          aToken,
		AccessTokenExpiresAt: aPayload.ExpiredAt,
	}
	return resp, nil
}

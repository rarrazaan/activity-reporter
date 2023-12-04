package repository

import (
	"context"
	"fmt"
	"mini-socmed/internal/cons"
	"mini-socmed/internal/dependency"
	"time"

	"github.com/redis/go-redis/v9"
)

type (
	RedisRepo interface {
		SetRefreshToken(ctx context.Context, rToken string, userID int64) error
		GetUserIDByRefreshToken(ctx context.Context, rToken string) (*int64, error)
		DeleteRefreshToken(ctx context.Context, refreshToken string) error
		SetVerifiyEmail(ctx context.Context, code string, userID int64) error
		GetUserIDByVerifiyEmail(ctx context.Context, code string) (*int64, error)
		DeleteVerifiyEmail(ctx context.Context, refreshToken string) error
	}
	redisRepo struct {
		cfg dependency.Config
		rd  *redis.Client
	}
)

func (r *redisRepo) SetRefreshToken(ctx context.Context, rToken string, userID int64) error {
	key := fmt.Sprintf(cons.RedisRefreshTokenTemplate, rToken)
	expiration := time.Duration(r.cfg.Jwt.RefreshTokenExpiration) * time.Minute

	cmd := r.rd.SetEx(ctx, key, userID, expiration)
	if err := cmd.Err(); err != nil {
		return err
	}

	return nil
}

func (r *redisRepo) GetUserIDByRefreshToken(ctx context.Context, rToken string) (*int64, error) {
	key := fmt.Sprintf(cons.RedisRefreshTokenTemplate, rToken)

	cmd := r.rd.Get(ctx, key)
	if cmd == nil {
		return nil, nil
	}

	if err := cmd.Err(); err != nil {
		return nil, err
	}

	id, err := cmd.Int64()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (r *redisRepo) DeleteRefreshToken(ctx context.Context, refreshToken string) error {
	key := fmt.Sprintf(cons.RedisRefreshTokenTemplate, refreshToken)
	cmd := r.rd.Del(ctx, key)
	if err := cmd.Err(); err != nil {
		return err
	}

	return nil
}

func (r *redisRepo) SetVerifiyEmail(ctx context.Context, code string, userID int64) error {
	key := fmt.Sprintf(cons.RedisVerifyEmailCodeTemplate, code)
	expiration := time.Duration(r.cfg.VerifyEmail.VerifyEmailCodeExpiration) * time.Minute

	cmd := r.rd.SetEx(ctx, key, userID, expiration)
	if err := cmd.Err(); err != nil {
		return err
	}

	return nil
}
func (r *redisRepo) GetUserIDByVerifiyEmail(ctx context.Context, code string) (*int64, error) {
	key := fmt.Sprintf(cons.RedisVerifyEmailCodeTemplate, code)

	cmd := r.rd.Get(ctx, key)
	if cmd == nil {
		return nil, nil
	}

	if err := cmd.Err(); err != nil {
		return nil, err
	}

	id, err := cmd.Int64()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (r *redisRepo) DeleteVerifiyEmail(ctx context.Context, refreshToken string) error {
	key := fmt.Sprintf(cons.RedisVerifyEmailCodeTemplate, refreshToken)
	cmd := r.rd.Del(ctx, key)
	if err := cmd.Err(); err != nil {
		return err
	}

	return nil
}

func NewRedisRepo(cfg dependency.Config, rd *redis.Client) RedisRepo {
	return &redisRepo{
		cfg: cfg,
		rd:  rd,
	}
}

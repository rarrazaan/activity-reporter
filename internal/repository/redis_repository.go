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
	}
	redisRepo struct {
		cfg dependency.Config
		rd  *redis.Client
	}
)

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

func (r *redisRepo) SetRefreshToken(ctx context.Context, rToken string, userID int64) error {
	key := fmt.Sprintf(cons.RedisRefreshTokenTemplate, rToken)
	expiration := time.Duration(r.cfg.Jwt.RefreshTokenExpiration) * time.Minute

	cmd := r.rd.SetEx(ctx, key, userID, expiration)
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

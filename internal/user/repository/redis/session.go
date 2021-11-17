package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Alexander272/astro-atlas/pkg/logger"
	"github.com/go-redis/redis/v8"
)

type SessionRepo struct {
	db redis.Cmdable
}

func NewSessionRepo(db redis.Cmdable) *SessionRepo {
	return &SessionRepo{db: db}
}

type SessionData struct {
	UserId string
	Email  string
	Role   string
	Ua     string
	Ip     string
	Exp    time.Duration
}

func (d SessionData) MarshalBinary() ([]byte, error) {
	return json.Marshal(d)
}

// func unmarshalBinary(str string) (s *SessionData) {
// 	json.Unmarshal([]byte(str), &s)
// 	return s
// }

func (r *SessionRepo) Create(ctx context.Context, token string, data SessionData) error {
	if strings.Trim(token, " ") == "" {
		return errors.New("empty token")
	}
	if (data == SessionData{}) {
		return errors.New("empty user data")
	}

	res := r.db.Set(ctx, token, data, data.Exp)
	if res.Err() != nil {
		return fmt.Errorf("failed to execute query. error: %w", res.Err())
	}
	logger.Debug(res.Result())

	return nil
}

func (r *SessionRepo) GetDel(ctx context.Context, key string) (data SessionData, err error) {
	if strings.Trim(key, " ") == "" {
		return data, errors.New("empty key")
	}

	cmd := r.db.GetDel(ctx, key)
	if cmd.Err() != nil {
		return data, fmt.Errorf("failed to execute query. error: %w", cmd.Err())
	}

	if err := cmd.Scan(&data); err != nil {
		// todo дописать ошибку
		return data, fmt.Errorf("failed to ... . error: %w", err)
	}
	logger.Info(data)
	return data, nil
}

func (r *SessionRepo) Delete(ctx context.Context, key string) error {
	if strings.Trim(key, " ") == "" {
		return errors.New("empty key")
	}

	res := r.db.Del(ctx, key)
	if res.Err() != nil {
		return fmt.Errorf("failed to execute query. error: %w", res.Err())
	}
	logger.Debug(res)
	return nil
}

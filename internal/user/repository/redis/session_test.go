package redis

import (
	"context"
	"os"
	"testing"
	"time"

	redisDB "github.com/Alexander272/astro-atlas/pkg/database/redis"
	"github.com/Alexander272/astro-atlas/pkg/logger"
	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock/v8"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

var client *redis.Client

func TestMain(m *testing.M) {
	mr, err := miniredis.Run()
	if err != nil {
		logger.Fatalf("failed to open miniredis. error %s", err.Error())
	}
	defer mr.Close()

	logger.Debug(mr.Addr())
	client, err = redisDB.NewRedisClient(redisDB.Config{
		Host: mr.Host(),
		Port: mr.Port(),
	})
	if err != nil {
		logger.Fatalf("failed to initialize redis %s", err.Error())
	}

	code := m.Run()
	os.Exit(code)
}

func TestSessionRepository_Create(t *testing.T) {
	testCases := []struct {
		name       string
		testToken  string
		testUserId string
		testExp    time.Duration
		isValid    bool
	}{
		{
			name:       "full data",
			testToken:  "qwerty1",
			testUserId: "userId",
			isValid:    true,
		},
		{
			name:      "only key",
			testToken: "qwerty2",
			isValid:   false,
		},
		{
			name:      "empty key",
			testToken: " ",
			isValid:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := redismock.NewNiceMock(client)
			mock.On("Set", context.Background(), tc.testToken, SessionData{UserId: tc.testUserId, Exp: tc.testExp}, tc.testExp).Return(redis.NewStatusResult("", nil))

			r := NewSessionRepo(mock)
			err := r.Create(context.Background(), tc.testToken, SessionData{UserId: tc.testUserId, Exp: tc.testExp})

			if tc.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

// func TestSessionRepository_GetDel(t *testing.T) {
// 	// mock := redismock.NewNiceMock(client)
// 	// mock.On("GetDel", context.Background(), testToken).Return(redis.NewStringResult("", nil))

// 	r := NewSessionRepo(client)
// 	r.Create(context.Background(), testToken, SessionData{UserId: testUserId, Exp: testExp})

// 	res, err := r.GetDel(context.Background(), testToken)
// 	assert.NoError(t, err)
// 	assert.Equal(t, SessionData{UserId: testUserId, Exp: testExp}, res)
// }

func TestSessionRepository_Delete(t *testing.T) {
	testCases := []struct {
		name      string
		testToken string
		isValid   bool
	}{
		{
			name:      "correct key",
			testToken: "qwerty",
			isValid:   true,
		},
		{
			name:      "incorrect key",
			testToken: " ",
			isValid:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// mock := redismock.NewNiceMock(client)
			// mock.On("Del", context.Background(), tc.testToken).Return(redis.NewStatusResult("", nil))

			r := NewSessionRepo(client)
			err := r.Delete(context.Background(), tc.testToken)

			if tc.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

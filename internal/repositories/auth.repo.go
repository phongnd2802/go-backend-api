package repositories

import (
	"database/sql"
	"github.com/phongnd2802/go-backend-api/global"
	"github.com/phongnd2802/go-backend-api/internal/database"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type IAuthRepository interface {
	GetAPIKeyDB(apiKey string) (*database.ApiKey, error)
	CreateAPIKey(apiKey string, desc string) error

	// Redis
	GetAPIKeyR(apiKey string) bool
	SetAPIKeyR(apiKey string) error

}

type authRepository struct {
	sqlc *database.Queries
}


// SetAPIKeyR implements IAuthRepository.
func (ar *authRepository) SetAPIKeyR(apiKey string) error {
	return global.Rdb.Set(ctx, apiKey, 1, 0).Err()
}

// GetAPIKeyR implements IAuthRepository.
func (ar *authRepository) GetAPIKeyR(apiKey string) bool {
	_, err := global.Rdb.Get(ctx, apiKey).Result()
	if err == redis.Nil {
		// Key does not exist
		return false
	} else if err != nil {
		// Handle other errors, such as connection issues
		global.Logger.Error("Error retrieving API key from Redis:", zap.Error(err))
		return false
	}
	return true
}

// CreateAPIKey implements IAuthRepository.
func (ar *authRepository) CreateAPIKey(apiKey string, desc string) error {
	err := ar.sqlc.CreateAPIKey(ctx, database.CreateAPIKeyParams{
		ApiKey: apiKey,
		Description: sql.NullString{
			String: desc,
			Valid:  true,
		},
	})

	return err
}

// GetAPIKey implements IAuthRepository.
func (ar *authRepository) GetAPIKeyDB(apiKey string) (*database.ApiKey, error) {
	obj, err := ar.sqlc.GetAPIKey(ctx, apiKey)
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func NewAuthRepository() IAuthRepository {
	return &authRepository{
		sqlc: database.New(global.Mdb),
	}
}

package repositories

import (
	"database/sql"

	"github.com/phongnd2802/go-backend-api/global"
	"github.com/phongnd2802/go-backend-api/internal/database"
)

type IAuthRepository interface {
	GetAPIKey(apiKey string) (*database.ApiKey, error)
	CreateAPIKey(apiKey string, desc string) error
}

type authRepository struct {
	sqlc *database.Queries
}

// CreateAPIKey implements IAuthRepository.
func (ar *authRepository) CreateAPIKey(apiKey string, desc string) error {
	err := ar.sqlc.CreateAPIKey(ctx, database.CreateAPIKeyParams{
		ApiKey: apiKey,
		Description: sql.NullString{
			String: desc,
			Valid: true,
		},
	})

	return err
}

// GetAPIKey implements IAuthRepository.
func (ar *authRepository) GetAPIKey(apiKey string) (*database.ApiKey, error) {
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

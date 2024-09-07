package repositories

import (
	"fmt"
	"time"

	"github.com/phongnd2802/go-backend-api/global"
	"github.com/phongnd2802/go-backend-api/internal/database"
)

type ITokenRepository interface {
	CreateToken(id, publicKey, refreshToken, userId string) error
	GetTokenByUserID(userID string) (*database.Token, error)

	SetAccessToken(accessToken string, userID string, expirationTime int64) error
	GetAccessToken(userID string) (string, error)
}

type tokenRepository struct {
	sqlc *database.Queries
}

// GetTokenByUserID implements ITokenRepository.
func (tr *tokenRepository) GetTokenByUserID(userID string) (*database.Token, error) {
	result, err := tr.sqlc.GetTokenByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &result, err
}

// GetAccessToken implements ITokenRepository.
func (tr *tokenRepository) GetAccessToken(userID string) (string, error) {
	key := fmt.Sprintf("accessToken:%s", userID)
	accessToken, err := global.Rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

// SetAccessToken implements ITokenRepository.
func (tr *tokenRepository) SetAccessToken(accessToken string, userID string, expirationTime int64) error {
	key := fmt.Sprintf("accessToken:%s", userID)
	return global.Rdb.SetEx(ctx, key, accessToken, time.Duration(expirationTime)).Err()
}

// CreateToken implements ITokenRepository.
func (tr *tokenRepository) CreateToken(id string, publicKey string, refreshToken string, userId string) error {
	err := tr.sqlc.CreateToken(ctx, database.CreateTokenParams{
		ID:           id,
		PublicKey:    publicKey,
		RefreshToken: refreshToken,
		UserID:       userId,
	})

	if err != nil {
		return err
	}

	return nil
}

func NewTokenRepository() ITokenRepository {
	return &tokenRepository{
		sqlc: database.New(global.Mdb),
	}
}

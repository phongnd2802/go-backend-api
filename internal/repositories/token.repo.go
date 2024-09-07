package repositories

import (
	"github.com/phongnd2802/go-backend-api/global"
	"github.com/phongnd2802/go-backend-api/internal/database"
)

type ITokenRepository interface {
	CreateToken(id, publicKey, refreshToken, userId string) error
}

type tokenRepository struct {
	sqlc *database.Queries
}

// CreateToken implements ITokenRepository.
func (tr *tokenRepository) CreateToken(id string, publicKey string, refreshToken string, userId string) error {
	err := tr.sqlc.CreateToken(ctx, database.CreateTokenParams{
		ID: id,
		PublicKey: publicKey,
		RefreshToken: refreshToken,
		UserID: userId,
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

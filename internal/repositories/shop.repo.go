package repositories

import (
	"github.com/phongnd2802/go-backend-api/global"
	"github.com/phongnd2802/go-backend-api/internal/database"
)

type IShopRepository interface {
	GetShopByEmail(email string) (*database.Shop, error)
	CreateShop(id, name, email, password string) error
	ActiveShopOTP(email string) error
	UpdatePassword(email string, newPassword string) error
}

type shopRepository struct {
	sqlc *database.Queries
}

// UpdatePassword implements IShopRepository.
func (sr *shopRepository) UpdatePassword(email string, newPassword string) error {
	err := sr.sqlc.UpdatePassword(ctx, database.UpdatePasswordParams{
		ShopEmail: email,
		ShopPassword: newPassword,
	})
	if err != nil {
		return err
	}

	return nil
}

func (sr *shopRepository) ActiveShopOTP(email string) error {
	err := sr.sqlc.ActiveShopOTP(ctx, email)
	if err != nil {
		return err
	}
	return nil
}

func (sr *shopRepository) CreateShop(id, name, email, password string) error {
	err := sr.sqlc.Create(ctx, database.CreateParams{
		ID:           id,
		ShopEmail:    email,
		ShopName:     name,
		ShopPassword: password,
	})

	if err != nil {
		return err
	}
	return nil
}

func (sr *shopRepository) GetShopByEmail(email string) (*database.Shop, error) {
	shop, err := sr.sqlc.GetShopByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &shop, nil
}

func NewShopRepository() IShopRepository {
	return &shopRepository{
		sqlc: database.New(global.Mdb),
	}
}

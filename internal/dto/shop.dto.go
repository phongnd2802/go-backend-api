package dto

type (
	ShopRegisterRequest struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	ShopLoginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	ShopSendOTPRequest struct {
		Email    string `json:"email" binding:"required"`
	}
)

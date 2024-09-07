package vo

type (
	ShopRegisterResponse struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	ShopLoginResponse struct {
		ShopRegisterResponse
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)

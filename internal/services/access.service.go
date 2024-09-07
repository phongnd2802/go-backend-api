package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/phongnd2802/go-backend-api/global"
	"github.com/phongnd2802/go-backend-api/internal/repositories"
	"github.com/phongnd2802/go-backend-api/internal/vo"
	"github.com/phongnd2802/go-backend-api/pkg/response"
	"github.com/phongnd2802/go-backend-api/pkg/utils/crypto"
	"github.com/phongnd2802/go-backend-api/pkg/utils/jwt"
	"github.com/phongnd2802/go-backend-api/pkg/utils/random"
	"github.com/phongnd2802/go-backend-api/pkg/utils/sendto"
	"go.uber.org/zap"
)

const (
	SIGNUP          = 0
	FORWARDPASSWORD = 1
)

const (
	ADMIN = 1
	SHOP = 2
	USER = 3
)

type IAccessService interface {
	SignUp(name, email, password string) (*vo.ShopRegisterResponse, int)
	SignIn(email, password string) (*vo.ShopLoginResponse, int)
	VerifyOTP(email string, otp int, purpose int) (*vo.ShopLoginResponse, int)
	SendOTP(email string) int
	ResetPassword(email, newPassword string) int
}

type accessService struct {
	otpRepo   repositories.IOTPRepository
	shopRepo  repositories.IShopRepository
	tokenRepo repositories.ITokenRepository
}

// ResetPassword implements IAccessService.
func (ac *accessService) ResetPassword(email string, newPassword string) int {
	foundShop, _ := ac.shopRepo.GetShopByEmail(email)
	if foundShop == nil {
		return response.ErrCodeNotFound
	}
	hashEmail := crypto.GetHash(email)
	exist := ac.otpRepo.GetOTPSetPassWord(hashEmail)
	if exist == 0 {
		return response.ErrCodeNotFound
	}

	passwordHash, err := crypto.HashPassword(newPassword)
	if err != nil {
		return response.ErrCodeInternalServerError
	}

	_ = ac.shopRepo.UpdatePassword(email, passwordHash)
	_ = ac.otpRepo.DeleteOTPSetPassword(hashEmail)

	return response.SuccessOK
}

func (ac *accessService) SendOTP(email string) int {
	foundShop, _ := ac.shopRepo.GetShopByEmail(email)
	if foundShop == nil {
		return response.ErrCodeNotFound
	}
	hashEmail := crypto.GetHash(email)
	otpExist := ac.otpRepo.GetOTP(hashEmail)
	if otpExist != 0 {
		return response.ErrCodeOTPExisting
	}
	fmt.Println("hashEmail::", hashEmail)
	otp := random.GenerateSixDigitOtp()
	fmt.Printf("otp::%d\n", otp)
	// Save OTP in Redis with expiration time
	err := ac.otpRepo.AddOTP(hashEmail, otp, int64(5*time.Minute))
	_ = ac.otpRepo.AddOTPCount(hashEmail, otp, int64(5*time.Minute))

	if err != nil {
		return response.ErrCodeInvalidOtp
	}
	// SendOTP to Email
	err = sendto.SendTextEmailOTP([]string{email}, "admin@gmail.com", strconv.Itoa(otp))
	if err != nil {
		return response.ErrCodeSendEmailFailed
	}
	return response.SuccessOK
}

func (ac *accessService) VerifyOTP(email string, otp int, purpose int) (*vo.ShopLoginResponse, int) {
	// Check Email Exist?
	foundShop, _ := ac.shopRepo.GetShopByEmail(email)
	if foundShop != nil {
		if foundShop.IsActive && purpose == SIGNUP {
			return nil, response.ErrCodeShopActive
		}
	} else {
		return nil, response.ErrCodeNotFound
	}

	// Hash Email
	hashEmail := crypto.GetHash(email)
	// GetOTP in Redis
	otpRedis := ac.otpRepo.GetOTP(hashEmail)
	if otpRedis == 0 {
		return nil, response.ErrCodeInvalidEmailOrOTP
	}

	// GetOTP Count
	otpCount := ac.otpRepo.GetOTPCount(hashEmail, otpRedis)
	if otpCount >= 5 {
		_ = ac.otpRepo.DeleteOTP(hashEmail)
		_ = ac.otpRepo.DeleteOTPCount(hashEmail, otpRedis)
		return nil, response.ErrCodeOtpSpam
	}

	if otp != otpRedis {
		ttl := ac.otpRepo.GetTTLOTPCount(hashEmail, otpRedis)
		if ttl > 0 {
			_ = ac.otpRepo.AddOTPCount(hashEmail, otpRedis, int64(ttl))
		}
		return nil, response.ErrCodeInvalidOtp
	}

	if purpose == FORWARDPASSWORD {
		_ = ac.otpRepo.AddOTPSetPassWord(hashEmail, int64(5*time.Minute))
		// Delete Infor OTP in Redis
		_ = ac.otpRepo.DeleteOTP(hashEmail)
		_ = ac.otpRepo.DeleteOTPCount(hashEmail, otpRedis)

		return nil, response.SuccessOK
	} else if purpose == SIGNUP {
		// Active Shop
		_ = ac.shopRepo.ActiveShopOTP(email)

		// Delete Infor OTP in Redis
		_ = ac.otpRepo.DeleteOTP(hashEmail)
		_ = ac.otpRepo.DeleteOTPCount(hashEmail, otpRedis)

		// Generate PublicKey, PrivateKey using RSA
		publicKeyStr, privateKeyStr, err := crypto.GenerateRSAKeyPair(1024)
		//fmt.Println(publicKeyStr)
		if err != nil {
			return nil, response.ErrCodeGenerateRSAFailed
		}

		// Payload (JWT)
		payload := map[string]interface{}{
			"id":    foundShop.ID,
			"email": foundShop.ShopEmail,
		}

		// Generate AccessToken, RefreshToken
		accessToken, _ := jwt.GenerateToken(payload, privateKeyStr, global.Config.JWT.ExpirationTimeAccessToken)
		refreshToken, _ := jwt.GenerateToken(payload, privateKeyStr, global.Config.JWT.ExpirationTimeRefreshToken)

		// Save infor token to DB
		tokenId := uuid.New().String()
		err = ac.tokenRepo.CreateToken(tokenId, publicKeyStr, refreshToken, foundShop.ID)
		if err != nil {
			return nil, response.ErrCodeInsertToken
		}

		// Response
		return &vo.ShopLoginResponse{
			ShopRegisterResponse: vo.ShopRegisterResponse{
				ID:    foundShop.ID,
				Name:  foundShop.ShopName,
				Email: foundShop.ShopEmail,
			},
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, response.SuccessOK
	}

	return nil, response.SuccessOK
}

func (ac *accessService) SignUp(name, email, password string) (*vo.ShopRegisterResponse, int) {
	// Check Email Exist?
	foundShop, _ := ac.shopRepo.GetShopByEmail(email)
	if foundShop != nil {
		return nil, response.ErrCodeShopExist
	}

	// Hash Email
	hashEmail := crypto.GetHash(email)
	fmt.Printf("hashEmail::%s\n", hashEmail)
	// Hash Password
	passwordHash, err := crypto.HashPassword(password)
	if err != nil {
		return nil, response.ErrCodeInternalServerError
	}
	// Generate UUID
	id := uuid.New()
	// Store to DB
	err = ac.shopRepo.CreateShop(id.String(), name, email, passwordHash)
	if err != nil {
		global.Logger.Error("Failed to create shop", zap.Error(err))
		return nil, response.ErrCodeCreateShop
	}
	err = ac.shopRepo.InsertRole(id.String(), SHOP)
	if err != nil {
		global.Logger.Error("Failed to create shop", zap.Error(err))
		return nil, response.ErrCodeCreateShop
	}


	// Generate OTP
	otp := random.GenerateSixDigitOtp()
	fmt.Printf("otp::%d\n", otp)
	// Save OTP in Redis with expiration time
	err = ac.otpRepo.AddOTP(hashEmail, otp, int64(5*time.Minute))
	_ = ac.otpRepo.AddOTPCount(hashEmail, otp, int64(5*time.Minute))
	if err != nil {
		return nil, response.ErrCodeInvalidOtp
	}

	// SendOTP to Email
	err = sendto.SendTextEmailOTP([]string{email}, "admin@gmail.com", strconv.Itoa(otp))
	if err != nil {
		return nil, response.ErrCodeSendEmailFailed
	}
	// Response
	return &vo.ShopRegisterResponse{
		ID:    id.String(),
		Name:  name,
		Email: email,
	}, response.CreatedOK
}

func (ac *accessService) SignIn(email, password string) (*vo.ShopLoginResponse, int) {
	foundShop, _ := ac.shopRepo.GetShopByEmail(email)
	if foundShop == nil {
		return nil, response.ErrCodeEmailOrPasswordInvalid
	}

	if !foundShop.IsActive {
		return nil, response.ErrCodeEmailNoActive
	}

	match := crypto.ComparePassword(password, foundShop.ShopPassword)
	if !match {
		return nil, response.ErrCodeEmailOrPasswordInvalid
	}

	publicKeyStr, privateKeyStr, err := crypto.GenerateRSAKeyPair(1024)
	fmt.Println(publicKeyStr)
	if err != nil {
		return nil, response.ErrCodeGenerateRSAFailed
	}

	payload := map[string]interface{}{
		"id":    foundShop.ID,
		"email": foundShop.ShopEmail,
	}
	accessToken, _ := jwt.GenerateToken(payload, privateKeyStr, global.Config.JWT.ExpirationTimeAccessToken)
	refreshToken, _ := jwt.GenerateToken(payload, privateKeyStr, global.Config.JWT.ExpirationTimeRefreshToken)

	tokenId := uuid.New().String()
	err = ac.tokenRepo.CreateToken(tokenId, publicKeyStr, refreshToken, foundShop.ID)
	if err != nil {
		return nil, response.ErrCodeInsertToken
	}

	return &vo.ShopLoginResponse{
		ShopRegisterResponse: vo.ShopRegisterResponse{
			ID:    foundShop.ID,
			Name:  foundShop.ShopName,
			Email: foundShop.ShopEmail,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, response.SuccessOK
}

func NewAccessService(
	shopRepo repositories.IShopRepository,
	otpRepo repositories.IOTPRepository,
	tokenRepo repositories.ITokenRepository,
) IAccessService {
	return &accessService{
		shopRepo:  shopRepo,
		otpRepo:   otpRepo,
		tokenRepo: tokenRepo,
	}
}

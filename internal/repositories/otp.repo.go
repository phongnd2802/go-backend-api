package repositories

import (
	"fmt"
	"github.com/phongnd2802/go-backend-api/global"
	"strconv"
	"time"
)

type IOTPRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
	GetOTP(email string) int
	DeleteOTP(email string) error

	SetOTPCount(email string, otp int, count int, expirationTime int64) error
	GetOTPCount(email string, otp int) int
	GetTTLOTPCount(email string, otp int) int
	DeleteOTPCount(email string, otp int) error

	AddOTPSetPassWord(email string, expirationTime int64) error
	GetOTPSetPassWord(email string) int
	DeleteOTPSetPassword(email string) error
}

type otpRepository struct{}

func (a *otpRepository) DeleteOTPSetPassword(email string) error {
	key := fmt.Sprintf("shop:%s:password", email)
	return global.Rdb.Del(ctx, key).Err()
}


// AddOTPSetPassWord implements IOTPRepository.
func (a *otpRepository) AddOTPSetPassWord(email string, expirationTime int64) error {
	key := fmt.Sprintf("shop:%s:password", email)
	return global.Rdb.SetEx(ctx, key, 1, time.Duration(expirationTime)).Err()
}


// GetOTPSetPassWord implements IOTPRepository.
func (a *otpRepository) GetOTPSetPassWord(email string) int {
	key := fmt.Sprintf("shop:%s:password", email)
	otp, err := global.Rdb.Get(ctx, key).Result()
	if err != nil {
		return 0
	}
	otpInt, _ := strconv.Atoi(otp)
	return otpInt
}

func (a *otpRepository) GetTTLOTPCount(email string, otp int) int {
	key := fmt.Sprintf("shop:%s:%d", email, otp)
	ttl, err := global.Rdb.TTL(ctx, key).Result()
	if err != nil {
		return -1
	}
	return int(ttl)
}

func (a *otpRepository) DeleteOTPCount(email string, otp int) error {
	key := fmt.Sprintf("shop:%s:%d", email, otp)
	return global.Rdb.Del(ctx, key).Err()
}

func (a *otpRepository) GetOTPCount(email string, otp int) int {
	key := fmt.Sprintf("shop:%s:%d", email, otp)
	countStr, err := global.Rdb.Get(ctx, key).Result()
	if err != nil {
		return -1
	}
	countInt, _ := strconv.Atoi(countStr)
	return countInt
}

func (a *otpRepository) SetOTPCount(email string, otp int, count int, expirationTime int64) error {
	key := fmt.Sprintf("shop:%s:%d", email, otp)
	return global.Rdb.SetEx(ctx, key, count, time.Duration(expirationTime)).Err()
}

func (a *otpRepository) DeleteOTP(email string) error {
	key := fmt.Sprintf("shop:%s:otp", email)
	return global.Rdb.Del(ctx, key).Err()
}

func (a *otpRepository) GetOTP(email string) int {
	key := fmt.Sprintf("shop:%s:otp", email)
	otp, err := global.Rdb.Get(ctx, key).Result()
	if err != nil {
		return 0
	}
	otpInt, _ := strconv.Atoi(otp)
	return otpInt
}

func (a *otpRepository) AddOTP(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("shop:%s:otp", email)
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}

func NewOTPRepository() IOTPRepository {
	return &otpRepository{}
}

package response

const (
	SuccessOK                     = 2000
	CreatedOK                     = 2001
	ErrCodeForbidden              = 4003
	ErrCodeInvalidParams          = 4010
	ErrCodeNotFound               = 4004
	ErrCodeShopExist              = 4011
	ErrCodeInternalServerError    = 5000
	ErrCodeCreateShop             = 4012
	ErrCodeInvalidOtp             = 4013
	ErrCodeInvalidEmailOrOTP      = 4014
	ErrCodeShopActive             = 4015
	ErrCodeOtpSpam                = 4016
	ErrCodeGenerateRSAFailed      = 4017
	ErrCodeEmailOrPasswordInvalid = 4018
	ErrCodeEmailNoActive          = 4019
	ErrCodeInsertToken            = 4020
	ErrCodeOTPExisting            = 4021
	ErrCodeSendEmailFailed        = 4022
)

var msg = map[int]string{
	SuccessOK: "success",
	CreatedOK: "created",

	ErrCodeInvalidParams:          "invalid params",
	ErrCodeShopExist:              "shop already exist",
	ErrCodeInternalServerError:    "internal server error",
	ErrCodeCreateShop:             "create shop error",
	ErrCodeInvalidOtp:             "otp error",
	ErrCodeNotFound:               "not found",
	ErrCodeInvalidEmailOrOTP:      "invalid email or otp",
	ErrCodeShopActive:             "shop is active",
	ErrCodeOtpSpam:                "otp is spam",
	ErrCodeGenerateRSAFailed:      "generate key pair error",
	ErrCodeEmailOrPasswordInvalid: "email or password incorrect",
	ErrCodeEmailNoActive:          "email no active",
	ErrCodeInsertToken:            "insert token error",
	ErrCodeOTPExisting:            "otp existing...",
	ErrCodeForbidden:              "forbidden error",
	ErrCodeSendEmailFailed:        "send email error",
}

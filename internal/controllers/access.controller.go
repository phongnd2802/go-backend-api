package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/phongnd2802/go-backend-api/internal/dto"
	"github.com/phongnd2802/go-backend-api/internal/services"
	"github.com/phongnd2802/go-backend-api/pkg/response"
	"strconv"
)

type AccessController struct {
	accessService services.IAccessService
}

func (ac *AccessController) ResetPassword(c *gin.Context) {
	var body dto.ShopLoginRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(c, response.ErrCodeInvalidParams, err)
		return
	}
	code := ac.accessService.ResetPassword(body.Email, body.Password)
	response.SuccessResponse(c, code, nil)
}


func (ac *AccessController) SendOTP(c *gin.Context) {
	var body dto.ShopSendOTPRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(c, response.ErrCodeInvalidParams, err)
		return
	}
	code := ac.accessService.SendOTP(body.Email)
	response.SuccessResponse(c, code, nil)
}


func (ac *AccessController) VerifyOTP(c *gin.Context) {
	email := c.DefaultQuery("email", "")
	otp := c.DefaultQuery("otp", "")
	purpose := c.DefaultQuery("purpose", "")
	if email == "" || otp == "" || purpose == "" {
		response.ErrorResponse(c, response.ErrCodeInvalidParams, errors.New("error"))
		return
	}
	otpInt, err := strconv.Atoi(otp)
	if err != nil {
		response.ErrorResponse(c, response.ErrCodeInvalidOtp, errors.New("error"))
		return
	}

	purposeInt, err := strconv.Atoi(purpose)
	if err != nil {
		response.ErrorResponse(c, response.ErrCodeInvalidParams, errors.New("error"))
		return
	}

	if purposeInt == 0 || purposeInt == 1 {
		res, code := ac.accessService.VerifyOTP(email, otpInt, purposeInt)
		response.SuccessResponse(c, code, res)
		return
	}
	response.ErrorResponse(c, response.ErrCodeInvalidParams, errors.New("error"))
}

func (ac *AccessController) SignUp(c *gin.Context) {
	var body dto.ShopRegisterRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(c, response.ErrCodeInvalidParams, err)
		return
	}
	res, code := ac.accessService.SignUp(body.Name, body.Email, body.Password)
	response.SuccessResponse(c, code, res)
}

func (ac *AccessController) SignIn(c *gin.Context) {
	var body dto.ShopLoginRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(c, response.ErrCodeInvalidParams, err)
		return
	}
	res, code := ac.accessService.SignIn(body.Email, body.Password)
	response.SuccessResponse(c, code, res)
}

func NewAccessController(
	accessService services.IAccessService,
) *AccessController {
	return &AccessController{
		accessService: accessService,
	}
}

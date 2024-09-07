package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phongnd2802/go-backend-api/global"
	"github.com/phongnd2802/go-backend-api/pkg/response"
	"github.com/phongnd2802/go-backend-api/pkg/utils/crypto"
	"github.com/phongnd2802/go-backend-api/pkg/utils/jwt"
)

func (m *Middleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := c.GetHeader("x-client-id")
		if clientID == "" {
			response.ErrorUnAuthorizedResponse(c, response.ErrCodeInvalidHeaderParam, "Header invalid!")
			return
		}
		
		accessToken := c.GetHeader("authorization")
		if accessToken == "" {
			response.ErrorUnAuthorizedResponse(c, response.ErrCodeInvalidHeaderParam, "Header invalid!")
			return
		}

		hashClientID := crypto.GetHash(clientID)
		accessTokenR, _ := m.tokenRepo.GetAccessToken(hashClientID)
		if accessTokenR == "" {
			keyToken, err := m.tokenRepo.GetTokenByUserID(clientID)
			if err != nil {
				response.ErrorInternalServerError(c, response.ErrCodeInternalServerError, err.Error())
				return
			}
			verifiedToken, err := jwt.VerifyToken(accessToken, keyToken.PublicKey)
			if err != nil {
				response.ErrorInternalServerError(c, response.ErrCodeInternalServerError, err.Error())
				return
			}

			verifiedTokenUserID, err := jwt.GetUserIDFromToken(verifiedToken)
			if err != nil {
				response.ErrorInternalServerError(c, response.ErrCodeInternalServerError, err.Error())
				return
			}

			if verifiedTokenUserID != clientID {
				response.ErrorUnAuthorizedResponse(c, response.ErrCodeInvalidUserID, "Error Decode UserID")
				return
			}

			_ = m.tokenRepo.SetAccessToken(accessToken, hashClientID, int64(global.Config.JWT.ExpirationTimeAccessToken * int(time.Hour)))

		} else {
			if accessTokenR != accessToken {
				response.ErrorUnAuthorizedResponse(c, response.ErrCodeTokenInvalid, "Header invalid!")
				return
			}
		}
		c.Next()
	}
}

func (m *Middleware) ApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("x-api-key")
		if apiKey == "" {
			response.ErrorForbiddenReponse(c, response.ErrCodeForbidden, "error")
			return
		}

		exist := m.authRepo.GetAPIKeyR(apiKey)
		if !exist {
			key, _ := m.authRepo.GetAPIKeyDB(apiKey)

			if key == nil {
				response.ErrorForbiddenReponse(c, response.ErrCodeForbidden, "apiKey not found!")
				return
			}
			if !key.IsActive {
				response.ErrorForbiddenReponse(c, response.ErrCodeForbidden, "apiKey invalid!")
				return
			}

			_ = m.authRepo.SetAPIKeyR(key.ApiKey)
		}

		c.Next()
	}
}

func (m *Middleware) PermissionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

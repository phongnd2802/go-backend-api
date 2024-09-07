package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func GenerateToken(payload map[string]interface{}, privateKeyPEM string, expirationTime int) (string, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyPEM))
	if err != nil {
		return "", err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": payload,
		"iss": "eCommerce-API",
		"exp": time.Now().Add(time.Hour * time.Duration(expirationTime)).Unix(),
	})

	tokenString, err := claims.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string, publicKeyPEM string) (*jwt.Token, error) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyPEM))
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
            return nil, err
        }

		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func GetUserIDFromToken(token *jwt.Token) (string, error) {
	// Truy cập claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Lấy payload, bạn có thể tìm "id" hoặc các trường khác từ payload
		if payload, ok := claims["sub"].(map[string]interface{}); ok {
			if userID, ok := payload["id"].(string); ok {
				return userID, nil
			}
			return "", fmt.Errorf("id not found in token claims")
		}
		return "", fmt.Errorf("invalid payload format")
	}

	return "", fmt.Errorf("invalid token claims")
}
package jwt

import (
	"auth/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

const (
	accessSecret  = "your-access-secret"
	refreshSecret = "your-refresh-secret"
)

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type Claims struct {
	UserId    uuid.UUID       `json:"user_id"`
	Role      models.RoleType `json:"role"`
	TokenType TokenType       `json:"token_type"`
	jwt.RegisteredClaims
}

func getTokenDetails(tokenType TokenType) (time.Time, string) {
	if tokenType == AccessToken {
		return time.Now().Add(15 * time.Minute), accessSecret
	}
	return time.Now().Add(7 * 24 * time.Hour), refreshSecret
}

func generateToken(userId uuid.UUID, role models.RoleType, tokenType TokenType) (string, error) {
	expirationTime, secret := getTokenDetails(tokenType)

	claims := &Claims{
		UserId:    userId,
		Role:      role,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
			Issuer:    "auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func GenerateTokenPair(userId uuid.UUID, role models.RoleType) (*TokenPair, error) {
	// Generate Access Token: expires in 15 minutes
	accessToken, err := generateToken(userId, role, AccessToken)
	if err != nil {
		return nil, err
	}

	// Generate Refresh Token: expires in 7 days
	refreshToken, err := generateToken(userId, role, RefreshToken)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func VerifyTokenAndGetClaims(tokenString string, tokenType TokenType) (*Claims, bool) {
	if tokenString == "" {
		return nil, false
	}

	_, secret := getTokenDetails(tokenType)
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, false
	}

	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid && claims.TokenType == tokenType {
		return claims, true
	}
	return nil, false
}

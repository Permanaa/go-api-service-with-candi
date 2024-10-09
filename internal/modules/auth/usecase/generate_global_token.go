package usecase

import (
	"go-api-service-with-candi/internal/modules/auth/domain"
	"go-api-service-with-candi/pkg/shared"
	"time"

	"github.com/golang-jwt/jwt"
)

func (uc *authUsecaseImpl) GenerateGlobalToken(req *domain.LoginRequest) (resp domain.LoginResponse, err error) {
	now := time.Now()
	accessTokenExp := now.Add(60 * time.Minute).Unix()
	refreshTokenExp := now.Add(time.Hour * 24 * 7).Unix()

	accessTokenClaims := jwt.MapClaims(map[string]interface{}{
		"iss": "trycandi",
		"exp": accessTokenExp,
	})

	accessTokenJwt := jwt.New(jwt.SigningMethodHS256)
	accessTokenJwt.Claims = accessTokenClaims

	accessToken, errAccessToken := accessTokenJwt.SignedString([]byte(shared.GetEnv().AccessTokenKey))
	if errAccessToken != nil {
		return resp, errAccessToken
	}

	refreshTokenClaims := jwt.MapClaims(map[string]interface{}{
		"iss": "trycandi",
		"exp": refreshTokenExp,
	})

	refreshTokenJwt := jwt.New(jwt.SigningMethodHS256)
	refreshTokenJwt.Claims = refreshTokenClaims

	refreshhToken, errRefreshToken := refreshTokenJwt.SignedString([]byte(shared.GetEnv().RefreshTokenKey))
	if errRefreshToken != nil {
		return resp, errRefreshToken
	}

	return domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshhToken,
	}, nil
}

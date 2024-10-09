package usecase

import (
	"errors"
	"go-api-service-with-candi/pkg/shared"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/golangid/candi/candihelper"
	"github.com/golangid/candi/middleware"
	"github.com/golangid/candi/wrapper"
)

func (uc *authUsecaseImpl) GenerateAnonymousToken() (token string, err error) {
	tokenHS := jwt.New(jwt.SigningMethodHS256)

	mapClaims := jwt.MapClaims(map[string]interface{}{
		"iss": "trycandi",
		"exp": time.Now().Add(60 * time.Minute).Unix(),
	})

	tokenHS.Claims = mapClaims

	token, err = tokenHS.SignedString([]byte(shared.GetEnv().AnonymousTokenKey))

	return token, err
}

func (uc *authUsecaseImpl) ValidateAnonymousToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		if err := func() error {
			authorization := req.Header.Get(candihelper.HeaderAuthorization)
			tokenValue, err := extractAuthType(middleware.BEARER, authorization)
			if err != nil {
				return err
			}

			parseToken, errParseToken := jwt.ParseWithClaims(tokenValue, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
				return []byte(shared.GetEnv().AnonymousTokenKey), nil
			})

			switch ve := errParseToken.(type) {
			case *jwt.ValidationError:
				if ve.Errors == jwt.ValidationErrorExpired {
					return errors.New("token is expired")
				} else {
					return errors.New("invalid token format")
				}
			}

			if errParseToken != nil {
				return errParseToken
			}

			if !parseToken.Valid {
				return errors.New("invalid token format")
			}

			return err
		}(); err != nil {
			wrapper.NewHTTPResponse(http.StatusUnauthorized, err.Error()).JSON(w)
			return
		}

		next.ServeHTTP(w, req.WithContext(ctx))
	})
}

func extractAuthType(prefix, authorization string) (string, error) {
	authValues := strings.Split(authorization, " ")
	if len(authValues) == 2 && strings.EqualFold(authValues[0], prefix) {
		return authValues[1], nil
	}

	return "", errors.New("invalid authorization")
}

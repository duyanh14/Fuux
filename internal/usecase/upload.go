package usecase

import (
	"fuux/internal/entity"
	errorEntity "fuux/internal/entity/error"
	resourceRepository "fuux/internal/repository/resource"
	"github.com/golang-jwt/jwt/v4"
)

type upload struct {
	config *entity.Config
}

var Upload *upload

func NewUpload(config *entity.Config) (*upload, error) {
	Upload = &upload{
		config: config,
	}
	return Upload, nil
}
func (s *upload) Auth(accessToken string) (*entity.ResourceAccess, error) {
	tokenParse, err := s.TokenParse(accessToken)
	if err != nil {
		return nil, err
	}

	model, err := resourceRepository.ResourceAccess.GetByID(tokenParse["id"].(string))
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *upload) TokenParse(accessToken string) (jwt.MapClaims, error) {
	hmacSecretString := s.config.JWT.Secret
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, errorEntity.Unknown.Error
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errorEntity.Unknown.Error
}

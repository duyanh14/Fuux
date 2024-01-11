package usecase

import (
	"fuux/internal/entity"
	errorEntity "fuux/internal/entity/error"
	"github.com/golang-jwt/jwt/v4"
)

type File struct {
	config *entity.Config
}

func NewFile(config *entity.Config) (*File, error) {
	return &File{
		config: config,
	}, nil
}

func (s *File) Auth(accessToken string) (*entity.ResourceAccess, error) {
	//tokenParse, err := s.TokenParse(accessToken)
	//if err != nil {
	//	return nil, err
	//}
	//
	//model, err := resourceRepository.ResourceAccess.GetByID(tokenParse["id"].(string))
	//if err != nil {
	//	return nil, err
	//}
	return nil, nil
}

func (s *File) TokenParse(accessToken string) (jwt.MapClaims, error) {
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

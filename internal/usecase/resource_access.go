package usecase

import (
	"fmt"
	"fuux/internal/entity"
	errorEntity "fuux/internal/entity/error"
	"fuux/internal/repository"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type ResourceAccess struct {
	config *entity.Config
	repo   *repository.ResourceAccess
}

func NewResourceAccess(config *entity.Config, repo *repository.ResourceAccess) (*ResourceAccess, error) {
	return &ResourceAccess{
		config: config,
		repo:   repo,
	}, nil
}

func (s *ResourceAccess) Create(payload *entity.ResourceAccess) (*entity.ResourceAccess, error) {
	_, err := s.repo.GetBy("name", payload.Name)
	if err == nil {
		return nil, errorEntity.NameAlreadyUse.Error
	}

	path, err := s.repo.GetByID(payload.ID)
	if err != nil {
		return nil, err
	}
	if path == nil {
		return nil, errorEntity.PathRecordNotFound.Error
	}

	fmt.Println(path)
	/////////////////

	//timeNow := time.Now()

	model := &entity.ResourceAccess{
		ID:            uuid.NewString(),
		Name:          payload.Name,
		ResourceRefer: payload.ResourceRefer,
		Status:        entity.ResourceAccessStatusEnable,
		Expire:        payload.Expire,
		Permission:    payload.Permission,
	}

	err = s.repo.Create(model)
	if err != nil {
		return nil, err
	}

	token, err := s.TokenGenerate(model)
	if err != nil {
		return nil, err
	}
	model.AccessToken = &token

	return model, nil
}

func (s *ResourceAccess) UpdatePath(payload *entity.ResourceAccess) (*entity.ResourceAccess, string, error) {
	//if pkg.IsStructContainNil(payload) {
	//	return nil, "", errorEntity.FieldRequired.Error
	//}

	// Name
	oldPathByID, err := s.repo.GetBy("id", payload.ID)
	if err != nil {
		return nil, "", err
	}

	oldPathByName, err := s.repo.GetBy("path", payload.ResourceRefer)
	if err != nil {
		if err != errorEntity.RecordNotFound.Error {
			return nil, "", err
		}
	}
	fmt.Println(oldPathByName)

	resourceAccessModel := &entity.ResourceAccess{
		ID:            oldPathByID.ID,
		Name:          oldPathByID.Name,
		ResourceRefer: oldPathByID.ResourceRefer,
		Status:        oldPathByID.Status,
		Expire:        oldPathByID.Expire,
	}
	resourceAccessSave := &entity.ResourceAccessSave{
		Name:       payload.Name,
		Status:     oldPathByID.Status,
		Permission: payload.Permission,
		Expire:     payload.Expire,
	}

	s.repo.Save(resourceAccessModel, resourceAccessSave)

	return &entity.ResourceAccess{
		ID:         oldPathByID.ID,
		Name:       resourceAccessSave.Name,
		Permission: resourceAccessSave.Permission,
		Status:     resourceAccessSave.Status,
		Expire:     resourceAccessSave.Expire,
	}, "", nil
}

func (s *ResourceAccess) RemovePath(id *entity.ResourceAccess) error {
	return s.repo.Delete(id)
}

func (s *ResourceAccess) List(list *entity.ResourceList) (*[]entity.ResourceAccess, int64, error) {
	return s.repo.List(list)
}

func (s *ResourceAccess) Get(id string) (*entity.ResourceAccess, error) {
	return s.repo.GetByID(id)
}

func (s *ResourceAccess) TokenGenerate(resourceAccess *entity.ResourceAccess) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = resourceAccess.ID

	//claims["exp"] = time.Now().Add(time.Hour * 8760).Unix()

	t, err := token.SignedString([]byte(s.config.JWT.Secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

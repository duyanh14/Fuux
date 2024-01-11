package usecase

import (
	"fuux/internal/entity"
	errorEntity "fuux/internal/entity/error"
	"fuux/internal/repository"
	"fuux/pkg"
	"github.com/google/uuid"
)

type Resource struct {
	config *entity.Config
	repo   *repository.Resource
}

func NewResource(config *entity.Config, repo *repository.Resource) (*Resource, error) {
	return &Resource{
		config: config,
		repo:   repo,
	}, nil
}

func (s *Resource) AddResource(payload *entity.Resource) (*entity.Resource, string, error) {
	if pkg.IsStructContainNil(payload) {
		return nil, "", errorEntity.FieldRequired.Error
	}

	var name string
	var path string
	name = payload.Name
	path = payload.Path

	/////////////////
	// Name
	_, err := s.repo.GetBy("name", name)
	if err == nil {
		return nil, "", errorEntity.NameAlreadyUse.Error
	}

	// Path
	_, err = s.repo.GetBy("path", path)
	if err == nil {
		return nil, "", errorEntity.PathAlreadyUse.Error
	}

	/////////////////

	//timeNow := time.Now()

	model := &entity.Resource{
		ID:   uuid.NewString(),
		Name: payload.Name,
		Path: payload.Path,
	}

	s.repo.Create(model)
	return model, "", nil
}

func (s *Resource) UpdatePath(payload *entity.Resource) (*entity.Resource, string, error) {
	if pkg.IsStructContainNil(payload) {
		return nil, "", errorEntity.FieldRequired.Error
	}

	// Name
	oldResourceByID, err := s.repo.GetBy("id", payload.ID)
	if err != nil {
		return nil, "", err
	}

	oldResourceByPath, err := s.repo.GetBy("path", payload.Path)
	if err != nil {
		if err != errorEntity.RecordNotFound.Error {
			return nil, "", err
		}
	}

	oldPathByName, err := s.repo.GetBy("name", payload.Name)
	if err != nil {
		if err != errorEntity.RecordNotFound.Error {
			return nil, "", err
		}
	}

	if oldResourceByPath != nil {
		if oldResourceByID.ID != oldResourceByPath.ID && payload.Path == oldResourceByPath.Path {
			return nil, "", errorEntity.PathExist.Error
		}
	}
	if oldPathByName != nil {
		if oldResourceByID.ID != oldPathByName.ID && payload.Name == oldPathByName.Name {
			return nil, "", errorEntity.NameExist.Error
		}
	}

	model := &entity.Resource{
		ID:   oldResourceByID.ID,
		Name: oldResourceByID.Name,
		Path: oldResourceByID.Path,
	}
	modelSave := &entity.ResourceSave{
		Name: payload.Name,
		Path: payload.Path,
	}

	s.repo.Save(model, modelSave)
	return &entity.Resource{
		ID:   oldResourceByID.ID,
		Name: modelSave.Name,
		Path: modelSave.Path,
		Date: oldResourceByID.Date,
	}, "", nil
}

func (s *Resource) RemovePath(id *entity.Resource) error {
	return s.repo.Delete(id)
}

func (s *Resource) List(list *entity.ResourceList) (*[]entity.Resource, int64, error) {
	return s.repo.List(list)
}

func (s *Resource) Get(id string) (*entity.Resource, error) {
	return s.repo.GetByID(id)
}

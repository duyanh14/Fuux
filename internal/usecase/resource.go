package usecase

import (
	"fuux/internal/entity"
	errorEntity "fuux/internal/entity/error"
	resourceRepository "fuux/internal/repository/resource"
	"fuux/pkg"
	"github.com/google/uuid"
)

type resource struct {
	config *entity.Config
}

var Resource *resource

func NewResource(config *entity.Config) (*resource, error) {
	Resource = &resource{
		config: config,
	}
	return Resource, nil
}
func (s *resource) AddResource(payload *entity.Resource) (*entity.Resource, string, error) {

	if pkg.IsStructContainNil(payload) {
		return nil, "", errorEntity.FieldRequired.Error
	}

	var name string
	var path string
	name = payload.Name
	path = payload.Path

	/////////////////
	// Name
	_, err := resourceRepository.Resource.GetBy("name", name)
	if err == nil {
		return nil, "", errorEntity.NameAlreadyUse.Error
	}

	// Path
	_, err = resourceRepository.Resource.GetBy("path", path)
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

	resourceRepository.Resource.Create(model)
	return model, "", nil
}

func (s *resource) UpdatePath(payload *entity.Resource) (*entity.Resource, string, error) {
	if pkg.IsStructContainNil(payload) {
		return nil, "", errorEntity.FieldRequired.Error
	}

	// Name
	oldResourceByID, err := resourceRepository.Resource.GetBy("id", payload.ID)
	if err != nil {
		return nil, "", err
	}

	oldResourceByPath, err := resourceRepository.Resource.GetBy("path", payload.Path)
	if err != nil {
		if err != errorEntity.RecordNotFound.Error {
			return nil, "", err
		}
	}

	oldPathByName, err := resourceRepository.Resource.GetBy("name", payload.Name)
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

	resourceRepository.Resource.Save(model, modelSave)
	return &entity.Resource{
		ID:   oldResourceByID.ID,
		Name: modelSave.Name,
		Path: modelSave.Path,
		Date: oldResourceByID.Date,
	}, "", nil
}

func (s *resource) RemovePath(id *entity.Resource) error {
	return resourceRepository.Resource.Delete(id)
}

func (s *resource) List(list *entity.ResourceList) (*[]entity.Resource, int64, error) {
	return resourceRepository.Resource.List(list)
}

func (s *resource) Get(id string) (*entity.Resource, error) {
	return resourceRepository.Resource.GetByID(id)
}

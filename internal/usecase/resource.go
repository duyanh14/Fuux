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
func (s *resource) AddPath(payload *entity.Path) (*entity.Path, string, error) {

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

	pathModel := &entity.Path{
		ID:   uuid.NewString(),
		Name: payload.Name,
		Path: payload.Path,
	}

	resourceRepository.Resource.Create(pathModel)
	return pathModel, "", nil
}

func (s *resource) UpdatePath(payload *entity.Path) (*entity.Path, string, error) {
	if pkg.IsStructContainNil(payload) {
		return nil, "", errorEntity.FieldRequired.Error
	}

	// Name
	oldPathByID, err := resourceRepository.Resource.GetBy("id", payload.ID)
	if err != nil {
		return nil, "", err
	}

	oldPathByPath, err := resourceRepository.Resource.GetBy("path", payload.Path)
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

	if oldPathByPath != nil {
		if oldPathByID.ID != oldPathByPath.ID && payload.Path == oldPathByPath.Path {
			return nil, "", errorEntity.PathExist.Error
		}
	}
	if oldPathByName != nil {
		if oldPathByID.ID != oldPathByName.ID && payload.Name == oldPathByName.Name {
			return nil, "", errorEntity.NameExist.Error
		}
	}

	pathModel := &entity.Path{
		ID:   oldPathByID.ID,
		Name: oldPathByID.Name,
		Path: oldPathByID.Path,
	}
	pathSave := &entity.PathSave{
		Name: payload.Name,
		Path: payload.Path,
	}

	resourceRepository.Resource.Save(pathModel, pathSave)
	return &entity.Path{
		ID:   oldPathByID.ID,
		Name: pathSave.Name,
		Path: pathSave.Path,
		Date: oldPathByID.Date,
	}, "", nil
}

func (s *resource) RemovePath(id *entity.Path) error {
	return resourceRepository.Resource.Delete(id)
}

func (s *resource) List(list *entity.PathList) (*[]entity.Path, int64, error) {
	return resourceRepository.Resource.List(list)
}

func (s *resource) Get(id string) (*entity.Path, error) {
	return resourceRepository.Resource.GetByID(id)
}

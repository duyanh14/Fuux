package usecase

import (
	"fuux/internal/entity"
	errorEntity "fuux/internal/entity/error"
	resourceRepository "fuux/internal/repository/resource"
	"fuux/pkg"
	"github.com/google/uuid"
)

type resourceAccess struct {
	config *entity.Config
}

var ResourceAccess *resourceAccess

func NewResourceAccess(config *entity.Config) (*resourceAccess, error) {
	ResourceAccess = &resourceAccess{
		config: config,
	}
	return ResourceAccess, nil
}
func (s *resourceAccess) AddPath(payload *entity.ResourceAccess) (*entity.ResourceAccess, string, error) {

	if pkg.IsStructContainNil(payload) {
		return nil, "", errorEntity.FieldRequired.Error
	}

	var name string
	var path string
	name = payload.Name
	path = payload.Path

	/////////////////
	// Name
	_, err := resourceRepository.ResourceAccess.GetBy("name", name)
	if err == nil {
		return nil, "", errorEntity.NameAlreadyUse.Error
	}

	// Path
	_, err = resourceRepository.ResourceAccess.GetBy("path", path)
	if err == nil {
		return nil, "", errorEntity.PathAlreadyUse.Error
	}

	/////////////////

	//timeNow := time.Now()

	resourceAccessModel := &entity.ResourceAccess{
		ID:     uuid.NewString(),
		Name:   payload.Name,
		Path:   payload.Path,
		Status: payload.Status, // edit sau
		Expire: payload.Expire, // edit sau
	}

	resourceRepository.ResourceAccess.Create(resourceAccessModel)
	return resourceAccessModel, "", nil
}

func (s *resourceAccess) UpdatePath(payload *entity.ResourceAccess) (*entity.ResourceAccess, string, error) {
	if pkg.IsStructContainNil(payload) {
		return nil, "", errorEntity.FieldRequired.Error
	}

	// Name
	oldPathByID, err := resourceRepository.ResourceAccess.GetBy("id", payload.ID)
	if err != nil {
		return nil, "", err
	}

	oldPathByPath, err := resourceRepository.ResourceAccess.GetBy("path", payload.Path)
	if err != nil {
		if err != errorEntity.RecordNotFound.Error {
			return nil, "", err
		}
	}

	oldPathByName, err := resourceRepository.ResourceAccess.GetBy("name", payload.Name)
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

	resourceAccessModel := &entity.ResourceAccess{
		ID:     oldPathByID.ID,
		Name:   oldPathByID.Name,
		Path:   oldPathByID.Path,
		Status: oldPathByID.Status,
		Expire: oldPathByID.Expire,
	}
	resourceAccessSave := &entity.ResourceAccessSave{
		Name:   payload.Name,
		Path:   payload.Path,
		Status: payload.Status,
		Expire: payload.Expire,
	}

	resourceRepository.ResourceAccess.Save(resourceAccessModel, resourceAccessSave)
	return &entity.ResourceAccess{
		ID:     oldPathByID.ID,
		Name:   resourceAccessSave.Name,
		Path:   resourceAccessSave.Path,
		Status: resourceAccessSave.Status,
		Expire: resourceAccessSave.Expire,
	}, "", nil
}

func (s *resourceAccess) RemovePath(id *entity.ResourceAccess) error {
	return resourceRepository.ResourceAccess.Delete(id)
}

func (s *resourceAccess) List(list *entity.ResourceList) (*[]entity.ResourceAccess, int64, error) {
	return resourceRepository.ResourceAccess.List(list)
}

func (s *resourceAccess) Get(id string) (*entity.ResourceAccess, error) {
	return resourceRepository.ResourceAccess.GetByID(id)
}

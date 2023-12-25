package resource

import (
	"fuux/internal/entity"
	errorEntity "fuux/internal/entity/error"
	"fuux/internal/repository"
)

type resource struct {
	database *entity.Database
}

var Resource *resource

func NewResource(database *entity.Database) (*resource, error) {
	Resource = &resource{
		database: database,
	}

	return Resource, nil
}

func (r *resource) Database() *entity.Database {
	return r.database
}

func (r *resource) Create(pathModel *entity.Path) error {
	result := r.database.Postgres.Create(pathModel)
	if result.RowsAffected == 0 {
		return errorEntity.Unknown.Error
	}
	return nil
}

func (r *resource) GetBy(by string, value string) (*entity.Path, error) {
	path := &entity.Path{}
	query := r.database.Postgres.Where(by+" = ?", value).Find(path)
	if query.RowsAffected == 0 {
		return nil, errorEntity.RecordNotFound.Error
	}
	return path, nil
}
func (r *resource) GetByID(id string) (*entity.Path, error) {
	return r.GetBy("id", id)
}

func (r *resource) Save(path *entity.Path, save *entity.PathSave) error {
	updateField := repository.UpdateField(save)

	query := r.database.Postgres.Model(path).Updates(updateField)

	if query.RowsAffected == 0 {
		return errorEntity.UserAccountSaveFailed.Error
	}
	return nil
}

func (r *resource) Delete(id *entity.Path) error {
	result := r.database.Postgres.Delete(id)
	if result.RowsAffected == 0 {
		return errorEntity.Unknown.Error
	}
	return nil
}

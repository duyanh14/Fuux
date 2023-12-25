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
func (r *resource) List(list *entity.PathList) (*[]entity.Path, int64, error) {
	rs := make([]entity.Path, 0)
	var count int64

	query := r.database.Postgres.Debug().Model(&entity.Path{})

	filterFindValue := list.Filter.Find.Value
	if filterFindValue != "" {
		for _, v := range list.Filter.Find.Field {
			mode := v + " = ?"
			value := filterFindValue
			if list.Filter.Find.Mode == 2 {
				mode = "LOWER(" + v + ") LIKE LOWER(?)"
				value = "%" + filterFindValue + "%"
			}
			query.Or(mode, value)
		}
	}

	query.Count(&count)
	if list.Limit == 0 {
		list.Limit = int(count)
	}

	query.Limit(list.Limit).Offset((list.Page - 1) * list.Limit)

	var paths []entity.Path
	query.Find(&paths)

	for _, path := range paths {
		rs = append(rs, path)
	}

	return &rs, count, nil
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

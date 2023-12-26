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

func (r *resource) Create(model *entity.Resource) error {
	result := r.database.Postgres.Create(model)
	if result.RowsAffected == 0 {
		return errorEntity.Unknown.Error
	}
	return nil
}

func (r *resource) GetBy(by string, value string) (*entity.Resource, error) {
	model := &entity.Resource{}
	query := r.database.Postgres.Where(by+" = ?", value).Find(model)
	if query.RowsAffected == 0 {
		return nil, errorEntity.RecordNotFound.Error
	}
	return model, nil
}
func (r *resource) GetByID(id string) (*entity.Resource, error) {
	return r.GetBy("id", id)
}
func (r *resource) List(list *entity.ResourceList) (*[]entity.Resource, int64, error) {
	rs := make([]entity.Resource, 0)
	var count int64

	query := r.database.Postgres.Debug().Model(&entity.Resource{})

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

	var resources []entity.Resource
	query.Find(&resources)

	for _, resource := range resources {
		rs = append(rs, resource)
	}

	return &rs, count, nil
}

func (r *resource) Save(path *entity.Resource, save *entity.ResourceSave) error {
	updateField := repository.UpdateField(save)

	query := r.database.Postgres.Model(path).Updates(updateField)

	if query.RowsAffected == 0 {
		return errorEntity.UserAccountSaveFailed.Error
	}
	return nil
}

func (r *resource) Delete(id *entity.Resource) error {
	result := r.database.Postgres.Delete(id)
	if result.RowsAffected == 0 {
		return errorEntity.Unknown.Error
	}
	return nil
}

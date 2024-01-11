package repository

import (
	"fuux/internal/entity"
	errorEntity "fuux/internal/entity/error"
)

type Resource struct {
	database *entity.Database
}

func NewResource(database *entity.Database) (*Resource, error) {
	return &Resource{
		database: database,
	}, nil
}

func (r *Resource) Database() *entity.Database {
	return r.database
}

func (r *Resource) Create(model *entity.Resource) error {
	result := r.database.Postgres.Create(model)
	if result.RowsAffected == 0 {
		return errorEntity.Unknown.Error
	}
	return nil
}

func (r *Resource) GetBy(by string, value string) (*entity.Resource, error) {
	model := &entity.Resource{}
	query := r.database.Postgres.Where(by+" = ?", value).Find(model)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, errorEntity.RecordNotFound.Error
	}
	return model, nil
}

func (r *Resource) GetByID(id string) (*entity.Resource, error) {
	return r.GetBy("id", id)
}

func (r *Resource) List(list *entity.ResourceList) (*[]entity.Resource, int64, error) {
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

func (r *Resource) Save(path *entity.Resource, save *entity.ResourceSave) error {
	//updateField := UpdateField(save)
	//
	//query := r.database.Postgres.Model(path).Updates(updateField)
	//
	//if query.RowsAffected == 0 {
	//	return errorEntity.UserAccountSaveFailed.Error
	//}
	return nil
}

func (r *Resource) Delete(id *entity.Resource) error {
	result := r.database.Postgres.Delete(id)
	if result.RowsAffected == 0 {
		return errorEntity.Unknown.Error
	}
	return nil
}

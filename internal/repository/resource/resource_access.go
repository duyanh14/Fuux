package resource

import (
	"fuux/internal/entity"
	errorEntity "fuux/internal/entity/error"
	"fuux/internal/repository"
)

type resourceAccess struct {
	database *entity.Database
}

var ResourceAccess *resourceAccess

func NewResourceAccess(database *entity.Database) (*resourceAccess, error) {
	ResourceAccess = &resourceAccess{
		database: database,
	}

	return ResourceAccess, nil
}

func (r *resourceAccess) Database() *entity.Database {
	return r.database
}

func (r *resourceAccess) Create(resourceAccess *entity.ResourceAccess) error {
	result := r.database.Postgres.Create(resourceAccess)
	if result.RowsAffected == 0 {
		return errorEntity.Unknown.Error
	}
	return nil
}

func (r *resourceAccess) GetBy(by string, value string) (*entity.ResourceAccess, error) {
	resourceAccessModel := &entity.ResourceAccess{}
	query := r.database.Postgres.Where(by+" = ?", value).Find(resourceAccessModel)
	if query.RowsAffected == 0 {
		return nil, errorEntity.RecordNotFound.Error
	}
	return resourceAccessModel, nil
}
func (r *resourceAccess) GetByID(id string) (*entity.ResourceAccess, error) {
	return r.GetBy("id", id)
}
func (r *resourceAccess) List(list *entity.ResourceList) (*[]entity.ResourceAccess, int64, error) {
	rs := make([]entity.ResourceAccess, 0)
	var count int64

	query := r.database.Postgres.Debug().Model(&entity.ResourceAccess{})

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

	var resourceAccesses []entity.ResourceAccess
	query.Find(&resourceAccesses)

	for _, resourceAccessModel := range resourceAccesses {
		rs = append(rs, resourceAccessModel)
	}

	return &rs, count, nil
}

func (r *resourceAccess) Save(resourceAccessModel *entity.ResourceAccess, save *entity.ResourceAccessSave) error {
	updateField := repository.UpdateField(save)

	query := r.database.Postgres.Model(resourceAccessModel).Updates(updateField)

	if query.RowsAffected == 0 {
		return errorEntity.UserAccountSaveFailed.Error
	}
	return nil
}

func (r *resourceAccess) Delete(id *entity.ResourceAccess) error {
	result := r.database.Postgres.Delete(id)
	if result.RowsAffected == 0 {
		return errorEntity.Unknown.Error
	}
	return nil
}

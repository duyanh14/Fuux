package entity

type ResourceAccess struct {
	ID     string `json:"id" gorm:"primaryKey"`
	Name   string `json:"name" gorm:"column:name"`
	Path   string `json:"path" gorm:"column:path"`
	Status string `json:"status" gorm:"column:status"`
	Expire string `json:"expire" gorm:"column:expire"`
}

func (ResourceAccess) TableName() string {
	return "Resource_access"
}

type ResourceAccessSave struct {
	Name   string `json:"name" gorm:"column:name"`
	Path   string `json:"path" gorm:"column:path"`
	Status string `json:"status" gorm:"column:status"`
	Expire string `json:"expire" gorm:"column:expire"`
}

type ResourceList struct {
	Filter ResourceListFilter `json:"filter"`
	Pagination
}
type ResourceListFilter struct {
	Find ResourceListFilterFind `json:"find"`
}

type ResourceListFilterFind struct {
	Mode  int      `json:"mode"`
	Field []string `json:"field"`
	Value string   `json:"value"`
}

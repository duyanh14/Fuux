package entity

//
//type Resource struct {
//	ID   string `json:"id" gorm:"primaryKey"`
//	Name string `json:"name"`
//	Path string `json:"path"`
//	Date
//}

//func (Resource) TableName() string {
//	return "resource"
//}
//
//type ResourceSecret struct {
//	ID            string   `json:"id" gorm:"primaryKey"`
//	Name          string   `json:"name"`
//	Key           string   `json:"key"`
//	Type          string   `json:"type"`
//	ResourceRefer string   `gorm:"column:resource"`
//	Resource      Resource `gorm:"foreignKey:ResourceRefer"`
//}

//func (ResourceSecret) TableName() string {
//	return "resource_secret"
//}

type Resource struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"column:name"`
	Path string `json:"path" gorm:"column:path"`
	Date
}

type ResourceSave struct {
	Name string `json:"name" gorm:"column:name"`
	Path string `json:"path" gorm:"column:path"`
}

func (Resource) TableName() string {
	return "resource"
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

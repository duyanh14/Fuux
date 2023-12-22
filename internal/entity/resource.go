package entity

type Resource struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Path string `json:"path"`
	Date
}

func (Resource) TableName() string {
	return "resource"
}

type ResourceSecret struct {
	ID            string   `json:"id" gorm:"primaryKey"`
	Name          string   `json:"name"`
	Key           string   `json:"key"`
	Type          string   `json:"type"`
	ResourceRefer string   `gorm:"column:resource"`
	Resource      Resource `gorm:"foreignKey:ResourceRefer"`
}

func (ResourceSecret) TableName() string {
	return "resource_secret"
}

type PathAccess struct {
	ID     string `json:"id" gorm:"primaryKey"`
	Name   string `json:"name" gorm:"column:name"`
	Path   string `json:"path" gorm:"column:path"`
	Status string `json:"status" gorm:"column:status"`
	Expire string `json:"expire" gorm:"column:expire"`
}

func (PathAccess) TableName() string {
	return "path_access"
}

type Path struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"column:name"`
	Path string `json:"path" gorm:"column:path"`
	Date
}

func (Path) TableName() string {
	return "path"
}

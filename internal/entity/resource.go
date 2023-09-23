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

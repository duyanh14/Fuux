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

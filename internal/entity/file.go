package entity

type File struct {
	ID            string   `json:"id" gorm:"primaryKey"`
	ResourceRefer string   `gorm:"column:resource"`
	Resource      Resource `gorm:"foreignKey:ReasonRefer"`
	Path          string   `json:"path" gorm:"column:path"`
	Date
}

func (File) TableName() string {
	return "file"
}

package entity

type File struct {
	ID            string `json:"id" gorm:"primaryKey"`
	ResourceRefer string `gorm:"column:resource"`
	//	Resource      Resource `gorm:"foreignKey:ResourceRefer"`
	Path  string `json:"path" gorm:"column:path"`
	URL   string `json:"url" gorm:"column:url"`
	Size  int64  `json:"size" gorm:"column:size"`
	Exist bool   `json:"exist"`
	Date
}

func (File) TableName() string {
	return "resource"
}

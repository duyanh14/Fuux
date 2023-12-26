package entity

const (
	ResourceAccessStatusEnable  = 1
	ResourceAccessStatusDisable = 2
)

type ResourceAccess struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name" gorm:"column:name"`
	ResourceRefer string    `gorm:"column:path" json:"path"`
	Resource      *Resource `gorm:"foreignKey:ResourceRefer" json:"-"`
	Status        int       `json:"status" gorm:"column:status"`
	Expire        int64     `json:"expire" gorm:"column:expire"`
	AccessToken   *string   `json:"access_token,omitempty" gorm:"-"`

	Permission *ResourceAccessPermission `json:"permission" gorm:"column:permission;serializer:json"`
	Date
}

type ResourceAccessPermission struct {
	Read  bool `json:"read"`
	Write bool `json:"write"`
}

func (ResourceAccess) TableName() string {
	return "resource_access"
}

type ResourceAccessSave struct {
	Name       string                    `json:"name" gorm:"column:name"`
	Status     int                       `json:"status" gorm:"column:status"`
	Permission *ResourceAccessPermission `json:"permission" gorm:"column:permission;serializer:json"`
	Expire     int64                     `json:"expire" gorm:"column:expire"`
}

type ResourceAccessList struct {
	Filter ResourceAccessListFilter `json:"filter"`
	Pagination
}
type ResourceAccessListFilter struct {
	Find ResourceAccessListFilterFind `json:"find"`
}

type ResourceAccessListFilterFind struct {
	Mode  int      `json:"mode"`
	Field []string `json:"field"`
	Value string   `json:"value"`
}

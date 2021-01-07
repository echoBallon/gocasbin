package models

type Routers struct {
	RouterId int `gorm:"column:user_id;primaryKey"`
	RouterName string `gorm:"column:r_name"`
	RouterUri string `gorm:"column:r_uri"`
	RouterMethod string `gorm:"column:r_method"`
	RoleName string
}

func (this *Routers)TableName()  string{
	return "roles"
}

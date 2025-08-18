package entity

type User struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Username string `json:"username" xml:"username" form:"username" gorm:"type:varchar(100)"`
	Password string `json:"password" xml:"password" form:"password" gorm:"type:varchar(255)"`
}

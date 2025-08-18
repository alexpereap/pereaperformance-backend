package entity

// Home image slides entity

type Slide struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Image    string `json:"image" xml:"image" form:"image" gorm:"type:varchar(255)"`
	Title    string `json:"title" xml:"title" form:"title" gorm:"type:varchar(255)"`
	TitlePos string `json:"title_pos" xml:"title_pos" form:"title_pos" gorm:"type:varchar(255)"`
}

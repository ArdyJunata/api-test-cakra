package models

type Car struct {
	ID    uint64 `json:"id" gorm:"primary_key:auto_increment"`
	Price uint64 `json:"price" gorm:"type:int" sql:"DEFAULT:0"`
	Brand string `json:"brand" gorm:"type:varchar(100)"`
	Type  string `json:"type" gorm:"type:varchar(100)"`
}

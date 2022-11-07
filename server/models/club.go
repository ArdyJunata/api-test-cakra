package models

type Club struct {
	ID       uint64 `json:"id" gorm:"primary_key:auto_increment"`
	ClubName string `json:"clubname" gorm:"type:varchar(100)"`
	Point    uint64 `json:"point" gorm:"type:int" sql:"DEFAULT:0"`
}

package model

type Group struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"group_name"`
	User []User `json:"-" gorm:"foreignKey:GroupID"`
}

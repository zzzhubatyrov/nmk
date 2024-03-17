package model

type User struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	RoleID      int    `json:"roleID"`
	GroupID     int    `json:"groupID"`
	Tag         string `json:"tag" gorm:"unique"`
	Email       string `json:"email" gorm:"unique;not null"`
	Password    string `json:"password" gorm:"not null"`
	Name        string `json:"name"`
	Photo       []byte `json:"photo"`
	Description string `json:"description"`
	Group       Group  `json:"group" gorm:"foreignKey:GroupID"`
	Role        Role   `json:"role" gorm:"foreignKey:RoleID"`
}

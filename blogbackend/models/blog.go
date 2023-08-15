package models

type Blog struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	UserID      string `json:"userid"`
	User        User   `gorm:"foreignKey:UserID"`
}

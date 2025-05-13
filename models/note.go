package models

type Note struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Title     string         `gorm:"size:200;not null" json:"title"`
    Content   string         `gorm:"type:text" json:"content"`
    UserID    uint           `json:"user_id"` // Relación con el modelo User
    User      User           `gorm:"foreignKey:UserID" json:"-"`
}
package model

type OrderItem struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	ChrtID      int    `gorm:"not null"`
	TrackNumber string `gorm:"type:varchar(50);not null"`
	Price       int    `gorm:"not null"`
	Rid         string `gorm:"type:varchar(50);not null"`
	Name        string `gorm:"type:varchar(100);not null"`
	Sale        int    `gorm:"not null"`
	Size        string `gorm:"type:varchar(10);not null"`
	TotalPrice  int    `gorm:"not null"`
	NmID        int    `gorm:"not null"`
	Brand       string `gorm:"type:varchar(100);not null"`
	Status      int    `gorm:"not null"`
	OrderID     string `gorm:"type:varchar(100);not null;index"`
}

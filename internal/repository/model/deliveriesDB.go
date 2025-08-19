package model

type Delivery struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	Name    string `gorm:"type:varchar(50);not null"`
	Phone   string `gorm:"type:varchar(50);not null;check:phone ~ '^\\+[0-9]{10,}$'"`
	Zip     string `gorm:"type:varchar(50);not null"`
	City    string `gorm:"type:varchar(50);not null"`
	Address string `gorm:"type:varchar(50);not null"`
	Region  string `gorm:"type:varchar(50);not null"`
	Email   string `gorm:"type:varchar(50);not null;check:email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}$'"`
}

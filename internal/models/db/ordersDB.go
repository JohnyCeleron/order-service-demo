package models_db

import "time"

type Order struct {
	Id                string `gorm:"type:varchar(100);primaryKey"`
	TrackNumber       string `gorm:"type:varchar(50);not null"`
	Entry             string `gorm:"type:varchar(50);not null"`
	DeliveryID        uint
	Delivery          Delivery `gorm:"foreignKey:DeliveryID"`
	PaymentID         uint
	Payment           Payments  `gorm:"foreignKey:PaymentID"`
	Locale            string    `gorm:"type:varchar(10);not null"`
	InternalSignature string    `gorm:"type:varchar(50);not null"`
	CustomerID        string    `gorm:"type:varchar(50);not null"`
	DeliveryService   string    `gorm:"type:varchar(50);not null"`
	Shardkey          string    `gorm:"type:varchar(50);not null"`
	SmID              int       `gorm:"not null"`
	DateCreated       time.Time `gorm:"type:timestamptz;not null"`
	OofShart          string    `gorm:"type:varchar(50);not null"`
}

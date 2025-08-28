package model

import "time"

type Order struct {
	ID                string    `gorm:"type:varchar(100);primaryKey;index"`
	TrackNumber       string    `gorm:"type:varchar(50);not null"`
	Entry             string    `gorm:"type:varchar(50);not null"`
	Locale            string    `gorm:"type:varchar(10);not null"`
	InternalSignature string    `gorm:"type:varchar(50);not null"`
	CustomerID        string    `gorm:"type:varchar(50);not null"`
	DeliveryService   string    `gorm:"type:varchar(50);not null"`
	Shardkey          string    `gorm:"type:varchar(50);not null"`
	SmID              int       `gorm:"not null"`
	DateCreated       time.Time `gorm:"type:timestamptz;not null"`
	OofShard          string    `gorm:"type:varchar(50);not null"`

	// Связи
	DeliveryID uint
	Delivery   Delivery `gorm:"foreignKey:DeliveryID"`
	PaymentID  uint
	Payment    Payment     `gorm:"foreignKey:PaymentID"`
	Items      []OrderItem `gorm:"foreignKey:OrderID"`
}

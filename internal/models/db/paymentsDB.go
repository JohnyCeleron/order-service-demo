package models_db

type Payment struct {
	Id           uint   `gorm:"primaryKey;autoIncrement"`
	Transaction  string `gorm:"type:varchar(100);not null;unique"`
	RequestID    string `gorm:"type:varchar(100)"`
	Currency     string `gorm:"type:varchar(10);not null"`
	Provider     string `gorm:"type:varchar(50);not null"`
	Amount       int    `gorm:"not null;check:amount >= 0"`
	PaymentDt    int    `gorm:"not null"`
	Bank         string `gorm:"type:varchar(50);not null"`
	DeliveryCost int    `gorm:"not null;check:delivery_cost >= 0"`
	GoodsTotal   int    `gorm:"not null;check:goods_total >= 0"`
	CustomFee    int    `gorm:"not null;check:custom_fee >= 0"`
}

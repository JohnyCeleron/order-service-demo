package models_db

type OrderItem struct {
	Id          uint   `gorm:"primaryKey;autoIncrement"`         // SERIAL PRIMARY KEY
	ChrtID      int    `gorm:"not null"`                         // chrt_id INT NOT NULL
	TrackNumber string `gorm:"type:varchar(50);not null"`        // track_number VARCHAR(50) NOT NULL
	Price       int    `gorm:"not null"`                         // price INT NOT NULL
	Rid         string `gorm:"type:varchar(50);not null"`        // rid VARCHAR(50) NOT NULL
	Name        string `gorm:"type:varchar(100);not null"`       // name VARCHAR(100) NOT NULL
	Sale        int    `gorm:"not null"`                         // sale INTEGER NOT NULL
	Size        string `gorm:"type:varchar(10);not null"`        // size VARCHAR(10) NOT NULL
	TotalPrice  int    `gorm:"not null"`                         // total_price INT NOT NULL
	NmID        int    `gorm:"not null"`                         // nm_id INT NOT NULL
	Brand       string `gorm:"type:varchar(100);not null"`       // brand VARCHAR(100) NOT NULL
	Status      int    `gorm:"not null"`                         // status INTEGER NOT NULL
	OrderID     string `gorm:"type:varchar(100);not null"`       // order_id VARCHAR(100) REFERENCES orders(id)
	Order       Order  `gorm:"foreignKey:OrderID;references:ID"` // Связь с Order
}

package domain

import (
	"time"

	domainDelivery "order-service/internal/domain/delivery"
	domainOrderItem "order-service/internal/domain/orderItem"
	domainPayment "order-service/internal/domain/payment"
)

type OrderValidateTestCase struct {
	name     string
	order    Order
	expected bool
	wantErr  bool
	errMsg   string
}

var orderRequiredValidateTestCases []OrderValidateTestCase = []OrderValidateTestCase{
	{
		name: "payment required",
		order: Order{
			OrderUID:    "asdfadsf",
			TrackNumber: "WBILMTESTTRACK",
			Entry:       "WBIL",
			Delivery: domainDelivery.Delivery{
				Name:    "Test Testov",
				Phone:   "+9720000000",
				Zip:     "2639809",
				City:    "Kiryat Mozkin",
				Address: "Ploshad Mira 15",
				Region:  "Kraiot",
				Email:   "test@gmail.com",
			},
			Items: []domainOrderItem.OrderItem{
				{
					ChrtID:      9934930,
					TrackNumber: "WBILMTESTTRACK",
					Price:       453,
					Rid:         "ab4219087a764ae0btest",
					Name:        "Mascaras",
					Sale:        30,
					Size:        "0",
					TotalPrice:  317,
					NmID:        2389212,
					Brand:       "Vivenne Sabo",
					Status:      202,
				},
			},
			Locale:            "en",
			InternalSignature: "",
			CustomerID:        "test",
			DeliveryService:   "meest",
			ShardKey:          "9",
			SmID:              99,
			DateCreated:       time.Now(),
			OofShard:          "1",
		},
		expected: false,
		wantErr:  true,
		errMsg:   "required",
	},
	{
		name: "delivery required",
		order: Order{
			OrderUID:    "asdfadsf",
			TrackNumber: "WBILMTESTTRACK",
			Entry:       "WBIL",
			Payment: domainPayment.Payment{
				Transaction:  "b563feb7b2b84b6test",
				RequestID:    "",
				Currency:     "USD",
				Provider:     "wbpay",
				Amount:       1817,
				PaymentDt:    1637907727,
				Bank:         "alpha",
				DeliveryCost: 1500,
				GoodsTotal:   317,
				CustomFee:    0,
			},
			Items: []domainOrderItem.OrderItem{
				{
					ChrtID:      9934930,
					TrackNumber: "WBILMTESTTRACK",
					Price:       453,
					Rid:         "ab4219087a764ae0btest",
					Name:        "Mascaras",
					Sale:        30,
					Size:        "0",
					TotalPrice:  317,
					NmID:        2389212,
					Brand:       "Vivenne Sabo",
					Status:      202,
				},
			},
			Locale:            "en",
			InternalSignature: "",
			CustomerID:        "test",
			DeliveryService:   "meest",
			ShardKey:          "9",
			SmID:              99,
			DateCreated:       time.Now(),
			OofShard:          "1",
		},
		expected: false,
		wantErr:  true,
		errMsg:   "required",
	},
	{
		name: "items required",
		order: Order{
			OrderUID:    "f",
			TrackNumber: "WBILMTESTTRACK",
			Entry:       "WBIL",
			Delivery: domainDelivery.Delivery{
				Name:    "Test Testov",
				Phone:   "+9720000000",
				Zip:     "2639809",
				City:    "Kiryat Mozkin",
				Address: "Ploshad Mira 15",
				Region:  "Kraiot",
				Email:   "test@gmail.com",
			},
			Payment: domainPayment.Payment{
				Transaction:  "b563feb7b2b84b6test",
				RequestID:    "",
				Currency:     "USD",
				Provider:     "wbpay",
				Amount:       1817,
				PaymentDt:    1637907727,
				Bank:         "alpha",
				DeliveryCost: 1500,
				GoodsTotal:   317,
				CustomFee:    0,
			},
			Locale:            "en",
			InternalSignature: "",
			CustomerID:        "test",
			DeliveryService:   "meest",
			ShardKey:          "9",
			SmID:              99,
			DateCreated:       time.Now(),
			OofShard:          "1",
		},
		expected: false,
		wantErr:  true,
		errMsg:   "required",
	},
}

package models

import (
	"time"
	db "web-api/utils/database"

	"gorm.io/gorm"
)

type ExchangeRate struct {
	ID       int     `json:"id"`
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
	Date     string  `json:"date"`
}

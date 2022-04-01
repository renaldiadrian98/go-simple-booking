package models

import (
	"go-simple-booking/helpers"
	"time"
)

type HotelType struct {
	ID        int       `json:"id" gorm:"primary_key"`
	HotelId   int       `json:"hotel_id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type HotelTypeWithRelation struct {
	ID        int       `json:"id" gorm:"primary_key"`
	HotelId   int       `json:"hotel_id"`
	Hotel     Hotel     `json:"hotel"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func HotelTypeGet(page int) ([]HotelType, error) {
	limit, offset := helpers.HelpersPaginate(page)
	var hotelType []HotelType
	err := DB.Raw(`
		SELECT *
		FROM hotel_types
		LIMIT ?
		OFFSET ?
	`, limit, offset).Scan(&hotelType).Error
	if err != nil {
		return hotelType, err
	}
	return hotelType, nil
}

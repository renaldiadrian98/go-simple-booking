package models

import (
	"go-simple-booking/helpers"
	"time"
)

type Hotel struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type HotelWithRelation struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	UserId    int       `json:"user_id"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func HotelGet(page int) ([]Hotel, error) {
	limit, offset := helpers.HelpersPaginate(page)
	var hotel []Hotel
	err := DB.Raw(`
		SELECT *
		FROM hotels
		LIMIT ? 
		OFFSET ?
	`, limit, offset).Scan(&hotel).Error
	if err != nil {
		return hotel, err
	}
	return hotel, nil
}

// func HotelUpdate(hotelId int, name string, userId int) (Hotel, error) {
// 	var hotel Hotel
// 	err := DB.Raw(`

// 	`)
// 	return hotel, err

// }

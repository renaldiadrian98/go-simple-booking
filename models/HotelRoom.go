package models

import "time"

type HotelRoom struct {
	ID          int       `json:"id" gorm:"primary_key"`
	HotelId     int       `json:"hotel_id"`
	HotelTypeId int       `json:"hotel_type_id"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type HotelRoomWithRelation struct {
	ID        int       `json:"id" gorm:"primary_key"`
	HotelId   int       `json:"hotel_id"`
	Hotel     Hotel     `json:"hotel,omitempty"`
	HotelType HotelType `json:"hotel_type,omitempty"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

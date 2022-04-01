package models

import (
	"database/sql"
	"go-simple-booking/helpers"
	"time"
)

type Transaction struct {
	ID           int          `json:"id" gorm:"primary_key"`
	UserId       int          `json:"user_id"`
	IsPaid       bool         `json:"is_paid"`
	HotelRoomId  int          `json:"hotel_room_id"`
	PaidPrice    int          `json:"paid_price"`
	PaidAt       sql.NullTime `json:"paid_at"`
	CheckinDate  time.Time    `json:"checkin_date"`
	CheckoutDate time.Time    `json:"checkout_date"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

func TransactionGet(page int, userId int) ([]Transaction, error) {
	limit, offset := helpers.HelpersPaginate(page)

	var transaction []Transaction
	err := DB.Raw(`
		SELECT *
		FROM transactions T
		WHERE
		T.user_id = ?
		LIMIT ?
		OFFSET ?
	`, userId, limit, offset).Scan(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func TransactionCheckAvailability(hotel_room_id int, checkin_date string, checkout_date string) bool {
	var transaction []Transaction
	err := DB.Raw(`
		SELECT * 
		FROM transactions T
		WHERE
		T.hotel_room_id = ?
		AND
		(T.checkin_date BETWEEN ? AND ?)
		AND
		(T.checkout_date BETWEEN ? AND ?)
		AND
		T.is_paid = true
	`, hotel_room_id, checkin_date, checkout_date, checkin_date, checkout_date).Scan(&transaction).Error
	if err != nil || len(transaction) == 0 {
		return true
	}
	return false
}

func TransactionStore(userId int, hotelRoomId int, checkinDateString string, checkoutDateString string) (Transaction, error) {
	var transaction Transaction
	checkinDate, _ := time.Parse("2006-01-02", checkinDateString)
	checkoutDate, _ := time.Parse("2006-01-02", checkoutDateString)
	transaction.UserId = userId
	transaction.HotelRoomId = hotelRoomId
	transaction.CheckinDate = checkinDate
	transaction.CheckoutDate = checkoutDate
	err := DB.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func TransactionUpdate(transactionId int, statusCode int, paidPrice int) (Transaction, error) {
	var transaction Transaction
	if statusCode == 200 {
		err := DB.Raw(`
			UPDATE transactions
			SET is_paid = true,
				paid_at = NOW(),
				paid_price = ?
			WHERE id = ?
			RETURNING *
		`, paidPrice, transactionId).Scan(&transaction).Error
		if err != nil {
			return transaction, err
		}
	}
	return transaction, nil
}

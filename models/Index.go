package models

import (
	"fmt"
	"go-simple-booking/configs"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Connect() {
	// START: Comment for heroku
	// err := godotenv.Load()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// END: Comment for heroku

	dbUserName := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// START: dataSourceName for heroku
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", dbHost, dbPort, dbUserName, dbPassword, dbName)
	// END: dataSourceName for heroku
	// dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUserName, dbPassword, dbName)

	// fmt.Println(dataSourceName)
	db, err := gorm.Open("postgres", dataSourceName)

	if err != nil {
		panic(err.Error())
	}

	// db.AutoMigrate(
	// 	&User{},
	// )

	DB = db
}

//Handle table name
type Tabler interface {
	TableName() string
}

// START: Handle pagination DATA

type PaginationData struct {
	NextPage     int `json:"next_page"`
	CurrentPage  int `json:"current_page"`
	PreviousPage int `json:"previous_page"`
	TotalPage    int `json:"total_page"`
	ItemPerPage  int `json:"item_per_page"`
}

func HelpersPaginationData(page int, tableName string) (PaginationData, error) {
	itemPerPage := configs.LimitPerPage
	var count int
	var paginationData PaginationData
	err := DB.Table(tableName).Count(&count).Error
	if err != nil {
		return paginationData, err
	}
	totalPage := count / itemPerPage
	currentPage := page
	previousPage := currentPage - 1
	if previousPage == 0 {
		previousPage = 1
	}
	nextPage := page + 1
	if (itemPerPage * page) >= count {
		nextPage = currentPage
	}

	paginationData.CurrentPage = currentPage
	paginationData.NextPage = nextPage
	paginationData.PreviousPage = previousPage
	paginationData.TotalPage = totalPage
	paginationData.ItemPerPage = itemPerPage
	return paginationData, nil
}

// END: Handle pagination DATA

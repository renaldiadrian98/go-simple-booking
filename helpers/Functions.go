package helpers

import (
	"fmt"
	"go-simple-booking/configs"

	"strconv"

	"github.com/gin-gonic/gin"
)

func HelpersPaginate(page int) (lmit int, offset int) {
	if page == 1 {
		return configs.LimitPerPage, 0
	}
	limit := page * configs.LimitPerPage
	offset = limit - configs.LimitPerPage
	return limit, offset
}

func HelpersPageQueryToInt(pageStr string) (int, error) {
	var pageInt int
	var err error

	if pageStr != "" {
		pageInt, err = strconv.Atoi(pageStr)
		if pageInt == 0 {
			return 1, nil
		}
		if err != nil {
			return pageInt, err
		}
	}

	return pageInt, nil
}

func HelpersGetUserIdInt(c *gin.Context) int {
	userId := c.MustGet("userId")
	userIdStr := fmt.Sprint(userId)
	userIdInt, _ := strconv.Atoi(userIdStr)

	return userIdInt
}

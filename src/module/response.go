package module

import (
	"math"

	"github.com/gin-gonic/gin"

	"task-management-system/src/util"
)

type ResponsePayload struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   *error      `json:"error"`
}

type responseDataPayload struct {
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
	TotalData *int64      `json:"total_data,omitempty"`
	Error     string      `json:"error,omitempty"`
}

type responsePaginatePayload struct {
	Data     interface{}     `json:"data"`
	Message  string          `json:"message"`
	Paginate paginatePayload `json:"paginate"`
}

type paginatePayload struct {
	CurrentPage int     `json:"current_page"`
	PerPage     int     `json:"per_page"`
	TotalPage   float64 `json:"total_page"`
	TotalData   int64   `json:"total_data"`
}

func ResponseData(c *gin.Context, res ResponsePayload) {
	var errMsg string
	if res.Error != nil {
		errMsg = (*res.Error).Error()
	}

	c.JSON(res.Code, responseDataPayload{
		Data:    res.Data,
		Message: res.Message,
		Error:   errMsg,
	})
}

func ResponsePaginate(c *gin.Context, paginate util.PaginationPayload, totalData int64, res ResponsePayload) {
	if paginate.Page <= 0 || paginate.Limit <= 0 {
		c.JSON(res.Code, responseDataPayload{
			Data:      res.Data,
			Message:   res.Message,
			TotalData: &totalData,
		})
		return
	}

	c.JSON(res.Code, responsePaginatePayload{
		Data:     res.Data,
		Message:  res.Message,
		Paginate: toPaginatePayload(paginate.Page, paginate.Limit, totalData),
	})
}

func toPaginatePayload(currentPage, perPage int, totalData int64) paginatePayload {
	totalPage := math.Ceil(float64(totalData) / float64(perPage))

	return paginatePayload{
		CurrentPage: currentPage,
		PerPage:     perPage,
		TotalPage:   totalPage,
		TotalData:   totalData,
	}
}

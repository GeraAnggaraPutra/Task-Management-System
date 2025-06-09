package service

import (
	"context"
	"fmt"
	"log"

	"task-management-system/src/domain/user/payload"
	"task-management-system/src/model"
)

func (s *Service) ReadUserListService(
	ctx context.Context,
	request payload.ReadUserListRequest,
) (data []model.UserDetail, totalData int64, err error) {
	statement := s.db.Table("users AS u").
		Select(`
			u.guid, u.username, u.email, u.role_guid, r.name AS role_name,
			u.created_at, u.created_by, u.updated_at, u.updated_by
		`).
		Joins("LEFT JOIN roles AS r ON u.role_guid = r.guid")

	if request.SetSearch {
		statement = statement.Where("u.username ILIKE ? OR u.email ILIKE ? OR r.name ILIKE ?", request.Search, request.Search, request.Search)
	}

	allowedSortFields := map[string]string{
		"guid":       "u.guid",
		"username":   "u.username",
		"email":      "u.email",
		"role":       "r.name",
		"created_at": "u.created_at",
		"updated_at": "u.updated_at",
	}

	if request.Sort != "" && request.Direction != "" {
		if column, ok := allowedSortFields[request.Sort]; ok {
			statement = statement.Order(fmt.Sprintf("%s %s", column, request.Direction))
		} else {
			statement = statement.Order("u.created_at DESC")
		}
	} else {
		statement = statement.Order("u.created_at DESC")
	}

	if err = statement.Count(&totalData).Error; err != nil {
		log.Printf("error count user : %v, request : %v", err, request)
		return
	}

	if request.SetPaginate {
		statement = statement.Limit(request.Limit).Offset(request.Offset)
	}

	if err = statement.Find(&data).Error; err != nil {
		log.Printf("error find user : %v, request : %v", err, request)
		return
	}

	return
}

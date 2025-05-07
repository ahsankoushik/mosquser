package dto_res

import (
	"time"
)

type (
	Response struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    any    `json:"data"`
	}
	AuthRes struct {
		ID        uint      `json:"id"`
		CreatedAt time.Time `json:"created"`
		UpdatedAt time.Time `json:"updated"`
		Email     string    `json:"email"`
		SuperUser bool      `json:"super_user"`
		Token     string    `json:"token"`
	}
	CollectionsRes struct {
		Status     int    `json:"status"`
		Message    string `json:"message"`
		Page       int    `json:"page"`
		PerPage    int    `json:"perPage"`
		TotalPages int    `json:"totalPages"`
		Data       any    `json:"data"`
	}
)

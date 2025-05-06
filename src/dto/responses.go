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
)

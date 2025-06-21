package dto_req

type (
	User struct {
		Username string `gorm:"unique,not null,index" validate:"required" json:"username"`
		Password string `validate:"required,min=8" json:"password"`
	}
)

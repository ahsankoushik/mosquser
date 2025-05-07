package dto_req

type (
	User struct {
		Email    string `gorm:"unique,not null,index" validate:"required,email" json:"email"`
		Password string `validate:"required,min=8" json:"password"`
	}
)

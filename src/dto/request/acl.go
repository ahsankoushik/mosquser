package dto_req

type (
	AclUpdate struct {
		ID    uint   `json:"id" validate:"required"`
		Topic string `json:"topic" validate:"required"`
		Acc   uint8  `json:"acc" validate:"required"`
	}
)

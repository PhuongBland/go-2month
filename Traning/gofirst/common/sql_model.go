package common

import "time"

type SQLModel struct {
	Id        int        `json:"page" form:"page"`
	Status    int        `json:"limit" form: "limit"`
	CreatedAt *time.Time `json:"create_at" form:"create_at"`
	UpdateAt  *time.Time `json:"update_at" form:"updated_at"`
}

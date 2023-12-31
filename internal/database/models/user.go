package models

import (
	"github.com/kamva/mgm/v3"
	"time"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string    `json:"email" bson:"email"`
	FullName         string    `json:"full_name" bson:"full_name"`
	Password         string    `json:"password" bson:"password"`
	CreatedAt        time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" bson:"updated_at"`
}

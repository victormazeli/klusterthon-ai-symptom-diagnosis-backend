package models

import (
	"github.com/kamva/mgm/v3"
	"time"
)

type Chat struct {
	mgm.DefaultModel `bson:",inline"`
	Message          string    `json:"message" bson:"message"`
	CreatedAt        time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" bson:"updated_at"`
}

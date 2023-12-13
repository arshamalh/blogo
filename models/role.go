package models

import (
	"github.com/uptrace/bun"
)

// Any user has a role, and any role has some permissions and accessories in the future (TODO) (logo, cool name and more)
type Role struct {
	bun.BaseModel `bun:"role"`
	Name          string `json:"name" bun:"unique"`
	Permissions   string `json:"permissions"`
}

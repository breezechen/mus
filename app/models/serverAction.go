package models

import (
	"github.com/breezechen/mus/app/utils"
	uuid "github.com/satori/go.uuid"
)

const (
	Stop    = "STOP"
	Start   = "START"
	Restart = "RESTART"
)

type ServerAction struct {
	Id         uuid.UUID  `json:"id"`
	CreateTime utils.Time `json:"create_at"`
	UpdateTime utils.Time `json:"update_at"`
	Port       string     `json:"port"`
	Action     string     `json:"action"`
}

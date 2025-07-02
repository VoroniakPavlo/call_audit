package model

import (
	"github.com/VoroniakPavlo/cases/api/cases"
)

type CaseAMQPMessage struct {
	Case *cases.Case `json:"case"`
}

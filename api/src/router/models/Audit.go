package models

import (
	"time"
)

//Auditoria para registrar passos em uma inscrição
type Audit struct {
	UserId 			uint64 `json:"user_id"`
	Registratrion 	uint64 `json:"registration"`
	ObjectType		string `json:"object_type"`
	CreateTimeStamp	time.Time `json:"create_timestamp"`
	Action			string `json:"action"`
	Message			string `json:"message"`
	Key				string `json:"key"`
	Value			string `json:"value"`
}	
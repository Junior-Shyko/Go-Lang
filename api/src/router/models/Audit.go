package models

import (
	"time"
)

//Auditoria para registrar passos em uma inscrição
type Audit struct {
	UserId 			uint64 `json:"user_id"`
	ObjectId 		string `json:"object_if"`
	ObjectType		string `json:"object_type"`
	CreateTimeStamp	time.Time `json:"create_timestamp"`
	Action			string `json:"action"`
	Message			string `json:"message"`
	Key				string `json:"key"`
	Value			string `json:"value"`
}	
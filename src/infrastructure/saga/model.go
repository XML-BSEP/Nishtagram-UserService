package saga

import (
	"encoding/json"
	"user-service/dto"
)

const (
	AuthChannel string = "AuthChannel"
	UserChannel string = "UserChannel"
	ReplyChannel string = "ReplyChannel"
	AuthService string = "Auth"
	UserService string = "User"
	ActionStart string = "Start"
	ActionDone string = "DoneMsg"
	ActionError string = "ErrorMsg"
	ActionRollback string = "RollbackMsg"
)

type Message struct {
	Service string `json:"service"`
	SenderService string `json:"sender_service"`
	Action string `json:"action"`
	Payload dto.NewUserDTO `json:"payload"`
	Confirm bool `json:"confirm"`
	Ok bool `json:"ok"`
}

func MarshalBinary(m *Message) ([]byte, error) {
	return json.Marshal(m)
}


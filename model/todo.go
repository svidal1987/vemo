package model

import (
	"errors"
	"strings"
)

const (
	StatusDone    = "done"
	StatusPending = "pending"
)

type ToDo struct {
	ID     string `json:"id"`
	User   string `json:"user"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

func (t *ToDo) IsValid() error {

	if t.Status != StatusDone && t.Status != StatusPending {
		return errors.New("Invalid Status")
	}
	if len(strings.TrimSpace(t.User)) == 0 {
		return errors.New("The user string is empty or null")
	}
	if len(strings.TrimSpace(t.Detail)) == 0 {
		return errors.New("The detail string is empty or null")
	}

	return nil
}

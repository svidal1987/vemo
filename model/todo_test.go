package model

import (
	"testing"
)

func TestToDo_IsValid(t *testing.T) {
	tests := []struct {
		name    string
		tr      ToDo
		wantErr bool
	}{
		{
			name:    "Test ok",
			tr:      ToDo{ID: "1", User: "user", Detail: "detail", Status: "done"},
			wantErr: false,
		},
		{
			name:    "Test ok",
			tr:      ToDo{ID: "1", User: "user", Detail: "detail", Status: "pending"},
			wantErr: false,
		},
		{
			name:    "Test Status not valid",
			tr:      ToDo{ID: "1", User: "user", Detail: "detail", Status: "no done"},
			wantErr: true,
		},
		{
			name:    "Test User not valid",
			tr:      ToDo{ID: "1", Detail: "detail", Status: "done"},
			wantErr: true,
		},
		{
			name:    "Test Detail not valid",
			tr:      ToDo{ID: "1", User: "user", Status: "done"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.IsValid(); (err != nil) != tt.wantErr {
				t.Errorf("ToDo.IsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package datastore

import (
	"testing"

	"github.com/vemo/model"
)

func TestToDo_Get(t *testing.T) {
	type args struct {
		id string
	}
	tr := NewDataStoreToDo()
	tests := []struct {
		name     string
		tr       *ToDo
		args     args
		wantTodo *model.ToDo
		wantOk   bool
	}{
		{
			name:     "Test OK",
			tr:       tr,
			args:     args{id: "1"},
			wantTodo: &model.ToDo{ID: "1", User: "bob", Detail: "Tarea pendiente de realizar", Status: model.StatusPending},
			wantOk:   true,
		},
		{
			name:     "Test Error: Al no existir",
			tr:       tr,
			args:     args{id: "2"},
			wantTodo: nil,
			wantOk:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTodo, gotOk := tt.tr.Get(tt.args.id)
			if tt.wantTodo != nil && gotTodo.ID != tt.wantTodo.ID {
				t.Errorf("ToDo.Get() gotTodo = %v, want %v", gotTodo, tt.wantTodo)
			}
			if gotOk != tt.wantOk {
				t.Errorf("ToDo.Get() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestToDo_Create(t *testing.T) {
	type args struct {
		todo *model.ToDo
	}
	tr := NewDataStoreToDo()
	tests := []struct {
		name string
		tr   *ToDo
		args args
	}{
		{
			name: "Test OK",
			tr:   tr,
			args: args{todo: &model.ToDo{User: "bob", Detail: "Tarea pendiente de realizar", Status: model.StatusPending}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Create(tt.args.todo)
			_, ok := tt.tr.Get(tt.args.todo.ID)
			if !ok {
				t.Errorf("ToDo.Create()")
			}
		})
	}
}

func TestToDo_Delete(t *testing.T) {
	type args struct {
		id string
	}
	tr := NewDataStoreToDo()
	tests := []struct {
		name  string
		tr    *ToDo
		args  args
		want  *model.ToDo
		want1 bool
	}{
		{
			name:  "Test OK",
			tr:    tr,
			args:  args{id: "1"},
			want:  &model.ToDo{ID: "1", User: "bob", Detail: "Tarea pendiente de realizar", Status: model.StatusPending},
			want1: true,
		},
		{
			name:  "Test Error: Al no existir",
			tr:    tr,
			args:  args{id: "2"},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.tr.Delete(tt.args.id)
			if tt.want != nil && got.ID != tt.want.ID {
				t.Errorf("ToDo.Delete() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToDo.Delete() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToDo_Update(t *testing.T) {
	type args struct {
		todo *model.ToDo
	}
	tr := NewDataStoreToDo()
	tests := []struct {
		name string
		tr   *ToDo
		args args
		want bool
	}{
		{
			name: "Test OK",
			tr:   tr,
			args: args{todo: &model.ToDo{ID: "1", User: "bob", Detail: "Tarea pendiente de realizar", Status: model.StatusPending}},
			want: true,
		},
		{
			name: "Test Error: al no existir el ToDo",
			tr:   tr,
			args: args{todo: &model.ToDo{ID: "2", User: "bob", Detail: "Tarea pendiente de realizar", Status: model.StatusPending}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Update(tt.args.todo); got != tt.want {
				t.Errorf("ToDo.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

package datastore

import (
	"strconv"
	"sync"

	"github.com/vemo/model"
)

type ToDo struct {
	m map[string]model.ToDo
	*sync.RWMutex
}

func NewDataStoreToDo() *ToDo {
	return &ToDo{
		m: map[string]model.ToDo{
			"1": {ID: "1", User: "bob", Detail: "Tarea pendiente de realizar", Status: model.StatusPending},
		},
		RWMutex: &sync.RWMutex{},
	}
}

func (t *ToDo) Get(id string) (todo model.ToDo, ok bool) {
	t.RLock()
	todo, ok = t.m[id]
	t.RUnlock()
	return todo, ok
}

func (t *ToDo) Create(todo *model.ToDo) {
	t.Lock()
	todo.ID = strconv.Itoa(len(t.m) + 1)
	t.m[todo.ID] = *todo
	t.Unlock()
}
func (t *ToDo) Delete(id string) (model.ToDo, bool) {
	t.Lock()
	todo, ok := t.m[id]
	if ok {
		delete(t.m, id)
	}
	t.Unlock()
	return todo, ok
}
func (t *ToDo) Update(todo *model.ToDo) bool {
	t.Lock()
	_, ok := t.m[todo.ID]
	if ok {
		t.m[todo.ID] = *todo
	}
	t.Unlock()
	return ok
}
func (t *ToDo) List() []model.ToDo {
	t.RLock()
	todos := make([]model.ToDo, 0, len(t.m))
	for _, v := range t.m {
		todos = append(todos, v)
	}
	t.RUnlock()
	return todos
}

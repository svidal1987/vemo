package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vemo/datastore"
	"github.com/vemo/model"
)

type ToDo struct {
	Store *datastore.ToDo
}

func HandlerToDo(mux *http.ServeMux) {
	handler := &ToDo{
		Store: datastore.NewDataStoreToDo(),
	}
	mux.HandleFunc("GET /todos", handler.list)
	mux.HandleFunc("POST /todos", handler.create)
	mux.HandleFunc("GET /todos/{id}", handler.get)
	mux.HandleFunc("PUT /todos/{id}", handler.update)
	mux.HandleFunc("DELETE /todos/{id}", handler.delete)

}

func (h *ToDo) list(w http.ResponseWriter, r *http.Request) {
	todos := h.Store.List()
	jsonBytes, err := json.Marshal(todos)
	if err != nil {
		InternalServerError(w)
		return
	}
	ResponseOk(w, jsonBytes)
}

func (h *ToDo) get(w http.ResponseWriter, r *http.Request) {

	u, ok := h.Store.Get(r.PathValue("id"))
	if !ok {
		NotFound(w)
		return
	}
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		InternalServerError(w)
		return
	}
	ResponseOk(w, jsonBytes)
}

func (h *ToDo) create(w http.ResponseWriter, r *http.Request) {
	var todo model.ToDo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		InternalServerError(w)
		return
	}
	if err := todo.IsValid(); err != nil {
		BadRquest(w, fmt.Sprint(err))
		return
	}

	h.Store.Create(&todo)

	jsonBytes, err := json.Marshal(todo)
	if err != nil {
		InternalServerError(w)
		return
	}
	ResponseOk(w, jsonBytes)
}

func (h *ToDo) delete(w http.ResponseWriter, r *http.Request) {

	u, ok := h.Store.Delete(r.PathValue("id"))

	if !ok {
		NotFound(w)
		return
	}
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		InternalServerError(w)
		return
	}
	ResponseOk(w, jsonBytes)
}
func (h *ToDo) update(w http.ResponseWriter, r *http.Request) {
	var todo model.ToDo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		InternalServerError(w)
		return
	}
	if err := todo.IsValid(); err != nil {
		BadRquest(w, fmt.Sprint(err))
		return
	}

	if r.PathValue("id") != todo.ID {
		BadRquest(w, "The ID parameters are different")
		return
	}

	ok := h.Store.Update(&todo)
	if !ok {
		NotFound(w)
		return
	}
	jsonBytes, err := json.Marshal(todo)
	if err != nil {
		InternalServerError(w)
		return
	}
	ResponseOk(w, jsonBytes)
}

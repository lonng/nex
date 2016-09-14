package nex

import (
	"encoding/json"
	"net/http"
	"reflect"
)

type HandlerAdapter interface {
	Invoke(http.ResponseWriter, *http.Request)
}

type getRequestAdapter struct {
	method reflect.Value
}

type postRequestAdapter struct {
	argType reflect.Type
	method  reflect.Value
}

func (g *getRequestAdapter) Invoke(w http.ResponseWriter, r *http.Request) {
	ret := g.method.Call([]reflect.Value{})
	if err := ret[1].Interface(); err != nil {
		fail(w, err.(error))
		return
	}

	json.NewEncoder(w).Encode(ret[0].Interface())
}

func (p *postRequestAdapter) Invoke(w http.ResponseWriter, r *http.Request) {
	data := reflect.New(p.argType.Elem()).Interface()
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		fail(w, err)
		return
	}

	ret := p.method.Call([]reflect.Value{reflect.ValueOf(data)})
	if err := ret[1].Interface(); err != nil {
		fail(w, err.(error))
		return
	}

	json.NewEncoder(w).Encode(ret[0].Interface())
}

package nex

import (
	"encoding/json"
	"net/http"
	"reflect"
)

type ErrorEncoder func(error) interface{}

type DefaultErrorMessage struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type handler struct {
	typ    reflect.Type
	method reflect.Value
}

var errorEncoder ErrorEncoder

func fail(w http.ResponseWriter, err error) {
	errMsg := errorEncoder(err)
	json.NewEncoder(w).Encode(errMsg)
	return
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := reflect.New(h.typ.In(0).Elem()).Interface()
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		fail(w, err)
		return
	}

	ret := h.method.Call([]reflect.Value{reflect.ValueOf(data)})
	if err := ret[1].Interface(); err != nil {
		fail(w, err.(error))
		return
	}

	json.NewEncoder(w).Encode(ret[0].Interface())
}

func Handler(f interface{}) http.Handler {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		panic("invalid parameter")
	}

	if t.NumIn() != 1 || t.In(0).Kind() != reflect.Ptr {
		panic("unsupport function type, function should accept one pointer parameter")
	}

	if t.NumOut() != 2 {
		panic("unsupport function type, function return values should contain response data or error")
	}
	return &handler{t, reflect.ValueOf(f)}
}

func SetErrorEncoder(c ErrorEncoder) {
	if c == nil {
		panic("nil pointer to error encoder")
	}
	errorEncoder = c
}

func init() {
	errorEncoder = func(err error) interface{} {
		println(err.Error())
		return &DefaultErrorMessage{
			Code:  -1,
			Error: err.Error(),
		}
	}
}

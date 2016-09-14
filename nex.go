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
	adapter HandlerAdapter
}

var errorEncoder ErrorEncoder

func fail(w http.ResponseWriter, err error) {
	errMsg := errorEncoder(err)
	json.NewEncoder(w).Encode(errMsg)
	return
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.adapter.Invoke(w, r)
}

func Handler(f interface{}) http.Handler {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		panic("invalid parameter")
	}

	if t.NumOut() != 2 {
		panic("unsupport function type, function return values should contain response data or error")
	}

	var adapter HandlerAdapter
	var num = t.NumIn()

	if num == 0 {
		adapter = &getRequestAdapter{reflect.ValueOf(f)}
	} else if num == 1 && t.In(0).Kind() == reflect.Ptr {
		adapter = &postRequestAdapter{t.In(0), reflect.ValueOf(f)}
	} else {
		panic("unsupport function type, function should accept one pointer parameter")
	}

	return &handler{adapter}
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

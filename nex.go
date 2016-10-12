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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(errMsg)
}

func succ(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
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
		adapter = &nonParameterAdapter{reflect.ValueOf(f)}
	} else if num == 1 && !isSupportType(t.In(0)) && t.In(0).Kind() == reflect.Ptr {
		adapter = &simpleUnaryAdapter{t.In(0), reflect.ValueOf(f)}
	} else {
		adapter = makeGenericAdapter(reflect.ValueOf(f))
	}

	return &handler{adapter}
}

func SetErrorEncoder(c ErrorEncoder) {
	if c == nil {
		panic("nil pointer to error encoder")
	}
	errorEncoder = c
}

func SetMultipartFormMaxMemery(m int64) {
	maxMemory = m
}

func init() {
	errorEncoder = func(err error) interface{} {
		return &DefaultErrorMessage{
			Code:  -1,
			Error: err.Error(),
		}
	}
}

package nex

import "reflect"

func Handler(f interface{}) *Nex {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		panic("invalid parameter")
	}

	numOut := t.NumOut()

	if numOut != 2 && numOut != 3 {
		panic("unsupport function type, function return values should contain response data & error")
	}

	if numOut == 3 {
		o0 := t.Out(0)
		if o0 != contextType {
			panic("unsupport function type")
		}
	}

	var (
		adapter    HandlerAdapter
		numIn      = t.NumIn()
		outContext = numOut == 3
		inContext  = false
	)

	if numIn > 0 {
		for i := 0; i < numIn; i++ {
			if t.In(i) == contextType {
				inContext = true
			}
		}
	}

	if numIn == 0 {
		adapter = &simplePlainAdapter{false, outContext, reflect.ValueOf(f)}
	} else if numIn == 1 && inContext {
		adapter = &simplePlainAdapter{true, outContext, reflect.ValueOf(f)}
	} else if numIn == 1 && !isSupportType(t.In(0)) && t.In(0).Kind() == reflect.Ptr {
		adapter = &simpleUnaryAdapter{outContext, t.In(0), reflect.ValueOf(f)}
	} else {
		adapter = makeGenericAdapter(reflect.ValueOf(f), inContext, outContext)
	}

	return &Nex{adapter: adapter}
}

func SetErrorEncoder(c ErrorEncoder) {
	if c == nil {
		panic("nil pointer to error encoder")
	}
	errorEncoder = c
}

func SetMultipartFormMaxMemory(m int64) {
	maxMemory = m
}



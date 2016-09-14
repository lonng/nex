# nex
This library aiming to simplify the construction of JSON API service, `nex.Handler`
wrap a function to `http.Handler`, which unmarshal POST data to struct automatically.

## Usage
```
package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/chrislonng/nex"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Result string `json:"result"`
}

type ErrorMessage struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func main() {
	// customize error encoder
	nex.SetErrorEncoder(func(err error) interface{} {
		return &ErrorMessage{Code: -1, Error: err.Error()}
	})

	mux := http.NewServeMux()
	// post request with data
	mux.Handle("/test1", nex.Handler(test1))
	mux.Handle("/test2", nex.Handler(test2))

	// get request without data
	mux.Handle("/test3", nex.Handler(test3))
	mux.Handle("/test4", nex.Handler(test4))

	http.ListenAndServe(":8080", mux)
}

// POST: regular response
func test1(m *LoginRequest) (*LoginResponse, error) {
	fmt.Printf("%+v\n", m)
	return &LoginResponse{Result: "success"}, nil
}

// POST: error response
func test2(m *LoginRequest) (*LoginResponse, error) {
	fmt.Printf("%+v\n", m)
	return nil, errors.New("error test")
}

// GET: regular response
func test3() (*LoginResponse, error) {
	return &LoginResponse{Result: "success"}, nil
}

// GET: error response
func test4() (*LoginResponse, error) {
	return nil, errors.New("error test")
}

```

```
curl -XPOST -d '{"username":"test", "password":"test"}' http://localhost:8080/test1
curl -XPOST -d '{"username":"test", "password":"test"}' http://localhost:8080/test2
curl  http://localhost:8080/test3
curl  http://localhost:8080/test4
```

## License
Copyright (c) <2016> <chris@lonng.org>


Permission is hereby granted, free of charge, to any person obtaining a copy of 
this software and associated documentation files (the "Software"), to deal in 
the Software without restriction, including without limitation the rights to use, 
copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the 
Software, and to permit persons to whom the Software is furnished to do so, subject 
to the following conditions:

The above copyright notice and this permission notice shall be included in all 
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, 
INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A 
PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT 
HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION 
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE 
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

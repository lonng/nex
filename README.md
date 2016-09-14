# nex
This library aims to simplify the construction of JSON API service,
`nex.Handler` is able to wrap any function to adapt the interface of
`http.Handler`, which unmarshals POST data to a struct automatically.

## Support types
```
io.ReadCloser      // request.Body
http.Header        // request.Header
nex.Form           // request.Form
nex.PostFrom       // request.PostFrom
*url.URL           // request.URL
*multipart.Form    // request.MultipartForm
```

## Usage
```
http.Handle("/test", nex.Handler(test))

func test(io.ReadCloser, http.Header, nex.Form, nex.PostFrom, *CustomizedRequestType, *url.URL, *multipart.Form) (*CustomizedResponseType, error)
```

## Example
```
package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

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
	mux.Handle("/test1", nex.Handler(test1))
	mux.Handle("/test2", nex.Handler(test2))
	mux.Handle("/test3", nex.Handler(test3))
	mux.Handle("/test4", nex.Handler(test4))
	mux.Handle("/test5", nex.Handler(test5))
	mux.Handle("/test6", nex.Handler(test6))
	mux.Handle("/test7", nex.Handler(test7))
	mux.Handle("/test8", nex.Handler(test8))

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

func test5(header http.Header) (*LoginResponse, error) {
	fmt.Printf("%#v\n", header)
	return &LoginResponse{Result: "success"}, nil
}

func test6(form nex.Form) (*LoginResponse, error) {
	fmt.Printf("%#v\n", form)
	return &LoginResponse{Result: "success"}, nil
}

func test7(header http.Header, form nex.Form, body io.ReadCloser) (*LoginResponse, error) {
	fmt.Printf("%#v\n", header)
	fmt.Printf("%#v\n", form)
	return &LoginResponse{Result: "success"}, nil
}

func test8(header http.Header, r *LoginResponse, url *url.URL) (*LoginResponse, error) {
	fmt.Printf("%#v\n", header)
	fmt.Printf("%#v\n", r)
	fmt.Printf("%#v\n", url)
	return &LoginResponse{Result: "success"}, nil
}
```

```
curl -XPOST -d '{"username":"test", "password":"test"}' http://localhost:8080/test1
curl -XPOST -d '{"username":"test", "password":"test"}' http://localhost:8080/test2
curl  http://localhost:8080/test3
curl  http://localhost:8080/test4
curl  http://localhost:8080/test5
curl  http://localhost:8080/test6\?test\=test
curl  http://localhost:8080/test7\?test\=tset
curl -XPOST -d '{"username":"test", "password":"test"}' http://localhost:8080/test8\?test\=test
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

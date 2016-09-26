package nex

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"testing"
)

func TestHandler(t *testing.T) {
	Handler(withBody)
	Handler(withReq)
	Handler(withHeader)
	Handler(withForm)
	Handler(withPostForm)
	Handler(withMultipartForm)
	Handler(withUrl)
	Handler(withAll)
}

type testRequest struct{}
type testResponse struct{}

func withBody(io.ReadCloser) (*testResponse, error) {
	return nil, nil
}

func withReq(*testRequest) (*testResponse, error) {
	return nil, nil
}

func withHeader(http.Header) (*testResponse, error) {
	return nil, nil
}

func withForm(Form) (*testResponse, error) {
	return nil, nil
}

func withPostForm(PostForm) (*testResponse, error) {
	return nil, nil
}

func withMultipartForm(*multipart.Form) (*testResponse, error) {
	return nil, nil
}

func withUrl(*url.URL) (*testResponse, error) {
	return nil, nil
}

func withAll(io.ReadCloser, *testRequest, Form, PostForm, http.Header, *multipart.Form, *url.URL) (*testResponse, error) {
	return nil, nil
}

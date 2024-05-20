package entity

import (
	"encoding/json"
	"errors"

	"github.com/valyala/fastjson"
)

var (
	ErrToken        = errors.New("error-token")
	ErrWrongToken   = errors.New("error-wrongToken")
	ErrNotPremium   = errors.New("error-notPremium")
	ErrorNotFound   = errors.New("error-notFound")
	ErrorContentsId = errors.New("error-contentsId")
	ErrorType       = errors.New("error-type")
	// Custom errors
	ErrEmptyStatus    = errors.New("error-emptyStatus")
	ErrPrivateContent = errors.New("error-privateContent")
	ErrAccount        = errors.New("error-account")
)

var responseParserPool fastjson.ParserPool

func init() {
	responseParserPool = fastjson.ParserPool{}
}

type AnyDataResponse Response[interface{}]

func (r AnyDataResponse) Json() (*fastjson.Value, error) {
	b, err := r.Marshal()
	if err != nil {
		return nil, err
	}
	p := responseParserPool.Get()
	defer responseParserPool.Put(p)
	return p.Parse(b)
}

func (r *AnyDataResponse) Marshal() (string, error) {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (r *AnyDataResponse) IsError() bool {
	return r.Status != "ok"
}

func (r *AnyDataResponse) Error() error {
	if r.IsError() {
		return ErrorResponseStatus(r.Status)
	}
	return nil
}

type EmptyDataResponse Response[struct{}]

func (r EmptyDataResponse) Json() (*fastjson.Value, error) {
	b, err := r.Marshal()
	if err != nil {
		return nil, err
	}
	p := responseParserPool.Get()
	defer responseParserPool.Put(p)
	return p.Parse(b)
}

func (r *EmptyDataResponse) Marshal() (string, error) {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (r *EmptyDataResponse) IsError() bool {
	return r.Status != "ok"
}

func (r *EmptyDataResponse) Error() error {
	if r.IsError() {
		return ErrorResponseStatus(r.Status)
	}
	return nil
}

type Response[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
}

func (r *Response[T]) Json() (*fastjson.Value, error) {
	b, err := r.Marshal()
	if err != nil {
		return nil, err
	}
	p := responseParserPool.Get()
	defer responseParserPool.Put(p)
	return p.Parse(b)
}

func (r *Response[T]) Marshal() (string, error) {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (r *Response[T]) IsError() bool {
	return r.Status != "ok"
}

func (r *Response[T]) Error() error {
	if r.IsError() {
		return ErrorResponseStatus(r.Status)
	}
	return nil
}

func ErrorResponseStatus(status string) error {
	if status == "ok" {
		return nil
	}
	switch status {
	case "error-token":
		return ErrToken
	case "error-wrongToken":
		return ErrWrongToken
	case "error-notPremium":
		return ErrNotPremium
	case "error-privateContent":
		return ErrPrivateContent
	case "error-notFound":
		return ErrorNotFound
	case "error-contentsId":
		return ErrorContentsId
	case "error-type":
		return ErrorType
	case "error-account":
		return ErrAccount
	case "":
		return ErrEmptyStatus
	}
	return errors.New(status)
}

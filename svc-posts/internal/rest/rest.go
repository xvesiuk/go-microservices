package rest

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrDecodeBody = errors.New("error while decoding body")
	ErrParseBody  = errors.New("error while parsing body")
	ErrInternal   = errors.New("internal server error")
)

type ResponseDefault struct {
	Msg   string `json:"msg"`
	Error bool   `json:"error"`
}

type Response[T any] struct {
	Data T `json:"data"`
	ResponseDefault
}

func WriteOk[T any](w http.ResponseWriter, status int, data T) error {
	res, err := json.Marshal(&Response[T]{
		Data: data,
	})
	if err != nil {
		return err
	}
	w.WriteHeader(status)
	_, err = w.Write(res)
	if err != nil {
		panic("That's so wrong!")
	}
	return nil
}

func WriteError(w http.ResponseWriter, status int, e error) {
	w.WriteHeader(status)
	_, err := w.Write([]byte(e.Error()))
	if err != nil {
		panic("That's so wrong!")
	}
}

func NotImplementedHandler(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented!")
}

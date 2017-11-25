package main

import (
	"errors"
	"net/http"
)

//http 通过自定义路由器 统一处理异常
/*
路由表接口
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)  // 路由实现器
}
*/

type appHandler func(http.ResponseWriter, *http.Request) *appError //接口类型定义
type appError struct {                                             //自定义异常
	Error   error
	Message string
	Code    int
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { //e is appError
		http.Error(w, e.Message, e.Code)
	}
}

func viewRecord(w http.ResponseWriter, r *http.Request) *appError {
	return &appError{errors.New("exception ...."), "not found", 404}
}

func init() {
	http.Handle("/error", appHandler(viewRecord)) //viewRecord 转换成appHandler
}

func main() {
	http.ListenAndServe(":8080", nil)
}

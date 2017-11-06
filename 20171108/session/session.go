package session

import (
	"fmt"
)

//https://github.com/astaxie/session

//session接口定义
type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

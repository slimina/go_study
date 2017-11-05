package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
Go目前支持redis的驱动有如下

https://github.com/garyburd/redigo (推荐)，https://github.com/astaxie/goredis
https://github.com/go-redis/redis
https://github.com/hoisie/redis
https://github.com/alphazero/Go-Redis
https://github.com/simonz05/godis
https://github.com/FZambia/go-sentinel
*/

var (
	Pool *redis.Pool
)

//初始化,go加载默认运行
func init() {
	redisHost := "127.0.0.1:6379"
	Pool = newPool(redisHost, "abc123", 0)
	Close()
}

func newPool(sever string, password string, db uint8) *redis.Pool {
	return &redis.Pool{
		MaxActive:   30,
		MaxIdle:     3,
		IdleTimeout: 3 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", sever)
			if err != nil {
				return nil, err
			}
			if _, err := conn.Do("AUTH", password); err != nil {
				conn.Close()
				return nil, err
			}
			if _, err := conn.Do("SELECT", db); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, nil
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}

//优雅退出
func Close() {
	c := make(chan os.Signal, 1)
	/*
		Golang 的系统信号处理主要涉及os包、os.signal包以及syscall包
		func Notify(c chan<- os.Signal, sig …os.Signal),将进程收到的系统Signal转发给channel c
		如果你没有传入sig参数，那么Notify会将系统收到的所有信号转发给c
		signal.Notify(c, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2)
		Go只会关注你传入的Signal类型，其他Signal将会按照默认方式处理，大多都是进程退出。
	*/
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c //接收到一个
		Pool.Close()
		os.Exit(0)
	}()
}

func Get(key string) ([]byte, error) {
	conn := Pool.Get()
	defer conn.Close()
	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}
	return data, nil
}
func main() {

	test, err := Get("test")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(test))

}

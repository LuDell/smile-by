package utils

import (
	"github.com/gomodule/redigo/redis"
	"time"
	"log"
)

var RedisPool *redis.Pool

func init()  {
	RedisPool = &redis.Pool{
		MaxIdle:2,
		MaxActive:5,
		IdleTimeout:6000 *time.Second,
		Dial: func() (redis.Conn, error) {
			con,err1 :=redis.Dial("tcp","45.77.82.85:6388")
			if err1 != nil {
				log.Fatal(err1)
				con.Close()
			}
			_,err2 := con.Do("AUTH","1q2w3e4r")
			if err2 != nil {
				log.Fatal(err2)
				con.Close()
			}
			return con,err1
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_,err := c.Do("PING")
			return err
		},
	}
}
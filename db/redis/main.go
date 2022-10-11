package main

import (
	"fmt"
	"time"

	redis "github.com/gomodule/redigo/redis"
)

func main() {
	set("maxm", 35)
	set("tq", 33)
	set("lqh", 30)
	age, err := get("maxm")
	if err != nil {
		fmt.Errorf("get err, %+v\n", err)
	} else {
		fmt.Printf("age = %+v\n", age)
	}
}

func set(key string, val interface{}) error {
	conn, err := getRedisConnFromPool()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Do("Set", key, val)
	if err != nil {
		fmt.Println(fmt.Errorf("set cmd err, %v", err))
		panic("set cmd err")
	}
	return nil
}

func get(key string) (interface{}, error) {
	conn, err := getRedisConnFromPool()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	r, err := redis.String(conn.Do("get", key))
	// if err != nil {
	// 	fmt.Errorf("redis get error, err=%+v\n", err)
	// 	return nil
	// }
	// return r
	return r, err
}

func getRedisConnFromPool() (redis.Conn, error) {
	pool := &redis.Pool{
		Dial:            func() (redis.Conn, error) { return getRedisConn() },
		MaxIdle:         5,
		MaxActive:       18,
		IdleTimeout:     240 * time.Second,
		MaxConnLifetime: 300 * time.Second,
	}
	return pool.Get(), nil
}

func getRedisConn() (redis.Conn, error) {
	conn, err := redis.Dial("tcp",
		"127.0.0.1:6379",
		redis.DialConnectTimeout(5*time.Second),
		redis.DialReadTimeout(2*time.Second),
		redis.DialWriteTimeout(2*time.Second),
		redis.DialPassword("Maxm#@!123"),
		redis.DialKeepAlive(2*time.Second),
	)
	return conn, err
}

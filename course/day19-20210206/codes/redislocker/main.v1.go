package main

import (
	"fmt"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/google/uuid"
)

func main() {
	addr := "10.0.0.2:6379"
	password := "18df72ec41b641aa402fd845c1f5ebea"
	dsn := fmt.Sprintf("redis://kk:%s@%s/0", password, addr)
	conn, err := redis.DialURL(dsn, redis.DialPassword(password))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// key
	key := fmt.Sprintf("locker:redis:%s", "test")
	// 标识
	tag := uuid.New().String()
	interval := 10 // >执行时间 (执行时间 * n > 2)
	go func() {
		for range time.Tick(time.Second) {
			conn.Do("SET", key, tag, "EX", interval, "NX")
			locker, err := redis.String(conn.Do("Get", key))
			if err == nil && locker == tag {
				redis.Do("EXPIRE", key, interval)
			}
		}
	}()

	for {
		locker, err := redis.String(conn.Do("GET", key))
		if err != nil {
			continue
		}
		if locker != tag {
			log.Printf("locker: %s", locker)
			time.Sleep(time.Duration(interval/3) * time.Second)
			continue
		}
		log.Println("exec")
		time.Sleep(time.Second * 2)
	}

}

package main

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
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

	name := fmt.Sprintf("queue:%s", "test")
	// 生产者 producer

	for {
		r, err := redis.Values(conn.Do("BRPOP", name, 3))
		if err != nil {
			if err != redis.ErrNil {
				break
			}
			continue
		}
		var (
			name string
			t    int
		)
		if _, err := redis.Scan(r, &name, &t); err == nil {
			fmt.Println(name, t)
		}

	}
}

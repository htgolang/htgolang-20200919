package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisQueue struct {
	conn redis.Conn
	name string
}

func NewRedisQueue(conn redis.Conn, name string) *RedisQueue {
	return &RedisQueue{conn, fmt.Sprintf("queue:%s", name)}
}

func (q *RedisQueue) Push(e string) {
	q.conn.Do("LPUSH", q.name, e)
}
func (q *RedisQueue) Pop(timeout int) string {
	values, err := redis.Values(q.conn.Do("BRPOP", q.name, timeout))
	if err != nil {
		return ""
	}
	var (
		name  string
		value string
	)
	if _, err := redis.Scan(values, &name, &value); err == nil {
		return value
	}
	return ""

}

func main() {
	addr := "10.0.0.2:6379"
	password := "18df72ec41b641aa402fd845c1f5ebea"
	conn, err := redis.Dial("tcp", addr, redis.DialPassword(password))

	if err != nil {
		log.Fatal(err)
	}
	queue := NewRedisQueue(conn, "time")
	go func() {
		for range time.Tick(time.Second * 3) {
			queue.Push(time.Now().Format("15:04:05"))
		}
	}()

	for {
		value := queue.Pop(3)
		if value != "" {
			log.Println("value: ", value)
		}
	}
}

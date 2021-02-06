package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	addr := "10.0.0.2:6379"
	password := "18df72ec41b641aa402fd845c1f5ebea"
	conn, err := redis.Dial("tcp", addr, redis.DialPassword(password))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	args := redis.Args{}
	args = args.Add("time2")
	args = args.Add(time.Now().Unix())
	args = args.Add("EX")
	args = args.Add(60)
	args = args.Add("NX")
	// conn.Do("SET", "time", time.Now().Unix(), "EX", 60, "NX")
	conn.Do("SET", args...)
	t, err := redis.Int(conn.Do("GET", "time2"))
	fmt.Println(t, err)
}

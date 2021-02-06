package main

import (
	"fmt"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/google/uuid"
)

type Locker struct {
	conn     redis.Conn
	name     string
	interval int
	tag      string
}

func NewLocker(conn redis.Conn, name string, interval int) *Locker {
	locker := &Locker{
		conn:     conn,
		name:     fmt.Sprintf("locker:redis:%s", name),
		interval: interval,
		tag:      uuid.New().String(),
	}
	go locker.boostrap()
	return locker
}

func (l *Locker) boostrap() {
	for range time.Tick(time.Second) {
		l.conn.Do("SET", l.name, l.tag, "EX", l.interval, "NX")
		locker, err := redis.String(l.conn.Do("Get", l.name))
		if err == nil && locker == l.tag {
			l.conn.Do("EXPIRE", l.name, l.interval)
		}
	}
}

func (l *Locker) Lock() bool {
	locker, err := redis.String(l.conn.Do("GET", l.name))
	if err != nil {
		return false
	}
	return locker == l.tag
}

func main() {
	addr := "10.0.0.2:6379"
	password := "18df72ec41b641aa402fd845c1f5ebea"
	dsn := fmt.Sprintf("redis://kk:%s@%s/0", password, addr)
	conn, err := redis.DialURL(dsn, redis.DialPassword(password))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	locker := NewLocker(conn, "test", 10)
	for {
		if !locker.Lock() {
			log.Println("not locker")
			time.Sleep(3 * time.Second)
			continue
		}
		log.Println("exec")
		time.Sleep(time.Second * 2)

	}
}

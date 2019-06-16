package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
)

const Port = 8080

var redisPool *redis.Pool

func NewRedisPool() *redis.Pool {
	if redisPool != nil {
		return redisPool
	}
	redisPool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", os.Getenv("REDIS_HOST"))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return redisPool
}

var sync *redsync.Redsync

func NewSync() *redsync.Redsync {
	if sync != nil {
		return sync
	}
	var pools []redsync.Pool
	pools = append(pools, NewRedisPool())
	sync = redsync.New(pools)
	return sync
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	sync := NewSync()
	mutex := sync.NewMutex("test-redsync")
	if err := mutex.Lock(); err != nil {
		http.Error(w, fmt.Sprintf("redsync lock failed. err = %v", err.Error()), 500)
		return
	}
	log.Println("lock")
	defer func() {
		log.Println("unlock")
		if ok := mutex.Unlock(); !ok {
			http.Error(w, "redsync unlock failed", 500)
			return
		}
	}()
}

func main() {
	log.Printf("listen: %s", os.Getenv("LISTEN_PORT"))
	r := chi.NewRouter()
	r.Get("/", indexHandler)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("LISTEN_PORT")), r); err != nil {
		panic(err)
	}
}

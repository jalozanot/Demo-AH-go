package database_client

import (
	"flag"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var (
	redisAddress   = flag.String("localhost", "localhost:6379", "Address to the Redis server")
	maxConnections = flag.Int("max-connections", 10, "Max connections to Redis")
)

type movie struct {
	Id        int64
	Nombre    string
	categoria string
}

var Con redis.Conn

func GetConnectionRedis() {

	//Redis Connection
	redisPool := redis.NewPool(func() (redis.Conn, error) {

		con, err := redis.Dial("tcp", "localhost:6379")
		con.Do("SELECT", 1)
		if err != nil {
			return nil, err
		}
		return con, err
	}, *maxConnections)

	fmt.Println("Redis Connecting...!")
	Con = redisPool.Get()

}

package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Person struct {
	Name string
	Age  int
	Id   int
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, //бд по умолчанию
	})
	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		return
	}
	fmt.Println("Redis response to ping:", ping)
	test(client)

	//

	ivan, _ := json.Marshal(Person{
		Name: "Ivan",
		Age:  25,
		Id:   0,
	})
	client.RPush(context.Background(), "person", ivan)
	val, _ := client.LRange(context.Background(), "person", 0, 0).Result()

	fmt.Println(val)

	res := Person{}
	json.Unmarshal([]byte(val[0]), &res)
	fmt.Println(res)
}

func test(client *redis.Client) {
	err := client.Set(context.Background(), "ferstKey", "Hellow", 0).Err()
	if err != nil {
		return
	}
	val1, err := client.Get(context.Background(), "ferstKey").Result()
	if err != nil {
		return
	}
	fmt.Println("1)", val1)

	val2, _ := client.Keys(context.Background(), "*").Result()
	fmt.Println("2)", val2)
}
